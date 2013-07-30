// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Node: Chapter 39, Passages through the Nodes.
package node

import (
	"math"

	"github.com/soniakeys/meeus/base"
)

// EllipticAscending computes time and distance of passage through the ascending node of a body in an elliptical orbit.
//
// Argument axis is semimajor axis in AU, ecc is eccentricity, argP is argument
// of perihelion in radians, timeP is time of perihelion as a jd.
//
// Result is jde of the event and distance from the sun in AU.
func EllipticAscending(axis, ecc, argP, timeP float64) (jde, r float64) {
	return el(-argP, axis, ecc, timeP)
}

// EllipticAscending computes time and distance of passage through the descending node of a body in an elliptical orbit.
//
// Argument axis is semimajor axis in AU, ecc is eccentricity, argP is argument
// of perihelion in radians, timeP is time of perihelion as a jd.
//
// Result is jde of the event and distance from the sun in AU.
func EllipticDescending(axis, ecc, argP, timeP float64) (jde, r float64) {
	return el(math.Pi-argP, axis, ecc, timeP)
}

func el(ν, axis, ecc, timeP float64) (jde, r float64) {
	E := 2 * math.Atan(math.Sqrt((1-ecc)/(1+ecc))*math.Tan(ν*.5))
	sE, cE := math.Sincos(E)
	M := E - ecc*sE
	n := base.K / axis / math.Sqrt(axis)
	jde = timeP + M/n
	r = axis * (1 - ecc*cE)
	return
}

// ParabolicAscending computes time and distance of passage through the ascending node of a body in a parabolic orbit.
//
// Argument q is perihelion distance in AU, argP is argument of perihelion
// in radians, timeP is time of perihelion as a jd.
//
// Result is jde of the event and distance from the sun in AU.
func ParabolicAscending(q, argP, timeP float64) (jde, r float64) {
	return pa(-argP, q, timeP)
}

// ParabolicDescending computes time and distance of passage through the descending node of a body in a parabolic orbit.
//
// Argument q is perihelion distance in AU, argP is argument of perihelion
// in radians, timeP is time of perihelion as a jd.
//
// Result is jde of the event and distance from the sun in AU.
func ParabolicDescending(q, argP, timeP float64) (jde, r float64) {
	return pa(math.Pi-argP, q, timeP)
}

func pa(ν, q, timeP float64) (jde, r float64) {
	s := math.Tan(ν * .5)
	jde = timeP + 27.403895*s*(s*s+3)*q*math.Sqrt(q)
	r = q * (1 + s*s)
	return
}
