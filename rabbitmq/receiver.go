package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("connect error: ", err)
	}
	defer conn.Close()
	channel ,err := conn.Channel()
	if err != nil {
		log.Fatal("Failed create channel")
	}
	defer channel.Close()

	if err = channel.Qos(1, 0, false); err != nil {
		log.Fatal("set worker qos error: ", err)
	}
	// message durability====> durable=true
	queue, err := channel.QueueDeclare("task_queue", true, false, false, false, nil)
	if err != nil {
		log.Fatal("Failed to declare queue")
	}
	msgs, err := channel.Consume(queue.Name, "customer", true, false, false, false, nil)
	if err != nil {
		log.Fatal("Consume msg error: ", err)
	}
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for {
			for msg := range msgs {
				log.Println("Received msg: ", string(msg.Body), " \nDone!!!")
				msg.Ack(false) // custom auto-ack=false. need do it by self.
			}
		}
	}()
	log.Println("[*] Wait for message, To exit press CTRL+C")
	<- signalChan
}
