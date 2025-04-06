package godash

import (
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	it := Iterable[int]{Val: []int{1, 2, 3}}
	got := it.Map(func(x int) int { return x * 2 })
	want := Iterable[int]{Val: []int{2, 4, 6}}
	if !reflect.DeepEqual(got, want) {
		t.Fail()
	}
}

func TestForEach(t *testing.T) {
	it := Iterable[int]{Val: []int{1, 2, 3}}
	sum := 0
	it.ForEach(func(x int) { sum += x })
	if sum != 6 {
		t.Fail()
	}
}

func TestFilter(t *testing.T) {
	it := Iterable[int]{Val: []int{1, 2, 3, 4}}
	got := it.Filter(func(x int) bool { return x%2 == 0 })
	want := Iterable[int]{Val: []int{2, 4}}
	if !reflect.DeepEqual(got, want) {
		t.Fail()
	}
}

func TestReduce(t *testing.T) {
	it := Iterable[int]{Val: []int{1, 2, 3}}
	got := it.Reduce(func(a, b int) int { return a + b }, 0)
	if got == nil || *got != 6 {
		t.Fail()
	}
}

func TestFind(t *testing.T) {
	it := Iterable[int]{Val: []int{1, 2, 3, 4}}
	got := it.Find(func(x int) bool { return x > 2 })
	if got == nil || *got != 3 {
		t.Fail()
	}
}

func TestFindIndexItr(t *testing.T) {
	it := Iterable[int]{Val: []int{1, 2, 3, 4}}
	got := it.FindIndex(func(x int) bool { return x == 3 })
	if got != 2 {
		t.Fail()
	}
}

func TestEvery(t *testing.T) {
	it := Iterable[int]{Val: []int{2, 4, 6}}
	if !it.Every(func(x int) bool { return x%2 == 0 }) {
		t.Fail()
	}
}

func TestSome(t *testing.T) {
	it := Iterable[int]{Val: []int{1, 3, 4}}
	if !it.Some(func(x int) bool { return x%2 == 0 }) {
		t.Fail()
	}
}

func TestPush(t *testing.T) {
	it := Iterable[int]{Val: []int{1, 2}}
	it.Push(3, 4)
	want := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(it.Val, want) {
		t.Fail()
	}
}

func TestFill(t *testing.T) {
	it := Iterable[int]{Val: []int{1, 2, 3, 4, 5}}
	it.Fill(9, 1, 3)
	want := []int{1, 9, 9, 9, 5}
	if !reflect.DeepEqual(it.Val, want) {
		t.Fail()
	}
}
