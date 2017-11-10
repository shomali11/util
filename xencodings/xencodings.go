package xencodings

import (
	"encoding/base32"
	"encoding/base64"
)

// Base32Encode base32 encode
func Base32Encode(data []byte) string {
	return base32.StdEncoding.EncodeToString(data)
}

// Base32Decode base32 decode
func Base32Decode(data string) ([]byte, error) {
	return base32.StdEncoding.DecodeString(data)
}

// Base64Encode base64 encode
func Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// Base64Decode base64 decode
func Base64Decode(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}
