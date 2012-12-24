// Copyright 2012 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package meeus

// Hints: Chapter 1, Hints and Tips.

import "math"

// DMSToDeg converts from parsed sexagesimal angle components to decimal
// degrees.
func DMSToDeg(neg bool, d, m int, s float64) float64 {
	s = (float64((d*60+m)*60) + s) / 3600
	if neg {
		return -s
	}
	return s
}

// DMSToRad converts from parsed sexagesimal angle components to radians.
//
// Trivially computed from degrees, but here as a one step convenience function
// since computations are normally done in radians.
func DMSToRad(neg bool, d, m int, s float64) float64 {
	return DMSToDeg(neg, d, m, s) * (math.Pi / 180)
}

// HAToRad converts from parsed hour angle components to radians.
//
// One hour = 15 degrees.
func HAToRad(neg bool, h, m int, s float64) float64 {
	return DMSToDeg(neg, h, m, s) * (15 * math.Pi / 180)
}

// HMSToHours converts from parsed sexagesimal time components to decimal
// hours.
func HMSToHours(neg bool, h, m int, s float64) float64 {
	// 60 minutes, 60 seconds, so code is the same as DMSToDeg.
	return DMSToDeg(neg, h, m, s)
}

/*
// RadToDMS is the inverse of DMSToRad.
//
// Note: The results of this function are not suitable for formatting.
// Formatting may round seconds leading to strange results like 1'60".
// Instead use formattable types such as SGAngle, found in package decsym.
func RadToDMS(x float64) (neg bool, d, m int, s float64) {
	if x < 0 {
		neg = true
		x = -x
	}
	x, y := math.Modf(x*60*180/math.Pi) // convert to minutes, then split
	m = int(x)
	return neg, m/60, m%60, y*60
}
*/

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
