package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	slice := []string{"one", "two", "three"}

	result1 := Contains(slice, "two")

	assert.True(t, result1)
}

func TestContainsString(t *testing.T) {
	slice := []string{"one", "two", "three"}

	result1 := ContainsString(slice, "two")
	result2 := ContainsString(slice, "xpto")

	assert.True(t, result1)

	assert.False(t, result2)
}

func TestContainsInt(t *testing.T) {
	slice := []int{1, 2, 3}

	result1 := ContainsInt(slice, 1)
	result2 := ContainsInt(slice, 4)

	assert.True(t, result1)

	assert.False(t, result2)
}

func TestContainsBool(t *testing.T) {
	slice := []bool{true}

	result1 := ContainsBool(slice, true)
	result2 := ContainsBool(slice, false)

	assert.True(t, result1)

	assert.False(t, result2)
}
