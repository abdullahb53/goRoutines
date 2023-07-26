package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbMQ struct {
	conn  *amqp.Connection
	close chan struct{}
}

func getRbbmQueue(q string, ch *amqp.Channel, interruptChan chan os.Signal) amqp.Queue {
	realQueue, err := ch.QueueDeclare(
		q,     // queue name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Printf("Failed to declare a queue [%v], error: %x", q, err)
		interruptChan <- os.Interrupt
	}
	return realQueue
}

func newRabbMQ() *RabbMQ {
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	if err != nil {
		log.Fatal("RabbitMQ connection is failed.")
	}

	return &RabbMQ{
		conn:  conn,
		close: make(chan struct{}, 1),
	}

}

type sendedConsumingFunc func(context.Context)

func injectCall(isNum int, messages <-chan amqp.Delivery) sendedConsumingFunc {
	return func(ctx context.Context) {

		go func(ctx context.Context) {
			count := 0
			ctx2, cancel := context.WithTimeout(ctx, time.Second*2)
			defer cancel()

		breakthere:
			for {
				select {
				case d := <-messages:
					log.Printf(" [x] %s", d.Body)
					count++
					if count >= isNum {
						break breakthere
					}
				case <-ctx2.Done():
					fmt.Printf("Context cancelled: [%v]\n", ctx2.Err())
					break breakthere
				}
			} // for-end

		}(ctx) // routine-end

	}
}

func main() {

	var p bool // produce
	var c bool // consume
	var queue string
	flag.BoolVar(&c, "c", false, "c")
	flag.BoolVar(&p, "p", false, "p")
	flag.StringVar(&queue, "q", "", "q")
	flag.Parse()

	if c && p {
		fmt.Println("You should select just one statement. (Consume -c / Produce -p)")
		os.Exit(0)
	}

	if queue == "" {
		fmt.Println("queue name is not exist.")
		os.Exit(0)
	}

	rmq := newRabbMQ()
	defer rmq.conn.Close()

	channel, err := rmq.conn.Channel()
	if err != nil {
		log.Fatalf("connection channel error: %X", err)
	}
	defer channel.Close()

	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-rmq.close
		channel.Close()
		rmq.conn.Close()
	}()

	go func() {
		sig := <-interruptChan
		fmt.Printf("Interrupt recevied, %v", sig)
		fmt.Println("The process will be interrupted after 2 seconds..")
		rmq.close <- struct{}{}
		time.Sleep(2 * time.Second)
		os.Exit(0)
	}()

	theQueue := getRbbmQueue(queue, channel, interruptChan)

	if c {

		messages, err := channel.Consume(
			theQueue.Name, // queue
			"",            // consumer
			true,          // auto-ack
			false,         // exclusive
			false,         // no-local
			false,         // no-wait
			nil,           // args
		)
		if err != nil {
			fmt.Printf("Consuming problem error: %v", err)
			interruptChan <- os.Interrupt
		}

		fmt.Println("Hi! i am consumer.")
		fmt.Println("Consume -> 1-10 integers.")
		fmt.Println("Exit -> 0.")
		scanner := bufio.NewScanner(os.Stdin)

		for {
			if scanner.Scan() {
				input := scanner.Text()
				fmt.Printf("You entered:%v\n", input)

				isNum, err := strconv.Atoi(input)
				if err != nil {
					fmt.Printf("%v is not valid value, it must be 'int'.", input)
					continue
				}
				if isNum < 0 || isNum > 10 {
					fmt.Printf("%v is not valid value, it must be between 0-10 values.", input)
					continue
				}

				if isNum == 0 {
					interruptChan <- os.Interrupt
				}
				ctx := context.TODO()

				injectedFun := injectCall(isNum, messages)
				injectedFun(ctx)

			}

		}

	} else if p {

		fmt.Println("Hi! i am producer.")
		fmt.Println("Produce string things -> Max string length is 15.")
		fmt.Println("Exit -> 0.")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		scanner := bufio.NewScanner(os.Stdin)
		for {
			if scanner.Scan() {
				input := scanner.Text()
				fmt.Printf("You entered:%v\n", input)

				if len(input) < 1 || len(input) > 15 {
					log.Println("Please enter string as 0-15 length")
					continue
				}

				theQueue := getRbbmQueue(queue, channel, interruptChan)

				err = channel.PublishWithContext(ctx,
					"",            // exchange
					theQueue.Name, // routing key
					false,         // mandatory
					false,         // immediate
					amqp.Publishing{
						ContentType: "text/plain",
						Body:        []byte(input),
					})
				if err != nil {
					log.Printf("Publishing [%v] error: %X", input, err)
					continue
				}
				log.Printf(" [x] Sent %s\n", input)

			}
		}

	} else {
		fmt.Println("You should provide *consuming or producing section. (-c / -p) -q queueName")
		os.Exit(0)
	}

}
