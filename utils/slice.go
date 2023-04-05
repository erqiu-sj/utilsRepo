package utils

import (
	"reflect"
)

type SliceTypes interface {
	~bool | ~int | ~int16 | ~int8 | ~int64 | ~float32 | ~float64 | ~string | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | any
}

type Slice[T SliceTypes] []T

// Join 参考js Array.join
func (s Slice[T]) Join(connector string) string {
	str := ""
	for index, t := range s {
		if index == len(s)-1 {
			str += reflect.ValueOf(t).String()
			break
		}
		str += reflect.ValueOf(t).String() + connector
	}
	return str
}

type SlideMap[T SliceTypes, R any] func(T) R

// Map 参考js Array.map
func (s Slice[T]) Map(fn SlideMap[T, interface{}]) []interface{} {
	var result []interface{}
	for _, t := range s {
		result = append(result, fn(t))
	}
	return result
}

type SlideFilter[T SliceTypes] func(T) bool

// Filter 过滤 参考js Array.filter
func (s Slice[T]) Filter(fn SlideFilter[T]) Slice[T] {
	var result Slice[T]
	for _, t := range s {
		if fn(t) {
			result = append(result, t)
		}
	}
	return result
}

// Push 会修改原数组 参考js Array.push
func (s *Slice[T]) Push(ele T) Slice[T] {
	//result := make(Slice[T], len(s))
	//copy(result, s)
	*s = append(*s, ele)
	return *s
}

// ImmutablePush 不会修改原数组 参考js Array.push
func (s Slice[T]) ImmutablePush(ele T) Slice[T] {
	result := make(Slice[T], len(s))
	copy(result, s)
	result = append(result, ele)
	return result
}
