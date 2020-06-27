FROM golang:1.14.4-buster

WORKDIR /app
COPY ./server.go /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server server.go
EXPOSE 8000
