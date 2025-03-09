package godash

import "errors"

var ErrInvalidArgs = errors.New("invalid arguments")

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

type Float interface {
	float32 | float64
}

type Interger interface {
	int | int8 | int16 | int32 | int64
}

func String(v string) *string {
	return &v
}

func Int(v int) *int {
	return &v
}

func Bool(v bool) *bool {
	return &v
}
