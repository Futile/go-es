package es

import "fmt"

type Component interface {
	reuseCount() uint64
	setReuseCount(count uint64)
}

type BaseComponent struct {
	reuseCounter uint64
}

func (b *BaseComponent) reuseCount() uint64 {
	return b.reuseCounter
}

func (b *BaseComponent) setReuseCount(count uint64) {
	b.reuseCounter = count
}

type ComponentFactory func() Component

type ComponentContainer struct {
	components       map[entityId]Component
	componentFactory ComponentFactory
}

func NewComponentContainer(componentFactory ComponentFactory) *ComponentContainer {
	return &ComponentContainer{components: make(map[entityId]Component), componentFactory: componentFactory}
}

// Get returns the component belonging to an entity from this container
func (cc *ComponentContainer) Get(e Entity) Component {
	c := cc.components[e.id]

	// check if there is no component, or the reuse component is wrong
	if c == nil || c.reuseCount() != e.reuseCount {
		return nil
	}

	return c
}

// Has returns whether the given entity has a component in this container
func (cc *ComponentContainer) Has(e Entity) bool {
	return cc.Get(e) != nil
}

// Create creates a new component for the entity.
// It panics if there is already an existing, valid Component for the given entity
func (cc *ComponentContainer) Create(e Entity) (Component, error) {
	// check if the entity already has a component
	if cc.Has(e) {
		return nil, fmt.Errorf("Trying to create component for entity which already has a component.")
	}

	c := cc.componentFactory()
	c.setReuseCount(e.reuseCount)

	// save the new component in the map
	cc.components[e.id] = c

	return c, nil
}

// GetOrCreate returns the component belonging to the entity, creating it if necessary
func (cc *ComponentContainer) GetOrCreate(e Entity) Component {
	c := cc.Get(e)

	if c == nil {
		c, _ = cc.Create(e)
	}

	return c
}

func (cc *ComponentContainer) Remove(e Entity) error {
	if !cc.Has(e) {
		return fmt.Errorf("Trying to remove Component even though it did not exist for entity!")
	}

	delete(cc.components, e.id)

	return nil
}
