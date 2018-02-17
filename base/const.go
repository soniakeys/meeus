// Copyright 2013 Sonia Keys
// License: MIT

package base

// K is the Gaussian gravitational constant.
const K = .01720209895

// K from ch 33, p. 228, for example

// AU is one astronomical unit in km.
const AU = 149597870

// from Appendix I, p, 407.

// SOblJ2000, COblJ2000 are sine and cosine of obliquity at J2000.
const (
	SOblJ2000 = .397777156
	COblJ2000 = .917482062
)

// SOblJ2000, COblJ2000 from ch 33, p. 228, for example

// LightTime returns time for light to travel a given distance.
//
// Δ is distance in AU.
//
// Result in days.
func LightTime(Δ float64) float64 {
	// Formula given as (33.3) p. 224.
	return .0057755183 * Δ
}
