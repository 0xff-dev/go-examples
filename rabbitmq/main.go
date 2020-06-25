package main

import (
	"fmt"
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
	// Message durability
	queue, err := channel.QueueDeclare("task_queue", true, false, false, false, nil)
	if err != nil {
		log.Fatal("Failed to declare a queue")
	}
	doneChan := make(chan struct{})
	go func() {
		UseQueue(channel, queue)
		doneChan <- struct{}{}
	}()
	<- doneChan
}

func UseQueue(channel *amqp.Channel, queue amqp.Queue) {
	count := 1
	for ; count <= 5; count ++ {
		message := "hello rabbitmq"
		log.Println("send message: ", count)
		if err := channel.Publish("", queue.Name, false, false, amqp.Publishing{
			ContentType: "text/plain",
			DeliveryMode: amqp.Persistent,
			Body:        []byte(fmt.Sprintf("%d-%s", count, message)),
		}); err != nil {
			log.Fatal("Failed to send message to queue")
		}
		<- time.NewTimer(time.Second * 5).C
	}
}
