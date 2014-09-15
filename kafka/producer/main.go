package main

import (
	"fmt"

	kafka "github.com/Shopify/sarama"
)

func main() {
	client, err := kafka.NewClient("client_id", []string{"localhost:9092"}, kafka.NewClientConfig())
	if err != nil {
		panic(err)
	} else {
		fmt.Println("> connected")
	}
	defer client.Close()

	producer, err := kafka.NewProducer(client, nil)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	err = producer.SendMessage("my_topic", nil, kafka.StringEncoder("testing 123"))
	if err != nil {
		panic(err)
	} else {
		fmt.Println("> message sent")
	}
}
