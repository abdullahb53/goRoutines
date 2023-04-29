package main

import "fmt"

func bridge(done <-chan interface{}, chanStream <-chan <-chan interface{}) <-chan interface{} {
	valStream := make(chan interface{})
}

func genVals() {

}

func main() {
	for v := range bridge(nil, genVals()) {
		fmt.Printf("%v", v)
	}
}
