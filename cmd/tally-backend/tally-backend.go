package main

import (
	"flag"
	"log"
	"net"
	"time"

	"github.com/emef/tally/backend"
	"github.com/emef/tally/lib"
	"github.com/emef/tally/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
		Ok:     true,
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
	reader := lib.CreateAndStartBlockReader(watcher.GetNewFilePaths(), 10)
	indexer := backend.CreateAndStartIndexer(reader.GetBlocks())
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
