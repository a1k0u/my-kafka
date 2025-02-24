package main

import (
	"context"
	"flag"
	"log"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

var (
	flagBootstrapServer = flag.String("bootstrap-server", "localhost:9094", "Kafka bootstrap server")
	flagTopic           = flag.String("topic", "test", "Kafka topic")
	flagProduceInterval = flag.Duration("produce-interval", time.Second, "Pause before next produce")
)

func main() {
	ctx := context.Background()
	ctx, _ = signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)

	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(-1)

	g.Go(RunPullConsumer(ctx))
	g.Go(RunPushConsumer(ctx))
	g.Go(RunProducer(ctx))

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
