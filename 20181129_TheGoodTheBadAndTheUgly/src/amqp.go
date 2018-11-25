package main

import (
	"time"

	"github.com/mantzas/patron"
	"github.com/mantzas/patron/async"
	"github.com/mantzas/patron/async/amqp"
	"github.com/mantzas/patron/log"
)

type amqpComponent struct {
	cmp patron.Component
}

func newAmqpComponent() (*amqpComponent, error) {
	cf, err := amqp.New("amqp://guest:guest@localhost:5672/", "test-queue", "test-exc")
	if err != nil {
		return nil, err
	}
	amqpCmp := amqpComponent{}

	cmp, err := async.New("amqp", amqpCmp.process, cf, async.ConsumerRetry(10, 10*time.Second))
	if err != nil {
		return nil, err
	}
	amqpCmp.cmp = cmp

	return &amqpCmp, nil
}

func (ac *amqpComponent) process(msg async.Message) error {
	var data map[string]string

	err := msg.Decode(&data)
	if err != nil {
		return err
	}

	data["amqp"] = time.Now().String()

	for k, v := range data {
		log.Infof("component: %s time: %s", k, v)
	}

	return nil
}
