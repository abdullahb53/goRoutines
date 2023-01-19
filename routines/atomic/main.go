package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var k int = 5
var howManyTimesIterate int = 2000

func IterFunc(wg *sync.WaitGroup, ops *uint64) {
	for i := 0; i < howManyTimesIterate; i++ {
		atomic.AddUint64(ops, 1)
	}
	wg.Done()

}

func main() {

	var ops uint64
	var wg sync.WaitGroup
	wg.Add(5)
	for m := 0; m < k; m++ {

		go IterFunc(&wg, &ops)
	}
	wg.Wait()

	fmt.Println("Expected :", k*howManyTimesIterate)
	fmt.Println("We have  :", ops)

}
