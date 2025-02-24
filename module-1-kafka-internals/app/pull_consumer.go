package main

import (
	"context"
	"time"
)

const PullConsumerPollTimeout = 16 * time.Second
const PullConsumerFetchMinBytes = 256

func RunPullConsumer(ctx context.Context) func() error {
	return RunConsumer(ctx, PullConsumer, PullConsumerFetchMinBytes, int(PullConsumerPollTimeout.Milliseconds()))
}
