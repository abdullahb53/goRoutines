package main

import (
	"testing"

	"github.com/anthdm/hollywood/actor"
)

func TestActorRestart(t *testing.T) {

	engine = actor.NewEngine()
	storageProcessId = engine.Spawn(newStorage, "my_name_is_storage!")
	fillStorage(5)

	showAllItems()
	// Kill the storage.
	engine.Poison(storageProcessId)
	showAllItems()

	// Create a new human.
	humanOnePid := engine.Spawn(newHuman, "humanOne")
	engine.Send(humanOnePid, &bobDoSomething{})

}
