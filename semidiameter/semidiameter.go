// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Semidiameter: Chapter 55, Semidiameters of the Sun, Moon, and Planets.
package semidiameter

import (
	"math"

	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/meeus/parallax"
	"github.com/soniakeys/unit"
)

// Standard semidiameters at unit distance of 1 AU.
var (
	Sun               = unit.AngleFromSec(959.63)
	Mercury           = unit.AngleFromSec(3.36)
	VenusSurface      = unit.AngleFromSec(8.34)
	VenusCloud        = unit.AngleFromSec(8.41)
	Mars              = unit.AngleFromSec(4.68)
	JupiterEquatorial = unit.AngleFromSec(98.44)
	JupiterPolar      = unit.AngleFromSec(92.06)
	SaturnEquatorial  = unit.AngleFromSec(82.73)
	SaturnPolar       = unit.AngleFromSec(73.82)
	Uranus            = unit.AngleFromSec(35.02)
	Neptune           = unit.AngleFromSec(33.50)
	Pluto             = unit.AngleFromSec(2.07)
	Moon              = unit.AngleFromSec(358473400 / base.AU)
)

// Semidiameter returns semidiameter at specified distance.
//
// Δ must be observer-body distance in AU.
func Semidiameter(s0 unit.Angle, Δ float64) unit.Angle {
	return s0.Div(Δ)
}

// SaturnApparentPolar returns apparent polar semidiameter of Saturn
// at specified distance.
//
// Argument Δ must be observer-Saturn distance in AU.  Argument B is
// Saturnicentric latitude of the observer as given by function saturnring.UB()
// for example.
func SaturnApparentPolar(Δ float64, B unit.Angle) unit.Angle {
	k := (SaturnPolar.Rad() / SaturnEquatorial.Rad())
	k = 1 - k*k
	cB := B.Cos()
	return SaturnEquatorial.Mul(math.Sqrt(1-k*cB*cB) / Δ)
}

// MoonTopocentric returns observed topocentric semidiameter of the Moon.
//
//	Δ is distance to Moon in AU.
//	δ is declination of Moon.
//	H is hour angle of Moon.
//	ρsφʹ, ρcφʹ are parallax constants as returned by
//	    globe.Ellipsoid.ParallaxConstants, for example.
func MoonTopocentric(Δ float64, δ unit.Angle, H unit.HourAngle, ρsφʹ, ρcφʹ float64) float64 {
	const k = .272481
	sπ := parallax.Horizontal(Δ).Sin()
	// q computed by (40.6, 40.7) p. 280, ch 40.
	sδ, cδ := δ.Sincos()
	sH, cH := H.Sincos()
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
// horizon.
func MoonTopocentric2(Δ float64, h unit.Angle) unit.Angle {
	return Moon.Mul((1 + h.Sin()*parallax.Horizontal(Δ).Sin()) / Δ)
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
// Result is semidiameter.
func Asteroid(d, Δ float64) unit.Angle {
	return unit.AngleFromSec(.0013788).Mul(d / Δ)
}
