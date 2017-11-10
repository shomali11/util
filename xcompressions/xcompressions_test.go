package xcompressions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompress(t *testing.T) {
	data, err := Compress([]byte("Raed Shomali"))
	assert.Nil(t, err)

	decompressedData, err := Decompress(data)
	assert.Nil(t, err)

	assert.Equal(t, string(decompressedData), "Raed Shomali")
}

func TestDecompress(t *testing.T) {
	_, err := Decompress([]byte("Kaka"))
	assert.NotNil(t, err)
}
