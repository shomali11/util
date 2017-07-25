package hashes

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
)

// MD5 hashes using md5 algorithm
func MD5(text string) string {
	algorithm := md5.New()
	return hasher(algorithm, text)
}

// SHA1 hashes using sha1 algorithm
func SHA1(text string) string {
	algorithm := sha1.New()
	return hasher(algorithm, text)
}

// SHA256 hashes using sha256 algorithm
func SHA256(text string) string {
	algorithm := sha256.New()
	return hasher(algorithm, text)
}

// SHA512 hashes using sha512 algorithm
func SHA512(text string) string {
	algorithm := sha512.New()
	return hasher(algorithm, text)
}

func hasher(algorithm hash.Hash, text string) string {
	algorithm.Write([]byte(text))
	return hex.EncodeToString(algorithm.Sum(nil))
}
