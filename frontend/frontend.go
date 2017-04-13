package frontend

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"

	"github.com/emef/tally/pb"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type FrontendConfig struct {
	StaticDir       string
	BackendEndpoint string
}

type FrontendServer struct {
	Config *FrontendConfig
	Router *mux.Router

	client pb.QueryCounterServiceClient
}

func (srv *FrontendServer) handleGetCsv(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	source := vars["source"]
	startEpochMinute, _ := strconv.Atoi(r.URL.Query().Get("startEpochMinute"))
	endEpochMinute, _ := strconv.Atoi(r.URL.Query().Get("endEpochMinute"))

	request := &pb.GetCounterRequest{
		Name:             name,
		Source:           source,
		StartEpochMinute: int32(startEpochMinute),
		EndEpochMinute:   int32(endEpochMinute)}
	resp, err := srv.client.GetCounter(context.Background(), request)

	if err != nil {
		println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else if !resp.Ok {
		println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(resp.Error))
	} else {
		timestamps := make([]int, len(resp.Values))
		i := 0
		for timestamp := range resp.Values {
			timestamps[i] = int(timestamp)
			i++
		}
		sort.Ints(timestamps)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("epochMinute,sum\n"))
		for _, timestamp := range timestamps {
			values := resp.Values[int32(timestamp)]
			w.Write([]byte(fmt.Sprint(timestamp)))
			w.Write([]byte(","))
			w.Write([]byte(fmt.Sprint(values.Sum)))
			w.Write([]byte("\n"))
		}
	}
}

func NewFrontendServer(config *FrontendConfig) (*FrontendServer, error) {
	conn, err := grpc.Dial(config.BackendEndpoint, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := pb.NewQueryCounterServiceClient(conn)

	router := mux.NewRouter()
	srv := &FrontendServer{
		Config: config, Router: router, client: client}

	router.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/", http.FileServer(http.Dir(config.StaticDir))))
	router.HandleFunc("/csv/{name}", srv.handleGetCsv)
	router.HandleFunc("/csv/{name}/", srv.handleGetCsv)
	router.HandleFunc("/csv/{name}/{source}", srv.handleGetCsv)

	return srv, nil
}
