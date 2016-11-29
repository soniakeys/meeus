// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Kepler: Chapter 30, Equation of Kepler.
package kepler

import (
	"math"

	"github.com/soniakeys/meeus/iterate"
	"github.com/soniakeys/unit"
)

// True returns true anomaly ν for given eccentric anomaly E.
//
// Argument e is eccentricity.  E must be in radians.
func True(E unit.Angle, e float64) unit.Angle {
	// (30.1) p. 195
	return unit.Angle(2 * math.Atan(math.Sqrt((1+e)/(1-e))*E.Mul(.5).Tan()))
}

// Radius returns radius distance r for given eccentric anomaly E.
//
// Argument e is eccentricity, a is semimajor axis.
//
// Result unit is the unit of semimajor axis a (typically AU.)
func Radius(E unit.Angle, e, a float64) float64 {
	// (30.2) p. 195
	return a * (1 - e*E.Cos())
}

// Kepler1 solves Kepler's equation by iteration.
//
// The iterated formula is
//
//	E1 = M + e * sin(E0)
//
// Argument e is eccentricity, M is mean anomaly,
// places is the desired number of decimal places in the result.
//
// Result E is eccentric anomaly.
//
// For some vaues of e and M it will fail to converge and the
// function will return an error.
func Kepler1(e float64, M unit.Angle, places int) (E unit.Angle, err error) {
	f := func(E0 float64) float64 {
		return M.Rad() + e*math.Sin(E0) // (30.5) p. 195
	}
	ea, err := iterate.DecimalPlaces(f, M.Rad(), places, places*5)
	return unit.Angle(ea), err
}

// Kepler2 solves Kepler's equation by iteration.
//
// The iterated formula is
//
//	E1 = E0 + (M + e * sin(E0) - E0) / (1 - e * cos(E0))
//
// Argument e is eccentricity, M is mean anomaly,
// places is the desired number of decimal places in the result.
//
// Result E is eccentric anomaly.
//
// The function converges over a wider range of inputs than does Kepler1
// but it also fails to converge for some values of e and M.
func Kepler2(e float64, M unit.Angle, places int) (E unit.Angle, err error) {
	f := func(E0 float64) float64 {
		se, ce := math.Sincos(E0)
		return E0 + (M.Rad()+e*se-E0)/(1-e*ce) // (30.7) p. 199
	}
	ea, err := iterate.DecimalPlaces(f, M.Rad(), places, places)
	return unit.Angle(ea), err
}

// Kepler2a solves Kepler's equation by iteration.
//
// The iterated formula is the same as in Kepler2 but a limiting function
// avoids divergence.
//
// Argument e is eccentricity, M is mean anomaly,
// places is the desired number of decimal places in the result.
//
// Result E is eccentric anomaly.
func Kepler2a(e float64, M unit.Angle, places int) (E unit.Angle, err error) {
	f := func(E0 float64) float64 {
		se, ce := math.Sincos(E0)
		// method of Leingärtner, p. 205
		return E0 + math.Asin(math.Sin((M.Rad()+e*se-E0)/(1-e*ce)))
	}
	ea, err := iterate.DecimalPlaces(f, M.Rad(), places, places*5)
	return unit.Angle(ea), err
}

// Kepler2b solves Kepler's equation by iteration.
//
// The iterated formula is the same as in Kepler2 but a (different) limiting
// function avoids divergence.
//
// Argument e is eccentricity, M is mean anomaly,
// places is the desired number of decimal places in the result.
//
// Result E is eccentric anomaly.
func Kepler2b(e float64, M unit.Angle, places int) (E unit.Angle, err error) {
	f := func(E0 float64) float64 {
		se, ce := math.Sincos(E0)
		d := (M.Rad() + e*se - E0) / (1 - e*ce)
		// method of Steele, p. 205
		if d > .5 {
			d = .5
		} else if d < -.5 {
			d = -.5
		}
		return E0 + d
	}
	ea, err := iterate.DecimalPlaces(f, M.Rad(), places, places)
	return unit.Angle(ea), err
}

// Kepler3 solves Kepler's equation by binary search.
//
// Argument e is eccentricity, M is mean anomaly.
//
// Result E is eccentric anomaly.
func Kepler3(e float64, M unit.Angle) (E unit.Angle) {
	// adapted from BASIC, p. 206
	MR := M.Mod1().Rad()
	f := 1
	if MR > math.Pi {
		f = -1
		MR = 2*math.Pi - MR
	}
	E0 := math.Pi * .5
	d := math.Pi * .25
	for i := 0; i < 53; i++ {
		M1 := E0 - e*math.Sin(E0)
		if MR-M1 < 0 {
			E0 -= d
		} else {
			E0 += d
		}
		d *= .5
	}
	if f < 0 {
		E0 = -E0
	}
	return unit.Angle(E0)
}

// Kepler4 returns an approximate solution to Kepler's equation.
//
// It is valid only for small values of e.
//
// Argument e is eccentricity, M is mean anomaly.
//
// Result E is eccentric anomaly.
func Kepler4(e float64, M unit.Angle) (E unit.Angle) {
	sm, cm := M.Sincos()
	return unit.Angle(math.Atan2(sm, cm-e)) // (30.8) p. 206
}
