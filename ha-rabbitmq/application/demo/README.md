# For demo step by step

## Starter

### Run RabbitMQ

```
cd ../..
docker-compose up
```

### Preparing iterm

- 3 vertical window for subscriber / worker
- 1 window for publisher
- 1 window for hit endpoint

## Basic Pubsub


### Run Publisher

```
cd application/demo/1.\ basic\ pubsub/publisher

docker build . -t zeihanaulia/rabbitmq-1-publisher
docker run -it --rm --net rabbitnet  -e RABBIT_HOST=rabbitmq-1 -e RABBIT_PORT=5672 -e RABBIT_USERNAME=guest -e RABBIT_PASSWORD=guest -p 80:80 zeihanaulia/rabbitmq-1-publisher
```

### Run Subscriber / Receiver / Worker


```
cd application/demo/1.\ basic\ pubsub/receiver

docker build . -t zeihanaulia/rabbitmq-1-receiver
docker run -it --rm --net rabbitnet  -e RABBIT_HOST=rabbitmq-1 -e RABBIT_PORT=5672 -e RABBIT_USERNAME=guest -e RABBIT_PASSWORD=guest zeihanaulia/rabbitmq-1-receiver
```

