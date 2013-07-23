// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Kepler: Chapter 30, Equation of Kepler.
package kepler

import (
	"math"

	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/meeus/iterate"
)

// True returns true anomaly ν for given eccentric anomaly E.
//
// Argument e is eccentricity.  E must be in radians.
//
// Result is in radians.
func True(E, e float64) float64 {
	// (30.1) p. 195
	return 2 * math.Atan(math.Sqrt((1+e)/(1-e))*math.Tan(E*.5))
}

// Radius returns radius distance r for given eccentric anomaly E.
//
// Argument e is eccentricity, a is semimajor axis.
//
// Result unit is the unit of semimajor axis a (typically AU.)
func Radius(E, e, a float64) float64 {
	// (30.2) p. 195
	return a * (1 - e*math.Cos(E))
}

// Kepler1 solves Kepler's equation by iteration.
//
// The iterated formula is
//
//	E1 = M + e * sin(E0)
//
// Argument e is eccentricity, M is mean anomoly in radians,
// places is the desired number of decimal places in the result.
//
// Result E is eccentric anomaly in radians.
//
// For some vaues of e and M it will fail to converge and the
// function will return an error.
func Kepler1(e, M float64, places int) (E float64, err error) {
	f := func(E0 float64) float64 {
		return M + e*math.Sin(E0) // (30.5) p. 195
	}
	return iterate.DecimalPlaces(f, M, places, places*5)
}

// Kepler2 solves Kepler's equation by iteration.
//
// The iterated formula is
//
//	E1 = E0 + (M + e * sin(E0) - E0) / (1 - e * cos(E0))
//
// Argument e is eccentricity, M is mean anomoly in radians,
// places is the desired number of decimal places in the result.
//
// Result E is eccentric anomaly in radians.
//
// The function converges over a wider range of inputs than does Kepler1
// but it also fails to converge for some values of e and M.
func Kepler2(e, M float64, places int) (E float64, err error) {
	f := func(E0 float64) float64 {
		se, ce := math.Sincos(E0)
		return E0 + (M+e*se-E0)/(1-e*ce) // (30.7) p. 199
	}
	return iterate.DecimalPlaces(f, M, places, places)
}

// Kepler2a solves Kepler's equation by iteration.
//
// The iterated formula is the same as in Kepler2 but a limiting function
// avoids divergence.
//
// Argument e is eccentricity, M is mean anomoly in radians,
// places is the desired number of decimal places in the result.
//
// Result E is eccentric anomaly in radians.
func Kepler2a(e, M float64, places int) (E float64, err error) {
	f := func(E0 float64) float64 {
		se, ce := math.Sincos(E0)
		// method of Leingärtner, p. 205
		return E0 + math.Asin(math.Sin((M+e*se-E0)/(1-e*ce)))
	}
	return iterate.DecimalPlaces(f, M, places, places*5)
}

// Kepler2b solves Kepler's equation by iteration.
//
// The iterated formula is the same as in Kepler2 but a (different) limiting
// function avoids divergence.
//
// Argument e is eccentricity, M is mean anomoly in radians,
// places is the desired number of decimal places in the result.
//
// Result E is eccentric anomaly in radians.
func Kepler2b(e, M float64, places int) (E float64, err error) {
	f := func(E0 float64) float64 {
		se, ce := math.Sincos(E0)
		d := (M + e*se - E0) / (1 - e*ce)
		// method of Steele, p. 205
		if d > .5 {
			d = .5
		} else if d < -.5 {
			d = -.5
		}
		return E0 + d
	}
	return iterate.DecimalPlaces(f, M, places, places)
}

// Kepler3 solves Kepler's equation by binary search.
//
// Argument e is eccentricity, M is mean anomoly in radians.
//
// Result E is eccentric anomaly in radians.
func Kepler3(e, M float64) (E float64) {
	// adapted from BASIC, p. 206
	M = base.PMod(M, 2*math.Pi)
	f := 1
	if M > math.Pi {
		f = -1
		M = 2*math.Pi - M
	}
	E0 := math.Pi * .5
	d := math.Pi * .25
	for i := 0; i < 53; i++ {
		M1 := E0 - e*math.Sin(E0)
		if M-M1 < 0 {
			E0 -= d
		} else {
			E0 += d
		}
		d *= .5
	}
	if f < 0 {
		return -E0
	}
	return E0
}

// Kepler4 returns an approximate solution to Kepler's equation.
//
// It is valid only for small values of e.
//
// Argument e is eccentricity, M is mean anomoly in radians.
//
// Result E is eccentric anomaly in radians.
func Kepler4(e, M float64) (E float64) {
	sm, cm := math.Sincos(M)
	return math.Atan2(sm, cm-e) // (30.8) p. 206
}
