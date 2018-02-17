// Copyright 2013 Sonia Keys
// License: MIT

// Circle: Chapter 20, Smallest Circle containing three Celestial Bodies.
package circle

import (
	"math"

	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/unit"
)

// Smallest finds the smallest circle containing three points.
//
// Arguments should represent coordinates in right ascension and declination
// or longitude and latitude.  Result Δ is the diameter of the circle, typeI
// is true if solution is of type I.
//
//	type I   Two points on circle, one interior.
//	type II  All three points on circle.
func Smallest(r1, d1, r2, d2, r3, d3 unit.Angle) (Δ unit.Angle, typeI bool) {
	// Using haversine formula, but reimplementing SepHav here to reuse
	// the computed cosines.
	cd1 := d1.Cos()
	cd2 := d2.Cos()
	cd3 := d3.Cos()
	a := 2 * math.Asin(math.Sqrt(base.Hav(d2-d1)+cd1*cd2*base.Hav(r2-r1)))
	b := 2 * math.Asin(math.Sqrt(base.Hav(d3-d2)+cd2*cd3*base.Hav(r3-r2)))
	c := 2 * math.Asin(math.Sqrt(base.Hav(d1-d3)+cd3*cd1*base.Hav(r1-r3)))
	if b > a {
		a, b = b, a
	}
	if c > a {
		a, c = c, a
	}
	if a*a >= b*b+c*c {
		return unit.Angle(a), true
	}
	// (20.1) p. 128
	return unit.Angle(2 * a * b * c /
		math.Sqrt((a+b+c)*(a+b-c)*(b+c-a)*(a+c-b))), false
}
