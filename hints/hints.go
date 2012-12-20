// Copyright 2012 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Hints has algorithms from Chapter 1, Hints and Tips.
package hints

import "math"

// DMSToDeg converts from parsed sexagesimal degree components to degrees.
func DMSToDeg(neg bool, d, m int, s float64) float64 {
	s = (float64((d*60+m)*60) + s) / 3600
	if neg {
		return -s
	}
	return s
}

// DMSToRad converts from parsed sexagesimal degree components to radians.
func DMSToRad(neg bool, d, m int, s float64) float64 {
	return DMSToDeg(neg, d, m, s) * (math.Pi / 180)
}

// HMSToDeg converts from parsed right ascension components to degrees.
func HMSToDeg(neg bool, h, m int, s float64) float64 {
	return DMSToDeg(neg, h, m, s) * 15
}

// HMSToRad converts from parsed right ascension components to radians.
func HMSToRad(neg bool, h, m int, s float64) float64 {
	return DMSToDeg(neg, h, m, s) * (15 * math.Pi / 180)
}

// Horner evaluates a polynomal with coefficients c at x.  The constant
// term is c[0].  The function panics with an empty coefficient list.
func Horner(c []float64, x float64) float64 {
	i := len(c) - 1
	y := c[i]
	for i > 0 {
		i--
		y = y*x + c[i] // sorry, no fused multiply-add in Go
	}
	return y
}
