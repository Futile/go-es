package es

import (
	"testing"

	"github.com/futile/go-lil-t"
)

func TestWorld(t *testing.T) {
	If := lilt.NewIf(t)

	world := NewWorld()

	e := world.NewEntity()
	e2 := world.NewEntity()
	If(e.id == e2.id).Errorf("Two entities with the same id created!")

	world.RemoveEntity(e2)
	err := world.RemoveEntity(e2)
	If(err == nil).Errorf("Double-remove not noticed!")

	e3 := world.NewEntity()
	If(e3.id != e2.id).Errorf("Id was not reused!")
	If(e3.reuseCount == e2.reuseCount).Errorf("reuseCount is the same!")
}
