BINARY_NAME=publisher

run:
	docker build . -t zeihanaulia/rabbitmq-watermill-publisher
	docker run -it --rm --net rabbitnet  -e RABBIT_HOST=rabbitmq-1 -e RABBIT_PORT=5672 -e RABBIT_USERNAME=guest -e RABBIT_PASSWORD=guest -p 8000:8000 zeihanaulia/rabbitmq-watermill-publisher

publish:
	curl -X POST http://localhost:8000/publish/hello