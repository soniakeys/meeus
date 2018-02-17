// Copyright 2013 Sonia Keys
// License: MIT

package base

import "github.com/soniakeys/unit"

// SmallAngle is threshold used by various routines for switching between
// trigonometric functions and Pythagorean approximations.
//
// In chapter 17, p. 109, Meeus recommends 10′.
var (
	SmallAngle    = unit.AngleFromMin(10) // about .003 radians
	CosSmallAngle = SmallAngle.Cos()      // about .999996
)

// Hav implements the haversine trigonometric function.
//
// See https://en.wikipedia.org/wiki/Haversine_formula.
func Hav(a unit.Angle) float64 {
	// (17.5) p. 115
	return .5 * (1 - a.Cos())
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
func FloorDiv(x, y int) (q int) {
	q = x / y
	if (x < 0) != (y < 0) && x%y != 0 {
		q--
	}
	return
}

// FloorDiv64 returns the integer floor of the fractional value (x / y).
//
// It uses integer math only, so is more efficient than using floating point
// intermediate values.  This function can be used in many places where INT()
// appears in AA.  As with built in integer division, it panics with y == 0.
func FloorDiv64(x, y int64) (q int64) {
	q = x / y
	if (x < 0) != (y < 0) && x%y != 0 {
		q--
	}
	return
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
