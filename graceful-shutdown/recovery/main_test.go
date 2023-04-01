package main

import "testing"

func TestRace(t *testing.T) {
	reg := newRegion()

	for i := 0; i < 1000; i++ {
		go reg.addHuman("aaa", 0)
	}
}
