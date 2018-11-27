package main

import (
	"fmt"
	"os"

	"github.com/mantzas/patron"
	"github.com/mantzas/patron/log"
	"github.com/mantzas/patron/sync/http"
)

var (
	version = "dev"
)

func main() {
	name := "service_a"

	err := patron.SetupLogging(name, version)
	if err != nil {
		fmt.Printf("failed to set up logging: %v", err)
		os.Exit(1)
	}

	httpCmp, err := newHTTPComponent()
	if err != nil {
		log.Fatalf("failed to create http component: %v", err)
	}

	routes := []http.Route{
		http.NewPostRoute("/", httpCmp.post, true),
	}

	srv, err := patron.New(name, version, patron.Routes(routes))
	if err != nil {
		log.Fatalf("failed to create service: %v", err)
	}

	err = srv.Run()
	if err != nil {
		log.Fatalf("failed to run service %v", err)
	}
}
