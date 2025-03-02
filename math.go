package godash

import "math"

// Adds two integers.
func Add[T int8 | int16 | int32 | int64 | int](a T, b T) T {
	return a + b
}

func Ceil(a float64) float64 {
	return math.Ceil(a)
}

// func Max
