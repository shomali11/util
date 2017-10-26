package xrunes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsMark(t *testing.T) {
	assert.False(t, IsMark([]rune(" ")[0]))
	assert.False(t, IsMark([]rune("a")[0]))
	assert.True(t, IsMark([]rune("\u0301")[0]))
}
