package godash

import (
	"reflect"
	"testing"
)

func TestChunk(t *testing.T) {
	tests := []struct {
		arr  []int
		size *int
		want [][]int
	}{
		{[]int{1, 2, 3, 4, 5}, intPtr(2), [][]int{{1, 2}, {3, 4}, {5}}},
		{[]int{1, 2, 3}, nil, [][]int{{1}, {2}, {3}}},
		{[]int{1, 2}, intPtr(0), [][]int{{1}, {2}}},
		{[]int{1, 2}, intPtr(1), [][]int{{1}, {2}}},
		{[]int{1, 2}, intPtr(3), [][]int{{1, 2}}},
		{nil, intPtr(2), nil},
		{[]int{}, intPtr(2), [][]int{}},
	}
	for _, tt := range tests {
		got := Chunk(tt.arr, tt.size)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Chunk(%v, %v) = %v, want %v", tt.arr, *tt.size, got, tt.want)
		}
	}
}

func TestConcat(t *testing.T) {
	tests := []struct {
		arr  []int
		args []int
		want []int
	}{
		{[]int{1, 2}, []int{3, 4}, []int{1, 2, 3, 4}},
		{[]int{}, []int{5, 6}, []int{5, 6}},
		{[]int{7, 8}, nil, []int{7, 8}},
		{nil, []int{9}, nil},
	}
	for _, tt := range tests {
		got := Concat(tt.arr, tt.args...)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Concat(%v, %v) = %v, want %v", tt.arr, tt.args, got, tt.want)
		}
	}
}

func TestDrop(t *testing.T) {
	tests := []struct {
		arr  []int
		n    *int
		want []int
	}{
		{[]int{1, 2, 3, 4}, intPtr(2), []int{3, 4}},
		{[]int{1, 2, 3}, intPtr(5), []int{}},
		{[]int{1}, nil, []int{}},
		{nil, intPtr(1), nil},
		{[]int{}, intPtr(1), []int{}},
	}
	for _, tt := range tests {
		got := Drop(tt.arr, tt.n)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Drop(%v, %v) = %v, want %v", tt.arr, tt.n, got, tt.want)
		}
	}
}

func TestDropRight(t *testing.T) {
	tests := []struct {
		arr  []int
		n    *int
		want []int
	}{
		{[]int{1, 2, 3, 4}, intPtr(2), []int{1, 2}},
		{[]int{1, 2, 3}, intPtr(5), []int{}},
		{[]int{1}, nil, []int{}},
		{nil, intPtr(1), nil},
		{[]int{}, intPtr(1), []int{}},
	}
	for _, tt := range tests {
		got := DropRight(tt.arr, tt.n)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("DropRight(%v, %v) = %v, want %v", tt.arr, tt.n, got, tt.want)
		}
	}
}

func TestIncludes(t *testing.T) {
	tests := []struct {
		arr  []int
		val  int
		want bool
	}{
		{[]int{1, 2, 3}, 2, true},
		{[]int{1, 2, 3}, 4, false},
		{nil, 1, false},
		{[]int{}, 0, false},
	}
	for _, tt := range tests {
		got := Includes(tt.arr, tt.val)
		if got != tt.want {
			t.Errorf("Includes(%v, %v) = %v, want %v", tt.arr, tt.val, got, tt.want)
		}
	}
}

func TestFindIndex(t *testing.T) {
	tests := []struct {
		arr  []int
		val  int
		want int
	}{
		{[]int{10, 20, 30}, 20, 1},
		{[]int{10, 20, 30}, 40, -1},
		{[]int{}, 1, -1},
		{nil, 1, -1},
	}
	for _, tt := range tests {
		got := FindIndex(tt.arr, tt.val)
		if got != tt.want {
			t.Errorf("FindIndex(%v, %v) = %v, want %v", tt.arr, tt.val, got, tt.want)
		}
	}
}

func TestHead(t *testing.T) {
	if got := Head([]int{5, 6}); got == nil || *got != 5 {
		t.Errorf("Head() = %v, want 5", got)
	}
	if got := Head([]int{}); got != nil {
		t.Errorf("Head() = %v, want nil", got)
	}
}

func TestLast(t *testing.T) {
	if got := Last([]int{5, 6}); got == nil || *got != 6 {
		t.Errorf("Last() = %v, want 6", got)
	}
	if got := Last([]int{}); got != nil {
		t.Errorf("Last() = %v, want nil", got)
	}
}

func TestInitial(t *testing.T) {
	tests := []struct {
		arr  []int
		want []int
	}{
		{[]int{1, 2, 3}, []int{1, 2}},
		{[]int{42}, []int{}},
		{[]int{}, []int{}},
	}
	for _, tt := range tests {
		got := Initial(tt.arr)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Initial(%v) = %v, want %v", tt.arr, got, tt.want)
		}
	}
}

func TestIntersection(t *testing.T) {
	tests := []struct {
		arrs [][]int
		want []int
	}{
		{[][]int{{1, 2, 3}, {2, 3, 4}, {3, 2, 5}}, []int{2, 3}},
		{[][]int{{1, 2}, {3, 4}}, []int{}},
		{[][]int{{1, 2}}, []int{1, 2}},
		{[][]int{}, []int{}},
	}
	for _, tt := range tests {
		got := Intersection(tt.arrs...)
		if !sameElements(got, tt.want) {
			t.Errorf("Intersection(%v) = %v, want %v", tt.arrs, got, tt.want)
		}
	}
}

func TestJoin(t *testing.T) {
	tests := []struct {
		arr  []string
		sep  string
		want string
	}{
		{[]string{"a", "b", "c"}, "-", "a-b-c"},
		{[]string{"x"}, "-", "x"},
		{[]string{}, "-", ""},
	}
	for _, tt := range tests {
		got := Join(tt.arr, tt.sep)
		if got != tt.want {
			t.Errorf("Join(%v, %v) = %v, want %v", tt.arr, tt.sep, got, tt.want)
		}
	}
}

func intPtr(i int) *int {
	return &i
}

func sameElements[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	m := make(map[T]int)
	for _, x := range a {
		m[x]++
	}
	for _, x := range b {
		m[x]--
		if m[x] < 0 {
			return false
		}
	}
	return true
}
