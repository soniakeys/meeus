// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Binary: Chapter 57, Binary Stars
package binary

import (
	"math"

	"github.com/soniakeys/meeus/base"
)

// M computes mean anomaly for the given date.
//
//	year is a decimal year specifying the date
//	T is time of periastron, as a decimal year
//	P is period of revolution in mean solar years
//
// Result is mean anomaly in radians.
func M(year, T, P float64) float64 {
	n := 2 * math.Pi / P
	return base.PMod(n*(year-T), 2*math.Pi)
}

// Position computes apparent position angle and angular distance of
// components of a binary star.
//
//	a is apparent semimajor axis in arc seconds
//	e is eccentricity of the true orbit
//	i is inclination relative to the line of sight
//	Ω is position angle of the ascending node
//	ω is longitude of periastron
//	E is eccentric anomaly, computed for example with package kepler
//	   and the mean anomaly as returned by function M in this package.
//
// Return value θ is the apparent position angle in radians, ρ is the
// angular distance in arc seconds.
func Position(a, e, i, Ω, ω, E float64) (θ, ρ float64) {
	r := a * (1 - e*math.Cos(E))
	ν := 2 * math.Atan(math.Sqrt((1+e)/(1-e))*math.Tan(E/2))
	sνω, cνω := math.Sincos(ν + ω)
	ci := math.Cos(i)
	num := sνω * ci
	θ = math.Atan2(num, cνω) + Ω
	if θ < 0 {
		θ += 2 * math.Pi
	}
	ρ = r * math.Sqrt(num*num+cνω*cνω)
	return
}

// ApparentEccentricity returns apparent eccenticity of a binary star
// given true orbital elements.
//
//  e is eccentricity of the true orbit
//  i is inclination relative to the line of sight
//  ω is longitude of periastron
func ApparentEccentricity(e, i, ω float64) float64 {
	ci := math.Cos(i)
	sω, cω := math.Sincos(ω)
	A := (1 - e*e*cω*cω) * ci * ci
	B := e * e * sω * cω * ci
	C := 1 - e*e*sω*sω
	d := A - C
	sD := math.Sqrt(d*d + 4*B*B)
	return math.Sqrt(2 * sD / (A + C + sD))
}
