package deepend

import (
	"github.com/emef/tally"
	"github.com/emef/tally/pb"
)

type RequestDispatcher struct {
	channels [](chan<- *pb.RecordCounterRequest)
}

func NewRequestDispatcher(
	channels [](chan<- *pb.RecordCounterRequest)) *RequestDispatcher {
	return &RequestDispatcher{channels}
}

func (dispatcher *RequestDispatcher) Dispatch(request *pb.RecordCounterRequest) {
	maxIndex := int64(len(dispatcher.channels) - 1)
	channelIndex := tally.HashToRange(request.Name, 0, maxIndex)
	dispatcher.channels[channelIndex] <- request
	println("dispatched request")
}
