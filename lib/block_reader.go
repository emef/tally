package lib

import (
	"io/ioutil"

	"github.com/emef/tally/pb"
	"github.com/golang/protobuf/proto"
)

type BlockReader struct {
	in   chan string
	out  chan *pb.RecordBlock
	done chan interface{}
}

func CreateAndStartBlockReader(paths chan string, queueSize int) *BlockReader {
	out := make(chan *pb.RecordBlock, queueSize)
	done := make(chan interface{})

	blockReader := &BlockReader{paths, out, done}
	go blockReader.start()

	return blockReader
}

func (reader *BlockReader) GetBlocks() chan *pb.RecordBlock {
	return reader.out
}

func (reader *BlockReader) Stop() {
	reader.done <- nil
}

func (reader *BlockReader) start() {
	defer close(reader.done)
	defer close(reader.out)

	for {
		select {
		case path := <-reader.in:
			data, err := ioutil.ReadFile(path)
			// TODO: error handling
			if err != nil {
				println("error reading file")
				continue
			}

			block := &pb.RecordBlock{}
			err = proto.Unmarshal(data, block)
			if err != nil {
				println("error unmarshalling")
			}

			reader.out <- block

		case <-reader.done:
			return
		}
	}
}
