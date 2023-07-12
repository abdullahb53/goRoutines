package main

const N = 1000

// a[1000] = 0..
var a [N]int

//go:noinline
func g0(a *[N]int) {
	for i := range a {
		//line 12
		(*a)[i] = i
	}
}

//go:noinline
func g1(a *[N]int) {
	_ = *a
	for i := range a {
		// line 20
		(*a)[i] = i
	}
}

//go:noinline
func g2(x *[N]int) {
	a := x[:]
	for i := range a {
		// line 20
		a[i] = i
	}
}
