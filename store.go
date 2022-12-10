package go_gcp_misc

import (
	"cloud.google.com/go/datastore"
	"context"
	"fmt"
	"log"
	"os"
)

type Cat1 struct {
	Id    int
	Name  string
	Age   int
	Color string
}

func CreateCat1(kind string, cat *Cat1) {
	ctx := context.Background()
	projectID := os.Getenv("GCP_PROJ_ID")
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Panicln(err)
	}
	defer client.Close()
	name := cat.Name // name ?
	catKey := datastore.NameKey(kind, name, nil)
	if _, err := client.Put(ctx, catKey, &cat); err != nil {
		log.Fatalf("Error when saving cat: %v", err)
	}
	fmt.Printf("Saved %v: %v\n", catKey, cat.Name)

}
