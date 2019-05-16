package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnsupportedTypeReturnsFalse(t *testing.T) {
	slice := []complex128{}

	result := Contains(slice, "two")

	assert.False(t, result)
}

func TestContains(t *testing.T) {
	slString := []string{"one", "two", "three"}
	slBool := []bool{true}
	slInt := []int{1, 2, 3}
	slInt64 := []int64{1, 2, 3}
	slUint := []uint{1, 2, 3}
	slFloat32 := []float32{10.0, 50.0, 30.0}
	float64slice := []float64{10.0, 50.0, 30.0}

	var res bool

	res = Contains(slString, "two")
	assert.True(t, res)

	res = Contains(slBool, true)
	assert.True(t, res)

	res = Contains(slInt, 1)
	assert.True(t, res)

	res = Contains(slInt64, 1)
	assert.True(t, res)

	res = Contains(slUint, 1)
	assert.True(t, res)

	res = Contains(slFloat32, 10.0)
	assert.True(t, res)

	res = Contains(float64slice, 10.0)
	assert.True(t, res)
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

func TestContainsInt32(t *testing.T) {
	slice := []int32{1, 2, 3}

	result1 := ContainsInt32(slice, 1)
	result2 := ContainsInt32(slice, 4)

	assert.True(t, result1)

	assert.False(t, result2)
}

func TestContainsIn64(t *testing.T) {
	slice := []int64{1, 2, 3}

	result1 := ContainsInt64(slice, 1)
	result2 := ContainsInt64(slice, 4)

	assert.True(t, result1)

	assert.False(t, result2)
}

func TestContainsUint(t *testing.T) {
	slice := []uint{1, 2, 3}

	result1 := ContainsUInt(slice, 1)
	result2 := ContainsUInt(slice, 4)

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

func TestContainsFloat32(t *testing.T) {
	slice := []float32{10.0, 50.0, 30.0}
	result1 := ContainsFloat32(slice, 10.0)
	result2 := ContainsFloat32(slice, 100.0)

	assert.True(t, result1)

	assert.False(t, result2)
}

func TestContainsFloat64(t *testing.T) {
	slice := []float64{10.0, 50.0, 30.0}
	result1 := ContainsFloat64(slice, 10.0)
	result2 := ContainsFloat64(slice, 100.0)

	assert.True(t, result1)

	assert.False(t, result2)
}
