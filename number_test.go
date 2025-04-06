package godash

import (
	"testing"
)

func TestClamp(t *testing.T) {
	if Clamp(5, 10, nil) != 5 {
		t.Fail()
	}
	if Clamp(15, 10, nil) != 10 {
		t.Fail()
	}
	lower := 3
	if Clamp(2, 10, &lower) != 3 {
		t.Fail()
	}
	if Clamp(12, 10, &lower) != 10 {
		t.Fail()
	}
	if Clamp(5, 10, &lower) != 5 {
		t.Fail()
	}
}

func TestInRange(t *testing.T) {
	ok, err := InRange(5, 10)
	if !ok || err != nil {
		t.Fail()
	}
	ok, err = InRange(10, 10)
	if ok || err != nil {
		t.Fail()
	}
	ok, err = InRange(5, 2, 10)
	if !ok || err != nil {
		t.Fail()
	}
	ok, err = InRange(1, 2, 10)
	if ok || err != nil {
		t.Fail()
	}
	_, err = InRange(5)
	if err == nil {
		t.Fail()
	}
	_, err = InRange(5, 1, 2, 3)
	if err == nil {
		t.Fail()
	}
}
