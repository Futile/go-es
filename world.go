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
	// check if there is a deleted entity available, if so and return it
	if len(w.deletedEntities) > 0 {
		lastIndex := len(w.deletedEntities) - 1

		e := w.deletedEntities[lastIndex]

		w.deletedEntities = w.deletedEntities[:lastIndex]

		return e
	}

	// panic if no more entity ids are available(probably much worse problems anyway..)
	if w.nextId == maxEntityId {
		panic("No new EntityId available.")
	}

	// increase next id of the world
	w.nextId++

	// return new entity with new id and zero reuse count
	return Entity{id: w.nextId - 1}
}

// RemoveEntity removes a given entity from the world, and stores it for reuse
func (w *World) RemoveEntity(e Entity) error {
	for _, e2 := range w.deletedEntities {
		if e == e2 {
			return fmt.Errorf("Entity double-removed: id: %v", e.id)
		}
	}

	w.deletedEntities = append(w.deletedEntities, e)

	for _, cc := range w.componentContainers {
		cc.Remove(e)
	}

	return nil
}

func (w *World) Components(componentType reflect.Type) *ComponentContainer {
	return w.componentContainers[componentType]
}

func (w *World) AddComponentType(componentType reflect.Type, componentFactory ComponentFactory) error {
	if w.Components(componentType) != nil {
		return fmt.Errorf("Component type '%v' is already registered for this world!", componentType)
	}

	w.componentContainers[componentType] = newComponentContainer(componentFactory)

	return nil
}

// NewWorld returns a new world
func NewWorld() *World {
	return &World{nextId: minEntityId, deletedEntities: make([]Entity, 0), componentContainers: make(map[reflect.Type]*ComponentContainer)}
}
