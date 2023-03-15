package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}
	fmt.Printf("slice address: %p Slice own address: %p \n", &slice[0], &slice)

	slice = append(slice, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20)
	fmt.Printf("slice address: %p Slice own address: %p \n", &slice[0], &slice)
}
