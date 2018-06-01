// Copyright 2013 Sonia Keys
// License: MIT

// Illum: Chapter 41, Illuminated Fraction of the Disk and Magnitude of a Planet.
//
// Also see functions Illuminated and Limb in package base.  While this
// chapter title includes "illumnated fraction," the function for computing
// illuminated fraction given a phase angle is base.Illuminated.
// Base.Limb is the function mentioned at the bottom of p. 283.
package illum

import (
	"math"

	"github.com/soniakeys/meeus/v3/base"
	"github.com/soniakeys/unit"
)

// PhaseAngle computes the phase angle of a planet.
//
// Argument r is planet's distance to Sun, Δ its distance to Earth, and R
// the distance from Sun to Earth.  All distances in AU.
func PhaseAngle(r, Δ, R float64) unit.Angle {
	return unit.Angle(math.Acos((r*r + Δ*Δ - R*R) / (2 * r * Δ)))
}

// Fraction computes the illuminated fraction of the disk of a planet.
//
// Argument r is planet's distance to Sun, Δ its distance to Earth, and R
// the distance from Sun to Earth.  All distances in AU.
func Fraction(r, Δ, R float64) float64 {
	// (41.2) p. 283
	s := r + Δ
	return (s*s - R*R) / (4 * r * Δ)
}

// PhaseAngle2 computes the phase angle of a planet.
//
// Arguments L, B, R are heliocentric ecliptical coordinates of the planet.
// L0, R0 are longitude and radius for Earth, Δ is distance from Earth to
// the planet.  All distances in AU.
func PhaseAngle2(L, B unit.Angle, R float64, L0 unit.Angle, R0, Δ float64) unit.Angle {
	// (41.3) p. 283
	return unit.Angle(math.Acos((R - R0*B.Cos()*(L-L0).Cos()) / Δ))
}

// PhaseAngle3 computes the phase angle of a planet.
//
// Arguments L, B are heliocentric ecliptical longitude and latitude of the
// planet.  x, y, z are cartesian coordinates of the planet, Δ is distance
// from Earth to the planet.  All distances in AU.
func PhaseAngle3(L, B unit.Angle, x, y, z, Δ float64) unit.Angle {
	// (41.4) p. 283
	sL, cL := L.Sincos()
	sB, cB := B.Sincos()
	return unit.Angle(math.Acos((x*cB*cL + y*cB*sL + z*sB) / Δ))
}

const p = math.Pi / 180

// FractionVenus computes an approximation of the illumanted fraction of Venus.
func FractionVenus(jde float64) float64 {
	T := base.J2000Century(jde)
	V := 261.51*p + 22518.443*p*T
	M := 177.53*p + 35999.05*p*T
	N := 50.42*p + 58517.811*p*T
	W := V + 1.91*p*math.Sin(M) + .78*p*math.Sin(N)
	Δ := math.Sqrt(1.52321 + 1.44666*math.Cos(W))
	s := .72333 + Δ
	return (s*s - 1) / 2.89332 / Δ
}

// Mercury computes the visual magnitude of Mercury.
//
// Argument r is the planet's distance from the Sun, Δ the distance from Earth,
// and i the phase angle.
func Mercury(r, Δ float64, i unit.Angle) float64 {
	s := i.Deg() - 50
	return 1.16 + 5*math.Log10(r*Δ) + (.02838+.0001023*s)*s
}

// Venus computes the visual magnitude of Venus.
//
// Argument r is the planet's distance from the Sun, Δ the distance from Earth,
// and i the phase angle.
func Venus(r, Δ float64, i unit.Angle) float64 {
	id := i.Deg()
	return -4 + 5*math.Log10(r*Δ) + (.01322+.0000004247*id*id)*id
}

// Mars computes the visual magnitude of Mars.
//
// Argument r is the planet's distance from the Sun, Δ the distance from Earth,
// and i the phase angle.
func Mars(r, Δ float64, i unit.Angle) float64 {
	return -1.3 + 5*math.Log10(r*Δ) + .01486*i.Deg()
}

// Jupiter computes the visual magnitude of Jupiter.
//
// Argument r is the planet's distance from the Sun, Δ the distance from Earth.
func Jupiter(r, Δ float64) float64 {
	return -8.93 + 5*math.Log10(r*Δ)
}

// Saturn computes the visual magnitude of Saturn.
//
// Argument r is the planet's distance from the Sun, Δ the distance from Earth.
// B is the Saturnicentric latitude of the Earth referred to the plane of
// Saturn's ring. ΔU is the difference between the Saturnicentric longitudes
// of the Sun and the Earth, measured in the plane of the ring.
// You can use saturndisk.Disk() to obtain B and ΔU.
func Saturn(r, Δ float64, B, ΔU unit.Angle) float64 {
	s := math.Abs(B.Sin())
	return -8.68 + 5*math.Log10(r*Δ) + .044*math.Abs(ΔU.Deg()) -
		2.6*s + 1.25*s*s
}

// Uranus computes the visual magnitude of Uranus.
//
// Argument r is the planet's distance from the Sun, Δ the distance from Earth.
func Uranus(r, Δ float64) float64 {
	return -6.85 + 5*math.Log10(r*Δ)
}

// Neptune computes the visual magnitude of Neptune.
//
// Argument r is the planet's distance from the Sun, Δ the distance from Earth.
func Neptune(r, Δ float64) float64 {
	return -7.05 + 5*math.Log10(r*Δ)
}

// Mercury84 computes the visual magnitude of Mercury.
//
// The formula is that adopted in "Astronomical Almanac" in 1984.
//
// Argument r is the planet's distance from the Sun, Δ the distance from Earth,
// and i the phase angle.
func Mercury84(r, Δ float64, i unit.Angle) float64 {
	return base.Horner(i.Deg(), -.42+5*math.Log10(r*Δ),
		.038, -.000273, .000002)
}

// Venus84 computes the visual magnitude of Venus.
//
// The formula is that adopted in "Astronomical Almanac" in 1984.
//
// Argument r is the planet's distance from the Sun, Δ the distance from Earth,
// and i the phase angle.
func Venus84(r, Δ float64, i unit.Angle) float64 {
	return base.Horner(i.Deg(), -4.4+5*math.Log10(r*Δ),
		.0009, .000239, -.00000065)
}

// Mars84 computes the visual magnitude of Mars.
//
// The formula is that adopted in "Astronomical Almanac" in 1984.
//
// Argument r is the planet's distance from the Sun, Δ the distance from Earth,
// and i the phase angle.
func Mars84(r, Δ float64, i unit.Angle) float64 {
	return -1.52 + 5*math.Log10(r*Δ) + .016*i.Deg()
}

// Jupiter84 computes the visual magnitude of Jupiter.
//
// The formula is that adopted in "Astronomical Almanac" in 1984.
//
// Argument r is the planet's distance from the Sun, Δ the distance from Earth,
// and i the phase angle.
func Jupiter84(r, Δ float64, i unit.Angle) float64 {
	return -9.4 + 5*math.Log10(r*Δ) + .005*i.Deg()
}

// Saturn84 computes the visual magnitude of Saturn.
//
// The formula is that adopted in "Astronomical Almanac" in 1984.
//
// Argument r is the planet's distance from the Sun, Δ the distance from Earth.
// B is the Saturnicentric latitude of the Earth referred to the plane of
// Saturn's ring. ΔU is the difference between the Saturnicentric longitudes
// of the Sun and the Earth, measured in the plane of the ring.
func Saturn84(r, Δ float64, B, ΔU unit.Angle) float64 {
	s := math.Abs(B.Sin())
	return -8.88 + 5*math.Log10(r*Δ) + .044*math.Abs(ΔU.Deg()) -
		2.6*s + 1.25*s*s
}

// Uranus84 computes the visual magnitude of Uranus.
//
// The formula is that adopted in "Astronomical Almanac" in 1984.
//
// Argument r is the planet's distance from the Sun, Δ the distance from Earth.
func Uranus84(r, Δ float64) float64 {
	return -7.19 + 5*math.Log10(r*Δ)
}

// Neptune84 computes the visual magnitude of Neptune.
//
// The formula is that adopted in "Astronomical Almanac" in 1984.
//
// Argument r is the planet's distance from the Sun, Δ the distance from Earth.
func Neptune84(r, Δ float64) float64 {
	return -6.87 + 5*math.Log10(r*Δ)
}

// Pluto84 computes the visual magnitude of Pluto.
//
// The formula is that adopted in "Astronomical Almanac" in 1984.
//
// Argument r is the planet's distance from the Sun, Δ the distance from Earth.
func Pluto84(r, Δ float64) float64 {
	return -1 + 5*math.Log10(r*Δ)
}
