package main

import "fmt"

const QUEUE_SIZE = 10

type Queue struct {
	data chan string
}

func (q *Queue) Enqueue(item string) {

	if len(q.data) >= QUEUE_SIZE {
		_ = q.Dequeue()
		// fmt.Println("Deleted item: ", deleted_item)
	}

	q.data <- item
	fmt.Println("Coming from main: ", item)
}

func (q *Queue) Dequeue() string {
	return <-q.data
}

func main() {

	myQueue := &Queue{data: make(chan string, QUEUE_SIZE)}

	myQueue.Enqueue("H")
	myQueue.Enqueue("e")
	myQueue.Enqueue("l")
	myQueue.Enqueue("l")
	myQueue.Enqueue("o")
	myQueue.Enqueue("!")
	myQueue.Enqueue(" ")
	myQueue.Enqueue("M")
	myQueue.Enqueue("y")
	myQueue.Enqueue(" ")
	myQueue.Enqueue("f")
	myQueue.Enqueue("r")
	myQueue.Enqueue("i")
	myQueue.Enqueue("e")
	myQueue.Enqueue("n")
	myQueue.Enqueue("d")
	myQueue.Enqueue("s")
	myQueue.Enqueue(".")

	fmt.Print("\nremainder that in queue: |``|``|.. -> \"")

breakthere:
	for {
		select {
		case item, ok := <-myQueue.data:
			if !ok {
				break breakthere
			}
			fmt.Print(item)
		default:
			fmt.Print("\"")
			fmt.Println("\nEnd of the queue..")
			break breakthere
		}
	}

}
