# BUILD
FROM golang:1.18-alpine AS build

WORKDIR /app

COPY . /app/

RUN go mod download

RUN go build -o publisher

# DEPLOY
FROM alpine:latest as certs

RUN apk --update add ca-certificates

WORKDIR /

COPY --from=build /app/publisher /publisher

EXPOSE 8000

CMD ["/publisher"]