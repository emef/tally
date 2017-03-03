package deepend

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"

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

	err := os.MkdirAll(writer.config.BaseDirectory, 0766)
	if err != nil {
		println("Couldn't mkdir -p")
	}

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
				close(writer.aggregators)
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
		log.Fatal("marshaling error: ", err)
		return
	}

	filepath := path.Join(config.BaseDirectory, uuid.NewV4().String())
	err = ioutil.WriteFile(filepath, data, 0766)
	if err != nil {
		println("could not flush to file: ", err.Error())
	}
}
