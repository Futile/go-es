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

	_, err := cc.Create(e)

	If(err != nil).Errorf("Create() returned an error")
	If(!cc.Has(e)).Errorf("Has() returned false")

	world.RemoveEntity(e)

	If(cc.Has(e)).Errorf("Entity was removed, but not the accompanying Component")
}

func TestEntitiesWith(t *testing.T) {
	world := NewWorld()

	world.AddComponentType(mockComponentType, mockComponentFactory)
	world.AddComponentType(mockComponentType2, mockComponentFactory2)
	world.AddComponentType(mockComponentType3, mockComponentFactory3)

	cc := world.Components(mockComponentType)
	cc2 := world.Components(mockComponentType2)
	cc3 := world.Components(mockComponentType3)

	e := world.NewEntity()

	cc.Create(e)
	cc2.Create(e)
	cc3.Create(e)

	success := false
	for en := range world.EntitiesWith(mockComponentType, mockComponentType2, mockComponentType3) {
		if en != e {
			t.Errorf("wrong entity returned")
		}

		success = true
	}

	if !success {
		t.Errorf("EntitiesWith did not return a necessary entity.")
	}

	cc3.Remove(e)

	for _ = range world.EntitiesWith(mockComponentType, mockComponentType2, mockComponentType3) {
		t.Errorf("Found an entity even though none was supposed to be found")
	}

}
