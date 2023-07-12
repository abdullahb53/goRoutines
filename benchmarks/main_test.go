package main

import "testing"

func Benchmark_g0(b *testing.B) {
	for i := 0; i < b.N; i++ { // 339.1 ns/op
		g0(&a)
	}
}

func Benchmark_g1(b *testing.B) {
	for i := 0; i < b.N; i++ { // 295.2 ns/op
		g1(&a)
	}
}

func Benchmark_g2(b *testing.B) {
	for i := 0; i < b.N; i++ { // 292.1 ns/op
		g2(&a)
	}
}
