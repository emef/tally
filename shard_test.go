package tally

import (
	"testing"
	"time"
	"github.com/emef/tally/pb"
)

func TestShard(t *testing.T) {
	config := &ShardConfig{
		numWorkers: 1,
		workerFlushEvery: time.Second,
		writerFlushEvery: time.Second * 5,
		flushBaseDirectory: "/tmp/shard"}

	shard := NewCounterShard(config)

	shard.RecordCounter(&pb.RecordCounterRequest{
		Name: "counter",
		Source: "source",
		EpochMinute: 0,
		Values: &pb.CounterValues{
			Count: 1,
			Sum: 10.0,
			Min: 10.0,
			Max: 10.0}})

	time.Sleep(time.Second * 20)
}
