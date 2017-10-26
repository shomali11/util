package xcompressions

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
)

const (
	empty = ""
)

// Compress returns a compressed string
func Compress(text string) (string, error) {
	var buffer bytes.Buffer
	gzipWriter := gzip.NewWriter(&buffer)

	_, err := gzipWriter.Write([]byte(text))
	if err != nil {
		return empty, err
	}

	err = gzipWriter.Flush()
	if err != nil {
		return empty, err
	}

	err = gzipWriter.Close()
	if err != nil {
		return empty, err
	}
	return string(buffer.Bytes()), nil
}

// Decompress returns the decompressed string
func Decompress(data string) (string, error) {
	dataBytes := []byte(data)
	byteReader := bytes.NewReader(dataBytes)
	gzipReader, err := gzip.NewReader(byteReader)
	if err != nil {
		return empty, err
	}

	textBytes, err := ioutil.ReadAll(gzipReader)
	if err != nil {
		return empty, err
	}

	err = gzipReader.Close()
	if err != nil {
		return empty, err
	}
	return string(textBytes), nil
}
