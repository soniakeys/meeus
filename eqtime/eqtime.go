// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Eqtime: Chapter 28, Equation of time.
package eqtime

import (
	"math"

	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/meeus/coord"
	"github.com/soniakeys/meeus/nutation"
	pp "github.com/soniakeys/meeus/planetposition"
	"github.com/soniakeys/meeus/solar"
)

// E computes the "equation of time" for the given JDE.
//
// Parameter e must be a planetposition.V87Planet object for Earth obtained
// with planetposition.LoadPlanet.
//
// Result is equation of time as an hour angle in radians.
func E(jde float64, e *pp.V87Planet) float64 {
	τ := base.J2000Century(jde) * .1
	L0 := l0(τ)
	// code duplicated from solar.ApparentEquatorialVSOP87 so that
	// we can keep Δψ and cε
	s, β, R := solar.TrueVSOP87(e, jde)
	Δψ, Δε := nutation.Nutation(jde)
	a := -20.4898 / 3600 * math.Pi / 180 / R
	λ := s + Δψ + a
	ε := nutation.MeanObliquity(jde) + Δε
	sε, cε := math.Sincos(ε)
	α, _ := coord.EclToEq(λ, β, sε, cε)
	// (28.1) p. 183
	E := L0 - .0057183*math.Pi/180 - α + Δψ*cε
	return base.PMod(E+math.Pi, 2*math.Pi) - math.Pi
}

// (28.2) p. 183
func l0(τ float64) float64 {
	return base.Horner(τ, 280.4664567, 360007.6982779, .03032028,
		1./49931, -1./15300, -1./2000000) * math.Pi / 180
}

// ESmart computes the "equation of time" for the given JDE.
//
// Result is equation of time as an hour angle in radians.
//
// Result is less accurate that E() but the function has the advantage
// of not requiring the V87Planet object.
func ESmart(jde float64) float64 {
	ε := nutation.MeanObliquity(jde)
	t := math.Tan(ε * .5)
	y := t * t
	T := base.J2000Century(jde)
	L0 := l0(T * .1)
	e := solar.Eccentricity(T)
	M := solar.MeanAnomaly(T)
	s2L0, c2L0 := math.Sincos(2 * L0)
	sM := math.Sin(M)
	// (28.3) p. 185
	return y*s2L0 - 2*e*sM + 4*e*y*sM*c2L0 -
		y*y*s2L0*c2L0 - 1.25*e*e*math.Sin(2*M)
}
