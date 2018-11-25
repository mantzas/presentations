package main

import (
	"context"
	"time"

	"github.com/mantzas/patron/log"
	"github.com/mantzas/patron/sync"
	"github.com/mantzas/patron/trace/kafka"
)

type httpComponent struct {
	prd kafka.Producer
}

func newHTTPComponent() (*httpComponent, error) {
	prd, err := kafka.NewAsyncProducer([]string{"localhost:9092"})
	if err != nil {
		return nil, err
	}
	return &httpComponent{prd: prd}, nil
}

func (hc *httpComponent) post(ctx context.Context, req *sync.Request) (*sync.Response, error) {
	data := map[string]string{"http": time.Now().String()}
	kafkaMsg, err := kafka.NewJSONMessage("test-topic", data)
	if err != nil {
		return nil, err
	}

	err = hc.prd.Send(ctx, kafkaMsg)
	if err != nil {
		return nil, err
	}
	log.Info("http request processed")
	return sync.NewResponse(data), nil
}
