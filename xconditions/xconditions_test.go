package xconditions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIfThen(t *testing.T) {
	assert.Equal(t, IfThen(1 == 1, "Yes"), "Yes")
	assert.Equal(t, IfThen(1 != 1, "Woo"), nil)
	assert.Equal(t, IfThen(1 < 2, "Less"), "Less")
}

func TestIfThenElse(t *testing.T) {
	assert.Equal(t, IfThenElse(1 == 1, "Yes", false), "Yes")
	assert.Equal(t, IfThenElse(1 != 1, nil, 1), 1)
	assert.Equal(t, IfThenElse(1 < 2, nil, "No"), nil)
}

func TestDefaultIfNil(t *testing.T) {
	assert.Equal(t, DefaultIfNil(nil, nil), nil)
	assert.Equal(t, DefaultIfNil(nil, ""), "")
	assert.Equal(t, DefaultIfNil("A", "B"), "A")
	assert.Equal(t, DefaultIfNil(true, "B"), true)
	assert.Equal(t, DefaultIfNil(1, false), 1)
}

func TestFirstNonNil(t *testing.T) {
	assert.Equal(t, FirstNonNil(nil, nil), nil)
	assert.Equal(t, FirstNonNil(nil, ""), "")
	assert.Equal(t, FirstNonNil("A", "B"), "A")
	assert.Equal(t, FirstNonNil(true, "B"), true)
	assert.Equal(t, FirstNonNil(1, false), 1)
	assert.Equal(t, FirstNonNil(nil, nil, nil, 10), 10)
	assert.Equal(t, FirstNonNil(nil, nil, nil, nil, nil), nil)
	assert.Equal(t, FirstNonNil(), nil)
}
