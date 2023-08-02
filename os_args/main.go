package main

import "fmt"

func do(b []int) {
	b[0] = 89

}

func main() {
	var b = []int{4, 5, 6}
	fmt.Println(b)
	do(b)
	fmt.Println(b)
}
