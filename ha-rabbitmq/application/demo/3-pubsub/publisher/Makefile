BINARY_NAME=publisher
IMAGE_NAME=zeihanaulia/rabbitmq-pubsub-publisher

run:
	docker build . -t ${IMAGE_NAME}
	docker run -it --rm --net rabbitnet  -e RABBIT_HOST=rabbitmq-1 -e RABBIT_PORT=5672 -e RABBIT_USERNAME=guest -e RABBIT_PASSWORD=guest -p 8000:8000 ${IMAGE_NAME}

publish:
	curl -X POST http://localhost:8000/publish/hello