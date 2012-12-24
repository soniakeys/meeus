// Copyright 2012 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Interp: Chapter 3, Interpolation.
package interp

import (
	"math"

	"github.com/soniakeys/meeus"
)

// Diff2 interpolates by taking second differences.
//
// AA p. 24.
//
// Argument y must be a table of y-values corresponding to evenly spaced
// increasing x values.  The function panics if Len(y)<3.
// Args x0 and xLast are the x values corresponding to the first and last
// entries in the table.  xLast must be be > x0.
// Function returns interpolated y value for argument x.
//
// The function extrapolates for values of x outside the table.
func Diff2(y []float64, x0, xLast, x float64) float64 {
	xRange := xLast - x0
	iMax := float64(len(y) - 1)
	i2 := int((x-x0)*iMax/xRange + .5)
	switch {
	case i2 < 1:
		i2 = 1
	case i2 > len(y)-2:
		i2 = len(y) - 2
	}
	y2 := y[i2]
	a := y2 - y[i2-1]
	b := y[i2+1] - y2
	c := b - a
	n := x - x0 - float64(i2)*xRange/iMax
	return y2 + (n/2)*(a+b+n*c)
}

// Extremum2 finds and interpolates the first extremum in the table,
// by second differences.
//
// It returns the x and y values at the extremum and ok=true if found,
// ok=false otherwise.
func Extremum2(y []float64, x0, dx float64) (xm, ym float64, ok bool) {
	// AA p. 25.
	y2 := y[1] // y2 name corresponding to book
	dir := meeus.Cmp(y[0], y2)
	for i, y3 := range y[2:] {
		if meeus.Cmp(y2, y3) != dir {
			// found "appropriate part" of table
			a := y2 - y[i]
			b := y3 - y2
			c := b - a
			apb := a + b
			ym = y2 - (apb*apb)/(8*c)
			nm := apb / (-2 * c)
			xm = x0 + (float64(i+1)+nm)*dx
			ok = true
			return
		}
		y2 = y3
	}
	return // fail
}

// Zero2 finds and interpolates the first zero in the table,
// by second differences.
//
// It returns the x and y values at the extremum and ok=true if found,
// ok=false otherwise.
func Zero2(y []float64, x0, dx float64) (xy0 float64, ok bool) {
	// AA p. 26.
	if len(y) < 3 {
		return // fail. table too small
	}
	y1 := y[0]
	if y1 == 0 {
		return x0, true
	}
	y2 := y[1]
	if y2 == 0 {
		return x0 + dx, true
	}
	s0 := math.Signbit(y1)
	var y3 float64
	i := 2
	for {
		y3 = y[i]
		if y3 == 0 {
			return x0 + float64(i)*dx, true
		}
		if math.Signbit(y3) != s0 {
			break
		}
		i++
		if i == len(y) {
			return // fail.  no zero
		}
		y1, y2 = y2, y3
	}
	a := y2 - y1
	b := y3 - y2
	c := b - a
	var n0, n1 float64
	for {
		n1 = -2 * y2 / (a + b + c*n0)
		if math.Abs((n1-n0)/n0) < 1e-15 {
			break
		}
		n0 = n1
	}
	return x0 + (float64(i-1)+n1)*dx, true
}
