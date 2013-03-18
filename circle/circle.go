// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Circle: Chapter 20, Smallest Circle containing three Celestial Bodies.
package circle

import (
	"math"
)

func hav(a float64) float64 {
	return .5 * (1 - math.Cos(a))
}

// Smallest finds the smallest circle containing three points.
//
// Arguments should represent coordinates in right ascension and declination
// or longitude and latitude.  Result Δ is the diameter of the circle, typeI
// is true if solution is of type I.
//
//	type I   Two points on circle, one interior.
//	type II  All three points on circle.
func Smallest(r1, d1, r2, d2, r3, d3 float64) (Δ float64, typeI bool) {
	// Using haversine formula
	cd1 := math.Cos(d1)
	cd2 := math.Cos(d2)
	cd3 := math.Cos(d3)
	a := 2 * math.Asin(math.Sqrt(hav(d2-d1)+cd1*cd2*hav(r2-r1)))
	b := 2 * math.Asin(math.Sqrt(hav(d3-d2)+cd2*cd3*hav(r3-r2)))
	c := 2 * math.Asin(math.Sqrt(hav(d1-d3)+cd3*cd1*hav(r1-r3)))
	if b > a {
		a, b = b, a
	}
	if c > a {
		a, c = c, a
	}
	if a*a >= b*b+c*c {
		return a, true
	}
	return 2 * a * b * c / math.Sqrt((a+b+c)*(a+b-c)*(b+c-a)*(a+c-b)), false
}
