// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Semidiameter: Chapter 55, Semidiameters of the Sun, Moon, and Planets.
package semidiameter

import (
	"math"

	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/meeus/parallax"
)

// Standard semidiameters at unit distance of 1 AU.
// Values are scaled here to radians.
var (
	Sun               = 959.63 / 3600 * math.Pi / 180
	Mercury           = 3.36 / 3600 * math.Pi / 180
	VenusSurface      = 8.34 / 3600 * math.Pi / 180
	VenusCloud        = 8.41 / 3600 * math.Pi / 180
	Mars              = 4.68 / 3600 * math.Pi / 180
	JupiterEquatorial = 98.44 / 3600 * math.Pi / 180
	JupiterPolar      = 92.06 / 3600 * math.Pi / 180
	SaturnEquatorial  = 82.73 / 3600 * math.Pi / 180
	SaturnPolar       = 73.82 / 3600 * math.Pi / 180
	Uranus            = 35.02 / 3600 * math.Pi / 180
	Neptune           = 33.50 / 3600 * math.Pi / 180
	Pluto             = 2.07 / 3600 * math.Pi / 180
	Moon              = 358473400 / base.AU / 3600 * math.Pi / 180
)

// Semidiameter returns semidiameter at specified distance.
//
// When used with S0 values provided, Δ must be observer-body distance in AU.
// Result will then be in radians.
func Semidiameter(s0, Δ float64) float64 {
	return s0 / Δ
}

// SaturnApparentPolar returns apparent polar semidiameter of Saturn
// at specified distance.
//
// Argument Δ must be observer-Saturn distance in AU.  Argument B is
// Saturnicentric latitude of the observer as given by function saturnring.UB()
// for example.
//
// Result is semidiameter in units of package variables SaturnPolar and
// SaturnEquatorial, nominally radians.
func SaturnApparentPolar(Δ, B float64) float64 {
	k := SaturnPolar / SaturnEquatorial
	k = 1 - k*k
	cB := math.Cos(B)
	return SaturnEquatorial / Δ * math.Sqrt(1-k*cB*cB)
}

// MoonTopocentric returns observed topocentric semidiameter of the Moon.
//
//	Δ is distance to Moon in AU.
//	δ is declination of Moon in radians.
//	H is hour angle of Moon in radians.
//	ρsφʹ, ρcφʹ are parallax constants as returned by
//	    globe.Ellipsoid.ParallaxConstants, for example.
//
// Result is semidiameter in radians.
func MoonTopocentric(Δ, δ, H, ρsφʹ, ρcφʹ float64) float64 {
	const k = .272481
	sπ := math.Sin(parallax.Horizontal(Δ))
	// q computed by (40.6, 40.7) p. 280, ch 40.
	sδ, cδ := math.Sincos(δ)
	sH, cH := math.Sincos(H)
	A := cδ * sH
	B := cδ*cH - ρcφʹ*sπ
	C := sδ - ρsφʹ*sπ
	q := math.Sqrt(A*A + B*B + C*C)
	return k / q * sπ
}

// MoonTopocentric2 returns observed topocentric semidiameter of the Moon
// by a less rigorous method.
//
// Δ is distance to Moon in AU, h is altitude of the Moon above the observer's
// horizon in radians.
//
// Result is semidiameter in radians.
func MoonTopocentric2(Δ, h float64) float64 {
	return Moon / Δ * (1 + math.Sin(h)*math.Sin(parallax.Horizontal(Δ)))
}

// AsteroidDiameter returns approximate diameter given absolute magnitude H
// and albedo A.
//
// Result is in km.
func AsteroidDiameter(H, A float64) float64 {
	return math.Pow(10, 3.12-.2*H-.5*math.Log10(A))
}

// Asteroid returns semidiameter of an asteroid with a given diameter
// at given distance.
//
// Argument d is diameter in km, Δ is distance in AU.
//
// Result is semidiameter in radians.
func Asteroid(d, Δ float64) float64 {
	return .0013788 * d / Δ / 3600 * math.Pi / 180
}
