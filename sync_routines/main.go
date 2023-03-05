package main

import (
	"fmt"
	"sync"
)

var GlobalTotal = 0

const NUMBER = 100

func Add(wg *sync.WaitGroup, mutex *sync.Mutex, num int) {

	mutex.Lock()
	GlobalTotal = GlobalTotal + num
	mutex.Unlock()
	wg.Done()
}

func main() {
	mutex := new(sync.Mutex)
	wg := new(sync.WaitGroup)

	wg.Add(NUMBER)

	for i := 0; i < 100; i++ {
		go Add(wg, mutex, i)
	}

	wg.Wait()

	fmt.Println("GlobalTotal:", GlobalTotal)

}
