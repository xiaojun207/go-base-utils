package utils

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestAesDecrypt(t *testing.T) {
	var aeskey = []byte("321423u9y8d2fwfl")

	src := []byte("11123444")

	xpass, err := AesEncrypt(src, aeskey)
	if err != nil {
		fmt.Println(err)
		return
	}

	pass64 := base64.StdEncoding.EncodeToString(xpass)
	fmt.Printf("加密后:%v\n", pass64)

	bytesPass, err := base64.StdEncoding.DecodeString(pass64)
	if err != nil {
		fmt.Println(err)
		return
	}

	tpass, err := AesDecrypt(bytesPass, aeskey)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("解密后:%s\n", tpass)
}

func TestAESEncrypt(t *testing.T) {
	src := "11123444"
	key := "321423u9y8d2fwfl"
	x1 := AESEncrypt(src, key)
	fmt.Println(x1)

	x2 := AESDecrypt(x1, key)
	fmt.Println(x2)
}
