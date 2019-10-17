package utils

import (
	"encoding/json"
	"fmt"
	"github.com/goinggo/mapstructure"
)

func MapToStruct(m interface{}, s interface{}) error {
	err := mapstructure.Decode(m, s)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func MapToJson(instance map[string]interface{}) string {
	mapInstances := []map[string]interface{}{}
	mapInstances = append(mapInstances, instance)

	jsonStr, err := json.Marshal(mapInstances)

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

func JsonToStruct(jsonStr string, result interface{}) {
	json.Unmarshal([]byte(jsonStr), &result)
}
