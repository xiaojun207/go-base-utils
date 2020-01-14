package utils

import (
	"encoding/binary"
	"math"
	"strconv"
	"unsafe"
)

//获取source的子串,如果start小于0或者end大于source长度则返回""
//start:开始index，从0开始，包括0
//end:结束index，以end结束，但不包括end
func Substring(source string, start int, end int) string {
	var r = []rune(source)
	length := len(r)

	if start < 0 || start > end {
		return ""
	}

	if end > length {
		end = length
	}

	if start == 0 && end >= length {
		return source
	}

	return string(r[start:end])
}

func ArrayContains(arr []string, s string) bool {
	for _, t := range arr {
		if t == s {
			return true
		}
	}
	return false
}

//float32 转 String工具类，保留6位小数
func Float32ToString(input_num float32) string {
	// to convert a float number to a string
	return strconv.FormatFloat(float64(input_num), 'f', 6, 64)
}

func Float64ToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}

func Float64sToStrings(input_nums []float64) []string {
	var res []string
	for _, f := range input_nums {
		res = append(res, Float64ToString(f))
	}
	return res
}

func StrToFloat64(input_num string) float64 {
	value, _ := strconv.ParseFloat(input_num, 64)
	return value
}

func StrToInt(input_num string) int {
	res, _ := strconv.Atoi(input_num)
	return res
}

func StrToBool(inputValue string) bool {
	res, _ := strconv.ParseBool(inputValue)
	return res
}

//string到int64
func StrToInt64(input_num string) int64 {
	i, _ := strconv.ParseInt(input_num, 10, 64)
	return i
}

// int64到string
func Int64ToStr(num int64) string {
	return strconv.FormatInt(num, 10)
}

func Int2Byte(data int) (ret []byte) {
	var len uintptr = unsafe.Sizeof(data)
	ret = make([]byte, len)
	var tmp int = 0xff
	var index uint = 0
	for index = 0; index < uint(len); index++ {
		ret[index] = byte((tmp << (index * 8) & data) >> (index * 8))
	}
	return ret
}

func Byte2Int(data []byte) int {
	var ret int = 0
	var len int = len(data)
	var i uint = 0
	for i = 0; i < uint(len); i++ {
		ret = ret | (int(data[i]) << (i * 8))
	}
	return ret
}

func Float32ToByte(float float32) []byte {
	bits := math.Float32bits(float)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits)

	return bytes
}

func ByteToFloat32(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	return math.Float32frombits(bits)
}

func Float64ToByte(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)
	return bytes
}

func ByteToFloat64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	return math.Float64frombits(bits)
}
