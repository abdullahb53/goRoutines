package main

func main() {
	var (
		a = 1 // moved to heap: a
		// b = true // moved to heap: b
		c = make(chan struct{})
	)

	// if const "b" be true; "a" escape to heap.
	// ... 			be false "a" doesn't escape.
	// if u pass "a" as a paramater to anony. it never
	// excape to heap.
	const b bool = false
	go func(a int) {
		if b {
			a++
		}
		close(c)
	}(a)
	<-c

	println(a, b) // 1 true
}
