package main

import (
	"context"
	"log"
	"os"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
)

var rabbit_host = os.Getenv("RABBIT_HOST")
var rabbit_port = os.Getenv("RABBIT_PORT")
var rabbit_user = os.Getenv("RABBIT_USERNAME")
var rabbit_password = os.Getenv("RABBIT_PASSWORD")

func main() {
	amqpURI := "amqp://" + rabbit_user + ":" + rabbit_password + "@" + rabbit_host + ":" + rabbit_port + "/"

	amqpConfig := amqp.NewDurableQueueConfig(amqpURI)
	amqpConfig.Exchange.Type = "topic"
	amqpConfig.Exchange.GenerateName = exchangeName()
	amqpConfig.Exchange.Durable = true
	amqpConfig.Exchange.AutoDeleted = false

	amqpConfig.Queue.Durable = true
	amqpConfig.Queue.Exclusive = false
	amqpConfig.QueueBind.GenerateRoutingKey = func(topic string) string {
		return "order.created"
	}

	subscriber, err := amqp.NewSubscriber(
		// This config is based on this example: https://www.rabbitmq.com/tutorials/tutorial-two-go.html
		// It works as a simple queue.
		//
		// If you want to implement a Pub/Sub style service instead, check
		// https://watermill.io/pubsubs/amqp/#amqp-consumer-groups
		amqpConfig,
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		panic(err)
	}

	messages, err := subscriber.Subscribe(context.Background(), "example.topic")
	if err != nil {
		panic(err)
	}
	var forever chan struct{}

	go process(messages)

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func process(messages <-chan *message.Message) {
	for msg := range messages {
		log.Printf("received message: %s, payload: %s", msg.UUID, string(msg.Payload))

		// we need to Acknowledge that we received and processed the message,
		// otherwise, it will be resent over and over again.
		msg.Ack()
	}
}
func exchangeName() func(topic string) string {
	return func(topic string) string {
		return "example_trx"
	}
}
