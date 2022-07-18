package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/julienschmidt/httprouter"
	amqp "github.com/rabbitmq/amqp091-go"
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
	message := p.ByName("message")

	fmt.Println("Received message: " + message)

	host1 := "amqp://" + rabbit_user + ":" + rabbit_password + "@" + rabbit_host + ":" + rabbit_port + "/"

	conn, err := amqp.Dial(host1)
	if err != nil {
		log.Fatalf("%s: %s", "Failed to connect to RabbitMQ", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("%s: %s", "Failed to open a channel", err)
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs",   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	split := strings.Split(message, " ")
	body := bodyFrom(split)

	err = ch.Publish(
		"logs", // exchange
		"",     // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	failOnError(err, "Failed to publish a message")

	fmt.Println("publish success!")
}

func bodyFrom(args []string) string {
	fmt.Println("len args", len(args))
	var s string
	if (len(args) < 2) || args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
