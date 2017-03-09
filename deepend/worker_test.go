package deepend

import (
	"github.com/emef/tally/pb"
	"testing"
	"time"
)

func TestAggregator(t *testing.T) {
	flushChannel := make(chan *CounterAggregator)
	config := &AggregatorConfig{250 * time.Millisecond}

	worker := CreateAndStartAggregatorWorker(flushChannel, config)
	requests := worker.GetRequestChannel()

	requests <- makeRequest("c1", "s1", 1)
	requests <- makeRequest("c1", "s1", 3)
	requests <- makeRequest("c1", "s2", 10)
	requests <- makeRequest("c2", "s1", 100)
	requests <- makeRequest("c2", "s3", 1000)

	aggregator := <-flushChannel
	worker.Stop()

	block := aggregator.AsBlock()

	if len(block.Entries) != 4 {
		t.Error("Block should contain 4 entries")
	}

	c1 := getAndAssertPresent(t, block.NameCodeMapping, "c1")
	c2 := getAndAssertPresent(t, block.NameCodeMapping, "c2")
	s1 := getAndAssertPresent(t, block.SourceCodeMapping, "s1")
	s2 := getAndAssertPresent(t, block.SourceCodeMapping, "s2")
	s3 := getAndAssertPresent(t, block.SourceCodeMapping, "s3")

	assertCounterEquals(t, block.Entries, c1, s1, 2, 4)
	assertCounterEquals(t, block.Entries, c1, s2, 1, 10)
	assertCounterEquals(t, block.Entries, c2, s1, 1, 100)
	assertCounterEquals(t, block.Entries, c2, s3, 1, 1000)
}

func assertCounterEquals(
	t *testing.T,
	entries []*pb.RecordEntry,
	nameCode int32,
	sourceCode int32,
	count int32,
	sum float32) {

	for _, entry := range entries {
		if entry.Key.NameCode == nameCode && entry.Key.SourceCode == sourceCode {
			if entry.Values.Count != count {
				t.Errorf("Mismatched count: %v != %v", entry.Values.Count, count)
			}

			if entry.Values.Sum != sum {
				t.Errorf("Mismatched sum: %v != %v", entry.Values.Sum, sum)
			}

			return
		}
	}

	t.Errorf("Key not found name = %v source = %v", nameCode, sourceCode)
}

func getAndAssertPresent(
	t *testing.T,
	codeMap map[int32]string,
	desiredValue string) int32 {

	for key, value := range codeMap {
		if value == desiredValue {
			return key
		}
	}

	t.Errorf("Value %v missing", desiredValue)
	return 0
}

func makeRequest(
	name string,
	source string,
	value float32) *pb.RecordCounterRequest {

	return &pb.RecordCounterRequest{
		Name:        name,
		Source:      source,
		EpochMinute: 0,
		Values: &pb.CounterValues{
			Count: 1,
			Sum:   value,
			Min:   value,
			Max:   value}}
}
