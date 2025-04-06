package godash

import (
	"errors"
	"math"
)

var (
	ErrDivideByZero = errors.New("attempting to divide by zero")
)

// Adds two integers.
func Add[T Number](a T, b T) T {
	return a + b
}

func Ceil(a float64) float64 {
	return math.Ceil(a)
}

func Divide[T Number](a T, b T) (T, error) {
	if b == T(0) {
		return -1, ErrDivideByZero
	}
	return a / b, nil
}

// Calculate the mean of all the elements
func Mean[T Number](arr []T) float64 {
	l := T(len(arr))
	if l == 0 {
		return 0
	}

	sum := 0.0
	for _, e := range arr {
		sum += float64(e)
	}
	return float64(sum / float64(l))
}

// This method is like Mean except that it accepts iteratee which is invoked for each element in array to generate the value to be averaged. The iteratee is invoked with one argument: (value).
func MeanBy[T Number](arr []T, iteratee func(T) T) float64 {
	l := T(len(arr))
	if l == 0 {
		return 0
	}

	sum := T(0)
	for _, e := range arr {
		sum += iteratee(e)
	}
	return float64(sum / l)
}

// Returns the max Number in the array
func Max[T Number](arr []T) *T {
	if len(arr) == 0 {
		return nil
	}
	if len(arr) == 1 {
		return &arr[0]
	}
	max := arr[0]
	for _, e := range arr {
		if e > max {
			max = e
		}
	}
	return &max
}

// This method is like Max except that it accepts iteratee which is invoked for each element in array to generate the criterion by which the value is ranked. The iteratee is invoked with one argument: (value).
func MaxBy[T Number](arr []T, iteratee func(T) T) *T {
	if len(arr) == 0 {
		return nil
	}
	if len(arr) == 1 {
		return &arr[0]
	}
	max := arr[0]
	for _, e := range arr {
		if e > iteratee(e) {
			max = e
		}
	}
	return &max

}

// Returns the min Number in the array
func Min[T Number](arr []T) *T {
	if len(arr) == 0 {
		return nil
	}
	if len(arr) == 1 {
		return &arr[0]
	}
	min := arr[0]
	for _, e := range arr {
		if e < min {
			min = e
		}
	}
	return &min
}

// This method is like Min except that it accepts iteratee which is invoked for each element in array to generate the criterion by which the value is ranked. The iteratee is invoked with one argument: (value).
func MinBy[T Number](arr []T, iteratee func(T) T) *T {
	if len(arr) == 0 {
		return nil
	}
	if len(arr) == 1 {
		return &arr[0]
	}
	min := arr[0]
	for _, e := range arr {
		if e < iteratee(e) {
			min = e
		}
	}
	return &min
}

// Multiply two Numbers
func Multiply[T Number](a, b T) T {
	return a * b
}

// Round off a Float to the specified precision
func RoundOff(value float64, precision int) float64 {
	multiplier := math.Pow(10, float64(precision))
	return math.Round(value*multiplier) / multiplier
}

// Sum of an array of Numbers
func Sum[T Number](arr []T) T {
	sum := T(0)
	for _, e := range arr {
		sum += e
	}
	return sum
}

// This method is like Sum except that it accepts iteratee which is invoked for each element in array to generate the value to be summed. The iteratee is invoked with one argument: (value).
func SumBy[T Number](arr []T, iteratee func(T) T) T {
	sum := T(0)
	for _, e := range arr {
		sum += iteratee(e)
	}
	return sum
}
