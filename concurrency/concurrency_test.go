package concurrency

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
