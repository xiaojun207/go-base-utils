package utils

import (
	"fmt"
	"testing"
)

func TestFloat64ToByte(t *testing.T) {
	f := 1099222222222.323
	b := Float64ToByte(f)
	fmt.Println(len(b), b)
}

func TestInt2Byte(t *testing.T) {
	f := 11122222222222
	b := Int2Byte(f)
	fmt.Println(len(b), b)
}
