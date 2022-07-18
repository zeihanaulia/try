BINARY_NAME=receiver
IMAGE_NAME=zeihanaulia/rabbitmq-routing-receiver

run:
	docker build . -t ${IMAGE_NAME}
	docker run -it --rm --net rabbitnet  -e RABBIT_HOST=rabbitmq-1 -e RABBIT_PORT=5672 -e RABBIT_USERNAME=guest -e RABBIT_PASSWORD=guest ${IMAGE_NAME}
