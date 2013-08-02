// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package base

// K is the Gaussian gravitational constant.
//
// From ch 33, p. 228, for example
const K = .01720209895

// SOblJ2000, COblJ2000 are sine and cosine of obliquity at J2000.
//
// From ch 33, p. 228, for example
const SOblJ2000 = .397777156
const COblJ2000 = .917482062

// LightTime returns time for light to travel a given distance.
//
// Formula given as (33.3) p. 224.
//
// Δ is distance in AU.
//
// Result in seconds of time.
func LightTime(Δ float64) float64 {
	return .0057755183 * Δ
}
