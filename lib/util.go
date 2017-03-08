package lib

import (
	"github.com/emef/tally/pb"
)

func MakeReverseMap(codeMap map[string]int32) map[int32]string {
	reverseMap := make(map[int32]string, len(codeMap))
	for key, code := range codeMap {
		reverseMap[code] = key
	}
	return reverseMap
}

func AddCountersInPlace(left, right *pb.CounterValues) {
		left.Count += right.Count
		left.Sum += right.Sum
		left.Min = min(left.Min, right.Min)
		left.Max = max(left.Max, right.Max)
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
