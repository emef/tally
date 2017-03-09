package main

import (
	"flag"
	"log"
	"net"
	"time"

	"github.com/emef/tally/deepend"
	"github.com/emef/tally/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type RecordCounterService struct {
	shard *deepend.CounterShard
}

var ok = &pb.RecordCounterResponse{Ok: true}
var bulkOk = &pb.BulkRecordCounterResponse{Ok: true}

func (svc *RecordCounterService) RecordCounter(
	ctx context.Context,
	request *pb.RecordCounterRequest) (*pb.RecordCounterResponse, error) {

	svc.shard.RecordCounter(request)

	return ok, nil
}

func (svc *RecordCounterService) BulkRecordCounter(
	ctx context.Context,
	request *pb.BulkRecordCounterRequest) (*pb.BulkRecordCounterResponse, error) {

	for _, request := range request.Requests {
		svc.shard.RecordCounter(request)
	}

	return bulkOk, nil
}

func main() {
	port := flag.String("port", ":5019", "Port that service will run on")
	numWorkers := flag.Int("workers", 1, "Number of worker threads")
	workerFlushEvery := flag.Int(
		"worker_flush_every", 60, "Seconds before worker thread flushes")
	writerFlushEvery := flag.Int(
		"writer_flush_every", 300, "Seconds before writer thread flushes")
	writeDirectory := flag.String(
		"write_directory", "", "Directory to write flushed data")

	flag.Parse()

	shard := deepend.NewCounterShard(&deepend.ShardConfig{
		Workers: *numWorkers,
		AggregatorConfig: &deepend.AggregatorConfig{
			FlushEvery: time.Second * time.Duration(*workerFlushEvery)},
		WriterConfig: &deepend.WriterConfig{
			FlushEvery:    time.Second * time.Duration(*writerFlushEvery),
			BaseDirectory: *writeDirectory}})

	service := &RecordCounterService{shard}

	lis, err := net.Listen("tcp", string(*port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterRecordCounterServiceServer(s, service)

	// Register reflection service on gRPC server.
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
