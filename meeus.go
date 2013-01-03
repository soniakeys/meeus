// Copyright 2012 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package meeus

// FloorDiv returns the floor of x / y.
//
// It uses integer math only, so is more efficient than using floating point
// intermediate values.  This function can be used in many places where INT()
// appears in AA.  As with built in integer division, it panics with y == 0.
func FloorDiv(x, y int) int {
	if (x < 0) == (y < 0) {
		return x / y
	}
	return x/y - 1
}

// FloorDiv64 returns the floor of x / y.
//
// It uses integer math only, so is more efficient than using floating point
// intermediate values.  This function can be used in many places where INT()
// appears in AA.  As with built in integer division, it panics with y == 0.
func FloorDiv64(x, y int64) int64 {
	if (x < 0) == (y < 0) {
		return x / y
	}
	return x/y - 1
}

// Cmp compares two float64s and returns -1, 0, or 1 if a is <, ==, or > b,
// respectively.
//
// The name and semantics are chosen to match big.Cmp in the Go standard
// library.
func Cmp(a, b float64) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	}
	return 0
}
