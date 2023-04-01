package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRingBuff(t *testing.T) {
	size := 4

	buf := newRingBuffer(size)

	for i := 0; i < size; i++ {
		buf.Add(newHuman("name", i))
	}

	for i := 0; i < size-2; i++ {
		a := buf.Get()
		println(a.age)
	}

	// x1x, @2@,  3,  4
	// x1x, x2x, @3@, 4 == (size'4' - 1)
	assert.Equal(t, size-1, buf.Get().age)
}
