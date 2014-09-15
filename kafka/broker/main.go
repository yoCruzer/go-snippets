package main

import (
	"fmt"

	kafka "github.com/Shopify/sarama"
)

func main() {
	broker := kafka.NewBroker("localhost:9092")
	err := broker.Open(nil)
	if err != nil {
		panic(err)
	}
	defer broker.Close()

	request := kafka.MetadataRequest{Topics: []string{"myTopic"}}
	response, err := broker.GetMetadata("myClient", &request)
	if err != nil {
		panic(err)
	}

	fmt.Println("There are", len(response.Topics), "topics active in the cluster.")
}
