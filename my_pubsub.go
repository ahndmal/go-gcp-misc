package go_gcp_misc

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"log"
	"os"
	"sync/atomic"
	"time"
)

func CreateMessage() {
	ctx := context.Background()

	projId := os.Getenv("GCP_PROJ_ID")

	client, err := pubsub.NewClient(ctx, projId)
	//client, err := pubsub.NewClient(ctx, option.WithCredentialsFile("path/to/keyfile.json"))
	if err != nil {
		fmt.Errorf(">>> ERROR :: pubsub.NewClient: %v", err)
	}

	defer client.Close()

	topic := client.Topic("wiki-pages-main")
	msg := "Hello from GO!"
	//msg := pubsub.Message{
	//	Data: []byte("Hello from GO!"),
	//}
	publishResult := topic.Publish(ctx, &pubsub.Message{
		Data: []byte(msg),
	})

	id, err := publishResult.Get(ctx)
	if err != nil {
		fmt.Errorf("pubsub: result.Get: %v", err)
	}
	log.Printf(">>> saved message with Id %s", id)
}

func Listen() {
	ctx := context.Background()
	projId := os.Getenv("GCP_PROJ_ID")
	client, err := pubsub.NewClient(ctx, projId)
	if err != nil {
		fmt.Errorf(">>> ERROR :: pubsub.NewClient: %v", err)
	}

	defer client.Close()

	sub := client.Subscription("sub1")
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var received int32
	err = sub.Receive(ctx, func(ctx context.Context, message *pubsub.Message) {
		fmt.Printf("Got message: %q\n", string(message.Data))
		atomic.AddInt32(&received, 1)
		message.Ack()
	})
	if err != nil {
		fmt.Errorf("sub.Receive: %v", err)
	}
}
