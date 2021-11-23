package utils

import (
	"fmt"
	"testing"
)

func TestRandomPassword(t *testing.T) {
	s := RandomPassword(16, "mix")
	fmt.Println(s)
}
