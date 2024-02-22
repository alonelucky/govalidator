package govalidator

import (
	"math"
	"reflect"
)

// Abs returns absolute value of number
func Abs(value float64) float64 {
	return math.Abs(value)
}

// Sign returns signum of number: 1 in case of value > 0, -1 in case of value < 0, 0 otherwise
func Sign(value float64) float64 {
	if value > 0 {
		return 1
	} else if value < 0 {
		return -1
	} else {
		return 0
	}
}

// IsNegative returns true if value < 0
func IsNegative(value float64) bool {
	return value < 0
}

// IsPositive returns true if value > 0
func IsPositive(value float64) bool {
	return value > 0
}

// IsNonNegative returns true if value >= 0
func IsNonNegative(value float64) bool {
	return value >= 0
}

// IsNonPositive returns true if value <= 0
func IsNonPositive(value float64) bool {
	return value <= 0
}

// InRangeInt returns true if value lies between left and right border
func InRangeInt(value64, left64, right64 int64) bool {
	if left64 > right64 {
		left64, right64 = right64, left64
	}
	return value64 >= left64 && value64 <= right64
}

// InRangeUint returns true if value lies between left and right border
func InRangeUint(value, left, right uint64) bool {
	if left > right {
		left, right = right, left
	}
	return value >= left && value <= right
}

// InRangeFloat64 returns true if value lies between left and right border
func InRangeFloat64(value, left, right float64) bool {
	if left > right {
		left, right = right, left
	}
	return value >= left && value <= right
}

// InRange returns true if value lies between left and right border, generic type to handle int, float32, float64 and string.
// All types must the same type.
// False if value doesn't lie in range or if it incompatible or not comparable
func InRange(value interface{}, left interface{}, right interface{}) bool {
	var v = reflect.ValueOf(value)
	var l = reflect.ValueOf(left)
	var r = reflect.ValueOf(right)
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return InRangeInt(v.Int(), l.Int(), r.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return InRangeUint(v.Uint(), l.Uint(), r.Uint())
	case reflect.Float32, reflect.Float64:
		return InRangeFloat64(v.Float(), l.Float(), r.Float())
	case reflect.String:
		return value.(string) >= left.(string) && value.(string) <= right.(string)
	default:
		return false
	}
}

// IsWhole returns true if value is whole number
func IsWhole(value float64) bool {
	return math.Remainder(value, 1) == 0
}

// IsNatural returns true if value is natural number (positive and whole)
func IsNatural(value float64) bool {
	return IsWhole(value) && IsPositive(value)
}
