package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

var rabbit_host = os.Getenv("RABBIT_HOST")
var rabbit_port = os.Getenv("RABBIT_PORT")
var rabbit_user = os.Getenv("RABBIT_USERNAME")
var rabbit_password = os.Getenv("RABBIT_PASSWORD")

func main() {
	router := httprouter.New()
	router.POST("/publish/:message", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		submit(w, r, p)
	})

	fmt.Println("Running...")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func submit(_ http.ResponseWriter, _ *http.Request, p httprouter.Params) {
	msg := p.ByName("message")

	fmt.Println("Received message: " + msg)

	amqpURI := "amqp://" + rabbit_user + ":" + rabbit_password + "@" + rabbit_host + ":" + rabbit_port + "/"

	amqpConfig := amqp.NewDurableQueueConfig(amqpURI)
	amqpConfig.Exchange.Type = "topic"
	amqpConfig.Exchange.GenerateName = exchangeName()
	amqpConfig.Exchange.Durable = true
	amqpConfig.Exchange.AutoDeleted = false
	amqpConfig.Publish.GenerateRoutingKey = func(topic string) string {
		return "order.created"
	}

	publisher, err := amqp.NewPublisher(amqpConfig, watermill.NewStdLogger(false, false))
	if err != nil {
		panic(err)
	}

	fmt.Println("here")

	message := message.NewMessage(watermill.NewUUID(), []byte("Hello, world!"))
	fmt.Println("publish message", msg)
	if err := publisher.Publish("", message); err != nil {
		panic(err)
	}

	err = publisher.Close()
	fmt.Println("publish success!", err)
}

func exchangeName() func(topic string) string {
	return func(topic string) string {
		return "example_trx"
	}
}
