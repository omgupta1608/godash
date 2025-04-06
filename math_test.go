package godash

import (
	"testing"
)

func TestAdd(t *testing.T) {
	if Add(2, 3) != 5 {
		t.Fail()
	}
	if Add(-4, 7) != 3 {
		t.Fail()
	}
}

func TestCeil(t *testing.T) {
	if Ceil(2.3) != 3 {
		t.Fail()
	}
	if Ceil(-2.3) != -2 {
		t.Fail()
	}
}

func TestDivide(t *testing.T) {
	res, err := Divide(10, 2)
	if res != 5 || err != nil {
		t.Fail()
	}
	_, err = Divide(10, 0)
	if err != ErrDivideByZero {
		t.Fail()
	}
}

func TestMean(t *testing.T) {
	if Mean([]int{1, 2, 3, 4}) != 2.5 {
		t.Fail()
	}
	if Mean([]int{}) != 0 {
		t.Fail()
	}
}

func TestMeanBy(t *testing.T) {
	arr := []int{1, 2, 3}
	got := MeanBy(arr, func(i int) int { return i * 2 })
	if got != 4 {
		t.Fail()
	}
	if MeanBy([]int{}, func(i int) int { return i }) != 0 {
		t.Fail()
	}
}

func TestMax(t *testing.T) {
	arr := []int{3, 5, 2, 9}
	if *Max(arr) != 9 {
		t.Fail()
	}
	if Max([]int{}) != nil {
		t.Fail()
	}
}

func TestMaxBy(t *testing.T) {
	arr := []int{1, 2, 3}
	res := MaxBy(arr, func(i int) int { return -i })
	if *res != 3 {
		t.Fail()
	}
	if MaxBy([]int{}, func(i int) int { return i }) != nil {
		t.Fail()
	}
}

func TestMin(t *testing.T) {
	arr := []int{3, 5, 2, 9}
	if *Min(arr) != 2 {
		t.Fail()
	}
	if Min([]int{}) != nil {
		t.Fail()
	}
}

func TestMinBy(t *testing.T) {
	arr := []int{1, 2, 3}
	res := MinBy(arr, func(i int) int { return -i })
	if *res != 1 {
		t.Fail()
	}
	if MinBy([]int{}, func(i int) int { return i }) != nil {
		t.Fail()
	}
}

func TestMultiply(t *testing.T) {
	if Multiply(2, 3) != 6 {
		t.Fail()
	}
	if Multiply(-1, 3) != -3 {
		t.Fail()
	}
}

func TestRoundOff(t *testing.T) {
	if RoundOff(3.14159, 2) != 3.14 {
		t.Fail()
	}
	if RoundOff(3.1459, 2) != 3.15 {
		t.Fail()
	}
}

func TestSum(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	if Sum(arr) != 10 {
		t.Fail()
	}
	if Sum([]int{}) != 0 {
		t.Fail()
	}
}

func TestSumBy(t *testing.T) {
	arr := []int{1, 2, 3}
	if SumBy(arr, func(i int) int { return i * 2 }) != 12 {
		t.Fail()
	}
	if SumBy([]int{}, func(i int) int { return i }) != 0 {
		t.Fail()
	}
}
