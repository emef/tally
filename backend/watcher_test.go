package backend

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestDirectoryWatcher(t *testing.T) {
	directory, err := ioutil.TempDir("", "TestDirectoryWatcher")

	if err != nil {
		t.Fatalf("Could not create temporary directory")
	}

	defer os.RemoveAll(directory)
	watcher := CreateAndStartDirectoryWatcher(
		[]string{directory}, time.Millisecond)

	filenames := make([]string, 10)
	for i := range filenames {
		filenames[i] = path.Join(directory, fmt.Sprint(i))
	}

	// Test with subdirectory as well
	subdir := path.Join(directory, "subdirectory")
	subdirPath := path.Join(subdir, "subfile")
	os.Mkdir(subdir, 0777)
	filenames = append(filenames, subdirPath)

	done := make(chan interface{})
	defer close(done)

	go func() {
		for _, filename := range filenames {
			time.Sleep(time.Millisecond)
			ioutil.WriteFile(filename, []byte(filename), 0666)
		}
		done <- nil
	}()

	// Give watcher enough time to check after all files written
	<-done
	time.Sleep(1 * time.Millisecond)
	watcher.Stop()

	actualFilenames := make([]string, 0)
	for actual := range watcher.GetNewFilePaths() {
		actualFilenames = append(actualFilenames, actual)
	}

	sort.Strings(filenames)
	sort.Strings(actualFilenames)

	if !reflect.DeepEqual(filenames, actualFilenames) {
		t.Errorf("%v != %v", filenames, actualFilenames)
	}
}
