package utils

import (
	"fmt"
	"reflect"
	"testing"
)

type Employee struct {
	Firstname string `json:"firstname"`
}

func TestNewInterface(t *testing.T) {
	data := "{\"firstName\": \"John\"}"
	obj := NewInterface(reflect.TypeOf(Employee{}), []byte(data))
	fmt.Println("Employee:", obj)
}
