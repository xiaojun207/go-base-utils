package utils

import (
	"fmt"
	"testing"
)

func TestEncrypt(t *testing.T) {
	src := "this is a text"
	key := "23456322" // 必须是8位
	x1 := Encrypt(src, key)
	fmt.Println(x1)
	x2 := Decrypt(x1, key)
	fmt.Println(x2)
}
