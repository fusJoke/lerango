package main

import (
	"kafka_demo/conf"
	"kafka_demo/helloworld/producer"
)

func main() {
	topic := conf.Topic
	producer.Produce(topic, 1000)
}
