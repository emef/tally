package backend

import (
	"reflect"
	"testing"
	"time"

	"github.com/emef/tally/pb"
)

func TestIndexer(t *testing.T) {
	blocks := make(chan *pb.RecordBlock, 2)

	blocks <- &pb.RecordBlock{
		NameCodeMapping: map[int32]string{
			0: "c1",
			1: "c2"},
		SourceCodeMapping: map[int32]string{0: "s"},
		Entries: []*pb.RecordEntry{
			entry(0, 0, 0, values(1, 1, 1, 1)),
			entry(0, 0, 1, values(2, 1, 2, 2)),
			entry(1, 0, 0, values(3, 2, 1, 2)),
			entry(1, 1, 3, values(4, 4, 1, 1))}}

	blocks <- &pb.RecordBlock{
		NameCodeMapping: map[int32]string{
			1: "c1",
			0: "c2",
			2: "c3"},
		SourceCodeMapping: map[int32]string{0: "s"},
		Entries: []*pb.RecordEntry{
			entry(0, 0, 0, values(1, 1, 1, 1)),
			entry(1, 0, 0, values(1, 1, 1, 1)),
			entry(2, 0, 0, values(8, 2, 1, 4))}}

	indexer := CreateAndStartIndexer(blocks)
	time.Sleep(time.Millisecond)

	empty := map[int32]*pb.CounterValues{}

	// out of range, expect empty
	actual := indexer.Get("c1", "s", 100, 200)
	expected := empty
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Query failed %v != %v", actual, expected)
	}

	// only includes first timestamp
	actual = indexer.Get("c1", "s", 0, 0)
	expected = map[int32]*pb.CounterValues{0: values(2, 2, 1, 1)}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Query failed %v != %v", actual, expected)
	}

	// include both timestamps for c1
	actual = indexer.Get("c1", "s", 0, 100)
	expected = map[int32]*pb.CounterValues{
		0: values(2, 2, 1, 1),
		1: values(2, 1, 2, 2)}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Query failed %v != %v", actual, expected)
	}
}

func TestTreeIndex(t *testing.T) {
	index := newTreeIndex()
	index.insert(0, values(1, 1, 1, 1))
	index.insert(2, values(2, 1, 2, 2))
	index.insert(1000, values(2, 2, 1, 1))
	index.insert(2, values(1, 1, 1, 1))
	index.insert(8, values(8, 1, 8, 8))
	index.insert(50, values(2, 2, 1, 1))
	index.insert(3, values(3, 2, 1, 2))
	index.insert(10, values(10, 5, 2, 1))

	actual := index.query(2, 49)

	expected := map[int32]*pb.CounterValues{
		2:  values(3, 2, 1, 2),
		3:  values(3, 2, 1, 2),
		8:  values(8, 1, 8, 8),
		10: values(10, 5, 2, 1)}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("%v != %v", actual, expected)
	}
}

func entry(
	nameCode, sourceCode, epochMinute int32,
	values *pb.CounterValues) *pb.RecordEntry {

	return &pb.RecordEntry{
		Key: &pb.RecordKey{
			NameCode:    nameCode,
			SourceCode:  sourceCode,
			EpochMinute: epochMinute},
		Values: values}
}

func values(sum float32, count int32, min, max float32) *pb.CounterValues {
	return &pb.CounterValues{
		Sum: sum, Count: count, Min: min, Max: max}
}
