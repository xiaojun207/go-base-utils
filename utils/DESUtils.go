package utils

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"log"
)

func padding(src []byte, blocksize int) []byte {
	n := len(src)
	padnum := blocksize - n%blocksize
	pad := bytes.Repeat([]byte{byte(padnum)}, padnum)
	dst := append(src, pad...)
	return dst
}

func unpadding(src []byte) []byte {
	n := len(src)
	unpadnum := int(src[n-1])
	dst := src[:n-unpadnum]
	return dst
}

func EncryptDES(src []byte, key []byte) []byte {
	block, _ := des.NewCipher(key)
	src = padding(src, block.BlockSize())
	blockmode := cipher.NewCBCEncrypter(block, key)
	blockmode.CryptBlocks(src, src)
	return src
}

func DecryptDES(src []byte, key []byte) []byte {
	block, _ := des.NewCipher(key)
	blockmode := cipher.NewCBCDecrypter(block, key)
	blockmode.CryptBlocks(src, src)
	src = unpadding(src)
	return src
}

func Encrypt(src, key string) string {
	byteSrc := []byte(src)
	byteKey := []byte(key)
	res := EncryptDES(byteSrc, byteKey)
	return base64.StdEncoding.EncodeToString(res)
}

func Decrypt(src, key string) string {
	byteSrc, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		log.Fatalln(err)
	}
	byteKey := []byte(key)
	res := DecryptDES(byteSrc, byteKey)
	return string(res)
}
