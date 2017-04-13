package deepend

import (
	"github.com/emef/tally/pb"
)

type CounterShard struct {
	dispatcher *RequestDispatcher
	workers    []*AggregatorWorker
	writer     *FlushWriter
	compactor  *Compactor
}

type ShardConfig struct {
	Workers          int
	AggregatorConfig *AggregatorConfig
	WriterConfig     *WriterConfig
	CompactorConfig  *CompactorConfig
}

func NewCounterShard(config *ShardConfig) *CounterShard {
	writer := CreateAndStartFlushWriter(config.WriterConfig)
	compactor := CreateAndStartCompactor(config.CompactorConfig)

	workers := make([]*AggregatorWorker, config.Workers)
	for i := range workers {
		workers[i] = CreateAndStartAggregatorWorker(
			writer.GetAggregatorChannel(), config.AggregatorConfig)
	}

	requestChannels := make([]chan<- *pb.RecordCounterRequest, 0)
	for _, worker := range workers {
		requestChannels = append(requestChannels, worker.GetRequestChannel())
	}

	dispatcher := NewRequestDispatcher(requestChannels)

	return &CounterShard{dispatcher, workers, writer, compactor}
}

func (shard *CounterShard) Stop() {
	// TODO: more graceful shutdown? potential data loss here
	shard.writer.Stop()
	shard.compactor.Stop()
	for _, worker := range shard.workers {
		worker.Stop()
	}
}

func (shard *CounterShard) RecordCounter(request *pb.RecordCounterRequest) {
	shard.dispatcher.Dispatch(request)
}
