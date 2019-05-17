// Package slice contains common used helper functions for dealing with slices.
package slice

// Contains checks if the "elem" argument argument is present in the specified "slice" argument.
// In order to make it generic, this function have to do some type assertions which makes it slower
// than calling the ContainsType function directly. If this magnitude of performance is critical
// in your application please call the typed functions directly.
func Contains(slice interface{}, elem interface{}) bool {

	// Determine what type the "elem" variable is, then act on it
	switch slice.(type) {
	case []bool:
		return ContainsBool(slice.([]bool), elem.(bool))
	case []int:
		return ContainsInt(slice.([]int), elem.(int))
	case []int64:
		var typedContains int64
		if _, ok := elem.(int); ok {
			typedContains = int64(elem.(int))
		} else {
			typedContains = elem.(int64)
		}
		return ContainsInt64(slice.([]int64), typedContains)
	case []uint:
		var typedContains uint
		if _, ok := elem.(int); ok {
			typedContains = uint(elem.(int))
		} else {
			typedContains = elem.(uint)
		}
		return ContainsUInt(slice.([]uint), typedContains)
	case []float32:
		var typedContains float32
		if _, ok := elem.(float64); ok {
			typedContains = float32(elem.(float64))
		} else {
			typedContains = elem.(float32)
		}
		return ContainsFloat32(slice.([]float32), typedContains)
	case []float64:
		if typedSlice, ok := slice.([]float64); ok {
			return ContainsFloat64(typedSlice, elem.(float64))
		}
	case []string:
		return ContainsString(slice.([]string), elem.(string))
	}

	return false
}

// ContainsBool checks if the slice has the elem value in it.
func ContainsBool(slice []bool, elem bool) bool {
	for _, value := range slice {
		if value == elem {
			return true
		}
	}
	return false
}

// ContainsInt checks if the slice has the elem value in it.
func ContainsInt(slice []int, elem int) bool {
	for _, value := range slice {
		if value == elem {
			return true
		}
	}
	return false
}

// ContainsInt32 checks if the slice has the elem value in it.
func ContainsInt32(slice []int32, elem int32) bool {
	for _, value := range slice {
		if value == elem {
			return true
		}
	}
	return false
}

// ContainsInt64 checks if the slice has the elem value in it.
func ContainsInt64(slice []int64, elem int64) bool {
	for _, value := range slice {
		if value == elem {
			return true
		}
	}
	return false
}

// ContainsUInt checks if the slice has the elem value in it.
func ContainsUInt(slice []uint, elem uint) bool {
	for _, value := range slice {
		if value == elem {
			return true
		}
	}
	return false
}

// ContainsFloat32 checks if the slice has the elem value in it.
func ContainsFloat32(slice []float32, elem float32) bool {
	for _, value := range slice {
		if value == elem {
			return true
		}
	}
	return false
}

// ContainsFloat64 checks if the slice has the elem value in it.
func ContainsFloat64(slice []float64, elem float64) bool {
	for _, value := range slice {
		if value == elem {
			return true
		}
	}
	return false
}

// ContainsString checks if the slice has the elem value in it.
func ContainsString(slice []string, elem string) bool {
	for _, value := range slice {
		if value == elem {
			return true
		}
	}
	return false
}
