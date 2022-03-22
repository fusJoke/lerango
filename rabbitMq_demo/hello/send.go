package main

import (
	"github.com/streadway/amqp"
	"log"
)

/**
发送消息
 */
func main() {
	//conn, err := amqp.Dial("amqp://admin:admin@localhost:5672")
	conn, err := amqp.Dial("amqp://fusjoke:123456@localhost:5672/my_vhost")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch ,err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
		)

	failOnError(err, "Failed to Queue Declare")

	body := "hello world"
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(body),
		})
	failOnError(err, "Failed to publish a message")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
	}
}