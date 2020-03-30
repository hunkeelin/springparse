package springparse

import (
	"bytes"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"os/exec"
)

// main
func SpringParse() error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	done := make(chan bool)

	go func() {
		for {
			select {
			// watch for events
			case event := <-watcher.Events:
				switch {
				//case event.Op == fsnotify.Create:
				//case event.Op == fsnotify.Remove:
				case event.Op == fsnotify.Write:
					result := shouldWatch(shouldWatchInput{
						logFile: event.Name,
					})
					if result.watch {
						acmd := exec.Command(tailBinary, "-n", "1", event.Name)
						out, _ := acmd.Output()
						out = bytes.Replace(out, []byte("\n"), []byte(""), -1)
						fmt.Println(string(out))
					}
				}
				// watch for errors
			case err := <-watcher.Errors:
				fmt.Println("ERROR", err)
			}
		}
	}()

	// out of the box fsnotify can watch a single file, or a single directory
	if err := watcher.Add(logDirectory); err != nil {
	}
	<-done
	return nil
}

func reverseByteArray(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
