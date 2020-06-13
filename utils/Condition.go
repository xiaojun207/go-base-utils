package utils

func If(condition bool, trueValue, falseValue interface{}) interface{} {
	if condition {
		return trueValue
	} else {
		return falseValue
	}
}
