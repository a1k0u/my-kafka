package main

import (
	"context"
	"time"
)

const PushConsumerPollTimeout = 2 * time.Second
const PushConsumerFetchMinBytes = 16

func RunPushConsumer(ctx context.Context) func() error {
	return RunConsumer(ctx, PushConsumer, PushConsumerFetchMinBytes, int(PushConsumerPollTimeout.Milliseconds()))
}
