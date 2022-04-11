package main

import (
	"kafka_demo/conf"
	"kafka_demo/helloworld/consumer"
)

func main() {
	topic := conf.Topic
	consumer.Consume(topic)
}
