package main

import (
	"github.com/streadway/amqp"
	"log"
	"os"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://fusjoke:123456@localhost:5672/my_vhost")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "failed to conn channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs_topic",
		"topic",
		true,
		false,
		false,
		false,
		nil,
		)
	failOnError(err, "Failed to declare an exchange")

	q, err := ch.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
		)
	failOnError(err, "Failed to declare a queue")

	if len(os.Args) < 2 {
		log.Printf("Usage: %s [binding_key]...", os.Args[0])
		os.Exit(0)
	}

	for _, s := range os.Args[1:] {
		log.Printf("Binding queue %s to exchange %s with routing key %s",
			q.Name, "logs_topic", s)
		err = ch.QueueBind(
			q.Name,
			s,
			"logs_topic",
			false,
			nil,
			)
	}
	failOnError(err, "Failed to register a consumer")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
