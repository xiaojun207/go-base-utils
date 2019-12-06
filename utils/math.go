package utils

import (
	"math"
	"math/rand"
	"strconv"
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
