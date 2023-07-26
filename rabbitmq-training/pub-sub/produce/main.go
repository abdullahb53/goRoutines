package main

import (
	"context"
	"log"
	"strconv"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbMQ struct {
	conn *amqp.Connection
}

func newRabbMQ() *RabbMQ {
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	if err != nil {
		log.Fatal("RabbitMQ connection is failed.")
	}

	return &RabbMQ{
		conn: conn,
	}

}

func main() {

	rbbmq := newRabbMQ()

	ch, err := rbbmq.conn.Channel()
	if err != nil {
		log.Fatalf("Creating channel error: %v ", err)
	}

	err = ch.ExchangeDeclare(
		"logs",   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		log.Printf("Exchange declare error: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for {

		body := strconv.Itoa(time.Now().Second())
		err = ch.PublishWithContext(ctx,
			"logs", // exchange
			"",     // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		if err != nil {
			log.Printf("Puslish  error: %v", err)
		}

		time.Sleep(time.Second * 2)
	}

}
