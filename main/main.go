package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"log"
)

type Cat struct {
	Id   int
	Name string
	Age  int
}

func main() {
	ctx := context.Background()

	projId := "silver-adapter-307718"
	//kind := "Cat1"

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
