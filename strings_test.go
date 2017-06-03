package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsEmpty(t *testing.T) {
	assert.True(t, IsEmpty(""))
	assert.False(t, IsEmpty("text"))
	assert.False(t, IsEmpty("	"))
}

func TestIsNotEmpty(t *testing.T) {
	assert.False(t, IsNotEmpty(""))
	assert.True(t, IsNotEmpty("text"))
	assert.True(t, IsNotEmpty("	"))
}

func TestIsBlank(t *testing.T) {
	assert.True(t, IsBlank(""))
	assert.True(t, IsBlank("	"))
	assert.False(t, IsBlank("text"))
}

func TestIsNotBlank(t *testing.T) {
	assert.False(t, IsNotBlank(""))
	assert.False(t, IsNotBlank("	"))
	assert.True(t, IsNotBlank("text"))
}
