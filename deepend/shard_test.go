package deepend

import (
	"github.com/emef/tally/pb"
	"testing"
	"time"
)

func TestShard(t *testing.T) {
	config := &ShardConfig{
		Workers:         1,
		AggregatorConfig: &AggregatorConfig{FlushEvery: time.Second},
		WriterConfig: &WriterConfig{
			FlushEvery:   time.Second * 5,
			BaseDirectory: "/tmp/shard"}}

	shard := NewCounterShard(config)

	shard.RecordCounter(&pb.RecordCounterRequest{
		Name:        "counter",
		Source:      "source",
		EpochMinute: 0,
		Values: &pb.CounterValues{
			Count: 1,
			Sum:   10.0,
			Min:   10.0,
			Max:   10.0}})

	shard.Stop()
}
