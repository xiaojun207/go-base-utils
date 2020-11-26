package utils

import (
	"math"
	"math/rand"
	"time"
)

const (
	FLOAT64_MAX      = math.MaxFloat64
	UINT_MIN    uint = 0
	UINT_MAX         = ^uint(0)
	INT_MAX          = int(^uint(0) >> 1)
	INT_MIN          = ^INT_MAX
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

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

func FloatFloor(f float64, prec int) float64 {
	l := math.Pow10(prec)
	return math.Floor(f*l) / l
}
