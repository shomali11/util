package xmanipulations

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func TestShuffles(t *testing.T) {
	source := rand.NewSource(time.Now().UnixNano())

	emptyArray := []interface{}{}
	Shuffle(emptyArray, source)
	assert.Equal(t, emptyArray, []interface{}{})

	oneElementArray := []interface{}{11}
	Shuffle(oneElementArray, source)
	assert.Equal(t, oneElementArray, []interface{}{11})

	array := []interface{}{"a", "b", "c"}
	Shuffle(array, source)
	assert.Contains(t, array, "a")
	assert.Contains(t, array, "b")
	assert.Contains(t, array, "c")
}
