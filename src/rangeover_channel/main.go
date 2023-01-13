package main

import (
	"fmt"
)

func SendDataToChannel(channel chan<- string, data string) {
	channel <- data

}

func main() {
	fmt.Println("RangeOverChannel")

	var queue = make(chan string, 5)

	SendDataToChannel(queue, "1")
	SendDataToChannel(queue, "2")

	SendDataToChannel(queue, "3")
	SendDataToChannel(queue, "4")
	// SendDataToChannel(queue, "5")

	close(queue)

	for channel_data := range queue {
		fmt.Println("Element:", channel_data)
	}

}
