package main

import (
	"fmt"
	"sync"
)

var do sync.Once
var Number int

func init() {
	Number = 0
}

func JustDoIt(s1 int) {
	Number += s1
}

func main() {

	// We are guaranteed to work once in the 'Just Do It' function.
	do.Do(func() { go JustDoIt(53) })
	fmt.Println("Number:", Number)

	do.Do(func() { JustDoIt(22) })
	fmt.Println("Number:", Number)

	do.Do(func() { JustDoIt(11) })
	fmt.Println("Number:", Number)

}
