package main

import (
	"fmt"
	"time"

	"github.com/anthdm/hollywood/actor"
)

type message struct {
	data string
}

type foo struct {
}

func newFoo() actor.Receiver {
	return &foo{}
}

func (f *foo) Receive(ctx *actor.Context) {
	switch msg := ctx.Message().(type) {
	case actor.Started:
		fmt.Println("foo has started")
	case *message:
		fmt.Println("foo has received", msg.data)
	}
}

func main() {
	engine := actor.NewEngine()
	pid := engine.Spawn(newFoo, "foo")
	engine.Send(pid, &message{
		data: "hello worldd!!",
	})
	time.Sleep(time.Second * 1000)
}
