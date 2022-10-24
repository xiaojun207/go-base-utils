package sort

import "sort"

// IntSlice attaches the methods of Interface to []T byte | int | int8 | int16 | int32 | int64 | uint | uintptr | uint16 | uint32 | uint64 | float32 | float64
// , sorting in increasing order.
type NumberSlice[T byte | int | int8 | int16 | int32 | int64 | uint | uintptr | uint16 | uint32 | uint64 | float32 | float64] []T

func (x NumberSlice[T]) Len() int           { return len(x) }
func (x NumberSlice[T]) Less(i, j int) bool { return x[i] < x[j] }
func (x NumberSlice[T]) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// Sort is a convenience method: x.Sort() calls Sort(x).
func (x NumberSlice[T]) Sort() { Sort(x) }

func Sort[T byte | int | int8 | int16 | int32 | int64 | uint | uintptr | uint16 | uint32 | uint64 | float32 | float64](arr []T) {
	sort.Sort(NumberSlice[T](arr))
}
