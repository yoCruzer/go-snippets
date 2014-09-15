package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func main() {
	broker := NewBroker("localhost:9092")
	err := broker.Open(nil)
	if err != nil {
		return err
	}
	defer broker.Close()

	request := MetadataRequest{Topics: []string{"myTopic"}}
	response, err := broker.GetMetadata("myClient", &request)
	if err != nil {
		return err
	}

	fmt.Println("There are", len(response.Topics), "topics active in the cluster.")

	return nil
}
