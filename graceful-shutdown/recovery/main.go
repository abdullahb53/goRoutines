package main

import (
	"context"
	"fmt"
	"time"
)

const SUPERVISOR supervisor = "SUPERVISOR"

type region struct {
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
		humans: new([]human),
	}
}

func (r *region) addHuman(name string, age int) {
	a := newHuman(name, age)
	(*r.humans) = append((*r.humans), (*a))
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
		reg := ctx.Value(SUPERVISOR).(*region)
		reg.showHumans()
		time.Sleep(time.Second * 2)
		panic(">>")

	}
}

type supervisor string

func main() {

	reg := newRegion()
	ctx := context.WithValue(context.Background(), SUPERVISOR, reg)

	go Actor(ctx)

	go func(ctx context.Context) {
		reg := ctx.Value(SUPERVISOR).(*region)
		reg.addHuman("Anthony", 30)
		time.Sleep(time.Second * 3)
		reg.addHuman("Invoker", 33)
		time.Sleep(time.Second * 3)
		reg.addHuman("Dede", 55)
		time.Sleep(time.Second * 3)
		reg.addHuman("Kezban", 11)
	}(ctx)

	select {}
}
