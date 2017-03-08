package backend

import (
	"io/ioutil"
	"flag"
	"path"
	"time"
)

var (
	watcherQueueSize = flag.Int(
		"watcher_queue_size", 100, "Watcher max queue size")
)

type DirectoryWatcher struct {
	directories []string
	checkEvery time.Duration
	newFilePaths chan string
	done chan interface{}
}

func CreateAndStartDirectoryWatcher(
	directories []string,
	checkEvery time.Duration) *DirectoryWatcher {

	newFilePaths := make(chan string, *watcherQueueSize)
	done := make(chan interface{})

	watcher := &DirectoryWatcher{
		directories, checkEvery, newFilePaths, done}
	go watcher.start()

	return watcher
}

func (watcher *DirectoryWatcher) Stop() {
	watcher.done <- nil
}

func (watcher *DirectoryWatcher) GetNewFilePaths() chan string {
	return watcher.newFilePaths
}

func (watcher *DirectoryWatcher) start() {
	ticker := time.NewTicker(watcher.checkEvery)
	defer ticker.Stop()

	seenFilePaths := make(map[string]bool)

	for {
		select {
		case <-ticker.C:
			for _, directory := range watcher.directories {
				files, err := ioutil.ReadDir(directory)
				if err != nil {
					// TODO: proper logging, handling
					println(err.Error())
				} else {
					for _, file := range files {
						path := path.Join(directory, file.Name())
						_, seenThisFile := seenFilePaths[path]
						if !seenThisFile {
							watcher.newFilePaths <- path
							seenFilePaths[path] = true
						}
					}
				}
			}

		case <-watcher.done:
			close(watcher.newFilePaths)
			return
		}
	}
}
