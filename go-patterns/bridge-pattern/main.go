package main

// import "fmt"

// func bridge(done <-chan interface{}, chanStream <-chan <-chan interface{}) <-chan interface{} {
// 	valStream := make(chan interface{})

// 	go func() {
// 		defer close(valStream)

// 		for {
// 			var stream <-chan interface{}
// 			select {
// 			case maybeStream, ok := <-chanStream:
// 			}
// 		}

// 	}()

// }

// func genVals() {

// }

// func main() {
// 	for v := range bridge(nil, genVals()) {
// 		fmt.Printf("%v", v)
// 	}
// }
