package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexOf(t *testing.T) {
	arr := []string{"abc", "123"}
	idx := IndexOf[string](arr, "abc")
	assert.Equal(t, idx, 0)
}

func TestContains(t *testing.T) {
	arr := []string{"abc", "123"}
	has := Contains[string](arr, "abc")
	assert.True(t, has)
}
