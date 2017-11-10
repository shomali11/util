package xencodings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBase32Encode(t *testing.T) {
	data := Base32Encode([]byte("Raed Shomali"))
	assert.Equal(t, data, "KJQWKZBAKNUG63LBNRUQ====")
}

func TestBase32Decode(t *testing.T) {
	data, err := Base32Decode("KJQWKZBAKNUG63LBNRUQ====")
	assert.Nil(t, err)
	assert.Equal(t, string(data), "Raed Shomali")
}

func TestBase64Encode(t *testing.T) {
	data := Base64Encode([]byte("Raed Shomali"))
	assert.Equal(t, data, "UmFlZCBTaG9tYWxp")
}

func TestBase64Decode(t *testing.T) {
	data, err := Base64Decode("UmFlZCBTaG9tYWxp")
	assert.Nil(t, err)
	assert.Equal(t, string(data), "Raed Shomali")
}
