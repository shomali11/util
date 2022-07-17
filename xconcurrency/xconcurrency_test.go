package xconcurrency

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestParallelize(t *testing.T) {
	value1 := 1
	value2 := 2

	fun1 := func() error {
		value1 = 11
		return nil
	}

	fun2 := func() error {
		value2 = 22
		return nil
	}

	err := Parallelize(fun1, fun2)
	assert.Nil(t, err)
	assert.Equal(t, value1, 11)
	assert.Equal(t, value2, 22)
}

func TestParallelizerLongTimeout(t *testing.T) {
	value1 := 1
	value2 := 2

	fun1 := func() error {
		value1 = 11
		return nil
	}

	fun2 := func() error {
		value2 = 22
		return nil
	}

	err := ParallelizeTimeout(time.Minute, fun1, fun2)
	assert.Nil(t, err)
	assert.Equal(t, value1, 11)
	assert.Equal(t, value2, 22)
}

func TestParallelizerShortTimeout(t *testing.T) {
	value1 := 1
	value2 := 2

	fun1 := func() error {
		time.Sleep(time.Minute)
		value1 = 11
		return nil
	}

	fun2 := func() error {
		time.Sleep(time.Minute)
		value2 = 22
		return nil
	}

	err := ParallelizeTimeout(time.Second, fun1, fun2)
	assert.NotNil(t, err)
	assert.Equal(t, value1, 1)
	assert.Equal(t, value2, 2)
}
