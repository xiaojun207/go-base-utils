package utils

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	NUmStr  = "0123456789"
	CharStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	SpecStr = "+=-@#~,.[]()!%^*$"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Random(start float64, end float64) float64 {
	return start + rand.Float64()*(end-start)
}

// RandomPassword
// 随机生成密码
// 参数说明：
//  length: 生成密码的长度
//  charset:
//    num:只使用数字[0-9],
//    char:只使用英文字母[a-zA-Z],
//    mix:使用数字和字母，
//    advance:使用数字、字母以及特殊字符
// 如：RandomPassword(16, "mix")
//    ->  bJZSD6YpSAvDM6DY
//
func RandomPassword(passwordLength int, charset string) string {
	//初始化密码切片
	var passwd []byte = make([]byte, passwordLength, passwordLength)
	//源字符串
	var sourceStr string
	//判断字符类型,如果是数字
	if charset == "num" {
		sourceStr = NUmStr
		//如果选的是字符
	} else if charset == "char" {
		sourceStr = CharStr
		//如果选的是混合模式
	} else if charset == "mix" {
		sourceStr = fmt.Sprintf("%s%s", NUmStr, CharStr)
		//如果选的是高级模式
	} else if charset == "advance" {
		sourceStr = fmt.Sprintf("%s%s%s", NUmStr, CharStr, SpecStr)
	} else {
		sourceStr = NUmStr
	}
	//fmt.Println("source:", sourceStr)

	//遍历，生成一个随机index索引,
	for i := 0; i < passwordLength; i++ {
		index := rand.Intn(len(sourceStr))
		passwd[i] = sourceStr[index]
	}
	return string(passwd)
}
