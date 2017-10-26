package xhashes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFNV32(t *testing.T) {
	data := FNV32("Raed Shomali")
	assert.Equal(t, data, uint32(0x194b953a))
}

func TestFNV32a(t *testing.T) {
	data := FNV32a("Raed Shomali")
	assert.Equal(t, data, uint32(0xdd36df08))
}

func TestFNV64(t *testing.T) {
	data := FNV64("Raed Shomali")
	assert.Equal(t, data, uint64(0xc03d55b5ff7722da))
}

func TestFNV64a(t *testing.T) {
	data := FNV64a("Raed Shomali")
	assert.Equal(t, data, uint64(0xf7fc847f1e6b4148))
}

func TestMD5(t *testing.T) {
	data := MD5("Raed Shomali")
	assert.Equal(t, data, "c313bc3b48fcfed9abc733429665b105")
}

func TestSHA1(t *testing.T) {
	data := SHA1("Raed Shomali")
	assert.Equal(t, data, "e0d66f6f09de72942e83289cc994b3c721ab34c5")
}

func TestSHA256(t *testing.T) {
	data := SHA256("Raed Shomali")
	assert.Equal(t, data, "75894b9be21065a833e57bfe4440b375fc216f120a965243c9be8b2dc36709c2")
}

func TestSHA512(t *testing.T) {
	data := SHA512("Raed Shomali")
	assert.Equal(t, data, "406e8d495140187a8b09893c30d054cf385ad7359855db0d2e0386c7189ac1c4667a4816d1b63a19f3d8ccdcbace7861ec4cc6ff5e2a1659c8f4360bda699b42")
}
