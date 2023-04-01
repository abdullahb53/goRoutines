package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRingBitWise(t *testing.T) {

	size := 64
	bufferMask := 63

	bufOne := newRingBuf(size)
	bufOne.bufFill()

	masked := 63 & bufferMask
	fmt.Printf("binary:%b normal:%v\n", masked, masked)
	fmt.Printf("~buf: %v\n", bufOne.buffer[masked])

	assert.Equal(t, masked, bufOne.buffer[63])

	masked = 64 & bufferMask
	fmt.Printf("binary:%b normal:%v\n", masked, masked)
	fmt.Printf("~buf: %v\n", bufOne.buffer[masked])

	assert.Equal(t, masked, bufOne.buffer[0])

}
