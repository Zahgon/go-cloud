package main

import (
	"context"
	"flag"
	"log"
)

var (
	requestTopicURL = flag.String("request-topic", "mem://requests", "gocloud.dev/pubsub URL for request topic")
	requestSubURL   = flag.String("request-sub", "mem://requests", "gocloud.dev/pubsub URL for request subscription")
	bucketURL       = flag.String("bucket", "", "gocloud.dev/blob URL for image bucket")
	collectionURL   = flag.String("collection", "mem://orders/ID", "gocloud.dev/docstore URL for order collection")

	port         = flag.Int("port", 10538, "HTTP port for frontend")
	runFrontend  = flag.Bool("frontend", true, "run the frontend")
	runProcessor = flag.Bool("processor", true, "run the image processor")
)

func main() {
	flag.Parse()
	conf := config{
		requestTopicURL: *requestTopicURL,
		requestSubURL:   *requestSubURL,
		bucketURL:       *bucketURL,
		collectionURL:   *collectionURL,
	}
	frontend, processor, cleanup, err := setup(conf)
	if err != nil {
		log.Fatal(err)
	}
	defer cleanup()

	errc := make(chan error, 2)
	if *runFrontend {
		go func() { errc <- frontend.run(context.Background(), *port) }()
		log.Printf("listening on port %d", *port)
	} else {
		errc <- nil
	}
	if *runProcessor {
		go func() { errc <- processor.run(context.Background()) }()
		log.Println("processing")
	} else {
		errc <- nil
	}

	for i := 0; i < 2; i++ {
		if err := <-errc; err != nil {
			log.Fatal(err)
		}
	}
}

type config struct {
	requestTopicURL string
	requestSubURL   string
	bucketURL       string
	collectionURL   string
}

func setup(conf config) (_ *frontend, _ *processor, cleanup func(), err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}
