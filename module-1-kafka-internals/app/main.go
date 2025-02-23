package main

import (
	"flag"
	"time"
)

var (
	flagBootstrapServer = flag.String("bootstrap-server", "localhost:9092", "Kafka bootstrap server")
	flagTopic           = flag.String("topic", "test", "Kafka topic")
	flagProduceInterval = flag.Duration("produce-interval", 300*time.Millisecond, "Pause before next produce")
)

func main() {

	// producer.Produce()
}
