package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKeys(t *testing.T) {
	assert := assert.New(t)
	result := Map[int, int]{
		1: 1,
		2: 2,
		3: 444,
	}
	assert.EqualValues(
		result.Keys(), Slice[int]{1, 2, 3}, "keys 有误",
	)
}

func TestValues(t *testing.T) {
	assert := assert.New(t)
	result := Map[int, int]{
		1: 1,
		2: 2,
		3: 444,
	}
	assert.EqualValues(
		result.Values(),
		Slice[int]{1, 2, 444},
		"values error",
	)
}

func TestHas(t *testing.T) {
	assert := assert.New(t)
	result := Map[int, int]{
		1: 1,
		2: 2,
		3: 444,
	}
	assert.EqualValues(false, result.Has(10), "has 10 不存在")
}

func TestGet(t *testing.T) {
	assert := assert.New(t)
	result := Map[int, int]{
		1: 1,
		2: 2,
		3: 444,
	}
	value, _ := result.Get(1)
	assert.EqualValues(1, value, fmt.Sprint("key 1 === ", value))
}
