package springparse

import (
	"bytes"
	"github.com/fsnotify/fsnotify"
	"github.com/hunkeelin/go-tail/follower"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// Client the struct that runs the program
type Client struct {
	tailedFiles map[string]int
}

// New creates a client struct with map initialized
func New() *Client {
	m := make(map[string]int)
	flushSig = make(chan bool)
	sendItems = make(chan elasticItem)
	go sendBatch(sendItems, flushSig)
	return &Client{
		tailedFiles: m,
	}
}

// Runner Runner spawn by go routine
type Runner struct {
	Buffer       *elasticOut
	BufferId     string
	sendbuffer   []elasticOutInfo
	sendbufferMu sync.Mutex
	bufferCount  int
}
type elasticOutInfo struct {
	record *elasticOut
	id     string
}

// SpringParse This is the main program
func (r *Client) SpringParse() {
	logFiles, err := r.listDirectory()
	if err != nil {
		log.Error("Unable to list directory: " + err.Error())
		return
	}
	go func() {
		for {
			time.Sleep(time.Duration(flushCycleInt) * time.Second)
			flushSig <- true
		}
	}()
	for _, fi := range logFiles {
		result := r.shouldWatch(shouldWatchInput{
			logFile: fi,
		})
		_, ok := r.tailedFiles[fi]
		if result.watch && !ok {
			log.Info("Tailing " + fi)
			r.tailedFiles[fi] = 0
			newRunner := Runner{}
			go newRunner.tailFile(fi)
		}
	}
	go r.tailDir()
	return
}

func (r *Runner) tailFile(fileName string) {
	t, err := follower.New(fileName, follower.Config{
		Whence: io.SeekEnd,
		Offset: 0,
		Reopen: true,
	})
	if err != nil {
		log.Error("Unable to create follower: " + err.Error())
		return
	}

	for line := range t.Lines() {
		err := r.sendElasticSearch(sendElasticSearchInput{
			rawLog:   bytes.Replace(line.Bytes(), []byte("\n"), []byte(""), -1),
			fileName: fileName,
		})
		if err != nil {
			putFailed.Inc()
			log.Error(err.Error())
			continue
		}
	}

	if t.Err() != nil {
		log.Error("Follower error: " + t.Err().Error())
		return
	}
}

func (r *Client) listDirectory() ([]string, error) {
	var files []string
	err := filepath.Walk(logDirectory, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		return files, err
	}
	return files, nil
}
func (r *Client) tailDir() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	defer watcher.Close()
	err = watcher.Add(logDirectory)
	if err != nil {
		panic(err)
	}
	errCh := make(chan error)
	for {
		select {
		case event := <-watcher.Events:
			if event.Op&fsnotify.Create == fsnotify.Create {
				result := r.shouldWatch(shouldWatchInput{
					logFile: event.Name,
				})
				_, ok := r.tailedFiles[event.Name]
				if result.watch && !ok {
					log.Info("Newly tailed file " + event.Name)
					r.tailedFiles[event.Name] = 0
					newRunner := Runner{}
					go newRunner.tailFile(event.Name)
				}
			}
		case err := <-watcher.Errors:
			errCh <- err
		}
	}
	<-errCh
}
