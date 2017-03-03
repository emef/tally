package tally

import (
	"time"

	"github.com/emef/tally/pb"
)

type CounterShard struct {
	dispatcher *RequestDispatcher
	workers []*AggregatorWorker
	writer *FlushWriter
}

type ShardConfig struct {
	NumWorkers int
	WorkerFlushEvery time.Duration
	WriterFlushEvery time.Duration
	FlushBaseDirectory string
}

func NewCounterShard(config *ShardConfig) *CounterShard {
	writer := CreateAndStartFlushWriter(&WriterConfig{
		config.WriterFlushEvery, config.FlushBaseDirectory})

	workerConfig := &AggregatorConfig{config.WorkerFlushEvery}
	workers := make([]*AggregatorWorker, config.NumWorkers)
	for i := range workers {
		workers[i] = CreateAndStartAggregatorWorker(
			writer.GetAggregatorChannel(), workerConfig)
	}

	requestChannels := make([]chan<- *pb.RecordCounterRequest, 0)
	for _, worker := range workers {
		requestChannels = append(requestChannels, worker.GetRequestChannel())
	}

	dispatcher := NewRequestDispatcher(requestChannels)

	return &CounterShard{dispatcher, workers, writer}
}

func (shard *CounterShard) RecordCounter(request *pb.RecordCounterRequest) {
	shard.dispatcher.Dispatch(request)
}
