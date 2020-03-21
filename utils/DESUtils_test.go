package utils

import (
	"fmt"
	"testing"
)

func TestEncrypt(t *testing.T) {
	src := "this is a test text"
	key := "12345678"
	x1 := Encrypt(src, key)
	x2 := Decrypt(x1, key)
	fmt.Print(string(x2))
}
