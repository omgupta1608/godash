# godash - Lodash in Go

`godash` is a utility package in Go that provides a collection of functions for common array, math, and functional-style operations inspired by libraries like Lodash.

## Installation

```bash
go get github.com/omgupta1608/godash
```


## Usage

```go
import "github.com/omgupta1608/godash"
```

---

## 📦 Array Utilities

### `Chunk(arr []T, size *int) [][]T`
Splits `arr` into chunks of `size`. If `size` is nil or 0, defaults to 1.

### `Concat(arr []T, args ...T) []T`
Concatenates `arr` with `args`.

### `Drop(arr []T, n *int) []T`
Drops the first `n` elements from `arr`. Defaults to 1 if `n` is nil.

### `DropRight(arr []T, n *int) []T`
Drops the last `n` elements from `arr`. Defaults to 1 if `n` is nil.

### `Includes(arr []T, val T) bool`
Checks if `val` exists in `arr`.

### `FindIndex(arr []T, val T) int`
Returns the index of the first match for `val` in `arr`.

### `Head(arr []T) *T`
Returns the first element of `arr`, or nil if empty.

### `Last(arr []T) *T`
Returns the last element of `arr`, or nil if empty.

### `Initial(arr []T) []T`
Returns all but the last element of `arr`.

### `Intersection(arrays ...[]T) []T`
Returns an array of unique values that are present in all arrays.

### `Join(arr []string, sep string) string`
Joins the elements in `arr` into a string with separator `sep`.

---

## 🧮 Math Utilities

### `Add(a, b T) T`
Returns the sum of two numbers.

### `Ceil(a float64) float64`
Returns the smallest integer value greater than or equal to `a`.

### `Divide(a, b T) (T, error)`
Returns the result of `a / b`. Returns an error if dividing by zero.

### `Mean(arr []T) float64`
Returns the average of elements in `arr`.

### `MeanBy(arr []T, iteratee func(T) T) float64`
Applies `iteratee` to elements before calculating the mean.

### `Max(arr []T) *T`
Returns the maximum number in `arr`.

### `MaxBy(arr []T, iteratee func(T) T) *T`
Applies `iteratee` to rank elements and returns the max.

### `Min(arr []T) *T`
Returns the minimum number in `arr`.

### `MinBy(arr []T, iteratee func(T) T) *T`
Applies `iteratee` to rank elements and returns the min.

### `Multiply(a, b T) T`
Returns the product of two numbers.

### `RoundOff(value float64, precision int) float64`
Rounds `value` to the specified number of decimal places.

### `Sum(arr []T) T`
Returns the sum of elements in `arr`.

### `SumBy(arr []T, iteratee func(T) T) T`
Applies `iteratee` to elements before summing.

---

## 🧱 Number Utilities

### `Clamp(num, upper T, lower *T) T`
Clamps `num` within `[lower, upper]` bounds. If `lower` is nil, only upper bound is enforced.

### `InRange(num T, args ...T) (bool, error)`
Checks if `num` is within range defined by one or two values. Behaves like Lodash `_.inRange`.

---

## 🔁 Functional Iterable

The `Iterable[T]` struct enables chaining-like behavior for collections.

```go
iter := godash.Iterable[int]{Val: []int{1, 2, 3, 4}}
```

### `Map(f func(T) T) Iterable[T]`
Returns a new iterable by applying `f` to each element.

### `ForEach(f func(T))`
Applies `f` to each element without returning a new iterable.

### `Filter(f func(T) bool) Iterable[T]`
Returns a new iterable with elements passing the predicate `f`.

### `Reduce(f func(T, T) T, initial T) *T`
Reduces elements to a single value using `f`.

### `Find(f func(T) bool) *T`
Returns the first element satisfying the predicate `f`.

### `FindIndex(f func(T) bool) int`
Returns the index of the first element satisfying `f`.

### `Every(f func(T) bool) bool`
Returns `true` if all elements satisfy `f`.

### `Some(f func(T) bool) bool`
Returns `true` if any element satisfies `f`.

### `Push(args ...T)`
Appends elements to the iterable.

### `Fill(value T, start, end int)`
Fills a range `[start:end]` with the given value.

---

## 🧪 Testing

You can run tests to ensure if everything is working.

```bash
make tests
```

---

