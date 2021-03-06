package main

import (
	"github.com/streadway/amqp"
	"log"
	"os"
	"strings"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://fusjoke:123456@localhost:5672/my_vhost")
	failOnError(err, "conn failed")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "channel failed to conn")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"log_direct",
		"direct",
		true,
		false,
		false,
		false,
		nil,
		)
	failOnError(err, "channel failed to conn")

	body := bodyForm(os.Args)
	err = ch.Publish(
		"log_direct",
		severityForm(os.Args),
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(body),
		})
	failOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", body)

}

func bodyForm(args []string) string {
	var s string
	if ( len(s) < 3) || os.Args[2] == "" {
		s ="hello"
	} else {
		s = strings.Join(args[2:], " ")
	}
	return s
}

func severityForm(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "info"
	} else {
		s = os.Args[1]
	}
	return s
}