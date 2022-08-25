package utils

import (
	"sync"
)

type SyncMap[K comparable, V comparable] struct {
	sync.Map
}

func (e *SyncMap[K, V]) ToMap() map[K]V {
	res := map[K]V{}
	e.ForEach(func(key K, value V) {
		res[key] = value
	})
	return res
}

func (e *SyncMap[K, V]) ToStrMap() map[K]V {
	res := map[K]V{}
	e.ForEach(func(key K, value V) {
		res[key] = value
	})
	return res
}

func (e *SyncMap[K, V]) Keys() []K {
	var res []K
	e.ForEach(func(key K, _ V) {
		res = append(res, key)
	})
	return res
}

func (e *SyncMap[K, V]) Values() []V {
	res := []V{}
	e.ForEach(func(_ K, value V) {
		res = append(res, value)
	})
	return res
}

func (e *SyncMap[K, V]) ToArr(f func(key K, value V) V) []V {
	res := []V{}
	e.Range(func(key, value interface{}) bool {
		res = append(res, f(key.(K), value.(V)))
		return true
	})
	return res
}

func (e *SyncMap[K, V]) ForEach(f func(key K, value V)) {
	e.Range(func(key, value interface{}) bool {
		f(key.(K), value.(V))
		return true
	})
}

func (e *SyncMap[K, V]) Store(key K, value V) {
	e.Map.Store(key, value)
}

func (e *SyncMap[K, V]) Load(key K) (value V, ok bool) {
	val, ok := e.Map.Load(key)
	value = val.(V)
	return
}

func (e *SyncMap[K, V]) LoadInit(key K, Init func() V) (V, bool) {
	val, ok := e.Load(key)
	if ok {
		return val, ok
	} else {
		value := Init()
		e.Store(key, value)
		return value, ok
	}
}

func (e *SyncMap[K, V]) Size() int {
	count := 0
	e.Range(func(key, value interface{}) bool {
		count++
		return true
	})
	return count
}

func (e *SyncMap[K, V]) ContainKey(key K) (ok bool) {
	_, ok = e.Load(key)
	return
}

func (e *SyncMap[K, V]) Remove(key string) {
	e.Delete(key)
}

func (e *SyncMap[K, V]) RemoveAll() {
	e.ForEach(func(key K, value V) {
		e.Delete(key)
	})
}
