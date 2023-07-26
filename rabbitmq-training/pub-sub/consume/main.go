package main

import (
	"log"

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
		log.Printf("Rbbitmq channel error: %v", err)
	}

	qu, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Printf("queue declare error: %v", err)
	}

	err = ch.QueueBind(
		qu.Name, // queue name
		"",      // routing key
		"logs",  // exchange
		false,
		nil,
	)
	if err != nil {
		log.Printf("queue bind error: %v", err)
	}

	msgs, err := ch.Consume(
		qu.Name, // queue
		"",      // consumer
		true,    // auto-ack
		false,   // exclusive
		false,   // no-local
		false,   // no-wait
		nil,     // args
	)
	if err != nil {
		log.Printf("Consume err: %v", err)
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever

}
