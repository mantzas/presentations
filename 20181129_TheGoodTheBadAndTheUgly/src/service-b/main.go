package main

import (
	"fmt"
	"os"

	"github.com/mantzas/patron"
	"github.com/mantzas/patron/log"
)

var (
	version = "dev"
)

func main() {
	name := "service_b"

	err := patron.SetupLogging(name, version)
	if err != nil {
		fmt.Printf("failed to set up logging: %v", err)
		os.Exit(1)
	}

	kafkaCmp, err := newKafkaComponent()
	if err != nil {
		log.Fatalf("failed to create kafka component: %v", err)
	}

	srv, err := patron.New(name, version, patron.Components(kafkaCmp.cmp))
	if err != nil {
		log.Fatalf("failed to create service: %v", err)
	}

	err = srv.Run()
	if err != nil {
		log.Fatalf("failed to run service %v", err)
	}
}
