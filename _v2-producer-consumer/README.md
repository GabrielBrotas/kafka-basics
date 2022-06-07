**librdkafka-dev**
lib created in C/C++ to communicate with kafka, we will use this lib to bind to our language;
https://github.com/edenhill/librdkafka

## How to run
1 - Create a topic on the kafka container
```bash
docker exec -it <kafka container name> bash

# create topic
kafka-topics --create --topic=first-topic --bootstrap-server=localhost:9092 --partitions=3
```

2 - Consumer messages
```bash
kafka-console-consumer --group=gp-x --topic=first-topic --bootstrap-server=localhost:9092
```