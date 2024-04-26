package godash

import (
	"reflect"
)

// Returns an array of elements split into groups the length of size. If array can't be split evenly, the final chunk will be the remaining elements.
func Chunk[T any](arr []T, size *int) (res [][]T) {
	if arr == nil {
		return nil
	}
	var DEFAULT_CHUNK_SIZE int = 1
	if size == nil || *size == 0 {
		size = &DEFAULT_CHUNK_SIZE
	}
	if len(arr) <= *size {
		res = append(res, arr)
		return res
	}
	i := 0
	for len(arr)-(i+1) >= *size {
		res = append(res, arr[i:i+*size])
		i += *size
	}
	if i < len(arr) {
		res = append(res, arr[i:])
	}
	return
}

// Returns a new array concatenating array with any additional values.
func Concat[T any](arr []T, args ...T) []T {
	if arr == nil {
		return nil
	}
	arr = append(arr, args...)
	return arr
}

// Returns a new slice with n elements dropped from the beginning.
func Drop[T any](arr []T, n *int) []T {
	if arr == nil {
		return nil
	}
	var DEFAULT_DROP_SIZE int = 1
	if n == nil {
		n = &DEFAULT_DROP_SIZE
	}

	if *n >= len(arr) {
		return []T{}
	}
	return arr[*n:]
}

// Returns a slice with n elements dropped from the end.
func DropRight[T any](arr []T, n *int) []T {
	if arr == nil {
		return nil
	}
	var DEFAULT_DROP_SIZE int = 1
	if n == nil {
		n = &DEFAULT_DROP_SIZE
	}
	if *n >= len(arr) {
		return []T{}
	}
	return arr[:len(arr)-*n]
}

// Checks whether a value is present in the given array or not. Returns a boolean.
func Includes[T any](arr []T, val T) bool {
	if arr == nil {
		return false
	}
	for _, v := range arr {
		if reflect.DeepEqual(v, val) {
			return true
		}
	}
	return false
}

// Returns the index of the first element that matches val.
func FindIndex[T any](arr []T, val T) int {
	if arr == nil {
		return -1
	}
	for i, v := range arr {
		if reflect.DeepEqual(v, val) {
			return i
		}
	}
	return -1
}

// Gets the first element of array. Returns nil with the array is empty or nil
func Head[T any](arr []T) *T {
	if len(arr) > 0 {
		return &arr[0]
	}
	return nil
}

// Gets the last element of array.  Returns nil with the array is empty or nil
func Last[T any](arr []T) *T {
	if len(arr) > 0 {
		return &arr[len(arr)-1]
	}
	return nil
}

// Gets all but the last element of array.
func Initial[T any](arr []T) []T {
	if len(arr) == 0 || len(arr) == 1 {
		return []T{}
	}
	return arr[:len(arr)-1]
}

// Creates an array of unique values that are included in all given arrays.
func Intersection[T comparable](arrays ...[]T) []T {
	valCounts := map[T]int{}
	for _, arr := range arrays {
		for _, val := range arr {
			valCounts[val]++
		}
	}
	res := []T{}
	for k, v := range valCounts {
		if v >= len(arrays) {
			res = append(res, k)
		}
	}
	return res
}

// Converts all elements in array into a string separated by separator.
func Join(arr []string, sep string) (res string) {
	for i, v := range arr {
		res += v
		if i != len(arr)-1 {
			res += sep
		}
	}
	return
}
