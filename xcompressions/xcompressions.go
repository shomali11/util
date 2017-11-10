package xcompressions

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
)

// Compress returns compressed bytes
func Compress(data []byte) ([]byte, error) {
	var buffer bytes.Buffer
	gzipWriter := gzip.NewWriter(&buffer)

	_, err := gzipWriter.Write(data)
	if err != nil {
		return nil, err
	}

	err = gzipWriter.Flush()
	if err != nil {
		return nil, err
	}

	err = gzipWriter.Close()
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

// Decompress returns the decompressed bytes
func Decompress(data []byte) ([]byte, error) {
	byteReader := bytes.NewReader(data)
	gzipReader, err := gzip.NewReader(byteReader)
	if err != nil {
		return nil, err
	}

	decompressedBytes, err := ioutil.ReadAll(gzipReader)
	if err != nil {
		return nil, err
	}

	err = gzipReader.Close()
	if err != nil {
		return nil, err
	}
	return decompressedBytes, nil
}
