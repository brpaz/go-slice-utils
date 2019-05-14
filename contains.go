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
	case []int8:
		var typedContains int8
		if _, ok := elem.(int); ok {
			typedContains = int8(elem.(int))
		} else {
			typedContains = elem.(int8)
		}
		return ContainsInt8(slice.([]int8), typedContains)
	case []int16:
		var typedContains int16
		if _, ok := elem.(int); ok {
			typedContains = int16(elem.(int))
		} else {
			typedContains = elem.(int16)
		}
		return ContainsInt16(slice.([]int16), typedContains)
	case []int32:
		var typedContains int32
		if _, ok := elem.(int); ok {
			typedContains = int32(elem.(int))
		} else {
			typedContains = elem.(int32)
		}
		return ContainsInt32(slice.([]int32), typedContains)
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
	case []uint8:
		var typedContains uint8
		if _, ok := elem.(int); ok {
			typedContains = uint8(elem.(int))
		} else {
			typedContains = elem.(uint8)
		}
		return ContainsUInt8(slice.([]uint8), typedContains)
	case []uint16:
		var typedContains uint16
		if _, ok := elem.(int); ok {
			typedContains = uint16(elem.(int))
		} else {
			typedContains = elem.(uint16)
		}
		return ContainsUInt16(slice.([]uint16), typedContains)
	case []uint32:
		var typedContains uint32
		if _, ok := elem.(int); ok {
			typedContains = uint32(elem.(int))
		} else {
			typedContains = elem.(uint32)
		}
		return ContainsUInt32(slice.([]uint32), typedContains)
	case []uint64:
		var typedContains uint64
		if _, ok := elem.(int); ok {
			typedContains = uint64(elem.(int))
		} else {
			typedContains = elem.(uint64)
		}
		return ContainsUInt64(slice.([]uint64), typedContains)
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
	default:
		return false

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

// ContainsInt8 checks if the slice has the elem value in it.
func ContainsInt8(slice []int8, elem int8) bool {
	for _, value := range slice {
		if value == elem {
			return true
		}
	}
	return false
}

// ContainsInt16 checks if the slice has the elem value in it.
func ContainsInt16(slice []int16, elem int16) bool {
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

// ContainsUInt8 checks if the slice has the elem value in it.
func ContainsUInt8(slice []uint8, elem uint8) bool {
	for _, value := range slice {
		if value == elem {
			return true
		}
	}
	return false
}

// ContainsUInt16 checks if the slice has the elem value in it.
func ContainsUInt16(slice []uint16, elem uint16) bool {
	for _, value := range slice {
		if value == elem {
			return true
		}
	}
	return false
}

// ContainsUInt32 checks if the slice has the elem value in it.
func ContainsUInt32(slice []uint32, elem uint32) bool {
	for _, value := range slice {
		if value == elem {
			return true
		}
	}
	return false
}

// ContainsUInt64 checks if the slice has the elem value in it.
func ContainsUInt64(slice []uint64, elem uint64) bool {
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

