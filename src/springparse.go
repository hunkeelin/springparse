package springparse

import (
	"bytes"
	"github.com/papertrail/go-tail/follower"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
)

// Client the struct that runs the program
type Client struct {
	tailedFiles map[string]int
}

// New creates a client struct with map initialized
func New() *Client {
	m := make(map[string]int)
	return &Client{
		tailedFiles: m,
	}
}

// runner runner spawn by go routine
type runner struct {
	buffer   *elasticOut
	bufferId string
}

// SpringParse This is the main program
func (r *Client) SpringParse() {
	logFiles, err := r.listDirectory()
	if err != nil {
		log.Error(err.Error())
		return
	}
	for _, fi := range logFiles {
		result := r.shouldWatch(shouldWatchInput{
			logFile: fi,
		})
		_, ok := r.tailedFiles[fi]
		if result.watch && !ok {
			r.tailedFiles[fi] = 0
			go func() {
				newRunner := runner{}
				newRunner.tailFile(fi)
			}()
		}
	}
	return
}

func (r *runner) tailFile(fileName string) {
	t, err := follower.New(fileName, follower.Config{
		Whence: io.SeekEnd,
		Offset: 0,
		Reopen: true,
	})
	if err != nil {
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
		putSuccess.Inc()
	}

	if t.Err() != nil {
		log.Error(t.Err())
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
