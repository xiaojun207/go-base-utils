package utils

import (
	"encoding/binary"
	"math"
	"strconv"
	"strings"
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

func SubStrStart(s, subStart string) string {
	a := strings.Index(s, subStart)
	return Substring(s, a+len(subStart), len(s))
}

func SubStrEnd(s, subEnd string) string {
	b := strings.Index(s, subEnd)
	return Substring(s, 0, b)
}

func SubStrBetween(s, subStart, subEnd string) string {
	a := strings.Index(s, subStart)
	b := strings.Index(s, subEnd)
	return Substring(s, a+len(subStart), b)
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
	return strconv.FormatFloat(float64(input_num), 'f', -1, 64)
}

func Float64ToString(input_num float64) string {
	return Float64ToStr(input_num, -1)
}

func Float64ToStr(input_num float64, prec int) string {
	return strconv.FormatFloat(input_num, 'f', prec, 64)
}

func Float64sToStrings(input_nums []float64) []string {
	var res []string
	for _, f := range input_nums {
		res = append(res, Float64ToString(f))
	}
	return res
}

func StrArr2ToFloat(arr *[][]interface{}) {
	for i, _ := range *arr {
		for j, _ := range (*arr)[i] {
			(*arr)[i][j] = StrToFloat64((*arr)[i][j].(string))
		}
	}
}

func StrArr2ToFloatArr2(arr [][]string) [][]float64 {
	res := [][]float64{}
	for i, _ := range arr {
		res = append(res, []float64{})
		for j, _ := range (arr)[i] {
			res[i] = append(res[i], StrToFloat64((arr)[i][j]))
		}
	}
	return res
}

func StrToFloat64(input_num string) float64 {
	value, _ := strconv.ParseFloat(input_num, 64)
	return value
}

func StrToFloat64Def(input_num string, defValue float64) float64 {
	if input_num == "" {
		return defValue
	}
	return StrToFloat64(input_num)
}

func StrToInt(input_num string) int {
	res, _ := strconv.Atoi(input_num)
	return res
}

func StrToIntDef(input_num string, defValue int) int {
	if input_num == "" {
		return defValue
	}
	return StrToInt(input_num)
}

func StrToBool(inputValue string) bool {
	res, _ := strconv.ParseBool(inputValue)
	return res
}

func StrToBoolDef(inputValue string, defValue bool) bool {
	if inputValue == "" {
		return defValue
	}
	return StrToBool(inputValue)
}

//string到int64
func StrToInt64(input_num string) int64 {
	i, _ := strconv.ParseInt(input_num, 10, 64)
	return i
}

func StrToInt64Def(inputValue string, defValue int64) int64 {
	if inputValue == "" {
		return defValue
	}
	return StrToInt64(inputValue)
}

//string到uint64
func StrToUint64(input_num string) uint64 {
	i, _ := strconv.ParseUint(input_num, 10, 64)
	return i
}

func StrToUint64Def(inputValue string, defValue uint64) uint64 {
	if inputValue == "" {
		return defValue
	}
	return StrToUint64(inputValue)
}

// int64到string
func Int64ToStr(num int64) string {
	return strconv.FormatInt(num, 10)
}

// uint64到string
func Uint64ToStr(num uint64) string {
	return strconv.FormatUint(num, 10)
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

func ByteToUint64(data []byte) uint64 {
	return binary.BigEndian.Uint64(data)
}

func Uint64ToByte(i uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, i)
	return b[:]
}

func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
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
