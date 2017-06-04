package util

import (
	"bytes"
	"encoding/json"
)

const (
	empty = ""
	tab   = "\t"
)

func PrettyJson(data interface{}) (string, error) {
	buffer := new(bytes.Buffer)
	enc := json.NewEncoder(buffer)
	enc.SetIndent(empty, tab)

	err := enc.Encode(data)
	if err != nil {
		return empty, err
	}
	return buffer.String(), nil
}
