package utils

import (
	"log"
	"reflect"
)

var MapType = reflect.TypeOf(map[string]interface{}{})
var ArrType = reflect.TypeOf([]interface{}{})

func CompareInterface(key string, a, b interface{}) (bool, interface{}, interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("PostContainers.err:", err)
		}
	}()
	vType := reflect.TypeOf(a)
	res := true
	if MapType == vType {
		bMap := b.(map[string]interface{})
		for k, v := range a.(map[string]interface{}) {
			r, _, _ := CompareInterface(key+"."+k, v, bMap[k])
			if !r {
				res = false
				//return res, ra, rb
			}
		}
	} else if ArrType == vType {
		arrA := a.([]interface{})
		arrB := b.([]interface{})
		num := len(arrA)
		if len(arrA) > len(arrB) {
			num = len(arrB)
			res = false
		}
		for i := 0; i < num; i++ {
			r, _, _ := CompareInterface(key+"["+Int64ToStr(int64(i))+"]", arrA[i], arrB[i])
			if !r {
				res = false
				//return res, ra, rb
			}
		}
	} else if a != b {
		log.Println("CompareInterface diff, key:{"+key+"} \t,value1:", a, "\t,value2:", b)
		return false, a, b
	}
	return res, a, b
}

func CompareInter(a, b interface{}) {
	var l []interface{}
	StructToMap(a, &l)
	var d []interface{}
	StructToMap(b, &d)
	CompareInterface("root", l, d)
}
