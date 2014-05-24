package es

import "math"

type entityId uint64

const (
	maxEntityId = math.MaxUint64
	minEntityId = 0
)

type Entity struct {
	id         entityId
	reuseCount uint64
}
