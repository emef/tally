package deepend

import (
	"flag"
	"time"

	"github.com/emef/tally/pb"
)

var (
	workerQueueSize = flag.Int(
		"worker_queue_size", 10, "Worker max queue size")
)

type AggregatorWorker struct {
	requests     chan *pb.RecordCounterRequest
	flushChannel chan<- *CounterAggregator
	done         chan interface{}
	config       *AggregatorConfig
}

type AggregatorConfig struct {
	FlushEvery time.Duration
}

func CreateAndStartAggregatorWorker(
	flushChannel chan<- *CounterAggregator,
	config *AggregatorConfig) *AggregatorWorker {

	done := make(chan interface{})
	requests := make(chan *pb.RecordCounterRequest, *workerQueueSize)

	worker := &AggregatorWorker{requests, flushChannel, done, config}
	go worker.start()

	return worker
}

func (worker *AggregatorWorker) Stop() {
	worker.done <- nil
}

func (worker *AggregatorWorker) GetRequestChannel() chan *pb.RecordCounterRequest {
	return worker.requests
}

func (worker *AggregatorWorker) start() {
	ticker := time.NewTicker(worker.config.FlushEvery)
	defer ticker.Stop()

	aggregator := NewCounterAggregator()

	for {
		select {
		case request := <-worker.requests:
			aggregator.AddInPlace(
				request.Name, request.Source, request.EpochMinute, request.Values)
			println("aggregated request")

		case <-ticker.C:
			if !aggregator.IsEmpty() {
				worker.flushChannel <- aggregator
				aggregator = NewCounterAggregator()
				println("flushed aggregator")
			}

		case <-worker.done:
			close(worker.requests)
			if !aggregator.IsEmpty() {
				worker.flushChannel <- aggregator
			}
			return
		}
	}
}
