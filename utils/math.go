package utils

import (
	"encoding/binary"
	"math"
	"math/rand"
	"unsafe"
)

const (
	FLOAT64_MAX      = math.MaxFloat64
	UINT_MIN    uint = 0
	UINT_MAX         = ^uint(0)
	INT_MAX          = int(^uint(0) >> 1)
	INT_MIN          = ^INT_MAX
)

func Random(start float64, end float64) float64 {
	return start + rand.Float64()*(end-start)
}

func MaxInt64(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}
func MinInt64(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}
