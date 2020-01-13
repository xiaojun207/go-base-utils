package utils

import (
	"encoding/binary"
	"math"
	"math/rand"
	"strconv"
	"unsafe"
)

const (
	FLOAT64_MAX      = math.MaxFloat64
	UINT_MIN    uint = 0
	UINT_MAX         = ^uint(0)
	INT_MAX          = int(^uint(0) >> 1)
	INT_MIN          = ^INT_MAX
)

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

func Random(start float64, end float64) float64 {
	return start + rand.Float64()*(end-start)
}

func Int64ToStr(i int64) string {
	return strconv.FormatInt(i, 10)
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

func MaxInt64(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}
func MinInt64(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}
