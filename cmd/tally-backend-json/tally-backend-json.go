package main

import (
	"flag"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	gw "github.com/emef/tally/pb"
)

var (
	port     = flag.String("port", ":8081", "port to run this proxy")
	endpoint = flag.String("endpoint", ":5020", "endpoint of tbackend")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterQueryCounterServiceHandlerFromEndpoint(
		ctx, mux, *endpoint, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(*port, mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
