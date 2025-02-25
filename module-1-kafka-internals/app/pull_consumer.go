package main

import (
	"context"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func RunPullConsumer(ctx context.Context) error {
	return Consumer{
		Name: "Pull Consumer",
		Config: &kafka.ConfigMap{
			"bootstrap.servers":  *flagBootstrapServer,
			"group.id":           "pull_consumer",
			"fetch.wait.max.ms":  8192,
			"fetch.min.bytes":    8192,
			"enable.auto.commit": false,
		},
		PollTimeout:         int((8 * time.Second).Milliseconds()),
		EnabledManualCommit: true,
	}.Run(ctx)
}
