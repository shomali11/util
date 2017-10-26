package xcompressions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompress(t *testing.T) {
	data, err := Compress("Raed Shomali")
	assert.Nil(t, err)

	text, err := Decompress(data)
	assert.Nil(t, err)

	assert.Equal(t, text, "Raed Shomali")
}

func TestDecompress(t *testing.T) {
	_, err := Decompress("Kaka")
	assert.NotNil(t, err)
}
