package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

//md5加密
func Md5(s string) string {
	m := md5.New()
	m.Write([]byte(s))
	res := hex.EncodeToString(m.Sum(nil))
	return res
}

//Sha256加密
func Sha256(s string) string {
	m := sha256.New()
	m.Write([]byte(s))
	res := hex.EncodeToString(m.Sum(nil))
	return res
}
