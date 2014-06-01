package es

import (
	"testing"
	"time"
)

type nilSystem struct {
}

func (s *nilSystem) Step(world *World, delta time.Duration) []func() {
	return nil
}

func nilSystemFunc(world *World, delta time.Duration) []func() {
	return nil
}

func TestSystemFunc(t *testing.T) {
	// make sure that a SystemFunc is also a System
	var _ System = SystemFunc(nilSystemFunc)
}
