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

	val := 2
	include := result.Keys().Include(val)
	assert.EqualValues(
		true,
		include, fmt.Sprint("keys 有误", include),
	)
}

func TestValues(t *testing.T) {
	assertions := assert.New(t)
	result := Map[int, int]{
		1: 1,
		2: 2,
		3: 444,
	}
	val := 444

	include := result.Values().Include(val)
	assertions.EqualValues(
		true,
		include,
		fmt.Sprint("values error", include),
	)
}

func TestHas(t *testing.T) {
	assertions := assert.New(t)
	result := Map[int, int]{
		1: 1,
		2: 2,
		3: 444,
	}

	assertions.EqualValues(false, result.Has(10), "has 10 不存在")
}

func TestGet(t *testing.T) {
	assertions := assert.New(t)
	result := Map[int, int]{
		1: 1,
		2: 2,
		3: 444,
	}
	value, _ := result.Get(1)
	assertions.EqualValues(1, value, fmt.Sprint("key 1 === ", value))
}

func TestAdd(t *testing.T) {
	assertions := assert.New(t)

	m := Map[string, int]{}
	get, _ := m.Add("200", 200).Add("300", 300).Get("200")

	assertions.EqualValues(200, get, fmt.Sprint("add value ====", get))
}

func TestDel(t *testing.T) {
	assertions := assert.New(t)

	m := Map[string, int]{}

	has := m.Add("200", 200).del("200").Has("200")

	assertions.EqualValues(false, has, fmt.Sprint("del value ====", has))
}
