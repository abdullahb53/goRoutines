package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const SUPERVISOR supervisor = "SUPERVISOR"

type region struct {
	muHms  *sync.RWMutex
	humans *[]human
}

type human struct {
	name string
	age  int
}

func newHuman(name string, age int) *human {
	return &human{
		name: name,
		age:  age,
	}
}

func newRegion() *region {
	return &region{
		muHms:  &sync.RWMutex{},
		humans: &[]human{*newHuman("Abdullah", 27)},
	}
}

func (r *region) addHuman(name string, age int) {
	// This solution is temporary. I'll change this.
	// TODO: Solve race-condition without mutexes.
	// Channel maybe or another way.. Search it.
	r.muHms.Lock()
	a := newHuman(name, age)
	(*r.humans) = append((*r.humans), (*a))
	r.muHms.Unlock()
}

func (r *region) showHumans() {
	for i := range *r.humans {
		fmt.Println(i, "->", (*r.humans)[i].age, (*r.humans)[i].name)
	}
}

func Actor(ctx context.Context) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			go Actor(ctx)
		}
	}()

	for {
		// Casting.
		switch reg := ctx.Value(SUPERVISOR).(type) {
		case *region:
			reg.showHumans()
		default:
			_ = reg
		}

		time.Sleep(time.Second * 2)
		panic(">>")
	}
}

type supervisor string

func main() {

	// reg := newRegion()
	// ctx := context.WithValue(context.Background(), SUPERVISOR, reg)

	// go Actor(ctx)

	// go func(ctx context.Context) {
	// 	reg := ctx.Value(SUPERVISOR).(*region)
	// 	reg.addHuman("Anthony", 30)
	// 	time.Sleep(time.Second * 3)
	// 	reg.addHuman("Invoker", 33)
	// 	time.Sleep(time.Second * 3)
	// 	reg.addHuman("Dede", 55)
	// 	time.Sleep(time.Second * 3)
	// 	reg.addHuman("Kezban", 11)
	// }(ctx)

	// select {}
}
