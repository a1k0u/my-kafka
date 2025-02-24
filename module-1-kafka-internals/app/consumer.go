package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type ConsumerType = string

const (
	PullConsumer ConsumerType = "Pull Consumer"
	PushConsumer              = "Push Consumer"
)

func RunConsumer(ctx context.Context, consumerType ConsumerType, fetchMinBytes, timeout int) func() error {
	return func() error {
		consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
			"bootstrap.servers":  *flagBootstrapServer,
			"group.id":           strings.Replace(strings.ToLower(consumerType), " ", "_", -1),
			"session.timeout.ms": 45000,
			"fetch.min.bytes":    fetchMinBytes,
			"enable.auto.commit": consumerType == PushConsumer,
		})
		if err != nil {
			return fmt.Errorf("%s: fail create consumer: %w", consumerType, err)
		}
		defer consumer.Close()

		if err := consumer.Subscribe(*flagTopic, nil); err != nil {
			return fmt.Errorf("%s: fail subscribe topic: %w", consumerType, err)
		}

		for {
			select {
			case <-ctx.Done():
				return nil
			default:
			}

			e := consumer.Poll(timeout)
			if e == nil {
				continue
			}

			switch ev := e.(type) {
			case *kafka.Message:
				var order Order

				if err := json.Unmarshal(ev.Value, &order); err != nil {
					return fmt.Errorf("%s: fail unmarshal: %w", consumerType, err)
				}

				fmt.Printf("%s: %v\n", consumerType, order)

				if consumerType == PullConsumer {
					if _, err := consumer.CommitMessage(ev); err != nil {
						return fmt.Errorf("%s: fail commit: %w", consumerType, err)
					}
				}
			case kafka.Error:
			default:
			}
		}
	}
}
