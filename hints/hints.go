// Copyright 2012 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Hints: Chapter 1, Hints and Tips.
package hints

// Horner evaluates a polynomal with coefficients c at x.  The constant
// term is c[0].  The function panics with an empty coefficient list.
func Horner(x float64, c []float64) float64 {
	i := len(c) - 1
	y := c[i]
	for i > 0 {
		i--
		y = y*x + c[i] // sorry, no fused multiply-add in Go
	}
	return y
}
