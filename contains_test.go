package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsString(t *testing.T) {
	slice := []string{"one", "two", "three"}

	result1 := ContainsString(slice, "two")
	result2 := ContainsString(slice, "xpto")

	assert.True(t, result1)

	assert.False(t, result2)
}
