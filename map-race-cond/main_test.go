package main

import (
	"testing"
)

func TestRace(t *testing.T) {

	newPidStorage := newPidStorage()
	newPidStorage.FillMap()

	for i := 3; i < 1000; i++ {

		go newPidStorage.handleNewMessage(i)
		go newPidStorage.handleGetMessage(i - 3)

	}

}
