package syncMap

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

type Itest interface {
	add(a, b int) int
}
type Stest struct {
}

type Stest1 struct {
	Stest
}

func (s *Stest1) add(a, b int) int {
	return a + b
}

type Stest2 struct {
	Stest
}

func (s *Stest2) add(a, b int) int {
	return a + b
}

func TestSyncMap2(t *testing.T) {

	testMap := SyncMap[string, *Stest]{}
	//testMap.Store("testKey", Stest1{})

	value, _ := testMap.Load("testKey")
	fmt.Println("value:", value)
}
