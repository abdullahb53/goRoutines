package main

import (
	"testing"

	"github.com/zeebo/assert"
)

func TestRingBitWise(t *testing.T) {

	// There is an incoming data.
	incomingBuf := make([]int, 16)
	fibonacciFill(&incomingBuf)

	// Our smart stack buffer. :size[8]
	ringBuff := newRingBuff()
	ringBuff.bufFill()

	// Put the incoming data into the ringBuffer.
	// Concurrent!
	go ringBuff.put(incomingBuf)

	// Read a stack buffer with the buffer size.
	res := ringBuff.read(len(incomingBuf))

	assert.Equal(t, len(incomingBuf), len(res))

	for i := 0; i < len(res); i++ {
		assert.Equal(t, res[i], incomingBuf[len(res)-i-1])
	}

}
