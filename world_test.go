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
}

func TestWorldRemoveEntity(t *testing.T) {
	If := lilt.NewIf(t)

	world := NewWorld()

	world.AddComponentType(mockComponentType, mockComponentFactory)

	e := world.NewEntity()

	cc := world.Components(mockComponentType)

	cc.Create(e)

	// If(err != nil).Errorf("Create() returned error")
	If(!cc.Has(e)).Errorf("Has() returned false")

	world.RemoveEntity(e)

	If(cc.Has(e)).Errorf("Entity was removed, but not the accompanying Component")
}
