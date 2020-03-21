package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"log"
)

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

/**
key :either 16, 24, or 32 bytes to select
*/
func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

/**
key :either 16, 24, or 32 bytes to select
*/
func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

/**
key :either 16, 24, or 32 bytes to select
*/
func AESEncrypt(src, key string) string {
	byteSrc := []byte(src)
	byteKey := []byte(key)

	xpass, err := AesEncrypt(byteSrc, byteKey)
	if err != nil {
		log.Fatalln(err)
	}
	return base64.StdEncoding.EncodeToString(xpass)
}

/**
key :either 16, 24, or 32 bytes to select
*/
func AESDecrypt(src, key string) string {
	byteSrc, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		log.Fatalln(err)
	}
	byteKey := []byte(key)
	res, _ := AesDecrypt(byteSrc, byteKey)
	return string(res)
}
