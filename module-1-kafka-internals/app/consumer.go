package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Consumer struct {
	Name                string
	Config              *kafka.ConfigMap
	PollTimeout         int
	EnabledManualCommit bool
}

func (c Consumer) Run(ctx context.Context) error {
	consumer, err := kafka.NewConsumer(c.Config)
	if err != nil {
		return fmt.Errorf("%s: fail create consumer: %w", c.Name, err)
	}
	defer consumer.Close()

	if err := consumer.Subscribe(*flagTopic, nil); err != nil {
		return fmt.Errorf("%s: fail subscribe topic: %w", c.Name, err)
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
		}

		e := consumer.Poll(c.PollTimeout)
		if e == nil {
			continue
		}

		switch ev := e.(type) {
		case *kafka.Message:
			var order Order

			if err := json.Unmarshal(ev.Value, &order); err != nil {
				return fmt.Errorf("%s: fail unmarshal: %w", c.Name, err)
			}

			fmt.Printf("%s: get from %s [%d] %v with %v offset\n",
				c.Name, *ev.TopicPartition.Topic, ev.TopicPartition.Partition, order, ev.TopicPartition.Offset)

			if c.EnabledManualCommit {
				if _, err := consumer.CommitMessage(ev); err != nil {
					return fmt.Errorf("%s: fail commit: %w", c.Name, err)
				}
			}
		case kafka.Error:
			return fmt.Errorf("%s: fail poll: %v", c.Name, ev)
		default:
		}
	}
}
