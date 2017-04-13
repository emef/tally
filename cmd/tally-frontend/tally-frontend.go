package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/emef/tally/frontend"
)

func main() {
	defaultStaticDir := path.Join(
		os.Getenv("GOPATH"), "src/github.com/emef/tally/frontend/static")
	staticDir := flag.String(
		"static_dir", defaultStaticDir, "Directory to static assets")
	port := flag.String("port", ":8000", "Port to run http frontend")
	endpoint := flag.String("endpoint", ":5020", "Endpoint to tally backend")
	flag.Parse()

	config := &frontend.FrontendConfig{
		StaticDir: *staticDir, BackendEndpoint: *endpoint}
	frontendServer, _ := frontend.NewFrontendServer(config)

	println("using static dir", *staticDir)

	srv := &http.Server{
		Handler:      frontendServer.Router,
		Addr:         *port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second}

	log.Fatal(srv.ListenAndServe())
}
