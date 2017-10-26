package xconcurrency

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestParallelize(t *testing.T) {
	value1 := 1
	value2 := 2

	fun1 := func() {
		value1 = 11
	}

	fun2 := func() {
		value2 = 22
	}

	Parallelize(fun1, fun2)

	assert.Equal(t, value1, 11)
	assert.Equal(t, value2, 22)
}

func TestParallelizerLongTimeout(t *testing.T) {
	value1 := 1
	value2 := 2

	fun1 := func() {
		value1 = 11
	}

	fun2 := func() {
		value2 = 22
	}

	ParallelizeTimeout(time.Minute, fun1, fun2)

	assert.Equal(t, value1, 11)
	assert.Equal(t, value2, 22)
}

func TestParallelizerShortTimeout(t *testing.T) {
	value1 := 1
	value2 := 2

	fun1 := func() {
		time.Sleep(time.Minute)
		value1 = 11
	}

	fun2 := func() {
		time.Sleep(time.Minute)
		value2 = 22
	}

	ParallelizeTimeout(time.Second, fun1, fun2)

	assert.Equal(t, value1, 1)
	assert.Equal(t, value2, 2)
}
