
## How to run
```bash
docker-compose up -d
```

We can access kafk through an SDK or CLI.

Inside kafka container

## First Example - Produce and Consume Message
```bash
docker exec -it <kafka container name> bash

# Create topics
kafka-topics # check the available flgas

kafka-topics --create --topic=first-topic --bootstrap-server=localhost:9092 --partitions=3

kafka-topics --list --bootstrap-server=localhost:9092 # list available topics

kafka-topics --topic=first-topic --describe --bootstrap-server=localhost:9092

# Consumer
kafka-console-consumer --topic=first-topic --bootstrap-server=localhost:9092
## if we want to read the messages from the beging we can pass the flag --from-beginning

# Producer
## Open another console access the container again 
kafka-console-producer --topic=first-topic --bootstrap-server=localhost:9092
## Now we can write any message and send to see if the consumer is receiving the message
```

## Second Example - Consumer Group
```bash
## Create 2 consumer cli and 1 Producer
docker exec -it <kafka container name> bash

# Create topic
kafka-topics --create --topic=first-topic --bootstrap-server=localhost:9092 --partitions=3

# Create 2 Consumer with the same group
kafka-console-consumer --group=x --topic=first-topic --bootstrap-server=localhost:9092

# Producer
## Open another console access the container again 
kafka-console-producer --topic=first-topic --bootstrap-server=localhost:9092
## Now when we send a message just one consumer will read, and then the next message the other consumer will read
## The first consumer will read the partition 1 and 2, and the other consumer is reading the partition 3

# Check the consumers connected to our partitions
kafka-consumer-groups --group=x --describe --bootstrap-server=localhost:9092
```

## Visualization with Control Center
Create by confluentinc company, the kafka creator.
url: localhost:9021


## Clean up
```bash
docker-compose down
```