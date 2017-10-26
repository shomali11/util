package xconversions

import (
	"bytes"
	"encoding/json"
)

const (
	empty = ""
	tab   = "\t"
)

// Stringify returns a string representation
func Stringify(data interface{}) (string, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return empty, err
	}
	return string(b), nil
}

// Structify returns the original representation
func Structify(data string, value interface{}) error {
	return json.Unmarshal([]byte(data), value)
}

// PrettyJson returns a pretty json string
func PrettyJson(data interface{}) (string, error) {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent(empty, tab)

	err := encoder.Encode(data)
	if err != nil {
		return empty, err
	}
	return buffer.String(), nil
}
