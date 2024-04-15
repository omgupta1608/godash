package godash

import "reflect"

func Chunk[T any](arr []T, size *int) [][]T {
	var DEFAULT_CHUNK_SIZE int = 1
	if size == nil {
		size = &DEFAULT_CHUNK_SIZE
	}
	return nil

	// TODO
}

func Concat[T any](arr []T, args ...T) []T {
	// TODO
	return arr
}

func Drop[T any](arr []T, n *int) []T {
	var DEFAULT_DROP_SIZE int = 1
	if n == nil {
		n = &DEFAULT_DROP_SIZE
	}
	return arr

	// TODO
}

func DropRight[T any](arr []T, n *int) []T {
	var DEFAULT_DROP_SIZE int = 1
	if n == nil {
		n = &DEFAULT_DROP_SIZE
	}
	return arr

	// TODO
}

func Includes[T any](arr []T, val T) bool {
	// TODO
	return true
}

func FindIndex[T any](arr []T, val T) int {
	for i, v := range arr {
		if reflect.DeepEqual(v, val) {
			return i
		}
	}
	return -1
}

func Head[T any](arr []T) *T {
	if len(arr) > 0 {
		return &arr[0]
	}
	return nil
}

func Last[T any](arr []T) *T {
	if len(arr) > 0 {
		return &arr[len(arr)-1]
	}
	return nil
}

func Initial[T any](arr []T) []T {
	if len(arr) == 0 || len(arr) == 1 {
		return []T{}
	}
	return arr[:len(arr)-1]
}

func Intersection[T comparable](arr1 []T, arr2 []T) []T {
	res := []T{}
	v1 := make(map[T]bool)

	for _, v := range arr1 {
		v1[v] = true
	}

	for _, v := range arr2 {
		if _, ok := v1[v]; ok {
			res = append(res, v)
		}
	}
	return res
}

func Join(arr []string, sep string) (res string) {
	for i, v := range arr {
		res += v
		if i != len(arr)-1 {
			res += sep
		}
	}
	return
}

