package utils

import (
	"fmt"
	"testing"
)

func TestSyncMap(t *testing.T) {
	testMap := SyncMap[string, string]{}
	testMap.Store("testKey", "testValue")

	value, _ := testMap.Load("testKey")
	fmt.Println("value:", value)
}
