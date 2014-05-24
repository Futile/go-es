package es

import "testing"

func TestWorld(t *testing.T) {
	world := NewWorld()

	e := world.NewEntity()
	e2 := world.NewEntity()

	if e.id == e2.id {
		t.Errorf("Two entities with the same id created!")
	}

	world.RemoveEntity(e2)

	err := world.RemoveEntity(e2)

	if err == nil {
		t.Errorf("Double-remove not noticed!")
	}

	e3 := world.NewEntity()

	if e3.id != e2.id {
		t.Errorf("Id was not reused!")
	}

	if e3.reuseCount == e2.reuseCount {
		t.Errorf("reuseCount is the same!")
	}
}
