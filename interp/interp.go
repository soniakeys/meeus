// Copyright 2012 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Interp: Chapter 3, Interpolation.
//
// Len3 and Len5 functions
//
// These functions interpolate from a table of equidistant x values
// and corresponding y values.  Since the x values are equidistant,
// only the first and last values are supplied as arguments.  The interior
// x values are implicit.
//
// All y values must be supplied however.  They are passed as a slice,
// and the length of y is fixed.  For Len3 functions it must be 3 and for
// Len5 functions it must be 5.
//
// For these Len3 and Len5 functions, Meeus notes the importance of choosing
// the 3 or 5 rows of a larger table that will minimize the interpolating
// factor n.  He does not provide algorithms for doing this however, so none
// are provided here.  For Interpolate, you must use your own algorithm to
// select 3 or 5 rows that will have a middle row with x value closest to
// the argument.  For Extremum and Zero, you must use your own algorithm to
// select 3 or 5 rows that will have a middle row with x value closest to
// the expected result.
package interp

import (
	"errors"
	"math"

	"github.com/soniakeys/meeus"
)

// Error values returned by functions in this package.  Defined here to help
// testing for specific errors.
var (
	ErrorNot3        = errors.New("Argument yTable must be length 3")
	ErrorNot4        = errors.New("Argument yTable must be length 4")
	ErrorNot5        = errors.New("Argument yTable must be length 5")
	ErrorNoXRange    = errors.New("Argument x3 (or x5) cannot equal x1")
	ErrorXOutOfRange = errors.New("Argument x outside of range x1 to x3 (or x5)")
	ErrorNoExtremum  = errors.New("No extremum in table")
	ErrorZeroOutside = errors.New("Zero falls outside of table")
	ErrorNoConverge  = errors.New("Failure to converge")
)

// Len3Interpolate interpolates from a table of three rows
// by taking second differences.
//
// Function returns interpolated y value for argument x, as long as x is
// within the table.
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
	var f func(n0 float64) (n1 float64)
	if strong {
		f = func(n0 float64) float64 {
			return n0 - (2*yTable[1]+n0*(a+b+c*n0))/(a+b+2*c*n0)
		}
	} else {
		f = func(n0 float64) float64 {
			return -2 * yTable[1] / (a + b + c*n0)
		}
	}
	n0, ok := iterate(0, f)
	if !ok {
		return 0, ErrorNoConverge
	}
	if n0 > 1 || n0 < -1 {
		return 0, ErrorZeroOutside
	}
	return .5 * ((x3 + x1) + (x3-x1)*n0), nil // success
}

func iterate(n0 float64, f func(n0 float64) (n1 float64)) (n1 float64, ok bool) {
	for limit := 0; limit < 50; limit++ {
		n1 = f(n0)
		if math.IsInf(n1, 0) || math.IsNaN(n1) {
			break // failure to converge
		}
		if math.Abs((n1-n0)/n0) < 1e-15 {
			return n1, true // success
		}
		n0 = n1
	}
	return 0, false // failure to converge
}

// Len4Half interpolates a center value from a table of four rows.
func Len4Half(yTable []float64) (float64, error) {
	if len(yTable) != 4 {
		return 0, ErrorNot4
	}
	return (9*(yTable[1]+yTable[2]) - yTable[0] - yTable[3]) / 16, nil
}

// Len5Interpolate interpolates from a table of five rows
// by taking fouth differences.
//
// Function returns interpolated y value for argument x, as long as x is
// within the table.
func Len5Interpolate(x, x1, x5 float64, yTable []float64, allowExtrapolate bool) (y float64, err error) {
	// AA p. 28.
	if len(yTable) != 5 {
		return 0, ErrorNot5
	}
	if x5 == x1 {
		return 0, ErrorNoXRange
	}
	if !allowExtrapolate {
		if x5 > x1 {
			// increasing x
			if x < x1 || x > x5 {
				return 0, ErrorXOutOfRange
			}
		} else {
			// decreasing x
			if x > x1 || x < x5 {
				return 0, ErrorXOutOfRange
			}
		}
	}
	l5 := newLen5diffs(yTable)
	n := (4*x - 2*(x1+x5)) / (x5 - x1)
	return l5.eval(n), nil
}

// Len5Extremum finds the extremum in a table of 5 rows, interpolating
// by second differences.
//
// It returns the x and y values at the extremum.
func Len5Extremum(x1, x5 float64, yTable []float64) (x, y float64, err error) {
	// AA p. 29.
	if len(yTable) != 5 {
		return 0, 0, ErrorNot5
	}
	if x5 == x1 {
		return 0, 0, ErrorNoXRange
	}
	l5 := newLen5diffs(yTable)
	nCoeff := []float64{
		6*(l5.B+l5.C) - l5.H - l5.J,
		0,
		3 * (l5.H + l5.K),
		2 * l5.K,
	}
	den := l5.K - 12*l5.F
	n0, ok := iterate(0, func(n0 float64) float64 {
		return meeus.Horner(n0, nCoeff) / den
	})
	if !ok {
		return 0, 0, ErrorNoConverge
	}
	x = .5*(x5+x1) + .25*(x5-x1)*n0
	return x, l5.eval(n0), nil
}

type len5diffs struct {
	y3, A, B, C, D, E, F, G, H, J, K float64
}

func newLen5diffs(y []float64) *len5diffs {
	l5 := new(len5diffs)
	l5.y3 = y[2]

	l5.A = y[1] - y[0]
	l5.B = y[2] - y[1]
	l5.C = y[3] - y[2]
	l5.D = y[4] - y[3]

	l5.E = l5.B - l5.A
	l5.F = l5.C - l5.B
	l5.G = l5.D - l5.C

	l5.H = l5.F - l5.E
	l5.J = l5.G - l5.F

	l5.K = l5.J - l5.H
	return l5
}

func (l5 *len5diffs) eval(n float64) float64 {
	return meeus.Horner(n, []float64{
		l5.y3,
		(l5.B+l5.C)/2 - (l5.H+l5.J)/12,
		l5.F/2 - l5.K/24,
		(l5.H + l5.J) / 12,
		l5.K / 24,
	})
}

// Len5Zero finds the first zero in a table of five rows, interpolating
// by fourth differences.
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
func Len5Zero(x1, x5 float64, yTable []float64, strong bool) (x float64, err error) {
	// AA p. 29.
	if len(yTable) != 5 {
		return 0, ErrorNot5
	}
	if x5 == x1 {
		return 0, ErrorNoXRange
	}
	l5 := newLen5diffs(yTable)
	var f func(n0 float64) (n1 float64)
	if strong {
		M := l5.K / 24
		N := (l5.H + l5.J) / 12
		P := l5.F/2 - M
		Q := (l5.B+l5.C)/2 - N
		numCoeff := []float64{l5.y3, Q, P, N, M}
		denCoeff := []float64{Q, 2 * P, 3 * N, 4 * M}
		f = func(n0 float64) float64 {
			return n0 - meeus.Horner(n0, numCoeff)/meeus.Horner(n0, denCoeff)
		}
	} else {
		numCoeff := []float64{
			-24 * l5.y3,
			0,
			l5.K - 12*l5.F,
			-2 * (l5.H + l5.J),
			-l5.K,
		}
		den := 12*(l5.B+l5.C) - 2*(l5.H+l5.J)
		f = func(n0 float64) float64 {
			return meeus.Horner(n0, numCoeff) / den
		}
	}
	n0, ok := iterate(0, f)
	if !ok {
		return 0, ErrorNoConverge
	}
	if n0 > 2 || n0 < -2 {
		return 0, ErrorZeroOutside
	}
	x = .5*(x5+x1) + .25*(x5-x1)*n0
	return x, nil
}

// Lagrange performs interpolation with unequally-spaced abscissae.
//
// Given a table of X and Y values, interpolate a new y value for argument x.
//
// X values in the table do not have to be equally spaced; they do not even
// have to be in order.  They must however, be distinct.
func Lagrange(x float64, table []struct{ X, Y float64 }) (y float64) {
	// method of BASIC program, p. 33.
	sum := 0.
	for i := range table {
		xi := table[i].X
		prod := 1.
		for j := range table {
			if i != j {
				xj := table[j].X
				prod *= (x - xj) / (xi - xj)
			}
		}
		sum += table[i].Y * prod
	}
	return sum
}

// LagrangePoly uses the formula of Lagrange to produce an interpolating
// polynomial.
//
// X values in the table do not have to be equally spaced; they do not even
// have to be in order.  They must however, be distinct.
//
// The returned polynomial will be of degree n-1 where n is the number of rows
// in the table.  It can be evaluated for x using meeus.Horner.
func LagrangePoly(table []struct{ X, Y float64 }) []float64 {
	// Method not fully described by Meeus, but needed for numerical solution
	// to Example 3.g.
	sum := make([]float64, len(table))
	prod := make([]float64, len(table))
	last := len(table) - 1
	for i := range table {
		xi := table[i].X
		yi := table[i].Y
		prod[last] = 1
		den := 1.
		n := last
		for j := range table {
			if i != j {
				xj := table[j].X
				prod[n-1] = prod[n] * -xj
				for k := n; k < last; k++ {
					prod[k] -= prod[k+1] * xj
				}
				n--
				den *= (xi - xj)
			}
		}
		for j, pj := range prod {
			sum[j] += yi * pj / den
		}
	}
	return sum
}
