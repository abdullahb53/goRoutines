package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	amqp "github.com/rabbitmq/amqp091-go"
)

var conn *amqp.Connection

var ch *amqp.Channel

var conn2WebSocket *websocket.Conn

var Messages = make(chan []byte)

var GlobalMessage []byte

type queList struct {
	Factory_id string
	Added      int
}

type CreatedQueueForMainPage struct {
	Factory_id string
	IsWorking  bool
}

var CancelChannel = make(chan string)

var QueueForMainPages = make([]CreatedQueueForMainPage, 0)

var queue = make([]queList, 0)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

// ------------------------------------------------------
// --------------------- WEB SOCKET ---------------------
// ------------------------------------------------------
func home(w http.ResponseWriter, r *http.Request) {
	log.Println(w, "Index Page")
}

var origins = []string{"http://127.0.0.1:3000", "http://localhost:3000"}

var upgrader = websocket.Upgrader{
	// Resolve cross-domain problems
	CheckOrigin: func(r *http.Request) bool {
		var origin = r.Header.Get("origin")
		log.Println("origin:", origin)
		for _, allowOrigin := range origins {
			if origin == allowOrigin {
				return true
			}
		}
		return false
	}}

func socketHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade our raw HTTP connection to a websocket based one
	var err error
	conn2WebSocket, err = upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Error during connection upgradation:", err)
		return
	}

	defer conn2WebSocket.Close()

	go func() {
		for {

			messageType, fromClientMessage, err := conn2WebSocket.ReadMessage()
			log.Println(messageType) //1
			if err != nil {
				log.Println("Error during message reading:", err)

				CancelChannel <- string(fromClientMessage)
				time.Sleep(1 * time.Second)
				// break
			}
			if err == nil {
				go CreateWithSocketMessageQueue(string(fromClientMessage))
				log.Printf("Received: %s", fromClientMessage)
			}

		}

	}()

	// The event loop
	for {

		sendMessageToFrontend := <-Messages
		err = conn2WebSocket.WriteMessage(1, sendMessageToFrontend)
		if err != nil {
			log.Println("Error during message reading:", err)
			// GETTING ID FOR FactoryCakkedID
			FactoryCancelledID := "123"
			CancelChannel <- FactoryCancelledID
			time.Sleep(2 * time.Second)
		}

	}
}

// ------------------------------------------------------
// --------------------- WEB SOCKET ---------------------
// ------------------------------------------------------

func CreateWithSocketMessageQueue(queueName string) {

	var ifNotExistForQueueDeclareCounter int = 0
	var Flag bool = false
	var lengthQueueForMainPages int = 0
	lengthQueueForMainPages = len(QueueForMainPages)

	for i := 0; i < lengthQueueForMainPages; i++ {
		if QueueForMainPages[i].Factory_id == queueName {
			if !QueueForMainPages[i].IsWorking {
				QueueForMainPages[i].IsWorking = true
				Flag = false
				break
			} else {
				Flag = true
				break
			}

		} else {
			ifNotExistForQueueDeclareCounter++
		}
	}

	if ifNotExistForQueueDeclareCounter == lengthQueueForMainPages {
		QueueForMainPages = append(QueueForMainPages, CreatedQueueForMainPage{
			Factory_id: queueName,
			IsWorking:  true,
		})
		Flag = false
	}
	ifNotExistForQueueDeclareCounter = 0

	if !Flag {

		QueueForMainPages = append(QueueForMainPages, CreatedQueueForMainPage{
			Factory_id: queueName,
			IsWorking:  true,
		})

		q, err := ch.QueueDeclare(
			queueName, // name
			false,     // durable
			false,     // delete when unused
			false,     // exclusive
			false,     // no-wait
			nil,       // arguments
		)
		failOnError(err, "Failed to declare a queue")

		msgs, err := ch.Consume(
			q.Name, // queue
			q.Name, // consumer
			true,   // auto-ack
			false,  // exclusive
			false,  // no-local
			false,  // no-wait
			nil,    // args
		)
		failOnError(err, "Failed to register a consumer")

		go func() {

			endofLifeName := queueName

			select {

			case a := <-CancelChannel:
				if a == endofLifeName {
					for i := 0; i < len(QueueForMainPages); i++ {
						if QueueForMainPages[i].Factory_id == endofLifeName {
							QueueForMainPages[i].IsWorking = false
						}
					}
					return
				}

			default:
				for d := range msgs {
					log.Printf("Received a message: %s", d.Body)
					Messages <- d.Body
					// TODO: append data to "Notification Collection" from RabbitMQ..

				}
			}

		}()

	} else {
		log.Println("Already created QueueName ->", queueName)
	}

}

func main() {

	var err error
	conn, err = amqp.Dial("amqp://admin:admin@127.0.0.1:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	log.Println("Connected to RabbitMQ - Notification for Navbar' Pages..")

	ch, err = conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
	log.Println("Connected to channel")

	var forever chan struct{}

	go func() {

		// ------------------------------------------------------
		// --------------------- WEB SOCKET ---------------------
		// ------------------------------------------------------
		http.HandleFunc("/socketSender", socketHandler)
		http.HandleFunc("/", home)
		log.Fatal(http.ListenAndServe("127.0.0.1:18085", nil))
		// ------------------------------------------------------
		// --------------------- WEB SOCKET ---------------------
		// ------------------------------------------------------
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}
