package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrettyJson(t *testing.T) {
	x := map[string]interface{}{"number": 1, "string": "cool", "bool": true, "float": 1.5}

	results, err := PrettyJson(x)
	assert.Nil(t, err)
	assert.Equal(t, results, "{\n\t\"bool\": true,\n\t\"float\": 1.5,\n\t\"number\": 1,\n\t\"string\": \"cool\"\n}")

	results, err = PrettyJson(func() {})
	assert.NotNil(t, err)
	assert.Equal(t, results, "")
}
