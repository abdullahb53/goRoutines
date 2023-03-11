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
	fmt.Println("-> ", item)
}

func (q *Queue) Dequeue() string {
	return <-q.data
}

func main() {

	myQueue := &Queue{data: make(chan string, QUEUE_SIZE)}

	msg := "Hello My friends!"
	// H,e,l,l,o, ,M,y,...
	for _, val := range msg {
		myQueue.Enqueue(string(val))
	}

breakthere:
	for {
		select {
		case item, ok := <-myQueue.data:
			if !ok {
				break breakthere
			}
			fmt.Print(item)
		default:
			fmt.Println("\nEnd of the queue..")
			break breakthere
		}
	}

}
