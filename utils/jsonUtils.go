package utils

import (
	"encoding/json"
	"fmt"
	"github.com/goinggo/mapstructure"
)

// 参数s必须是指针
func MapToStruct(m interface{}, s interface{}) error {
	jsonBytes, err := json.Marshal(m)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonBytes, s)
	return err
}

func MapToStruct2(m interface{}, s interface{}) error {
	err := mapstructure.Decode(m, s)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// 参数m必须是指针
func StructToMap(s interface{}, m interface{}) error {
	j, _ := json.Marshal(s)
	return json.Unmarshal(j, &m)
}

func MapToJson(instance interface{}) string {
	jsonStr, err := json.Marshal(instance)
	if err != nil {
		fmt.Println("MapToJson err: ", err)
	}
	return string(jsonStr)
}

func JsonToMap(jsonStr string, mapResult map[string]interface{}) {
	err := json.Unmarshal([]byte(jsonStr), &mapResult)
	if err != nil {
		fmt.Println("JsonToMap err: ", err)
	}
	fmt.Println(mapResult)
}

func StructToJson(result interface{}) string {
	jsonBytes, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}
	return string(jsonBytes)
}

// 参数result必须是指针
func JsonToStruct(jsonStr string, result interface{}) {
	json.Unmarshal([]byte(jsonStr), result)
}

// 参数result必须是指针
func JsonBytesToStruct(jsonByte []byte, result interface{}) {
	json.Unmarshal(jsonByte, result)
}
