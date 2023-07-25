package main

import (
	"context"
	"os"
	"testing"

	amqp "github.com/rabbitmq/amqp091-go"
	assert "github.com/stretchr/testify/assert"
)

func Produce(t *testing.T, ch *amqp.Channel, rbmqQueue amqp.Queue, ctx context.Context, input string) {
	err := ch.PublishWithContext(ctx,
		"",             // exchange
		rbmqQueue.Name, // routing key
		false,          // mandatory
		false,          // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(input),
		})
	assert.Nil(t, err)
}

func TestMain(t *testing.T) {
	rmq := newRabbMQ()

	ch, err := rmq.conn.Channel()

	assert.Nil(t, err)

	interruptSignal := make(chan os.Signal, 1)

	rbmqQueue := getRbbmQueue("queue1", ch, interruptSignal)

	ctx := context.Background()

	var (
		TIMES    int
		inputStr string
	)
	TIMES = 10
	datas := make([]string, 0)
	for i := 0; i < TIMES; i++ {
		inputStr = string(rune(i))
		datas = append(datas, inputStr)

		Produce(t, ch, rbmqQueue, ctx, inputStr)
	}

	messages, err := ch.Consume(
		rbmqQueue.Name, // queue
		"",             // consumer
		true,           // auto-ack
		false,          // exclusive
		false,          // no-local
		false,          // no-wait
		nil,            // args
	)
	assert.Nil(t, err)
	count := 0
	for d := range messages {
		count++
		assert.Equal(t, string(d.Body), datas[count-1])
		if count >= TIMES {
			break
		}
	}

}
