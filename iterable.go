package godash

type Iterable[T any] struct {
	Val []T
}

// Returns a new iterable with each element transformed based on the callback.
func (v *Iterable[T]) Map(f func(T) T) Iterable[T] {
	if v.Val == nil {
		return Iterable[T]{Val: nil}
	}
	arr := []T{}
	for _, e := range v.Val {
		arr = append(arr, f(e))
	}
	return Iterable[T]{Val: arr}
}

// Iterates over each element but does not return a new interable.
func (v *Iterable[T]) ForEach(f func(T)) {
	if v.Val == nil {
		return
	}
	for _, e := range v.Val {
		f(e)
	}
}

func (v *Iterable[T]) Filter(f func(T) bool) Iterable[T] {
	if v.Val == nil {
		return Iterable[T]{Val: nil}
	}

	arr := []T{}
	for _, e := range v.Val {
		if f(e) {
			arr = append(arr, e)
		}
	}
	return Iterable[T]{Val: arr}
}

// Reduces the iterable to a single value.
func (v *Iterable[T]) Reduce(f func(T, T) T, initalVal T) *T {
	if v.Val == nil {
		return nil
	}
	for _, e := range v.Val {
		initalVal = f(initalVal, e)
	}

	return &initalVal
}

// Returns the first matching element.
func (v *Iterable[T]) Find(f func(T) bool) *T {
	if v.Val == nil {
		return nil
	}

	for _, e := range v.Val {
		if f(e) {
			return &e
		}
	}
	return nil
}

// Returns the index of the first matching element.
func (v *Iterable[T]) FindIndex(f func(T) bool) int {
	if v.Val == nil {
		return -1
	}

	for i, e := range v.Val {
		if f(e) {
			return i
		}
	}
	return -1
}

// Returns true if all elements satisfy the condition.
func (v *Iterable[T]) Every(f func(T) bool) bool {
	if v.Val == nil {
		return false
	}

	for _, e := range v.Val {
		if !f(e) {
			return false
		}
	}
	return true
}

// Returns true if at least one element satisfies the condition.
func (v *Iterable[T]) Some(f func(T) bool) bool {
	if v.Val == nil {
		return false
	}

	for _, e := range v.Val {
		if f(e) {
			return true
		}
	}
	return false
}
