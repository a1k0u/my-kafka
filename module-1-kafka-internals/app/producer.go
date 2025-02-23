package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func RunProducer(ctx context.Context) error {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": *flagBootstrapServer,
	})
	if err != nil {
		return fmt.Errorf("fail create producer: %w", err)
	}
	defer producer.Close()

	events := make(chan kafka.Event, 32)
	defer close(events)

	var wg sync.WaitGroup
	defer wg.Wait()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case e, ok := <-events:
				if !ok {
					return
				}

				switch ev := e.(type) {
				case *kafka.Message:
					if ev.TopicPartition.Error != nil {
						fmt.Printf("Delivery failed: %v\n", ev.TopicPartition.Error)
					} else {
						fmt.Printf("Delivered message to %s [%d] with %v offset\n",
							*ev.TopicPartition.Topic, ev.TopicPartition.Partition, ev.TopicPartition.Offset)
					}
				}
			}
		}
	}()

producing:
	for {
		order := Order{Price: rand.Uint64()}
		orderRaw, err := json.Marshal(order)
		if err != nil {
			return fmt.Errorf("fail marhal order: %w", err)
		}

		err = producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     flagTopic,
				Partition: kafka.PartitionAny,
			},
			Value: orderRaw,
		}, events)
		if err != nil {
			return fmt.Errorf("fail produce message")
		}

		select {
		case <-ctx.Done():
			break producing
		case <-time.After(*flagProduceInterval):
		}
	}

	return nil
}
