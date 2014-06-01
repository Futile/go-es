package es

import "time"

type System interface {
	Step(world *World, delta time.Duration) []func()
}

type SystemFunc func(world *World, delta time.Duration) []func()

func (s SystemFunc) Step(world *World, delta time.Duration) []func() {
	return s(world, delta)
}
