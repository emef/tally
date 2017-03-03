package deepend

import (
	"github.com/emef/tally/pb"
	"testing"
)

func TestRequestDispatcher(t *testing.T) {
	n := 1000
	chan1 := make(chan *pb.RecordCounterRequest, n)
	chan2 := make(chan *pb.RecordCounterRequest, n)
	chan3 := make(chan *pb.RecordCounterRequest, n)

	dispatcher := NewRequestDispatcher([]chan<- *pb.RecordCounterRequest{
		chan1, chan2, chan3})

	for i := 0; i < n; i++ {
		request := &pb.RecordCounterRequest{Name: string(i)}
		dispatcher.Dispatch(request)
	}

	lenChan1 := len(chan1)
	lenChan2 := len(chan2)
	lenChan3 := len(chan3)

	if lenChan1+lenChan2+lenChan3 != n {
		t.Errorf("All requests were not dispatched")
	}

	// we expect about 333 requests per channel, but give a 10% margin of error
	assertGreater(t, lenChan1, 300)
	assertGreater(t, lenChan2, 300)
	assertGreater(t, lenChan3, 300)
}

func assertGreater(t *testing.T, v int, minValue int) {
	if v < minValue {
		t.Errorf("%v must be greater than %v", v, minValue)
	}
}
