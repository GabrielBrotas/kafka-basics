package main

import (
	"fmt"
	"log"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	deliveryChannel := make(chan kafka.Event) // create a channel that will work like a callback to listen our message response
	
	producer := NewKafkaProducer()

	msg := "Legen... wait for it"
	topic := "first-topic"
	PublishMessage(msg, topic, producer, nil, deliveryChannel)
	// 1 --
	// producer.Flush(1000) // wait for the message response otherwise our program will finish and the message will not be send 

	// 2 --
	// synchronous way to wait the response from the message ------------------
	/* // but this a bad implementation because we are waiting a response to execute the rest of the code
	// and consenquently will affect our performance;
	event := <-deliveryChannel // event receiver
	msg_response := event.(*kafka.Message)
	if msg_response.TopicPartition.Error != nil {
		fmt.Println("Error on send message")
	} else {
		fmt.Println("Message sent: ", msg_response.TopicPartition)
		//message response 
		//Message sent:  first-topic[0]@8
		//<topic name>[Partition number]@<Parition offset>
	}
	*/

	// 3 --
	// asynchronous callback
	// go routing will create another thread for this function
	// go DeliveryReport(deliveryChannel)
	DeliveryReport(deliveryChannel) // syncronous
	// if the program finish before send the message it will not be sent
}

func NewKafkaProducer() (*kafka.Producer) {
	configMap := &kafka.ConfigMap{
		// ? dev
		// all brokers on our cluster
		// you can get the brokers by using 'docker-compose ps' and get the broker name
		// "bootstrap.servers": "kafka_kafka_1:9092",
		// you can check all availables options on https://github.com/edenhill/librdkafka/blob/master/CONFIGURATION.md
		// C/P = Consumer, Producer, * = both

		"delivery.timeout.ms": "0", // max time we want to wait a message to be delivered, 0 = infinity
		"acks": "1", // all = leader and followers, 0 = no return, 1 = just leader
		"enable.idempotence": "false", // the produtor is not idempotence, sometimes the message can be duplicated or not be in order, if true the ackes must be all

		// ? prod
		"bootstrap.servers": "pkc-2396y.us-east-1.aws.confluent.cloud:9092",
		"security.protocol": "SASL_SSL",
		"sasl.mechanisms": "PLAIN",
		"sasl.username": "",
		"sasl.password": "",

	}

	producer, err := kafka.NewProducer(configMap)

	if err != nil {
		log.Println(err.Error())
	}

	return producer
}

func PublishMessage(msg string, topic string, producer *kafka.Producer, key []byte, deliveryChannel chan kafka.Event) error {
	message := &kafka.Message{
		Value: []byte(msg), // convert message to byte because the message can be a binary file or other thing
		TopicPartition: kafka.TopicPartition{
			Topic: &topic, 
			Partition: kafka.PartitionAny, // any partition, kafka take care of it
		},
		Key: key,
	}

	// err := producer.Produce(message, nil) // message without callback
	err := producer.Produce(message, deliveryChannel) // send the response to the channel


	if err != nil {
		return err
	}

	return nil
}

func DeliveryReport(deliveryChannel chan kafka.Event) {
	// infinity loop that will receive our channel message,
	// every time we receive a message this loop will execute
	for event := range deliveryChannel {
		switch event_type := event.(type) {
			case *kafka.Message:
				if event_type.TopicPartition.Error != nil {
					fmt.Println("Error on send message", event_type.TopicPartition.Error)
				} else {
					fmt.Println("Message sent: ", event_type.TopicPartition)
					//message response 
					//Message sent:  first-topic[0]@8
					//<topic name>[Partition number]@<Parition offset>

					// save on the database the message was processed
					// send email
					// ...
				}

		}
	}
}