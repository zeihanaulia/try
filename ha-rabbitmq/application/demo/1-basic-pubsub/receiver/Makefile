BINARY_NAME=receiver

run:
	docker build . -t zeihanaulia/rabbitmq-receiver
	docker run -it --rm --net rabbitnet  -e RABBIT_HOST=rabbitmq-1 -e RABBIT_PORT=5672 -e RABBIT_USERNAME=guest -e RABBIT_PASSWORD=guest zeihanaulia/rabbitmq-receiver
