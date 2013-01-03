// Copyright 2012 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Interp: Chapter 3, Interpolation.
//
// Len3 and Len5 functions
//
// These functions interpolate from a table of equidistant x values
// and corresponding y values.  Since the x values are equidistant,
// only the first and last values are supplied as arguments.  The interior
// values are implicit.
//
// All y values must be supplied however.  They are passed as a slice,
// so the length of y is fixed.  For Len3 functions it must be 3 and for
// Len5 functions it must be 5.
package interp

import (
	"errors"
	"math"
)

var (
	ErrorNot3        = errors.New("Argument yTable must be length 3")
	ErrorNoXRange    = errors.New("Argument x3 cannot equal x1")
	ErrorXOutOfRange = errors.New("Argument x outside of range x1 to x3")
	ErrorNoExtremum  = errors.New("No extremum in table")
	ErrorZeroOutside = errors.New("Zero falls outside of table")
	ErrorNoConverge  = errors.New("Failure to converge")
)

// Len3Interpolate interpolates from a table of three rows
// by taking second differences.
//
// Function returns interpolated y value for argument x, as long as x is
// within the table.  X3 must be > x1,
func Len3Interpolate(x, x1, x3 float64, yTable []float64, allowExtrapolate bool) (y float64, err error) {
	// AA p. 24.
	if len(yTable) != 3 {
		return 0, ErrorNot3
	}
	if x3 == x1 {
		return 0, ErrorNoXRange
	}
	if !allowExtrapolate {
		if x3 > x1 {
			// increasing x
			if x < x1 || x > x3 {
				return 0, ErrorXOutOfRange
			}
		} else {
			// decreasing x
			if x > x1 || x < x3 {
				return 0, ErrorXOutOfRange
			}
		}
	}
	a := yTable[1] - yTable[0]
	b := yTable[2] - yTable[1]
	c := b - a
	n := x - (x1+x3)*.5
	return yTable[1] + n*.5*(a+b+n*c), nil
}

// Len3Extremum finds the extremum in a table of 3 rows, interpolating
// by second differences.
//
// It returns the x and y values at the extremum.
func Len3Extremum(x1, x3 float64, yTable []float64) (x, y float64, err error) {
	// AA p. 25.
	if len(yTable) != 3 {
		return 0, 0, ErrorNot3
	}
	if x3 == x1 {
		return 0, 0, ErrorNoXRange
	}
	a := yTable[1] - yTable[0]
	b := yTable[2] - yTable[1]
	c := b - a
	if c == 0 {
		return 0, 0, ErrorNoExtremum
	}
	apb := a + b
	y = yTable[1] - (apb*apb)/(8*c)
	n := apb / (-2 * c)
	x = .5 * ((x3 + x1) + (x3-x1)*n)
	return x, y, nil
}

// Len3Zero finds the first zero in a table of three rows, interpolating
// by second differences.
//
// It returns the x value that yields y=0.
//
// Argument strong switches between two strategies for the estimation step.
// when iterating to converge on the zero.
// Strong=false specifies a quick and dirty estimate that works well
// for gentle curves, but can work poorly or fail on more dramatic curves.
//
// Strong=true specifies a more sophisticated and thus somewhat more
// expensive estimate.  However, if the curve has quick changes, This estimate
// will converge more reliably and in fewer steps, making it a better choice.
func Len3Zero(x1, x3 float64, yTable []float64, strong bool) (x float64, err error) {
	// AA p. 26.
	if len(yTable) != 3 {
		return 0, ErrorNot3
	}
	if x3 == x1 {
		return 0, ErrorNoXRange
	}
	a := yTable[1] - yTable[0]
	b := yTable[2] - yTable[1]
	c := b - a
	var n0, n1 float64
	for limit := 0; limit < 50; limit++ {
		if strong {
			// AA p. 27.
			n1 = n0 - (2*yTable[1]+n0*(a+b+c*n0))/(a+b+2*c*n0)
		} else {
			n1 = -2 * yTable[1] / (a + b + c*n0)
		}
		if math.IsInf(n1, 0) || math.IsNaN(n1) {
			break // failure to converge
		}
		if math.Abs((n1-n0)/n0) < 1e-15 {
			if n1 > 1 || n1 < -1 {
				return 0, ErrorZeroOutside
			}
			return .5 * ((x3 + x1) + (x3-x1)*n1), nil // success
		}
		n0 = n1
	}
	return 0, ErrorNoConverge
}
