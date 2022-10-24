package sort

import (
	"log"
	"testing"
)

func TestSort(t *testing.T) {
	ints := []int{0, 1, 23, 4, 5, 6, 3}
	Sort[int](ints)
	log.Println(ints)

	uints := []uint{0, 1, 23, 4, 5, 6, 3}
	Sort[uint](uints)
	log.Println(uints)

	uint64s := []uint64{0, 1, 23, 4, 5, 6, 3}
	Sort[uint64](uint64s)
	log.Println(uint64s)
}
