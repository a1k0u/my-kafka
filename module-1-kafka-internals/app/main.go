package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var (
	flagBootstrapServer = flag.String("bootstrap-server", "localhost:9094", "Kafka bootstrap server")
	flagTopic           = flag.String("topic", "yandex-practicum", "Kafka topic")
	flagProduceInterval = flag.Duration("produce-interval", time.Second, "Pause before next produce")
)

func main() {
	flag.Parse()

	ctx := context.Background()
	ctx, _ = signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)

	services := []struct {
		service func(context.Context) error
		enabled bool
	}{
		{RunPullConsumer, true},
		{RunPushConsumer, true},
		{RunProducer, true},
	}

	var wg sync.WaitGroup
	errs := make(chan error, len(services))

	runningServices := 0
	for _, s := range services {
		if !s.enabled {
			continue
		}

		wg.Add(1)
		go func(service func(context.Context) error) {
			defer wg.Done()
			errs <- service(ctx)
		}(s.service)

		runningServices++
	}

	wg.Wait()

	for _ = range runningServices {
		if err := <-errs; err != nil && !errors.Is(err, context.Canceled) {
			log.Fatal(err)
		}
	}
}
