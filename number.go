package godash

// Clamps number within the inclusive lower and upper bounds.
func Clamp[T Number](num T, upper T, lower *T) T {
	if lower == nil {
		if num < upper {
			return num
		} else {
			return upper
		}
	} else {
		if num < *lower {
			return *lower
		} else if num > upper {
			return upper
		} else {
			return num
		}
	}
}

// Checks if n is between start and up to, but not including, end. If end is not specified, it's set to start with start then set to 0. If start is greater than end the params are swapped to support negative ranges.
func InRange[T Number](num T, args ...T) (bool, error) {
	if len(args) == 0 || len(args) > 2 {
		return false, ErrInvalidArgs
	}

	start := T(0)
	end := args[1]
	if len(args) == 2 {
		start = end
		end = args[1]
	}

	return num >= start && num < end, nil
}
