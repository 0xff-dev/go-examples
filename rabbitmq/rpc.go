package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"math/rand"
	"strconv"
)

func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func fibRpc()  {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("connect rabbitmq error: %s", err)
		return
	}
	defer conn.Close()
	channel, err := conn.Channel()
	if err != nil {
		log.Fatal("connect channel error")
	}
	defer channel.Close()
	q, err := channel.QueueDeclare("", false, false, true, false, nil)
	if err != nil {
		log.Fatal("declare queue error")
	}
	msgs, err := channel.Consume("", q.Name, true, false, false, false, nil)
	if err != nil {
		log.Fatal("faield to register consumer")
	}
	key := "123456"
	for cnt := 0; cnt < 10; cnt ++ {
		fibN := randInt(10, 30)
		log.Printf("calculate %d", fibN)
		err = channel.Publish("", "rpc_queue", false, false, amqp.Publishing{
			ContentType: "text/plain",
			CorrelationId: key,
			ReplyTo: q.Name,
			Body: []byte(fmt.Sprintf("%d", fibN)),
		})
	}
	for d := range msgs {
		if key == d.CorrelationId {
			res, _ := strconv.Atoi(string(d.Body))
			log.Println("received ", res)
		}
	}
}
func main() {
	fibRpc()
}
