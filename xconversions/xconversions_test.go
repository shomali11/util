package xconversions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrettyJson(t *testing.T) {
	x := map[string]interface{}{"number": 1, "string": "cool", "bool": true, "float": 1.5}
	results, err := PrettyJson(x)
	assert.Nil(t, err)
	assert.Equal(t, results, "{\n\t\"bool\": true,\n\t\"float\": 1.5,\n\t\"number\": 1,\n\t\"string\": \"cool\"\n}\n")

	results, err = PrettyJson(func() {})
	assert.NotNil(t, err)
	assert.Equal(t, results, "")
}

func TestStringify(t *testing.T) {
	x := map[string]interface{}{"number": 1, "string": "cool", "bool": true, "float": 1.5}
	results, err := Stringify(x)
	assert.Nil(t, err)
	assert.Equal(t, results, "{\"bool\":true,\"float\":1.5,\"number\":1,\"string\":\"cool\"}")

	results, err = Stringify(func() {})
	assert.NotNil(t, err)
	assert.Equal(t, results, "")
}

func TestStructify(t *testing.T) {
	x := "{\"bool\":true,\"float\":1.5,\"number\":1,\"string\":\"cool\"}"
	var results map[string]interface{}
	err := Structify(x, &results)
	assert.Nil(t, err)
	assert.Equal(t, results, map[string]interface{}{"number": float64(1), "string": "cool", "bool": true, "float": 1.5})
}
