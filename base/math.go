// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package base

import "math"

// SmallAngle is threshold used by various routines for switching between
// trigonometric functions and Pythagorean approximations.
//
// In chapter 17, p. 109, Meeus recommends 10′.
var (
	SmallAngle    = 10 * math.Pi / 180 / 60 // about .003 radians
	CosSmallAngle = math.Cos(SmallAngle)    // about .999996
)

// PMod returns a positive floating-point x mod y.
//
// For a positive argument y, it returns a value in the range [0,y).
//
// The result may not be useful if y is negative.
func PMod(x, y float64) float64 {
	r := math.Mod(x, y)
	if r < 0 {
		r += y
	}
	return r
}

// Horner evaluates a polynomal with coefficients c at x.  The constant
// term is c[0].  The function panics with an empty coefficient list.
func Horner(x float64, c ...float64) float64 {
	i := len(c) - 1
	y := c[i]
	for i > 0 {
		i--
		y = y*x + c[i] // sorry, no fused multiply-add in Go
	}
	return y
}

// FloorDiv returns the integer floor of the fractional value (x / y).
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

// FloorDiv64 returns the integer floor of the fractional value (x / y).
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
