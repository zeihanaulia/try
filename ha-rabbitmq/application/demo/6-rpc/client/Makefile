BINARY_NAME=receiver
IMAGE_NAME=zeihanaulia/rabbitmq-rpc-receiver

run:
	docker build . -t ${IMAGE_NAME}
	docker run -it -p 8000:8000 --rm --net rabbitnet  -e RABBIT_HOST=rabbitmq-1 -e RABBIT_PORT=5672 -e RABBIT_USERNAME=guest -e RABBIT_PASSWORD=guest ${IMAGE_NAME}
