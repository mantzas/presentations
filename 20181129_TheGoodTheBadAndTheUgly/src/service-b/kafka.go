package main

import (
	"time"

	"github.com/mantzas/patron"
	"github.com/mantzas/patron/async"
	"github.com/mantzas/patron/async/kafka"
	"github.com/mantzas/patron/encoding/json"
	"github.com/mantzas/patron/log"
	"github.com/mantzas/patron/trace/amqp"
)

type kafkaComponent struct {
	cmp patron.Component
	pub amqp.Publisher
}

func newKafkaComponent() (*kafkaComponent, error) {
	cf, err := kafka.New("test", json.Type, "test-topic", []string{"localhost:9092"})
	if err != nil {
		return nil, err
	}

	kafkaCmp := kafkaComponent{}

	cmp, err := async.New("kfk", kafkaCmp.process, cf, async.ConsumerRetry(10, 5*time.Second))
	if err != nil {
		return nil, err
	}

	kafkaCmp.cmp = cmp

	pub, err := amqp.NewPublisher("amqp://guest:guest@localhost:5672/", "test-exc")
	if err != nil {
		return nil, err
	}
	kafkaCmp.pub = pub

	return &kafkaCmp, nil
}

func (kc *kafkaComponent) process(msg async.Message) error {
	var data map[string]string

	err := msg.Decode(&data)
	if err != nil {
		return err
	}

	data["kafka"] = time.Now().String()

	amqpMsg, err := amqp.NewJSONMessage(data)
	if err != nil {
		return err
	}

	err = kc.pub.Publish(msg.Context(), amqpMsg)
	if err != nil {
		return err
	}

	log.Info("kafka message processed")
	return nil
}
