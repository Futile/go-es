package es

import (
	"fmt"
	"reflect"
)

type World struct {
	nextId              entityId
	deletedEntities     []Entity
	componentContainers map[reflect.Type]*ComponentContainer
}

// NewEntity returns a new entity belonging to the given world
func (w *World) NewEntity() Entity {
	// check if there is a deleted entity available, if so increase reuse count and return it
	if len(w.deletedEntities) > 0 {
		lastIndex := len(w.deletedEntities) - 1

		e := w.deletedEntities[lastIndex]
		e.reuseCount++

		w.deletedEntities = w.deletedEntities[:lastIndex]

		return e
	}

	// panic if no more entity ids are available(probably much worse problem anyway..)
	if w.nextId == maxEntityId {
		panic("No new EntityId available.")
	}

	// increase next id of the world
	w.nextId++

	// return new entity with new id and zero reuse count
	return Entity{id: w.nextId - 1, reuseCount: 0}
}

// RemoveEntity removes a given entity from the world, and stores it for reuse
func (w *World) RemoveEntity(e Entity) error {
	for _, e2 := range w.deletedEntities {
		if e == e2 {
			return fmt.Errorf("Entity double-removed: id: %v", e.id)
		}
	}

	// if len == cap, just append the entity
	if len(w.deletedEntities) == cap(w.deletedEntities) {
		w.deletedEntities = append(w.deletedEntities, e)

		return nil
	}

	// else, grow the len of the slice and insert it at last position
	newIndex := len(w.deletedEntities) + 1

	w.deletedEntities = w.deletedEntities[:newIndex]
	w.deletedEntities[newIndex] = e

	return nil
}

// NewWorld returns a new world
func NewWorld() *World {
	return &World{nextId: minEntityId, deletedEntities: make([]Entity, 0), componentContainers: make(map[reflect.Type]*ComponentContainer)}
}
