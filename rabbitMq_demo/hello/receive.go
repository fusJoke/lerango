package main

import (
	"log"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
	}
}

func main() {
	//conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	conn, err := amqp.Dial("amqp://fusjoke:123456@localhost:5672/my_vhost")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when usused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
		)
	failOnError(err, "Failed to register a consumer")
	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Receive a message： %s", d.Body)
		}
	}()

	log.Printf(" [*]waiting for messages. to exit prees ctrl+c")
	<-forever
}


