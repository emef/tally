package deepend

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"

	"github.com/emef/tally/lib"
	"github.com/golang/protobuf/proto"
	"github.com/satori/go.uuid"
)

var (
	writerQueueSize = flag.Int(
		"writer_queue_size", 10, "Writer max queue size")
)

type FlushWriter struct {
	aggregators chan *CounterAggregator
	done        chan interface{}
	config      *WriterConfig
}

type WriterConfig struct {
	FlushEvery    time.Duration
	BaseDirectory string
}

func CreateAndStartFlushWriter(config *WriterConfig) *FlushWriter {
	aggregators := make(chan *CounterAggregator, *writerQueueSize)
	done := make(chan interface{})

	writer := &FlushWriter{aggregators, done, config}
	go writer.start()

	return writer
}

func (writer *FlushWriter) Stop() {
	writer.done <- nil
}

func (writer *FlushWriter) GetAggregatorChannel() chan *CounterAggregator {
	return writer.aggregators
}

func (writer *FlushWriter) start() {
	ticker := time.NewTicker(writer.config.FlushEvery)
	defer ticker.Stop()
	defer close(writer.done)
	defer close(writer.aggregators)

	combinedAggregator := NewCounterAggregator()

	for {
		select {
		case aggregator := <-writer.aggregators:
			combinedAggregator.CombineInPlace(aggregator)

		case <-ticker.C:
			aggregatorToFlush := combinedAggregator
			go flushAggregator(aggregatorToFlush, writer.config)
			combinedAggregator = NewCounterAggregator()

		case <-writer.done:
			if len(writer.aggregators) == 0 {
				flushAggregator(combinedAggregator, writer.config)
				return
			}
		}
	}
}

func flushAggregator(aggregator *CounterAggregator, config *WriterConfig) {
	if aggregator.IsEmpty() {
		return
	}

	block := aggregator.AsBlock()

	data, err := proto.Marshal(block)
	if err != nil {
		// TODO proper logging; don't log if directory exists
		log.Fatal("marshaling error: ", err)
		return
	}

	directory := path.Join(
		config.BaseDirectory, fmt.Sprint(lib.NowEpochMinute()))
	os.MkdirAll(directory, 0766)

	filepath := path.Join(directory, uuid.NewV4().String())
	err = ioutil.WriteFile(filepath, data, 0766)
	if err != nil {
		println("could not flush to file: ", err.Error())
	}
}
