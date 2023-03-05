package main

import (
	"context"
	"fmt"
	"time"
)

type ContextPool struct {
	Type string
	Port string
}

func BobDoSomething(ctx context.Context) error {

	for {
		fmt.Println("323!")
		// do something
		select {
		case <-ctx.Done(): // closes when the caller cancels the ctx
			fmt.Println("Ctx Done.!")
			return ctx.Err() // has a value on context cancellation
		default:

			fmt.Println("bob!")

		}
	}

}
func AsheDoSomething(ctx context.Context) error {

	for {

		select {
		case <-ctx.Done():
			fmt.Println("Ctx Done.! ASHE")
			return ctx.Err()
		default:
			fmt.Println("ashe!")
		}

	}

}

func main() {

	ctx1, cancel_ctx1 := context.WithCancel(context.Background())
	fmt.Println("cancel_ctx1:", cancel_ctx1, " ctx1:", ctx1)

	ctx2, cancel_ctx2 := context.WithCancel(context.Background())
	fmt.Println("cancel_ctx2:", cancel_ctx1, " ctx2:", ctx2)

	time.Sleep(3 * time.Second)

	go BobDoSomething(ctx1)
	go AsheDoSomething(ctx2)

	time.Sleep(3 * time.Second)
	cancel_ctx1()
	time.Sleep(2 * time.Second)
	cancel_ctx2()
	time.Sleep(12 * time.Second)

}
