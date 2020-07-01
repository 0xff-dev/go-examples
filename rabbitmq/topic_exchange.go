package main

import (
	"log"
	"sync"
	"time"

	"github.com/streadway/amqp"
)


func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("connect error: ", err)
	}
	bindingKeys := []string{"kern.*"}
	defer conn.Close()
	var wg sync.WaitGroup
	for goRoutine := 0; goRoutine < 5; goRoutine++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			// DO receive message
			ch, err := conn.Channel()
			if err != nil {
				log.Fatalf("%d create channel error: %s", index, err)
			}
			if err = ch.ExchangeDeclare("logs_topic", "topic", true, false, false, false, nil); err != nil {
				log.Fatalf("%d declare exchange error: %s", index, err)
			}
			q, err := ch.QueueDeclare("", false, false, true, false, nil)
			if err != nil {
				log.Fatalf("%d declare queue error: %s", index, err)
			}
			for _, key := range bindingKeys {
				if err = ch.QueueBind(q.Name, key, "logs_topic", false, nil); err != nil {
					log.Fatalf("%d bind queue with logs error: %s", index, err)
				}
			}
			msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
			if err != nil {
				log.Fatalf("%d failed to register consumer", index)
			}
			count := 1
			for {
				select {
				case d := <-msgs:
					log.Printf("%d received [%d] data: %s\n", index, count, string(d.Body))
					count++
				case <-time.NewTimer(time.Second * 20).C:
					log.Printf("[%d]it's time to logout.", index)
					goto forBreak
				}
			}
		forBreak:
			log.Println("stop received message")
		}(goRoutine)
	}
	wg.Wait()
}
