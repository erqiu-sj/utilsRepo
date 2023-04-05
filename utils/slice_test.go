package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJoin(t *testing.T) {
	assert := assert.New(t)
	got := Slice[string]{"1", "2", "3"}
	assert.EqualValues(got.Join(","), "1,2,3", "slide join fn result is error!")
}

type gotMap struct {
	types string
}

func TestMap(t *testing.T) {
	assert := assert.New(t)
	got := Slice[string]{"1", "2", "3"}
	var gotsMap []interface{}
	gotsMap = got.Map(func(s string) interface{} {
		return gotMap{types: s}
	})

	firstTypes := gotsMap[0].(gotMap).types

	assert.EqualValues(
		firstTypes, "1", "firstTypes 不等于 1",
	)
}

func TestFilter(t *testing.T) {
	assert := assert.New(t)
	got := Slice[int]{1, 1, 3, 4}
	filterResult := got.Filter(func(i int) bool {
		return i > 2
	})
	assert.EqualValues(
		filterResult, Slice[int]{3, 4}, "filterResult 不等于 [3,4]",
	)
}

func TestPush(t *testing.T) {
	assert := assert.New(t)

	got := Slice[int]{1, 1, 2}
	got.Push(4)
	assert.EqualValues(
		got, Slice[int]{1, 1, 2, 4}, "Push 没有修改到原数组",
	)

	ImmutablePush := got.ImmutablePush(9)
	assert.EqualValues(ImmutablePush, Slice[int]{1, 1, 2, 4, 9}, "ImmutablePush 没有返回一个新数组")
}
