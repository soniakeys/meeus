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
var (
	Sun               = base.AngleFromSec(959.63)
	Mercury           = base.AngleFromSec(3.36)
	VenusSurface      = base.AngleFromSec(8.34)
	VenusCloud        = base.AngleFromSec(8.41)
	Mars              = base.AngleFromSec(4.68)
	JupiterEquatorial = base.AngleFromSec(98.44)
	JupiterPolar      = base.AngleFromSec(92.06)
	SaturnEquatorial  = base.AngleFromSec(82.73)
	SaturnPolar       = base.AngleFromSec(73.82)
	Uranus            = base.AngleFromSec(35.02)
	Neptune           = base.AngleFromSec(33.50)
	Pluto             = base.AngleFromSec(2.07)
	Moon              = base.AngleFromSec(358473400 / base.AU)
)

// Semidiameter returns semidiameter at specified distance.
//
// Δ must be observer-body distance in AU.
func Semidiameter(s0 base.Angle, Δ float64) base.Angle {
	return s0.Div(Δ)
}

// SaturnApparentPolar returns apparent polar semidiameter of Saturn
// at specified distance.
//
// Argument Δ must be observer-Saturn distance in AU.  Argument B is
// Saturnicentric latitude of the observer as given by function saturnring.UB()
// for example.
func SaturnApparentPolar(Δ float64, B base.Angle) base.Angle {
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
func MoonTopocentric(Δ float64, δ base.Angle, H base.HourAngle, ρsφʹ, ρcφʹ float64) float64 {
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
func MoonTopocentric2(Δ float64, h base.Angle) base.Angle {
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
func Asteroid(d, Δ float64) base.Angle {
	return base.AngleFromSec(.0013788).Mul(d / Δ)
}
