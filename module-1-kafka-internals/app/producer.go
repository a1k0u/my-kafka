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
		"acks":              1,
		"retries":           3,
	})
	if err != nil {
		return fmt.Errorf("fail create producer: %w", err)
	}

	var wg sync.WaitGroup
	defer wg.Wait()

	events := make(chan kafka.Event, 32)
	defer close(events)

	defer producer.Close()

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
						fmt.Printf("Producer: delivery failed: %v\n", ev.TopicPartition.Error)
					} else {
						fmt.Printf("Producer: delivered message to %s [%d] %v with %v offset\n",
							*ev.TopicPartition.Topic, ev.TopicPartition.Partition, ev.Value, ev.TopicPartition.Offset)
					}
				}
			}
		}
	}()

	time.Sleep(4 * time.Second)

	for {
		order := Order{Price: rand.Uint64()}
		orderRaw, err := json.Marshal(order)
		if err != nil {
			return fmt.Errorf("Producer: fail marshal order: %w", err)
		}

		err = producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     flagTopic,
				Partition: kafka.PartitionAny,
			},
			Value: orderRaw,
		}, events)
		if err != nil {
			return fmt.Errorf("Producer: fail produce message")
		}

		select {
		case <-ctx.Done():
			return nil
		case <-time.After(*flagProduceInterval):
		}
	}
}
