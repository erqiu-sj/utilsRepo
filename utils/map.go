package utils

import (
	"errors"
	"fmt"
)

type MapKeys interface {
	~bool | ~int | ~int16 | ~int8 | ~int64 | ~float32 | ~float64 | ~string | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Map[T MapKeys, V any] map[T]V

// Keys 参考js Object.keys
func (receiver Map[T, V]) Keys() Slice[T] {
	keys := make(Slice[T], 0, len(receiver))
	for k := range receiver {
		keys.Push(k)
	}
	return keys
}

// Values 参考js Object.values
func (receiver Map[T, V]) Values() Slice[V] {
	keys := make(Slice[V], 0, len(receiver))
	for _, val := range receiver {
		keys.Push(val)
	}
	return keys
}

func (receiver Map[T, V]) Has(k T) bool {
	_, ok := receiver[k]
	return ok
}

func (receiver Map[T, V]) Get(k T) (V, error) {
	value, ok := receiver[k]
	if ok != false {
		return value, nil
	}
	return value, errors.New(fmt.Sprint(k, "non-existent"))
}
