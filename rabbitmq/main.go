package main

import (
	"github.com/streadway/amqp"
	"log"
	"time"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("connect error: ", err)
	}
	defer conn.Close()

	// a channel session
	channel, err := conn.Channel()
	if err != nil {
		log.Fatal("create a channel error: ", err)
	}
	defer channel.Close()
	queue, err := channel.QueueDeclare("hello", false, false, false, false, nil)
	if err != nil {
		log.Fatal("Failed to declare a queue")
	}
	count := 1
	for ; count <= 5; count ++ {
		message := "hello rabbitmq"
		log.Println("send message: ", count)
		if err = channel.Publish("", queue.Name, false, false, amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		}); err != nil {
			log.Fatal("Failed to send message to queue")
		}
		<- time.NewTimer(time.Second * 5).C
	}
}
