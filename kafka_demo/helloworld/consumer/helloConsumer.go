package consumer

import (
	"github.com/Shopify/sarama"
	"kafka_demo/conf"
	"log"
)

func Consume(topic string) {
	config := sarama.NewConfig()
	consume, err := sarama.NewConsumer([]string{conf.HOST}, config)

	if err != nil {
		log.Fatal("failed to new cunsume :", err)
	}
	defer consume.Close()

	partitionConsumer, err := consume.ConsumePartition(topic,0, sarama.OffsetNewest)
	defer partitionConsumer.Close()
	for message := range partitionConsumer.Messages() {
		log.Printf("[Consumer] partitionid: %d; offset:%d, value: %s\n", message.Partition, message.Offset, string(message.Value))
	}
}
