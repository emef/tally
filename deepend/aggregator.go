package deepend

import (
	"github.com/emef/tally/pb"
)

type CounterAggregator struct {
	counters map[pb.RecordKey]*pb.CounterValues
	nameCodeMapping map[string]int32
	sourceCodeMapping map[string]int32
}

func NewCounterAggregator() *CounterAggregator {
	return &CounterAggregator{
		make(map[pb.RecordKey]*pb.CounterValues),
		make(map[string]int32),
		make(map[string]int32)}
}

func (aggregator *CounterAggregator) IsEmpty() bool {
	return len(aggregator.counters) == 0
}

func (aggregator *CounterAggregator) CombineInPlace(other *CounterAggregator) {
	reverseNameMap := makeReverseMap(other.nameCodeMapping)
	reverseSourceMap := makeReverseMap(other.sourceCodeMapping)

	for key, values := range other.counters {
		name := reverseNameMap[key.NameCode]
		source := reverseSourceMap[key.SourceCode]
		aggregator.AddInPlace(name, source, key.EpochMinute, values)
	}
}

func (aggregator *CounterAggregator) AddInPlace(
	name string,
	source string,
	epochMinute int32,
	values *pb.CounterValues) {

	key := pb.RecordKey{
		NameCode: getOrSetCode(name, aggregator.nameCodeMapping),
		SourceCode: getOrSetCode(source, aggregator.sourceCodeMapping),
		EpochMinute: epochMinute}

	existing, ok := aggregator.counters[key]
	if ok {
		existing.Count += values.Count
		existing.Sum += values.Sum
		existing.Min = min(existing.Min, values.Min)
		existing.Max = max(existing.Max, values.Max)
	} else {
		aggregator.counters[key] = values
	}
}

func (aggregator *CounterAggregator) AsBlock() *pb.RecordBlock {
	entries := make([]*pb.RecordEntry, 0, len(aggregator.counters))
	for key, counter := range aggregator.counters {
		// NOTE: the problem here is that we can't take &key directly
		// because it is the loop variable but we want a reference to the
		// actual key... there is probably a better way to do this other
		// than taking a copy.
		keyCopy := key
		entry := pb.RecordEntry{Key: &keyCopy, Values: counter}
		entries = append(entries, &entry)
	}

	return &pb.RecordBlock{
		NameCodeMapping: aggregator.nameCodeMapping,
		SourceCodeMapping: aggregator.sourceCodeMapping,
		Entries: entries}
}

func getOrSetCode(
	key string,
	codeMap map[string]int32) int32 {

	code, ok := codeMap[key]
	if !ok {
		code = int32(len(codeMap))
		codeMap[key] = code
	}

	return code
}

func makeReverseMap(codeMap map[string]int32) map[int32]string {
	reverseMap := make(map[int32]string, len(codeMap))
	for key, code := range codeMap {
		reverseMap[code] = key
	}
	return reverseMap
}

func min(x, y float32) float32 {
    if x < y {
        return x
    }
    return y
}

func max(x, y float32) float32 {
    if x > y {
        return x
    }
    return y
}
