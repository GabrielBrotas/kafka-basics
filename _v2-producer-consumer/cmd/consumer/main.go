package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	
	consumer := NewKafkaConsumer()

	topics := []string{"first-topic"}
	consumer.SubscribeTopics(topics, nil)
	
	for {
		msg, err := consumer.ReadMessage(-1) // -1 = stay connected forever

		if err == nil {
			fmt.Println(string(msg.Value), msg.TopicPartition)
		}

	}
}

func NewKafkaConsumer() (*kafka.Consumer) {
	configMap := &kafka.ConfigMap{
		// all brokers on our cluster
		// you can get the brokers by using 'docker-compose ps' and get the broker name
		// "bootstrap.servers": "kafka_kafka_1:9092",
		// you can check all availables options on https://github.com/edenhill/librdkafka/blob/master/CONFIGURATION.md
		// C/P = Consumer, Producer, * = both
		"group.id": "Group-X",
		"client.id": "goapp-consumer",
		"auto.offset.reset": "earliest",

		// ? prod
		"bootstrap.servers": "pkc-2396y.us-east-1.aws.confluent.cloud:9092",
		"security.protocol": "SASL_SSL",
		"sasl.mechanisms": "PLAIN",
		"sasl.username": "TOBVVEUMMYOII5BD",
		"sasl.password": "0gowmcCv+F8XOrM+CfdLqawPpXVJM3k/659KDKa4ZZ8pniTEpyiids0t4yp9BpJZ",

		"session.timeout.ms":45000,
	}

	consumer, err := kafka.NewConsumer(configMap)

	if err != nil {
		fmt.Println("error consumer = ", err.Error())
	}

	return consumer
}
