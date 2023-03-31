package main

import (
	"fmt"
	"time"

	"github.com/anthdm/hollywood/actor"
)

type storageType struct {
	storage map[*int]*int
}

type setItem struct {
	key *int
	val *int
}

type showItems struct {
}

type showItem struct {
	key *int
}

type human struct {
}

type bobDoSomething struct {
}

// Human receiver.
func (h *human) Receive(ctx *actor.Context) {
	switch msg := ctx.Message().(type) {
	case *bobDoSomething:
		ctx.Engine().Send(storageProcessId, &showItems{})
	default:
		_ = msg
	}
}

func (s *storageType) Receive(ctx *actor.Context) {
	switch msg := ctx.Message().(type) {
	case actor.Started:
		fmt.Println("[STORAGE] 000 i'm spawned:", ctx.PID().String())

		// I want to get and show storage item.
		fmt.Println("doesn't matter get item..:", msg)

	// Add item to storage.
	case *setItem:
		s.storage[msg.key] = msg.val

	case *showItems:
		fmt.Println("[SHOW-ITEMS] Executing..")
		for key, val := range s.storage {
			fmt.Println("key:", key, "--> val", val)
		}

	case *showItem:
		fmt.Println("[SHOW-ITEM-ONE] Executing..")
		fmt.Println("key:", msg.key, "val:", s.storage[msg.key])

	case actor.Stopped:
		fmt.Println("[STORAGE] XXX i'm dead...:", ctx.PID().String())
	default:
		_ = msg

	}

}

func fillStorage(n int) {

	for i := 0; i < n; i++ {
		val := i + 13
		engine.Send(storageProcessId, &setItem{
			key: new(int),
			val: &val,
		})
	}
}

func showAllItems() {
	engine.Send(storageProcessId, &showItems{})
}

func newStorage() actor.Receiver {
	return &storageType{
		storage: map[*int]*int{},
	}

}

func newHuman() actor.Receiver {
	return &human{}
}

var (
	engine           *actor.Engine
	storageProcessId *actor.PID
)

func main() {
	engine = actor.NewEngine()
	storageProcessId = engine.Spawn(newStorage, "my_name_is_storage!")
	fillStorage(5)

	showAllItems()

	// Kill the storage.
	engine.Poison(storageProcessId)

	// time.Sleep(time.Second * 3)

	showAllItems()

	// Create a new human.
	humanOnePid := engine.Spawn(newHuman, "humanOne")
	engine.Send(humanOnePid, &bobDoSomething{})

	time.Sleep(time.Second * 888)

}
