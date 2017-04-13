package deepend

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path"
	"sort"
	"time"

	"github.com/emef/tally/lib"
	"github.com/golang/protobuf/proto"
)

var (
	targetSizeMB = flag.Int(
		"compactor_target_size_mb", 64, "Target compacted file size")
	compactFactor = flag.Float64(
		"compactor_compact_factor", 0.5,
		"Assumed reduction in file size due to compaction")
	compactTolerance = flag.Float64(
		"compactor_tolerance", 0.9,
		"Don't compact files within this percent of target")
)

type Compactor struct {
	done   chan interface{}
	config *CompactorConfig
}

type CompactorConfig struct {
	RunEvery      time.Duration
	BaseDirectory string
}

func CreateAndStartCompactor(config *CompactorConfig) *Compactor {
	done := make(chan interface{})

	compactor := &Compactor{done, config}
	go compactor.start()

	return compactor
}

func (compactor *Compactor) Stop() {
	compactor.done <- nil
}

func (compactor *Compactor) start() {
	ticker := time.NewTicker(compactor.config.RunEvery)
	defer ticker.Stop()
	defer close(compactor.done)

	for {
		select {
		case <-ticker.C:
			compactor.compact()

		case <-compactor.done:
			return
		}
	}
}

type byNameDesc []os.FileInfo

func (f byNameDesc) Len() int           { return len(f) }
func (f byNameDesc) Less(i, j int) bool { return f[i].Name() > f[i].Name() }
func (f byNameDesc) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }

func (comparator *Compactor) compact() {
	// target compacted file size
	targetSizeBytes := int64(*targetSizeMB * 1e9)

	// threshold file size to trigger a compaction
	thresholdBytesToCompact := int64(float64(targetSizeBytes) / *compactFactor)

	// compaction proposal, will compact if sum of file size > threshold
	var proposedFiles []string = nil
	var proposedSize int64 = 0

	directories, err := ioutil.ReadDir(comparator.config.BaseDirectory)
	if err != nil {
		// TODO: proper logging, handling
		println(err.Error())
		return
	}

	sort.Sort(byNameDesc(directories))

	for _, directory := range directories {
		if !directory.IsDir() {
			continue
		}

		directoryPath := path.Join(
			comparator.config.BaseDirectory, directory.Name())

		files, err := ioutil.ReadDir(directoryPath)
		if err != nil {
			// TODO: proper logging, handling
			println(err.Error())
			continue
		}

		sort.Sort(byNameDesc(files))

		for _, file := range files {
			// skip directories
			if file.IsDir() {
				continue
			}

			// skip files within tolerance
			if float64(file.Size())/float64(targetSizeBytes) >= *compactTolerance {
				continue
			}

			filePath := path.Join(directoryPath, file.Name())
			proposedFiles = append(proposedFiles, filePath)
			proposedSize += file.Size()

			if proposedSize >= thresholdBytesToCompact {
				compactFiles(proposedFiles)
				proposedFiles = nil
				proposedSize = 0
			}
		}
	}
}

func compactFiles(pathsToCompact []string) {
	numPaths := len(pathsToCompact)
	paths := make(chan string, numPaths)
	defer close(paths)
	for _, path := range pathsToCompact {
		paths <- path
	}

	aggregator := NewCounterAggregator()
	reader := lib.CreateAndStartBlockReader(paths, 10)
	defer reader.Stop()

	blocks := reader.GetBlocks()
	for i := 0; i < numPaths; i++ {
		block := <-blocks
		aggregator.AddBlockInPlace(block)
	}

	compactedBlock := aggregator.AsBlock()
	data, err := proto.Marshal(compactedBlock)
	if err != nil {
		// TODO proper logging; don't log if directory exists
		log.Fatal("marshaling error: ", err)
		return
	}

	for _, path := range pathsToCompact {
		if err := os.Remove(path); err != nil {
			log.Fatal("Could not remove file: ", path)
		}
	}

	if err := ioutil.WriteFile(pathsToCompact[0], data, 0766); err != nil {
		log.Fatal("Could not write compacted file.. DATA LOSS")
	}
}
