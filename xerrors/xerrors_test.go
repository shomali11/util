package xerrors

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultErrorIfNil(t *testing.T) {
	assert.Equal(t, DefaultErrorIfNil(nil, "Cool"), "Cool")
	assert.Equal(t, DefaultErrorIfNil(errors.New("Oops"), "Cool"), "Oops")
}
