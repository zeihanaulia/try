# BUILD
FROM golang:1.18-alpine AS build

WORKDIR /app

COPY . /app/

RUN go mod download

RUN go build -o receiver

# DEPLOY
FROM alpine:latest as certs

RUN apk --update add ca-certificates

WORKDIR /

COPY --from=build /app/receiver /receiver

EXPOSE 80

# CMD ["/receiver", "#"]
# CMD ["/receiver", "order.*"]
# CMD ["/receiver", "*.created"]
CMD ["/receiver", "order.*", "*.created"]