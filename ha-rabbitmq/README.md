# HA RabbitMQ

## Running

```bash
docker-compose restart rabbitmq-1
docker-compose restart rabbitmq-2
docker-compose restart rabbitmq-3

# ambil .erlang.cookie
docker exec -it rabbitmq-1 cat /var/lib/rabbitmq/.erlang.cookie
```

## Tambahkan pada env

```bash
- RABBITMQ_ERLANG_COOKIE=SWBPIUIBKTJGWOURYHCD

rabbitmq-1:
    image: rabbitmq:3.8-management-alpine
    container_name: rabbitmq-1
    hostname: rabbitmq-1
    environment:
      - RABBITMQ_ERLANG_COOKIE=SWBPIUIBKTJGWOURYHCD
      - RABBITMQ_DEFAULT_USER=myuser
      - RABBITMQ_DEFAULT_PASS=mypassword
    ports:
        - 5672:5672
        - 15672:15672
    networks:
        - rabbitmq_go_net

docker-compose restart rabbitmq-1
docker-compose restart rabbitmq-2
docker-compose restart rabbitmq-3
```

## Setting cluster manual

```bash
docker exec -it rabbitmq-1 rabbitmqctl cluster_status

# Join cluster rabbitmq-2

docker exec -it rabbitmq-2 rabbitmqctl stop_app
docker exec -it rabbitmq-2 rabbitmqctl reset
docker exec -it rabbitmq-2 rabbitmqctl join_cluster rabbit@rabbitmq-1
docker exec -it rabbitmq-2 rabbitmqctl start_app
docker exec -it rabbitmq-2 rabbitmqctl cluster_status

# Join cluster rabbitmq-2

docker exec -it rabbitmq-3 rabbitmqctl stop_app
docker exec -it rabbitmq-3 rabbitmqctl reset
docker exec -it rabbitmq-3 rabbitmqctl join_cluster rabbit@rabbitmq-1
docker exec -it rabbitmq-3 rabbitmqctl start_app
docker exec -it rabbitmq-3 rabbitmqctl cluster_status
```

## Test application

### Publisher

```bash
cd application/publisher

docker build . -t zeihanaulia/rabbitmq-publisher
docker run -it --rm --net rabbitnet  -e RABBIT_HOST=rabbitmq-1 -e RABBIT_PORT=5672 -e RABBIT_USERNAME=guest -e RABBIT_PASSWORD=guest -p 80:80 zeihanaulia/rabbitmq-publisher
```

### Receiver

```
cd application/receiver

docker build . -t zeihanaulia/rabbitmq-receiver
docker run -it --rm --net rabbitnet  -e RABBIT_HOST=rabbitmq-1 -e RABBIT_PORT=5672 -e RABBIT_USERNAME=guest -e RABBIT_PASSWORD=guest zeihanaulia/rabbitmq-receiver
```

## Basic Mirroring

```bash
docker exec -it rabbitmq-1 bash

rabbitmqctl set_policy ha-fed \
    ".*" '{"federation-upstream-set":"all", "ha-mode":"nodes", "ha-params":["rabbit@rabbitmq-1","rabbit@rabbitmq-2","rabbit@rabbitmq-3"]}' \
    --priority 1 \
    --apply-to queues
```


## Enable federation plugin
docker exec -it rabbitmq-1 rabbitmq-plugins enable rabbitmq_federation 
docker exec -it rabbitmq-2 rabbitmq-plugins enable rabbitmq_federation
docker exec -it rabbitmq-3 rabbitmq-plugins enable rabbitmq_federation