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
// Result is mean anomaly.
func M(year, T, P float64) base.Angle {
	n := 2 * math.Pi / P
	return base.Angle(n * (year - T)).Mod1()
}

// Position computes apparent position angle and angular distance of
// components of a binary star.
//
//	e is eccentricity of the true orbit
//	a is angular apparent semimajor axis
//	i is inclination relative to the line of sight
//	Ω is position angle of the ascending node
//	ω is longitude of periastron
//	E is eccentric anomaly, computed for example with package kepler
//	   and the mean anomaly as returned by function M in this package.
//
// Return value θ is the apparent position angle, ρ is the angular distance.
func Position(e float64, a, i, Ω, ω, E base.Angle) (θ, ρ base.Angle) {
	r := a.Mul(1 - e*math.Cos(E.Rad()))
	ν := 2 * math.Atan(math.Sqrt((1+e)/(1-e))*math.Tan(E.Rad()/2))
	sνω, cνω := math.Sincos(ν + ω.Rad())
	ci := math.Cos(i.Rad())
	num := sνω * ci
	θ = (base.Angle(math.Atan2(num, cνω)) + Ω).Mod1()
	ρ = r.Mul(math.Sqrt(num*num + cνω*cνω))
	return
}

// ApparentEccentricity returns apparent eccenticity of a binary star
// given true orbital elements.
//
//  e is eccentricity of the true orbit
//  i is inclination relative to the line of sight
//  ω is longitude of periastron
func ApparentEccentricity(e float64, i, ω base.Angle) float64 {
	ci := math.Cos(i.Rad())
	sω, cω := math.Sincos(ω.Rad())
	A := (1 - e*e*cω*cω) * ci * ci
	B := e * e * sω * cω * ci
	C := 1 - e*e*sω*sω
	d := A - C
	sD := math.Sqrt(d*d + 4*B*B)
	return math.Sqrt(2 * sD / (A + C + sD))
}
