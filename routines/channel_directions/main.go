package main

import (
	"fmt"
)

const Message = "Hi Dev."

func SendValue(from chan<- string, value string) {
	from <- value

}

func ReceiveValue(from <-chan string, to chan<- string) {
	message := <-from
	to <- message

}

func main() {

	from := make(chan string, 1)
	to := make(chan string, 1)

	SendValue(from, Message)
	ReceiveValue(from, to)

	fmt.Println("Received:", <-to)

}
