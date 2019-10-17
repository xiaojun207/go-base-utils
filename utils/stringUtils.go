package utils

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
