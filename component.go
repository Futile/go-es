package es

import (
	"fmt"
	"reflect"
)

type Component struct {
	reuseCount uint64
}

type ComponentContainer struct {
	components    map[entityId]*Component
	componentType reflect.Type
}

// Get returns the component belonging to an entity from this container
func (cc *ComponentContainer) Get(e *Entity) *Component {
	c := cc.components[e.id]

	// check if there is no component, or the reuse component is wrong
	if c == nil || c.reuseCount != e.reuseCount {
		return nil
	}

	// assert the type
	if reflect.TypeOf(c) != cc.componentType {
		panic(fmt.Errorf("ComponentContainer contains wrong component type: expected: %v, but got: %T\n", cc.componentType, c))
	}

	return c
}

// Has returns whether the given entity has a component in this container
func (cc *ComponentContainer) Has(e *Entity) bool {
	return cc.Get(e) != nil
}

// Create creates a new component for the entity.
// It panics if there is already an existing, valid Component for the given entity
func (cc *ComponentContainer) Create(e *Entity) *Component {
	// check if the entity already has a component
	if cc.Has(e) {
		panic("Trying to create component for entity which already has a component.")
	}

	// create a new component of the correct type for the entity, and set the correct reuseCount
	// crazy reflection-shit: https://stackoverflow.com/questions/10210188/instance-new-type-golang
	c := reflect.New(cc.componentType).Elem().Interface().(Component)
	c.reuseCount = e.reuseCount

	// save the new component in the map
	cc.components[e.id] = &c

	return &c
}

// GetOrCreate returns the component belonging to the entity, creating it if necessary
func (cc *ComponentContainer) GetOrCreate(e *Entity) *Component {
	c := cc.Get(e)

	if c == nil {
		c = cc.Create(e)
	}

	return c
}
