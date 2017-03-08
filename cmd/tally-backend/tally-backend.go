package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net"
	"time"

	"github.com/emef/tally/backend"
	"github.com/emef/tally/pb"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/reflection"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type QueryCounterService struct {
	indexer *backend.Indexer
}

func (svc *QueryCounterService) GetCounter(
	ctx context.Context,
	request *pb.GetCounterRequest) (*pb.GetCounterResponse, error) {

	println((*request).String())
	values := svc.indexer.Get(
		request.Name,
		request.Source,
		request.StartEpochMinute,
		request.EndEpochMinute)

	resp := &pb.GetCounterResponse{
		Ok: true,
		Values: values}

	return resp, nil
}

func main() {
	port := flag.String("port", ":5020", "Port that service will run on")
	indexDirectory := flag.String(
		"index_directory", "", "Directory to index data")

	flag.Parse()

	watcher := backend.CreateAndStartDirectoryWatcher(
		[]string{*indexDirectory}, time.Second)

	println("starting...")

	blocks := make(chan *pb.RecordBlock)
	go func() {
		for path := range watcher.GetNewFilePaths() {
			println(path)
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

			println("ingested block")
			blocks <- block
		}
	}()

	indexer := backend.CreateAndStartIndexer(blocks)

	service := &QueryCounterService{indexer}

	lis, err := net.Listen("tcp", string(*port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterQueryCounterServiceServer(s, service)

	// Register reflection service on gRPC server.
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
