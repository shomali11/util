package util

import (
	"encoding/json"
)

const (
	empty = ""
	tab   = "\t"
)

func PrettyJson(data interface{}) (string, error) {
	bytes, err := json.MarshalIndent(data, empty, tab)
	if err != nil {
		return empty, err
	}
	return string(bytes), nil
}
