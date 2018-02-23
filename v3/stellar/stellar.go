// Copyright 2013 Sonia Keys
// License: MIT

// Stellar: Chapter 56, Stellar Magnitudes.
package stellar

import (
	"math"

	"github.com/soniakeys/unit"
)

// Sum returns the combined apparent magnitude of two stars.
func Sum(m1, m2 float64) float64 {
	x := .4 * (m2 - m1)
	return m2 - 2.5*math.Log10(math.Pow(10, x)+1)
}

// SumN returns the combined apparent magnitude of a number of stars.
func SumN(m ...float64) float64 {
	var s float64
	for _, mi := range m {
		s += math.Pow(10, -.4*mi)
	}
	return -2.5 * math.Log10(s)
}

// Ratio returns the brightness ratio of two stars.
//
// Arguments m1, m2 are apparent magnitudes.
func Ratio(m1, m2 float64) float64 {
	x := .4 * (m2 - m1)
	return math.Pow(10, x)
}

// Difference returns the difference in apparent magnitude of two stars
// given their brightness ratio.
func Difference(ratio float64) float64 {
	return 2.5 * math.Log10(ratio)
}

// AbsoluteByParallax returns absolute magnitude given annual parallax.
//
// Argument m is apparent magnitude, π is annual parallax.
func AbsoluteByParallax(m float64, π unit.Angle) float64 {
	return m + 5 + 5*math.Log10(π.Sec())
}

// AbsoluteByDistance returns absolute magnitude given distance.
//
// Argument m is apparent magnitude, d is distance in parsecs.
func AbsoluteByDistance(m, d float64) float64 {
	return m + 5 - 5*math.Log10(d)
}
