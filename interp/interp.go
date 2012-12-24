// Copyright 2012 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Interp: Chapter 3, Interpolation.
package interp

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

func sgn(a, b float64) int {
	switch {
	case b > a:
		return 1
	case b < a:
		return -1
	}
	return 0
}

// Extremum2 finds the first extremum in the table, by second differences.
//
// AA p. 25.
//
// It returns the x and y values at the extremum and ok=true if found,
// ok=false otherwise.
func Extremum2(y []float64, x0, dx float64) (xm, ym float64, ok bool) {
	y2 := y[1] // y2 name corresponding to book
	dir := sgn(y[0], y2)
	for i, y3 := range y[2:] {
		if sgn(y2, y3) != dir {
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

/*
func Zero2(y []float64, x0, dx float64) (xy0 float64, ok bool) {
	if y[0] == 0 {
		return x0, true
	}
	neg := y[0] < 0
	for i, y := range y[1:] {
		if y == 0 {
			return x0+float64(i+1)*dx, true
	}
	y2 := y[1]
*/
