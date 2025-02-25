package main

import (
	"context"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func RunPushConsumer(ctx context.Context) error {
	return Consumer{
		Name: "Push Consumer",
		Config: &kafka.ConfigMap{
			"bootstrap.servers":  *flagBootstrapServer,
			"group.id":           "push_consumer",
			"fetch.wait.max.ms":  100,
			"fetch.min.bytes":    32,
			"enable.auto.commit": true,
		},
		PollTimeout:         int((100 * time.Millisecond).Milliseconds()),
		EnabledManualCommit: false,
	}.Run(ctx)
}
