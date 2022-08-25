package utils

import (
	"fmt"
	"reflect"
)

func GetTag(u interface{}, key string) string {
	t := reflect.TypeOf(u)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get(key)
		fmt.Println("get tag is ", tag)
		return tag
	}
	return ""
}
