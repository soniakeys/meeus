// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Planetposition: Chapter 32 Positions of the Planets
package planetposition

import (
	"math"

	"github.com/soniakeys/meeus/base"
)

// Mercury-Neptune constants suitable for first argument to VSOP87.
const (
	Mercury = iota
	Venus
	Earth
	Mars
	Jupiter
	Saturn
	Uranus
	Neptune
	nPlanets // sad practicality
)

// VSOP87 returns ecliptic position of planets.
//
//	Planet is one of constants Mercury-Neptune.
//	Jde is the jde for which positions are desired.
//
//	L is heliocentric longitude in radians.
//	B is heliocentric latitude in radians.
//	R is heliocentric range in AU.
func VSOP87(planet int, jde float64) (L, B, R float64) {
	vt := appendixIII[planet]
	τ := base.J2000Century(jde) * .1
	cf := make([]float64, 6)
	sum := func(series [][]abc) float64 {
		for x, terms := range series {
			cf[x] = 0
			// sum terms in reverse order to preserve accuracy
			for y := len(terms) - 1; y >= 0; y-- {
				term := &terms[y]
				cf[x] += term.a * math.Cos(term.b+term.c*τ)
			}
		}
		return base.Horner(τ, cf[:len(series)]...)
	}
	L = base.PMod(sum(vt.l)*1e-8, 2*math.Pi)
	B = sum(vt.b) * 1e-8
	R = sum(vt.r) * 1e-8
	return
}
