package utils

import (
	"strconv"
	"strings"
)

// 获取source的子串,如果start小于0或者end大于source长度则返回""
// start:开始index，从0开始，包括0
// end:结束index，以end结束，但不包括end
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

func SubstrAfter(s, substr string) string {
	a := strings.Index(s, substr)
	return Substring(s, a+len(substr), len(s))
}

func SubstrBefore(s, substr string) string {
	b := strings.Index(s, substr)
	return Substring(s, 0, b)
}

func SubstrBetween(s, afterStr, beforeStr string) string {
	a := strings.Index(s, afterStr)
	b := strings.Index(s, beforeStr)
	return Substring(s, a+len(afterStr), b)
}

func ArrayContains(arr []string, s string) bool {
	for _, t := range arr {
		if t == s {
			return true
		}
	}
	return false
}

// float32 转 String工具类，保留6位小数
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

// string到int64
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

// string到uint64
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
