// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Moonillum: Chapter 48, Illuminated Fraction of the Moon's Disk
//
// Also see functions Illuminated and Limb in package base.  The function
// for computing illuminated fraction given a phase angle (48.1) is
// base.Illuminated.  Formula (48.3) is implemented as base.Limb.
package moonillum

import (
	"math"

	"github.com/soniakeys/meeus/base"
)

// PhaseAngleEQ computes the phase angle of the Moon given equatorial coordinates.
//
// Arguments α, δ, Δ are geocentric right ascension, declination, and distance
// to the Moon; α0, δ0, R  are coordinates of the Sun.  Angles must be in
// radians.  Distances must be in the same units as each other.
//
// Result in radians.
func PhaseAngleEq(α, δ, Δ, α0, δ0, R float64) float64 {
	return pa(Δ, R, cψEq(α, δ, α0, δ0))
}

// cos elongation from equatorial coordinates
func cψEq(α, δ, α0, δ0 float64) float64 {
	sδ, cδ := math.Sincos(δ)
	sδ0, cδ0 := math.Sincos(δ0)
	return sδ0*sδ + cδ0*cδ*math.Cos(α0-α)
}

// phase angle from cos elongation and distances
func pa(Δ, R, cψ float64) float64 {
	sψ := math.Sin(math.Acos(cψ))
	i := math.Atan(R * sψ / (Δ - R*cψ))
	if i < 0 {
		i += math.Pi
	}
	return i
}

// PhaseAngleEQ2 computes the phase angle of the Moon given equatorial coordinates.
//
// Less accurate than PhaseAngleEq.
//
// Arguments α, δ are geocentric right ascension and declination of the Moon;
// α0, δ0  are coordinates of the Sun.  Angles must be in radians.
//
// Result in radians.
func PhaseAngleEq2(α, δ, α0, δ0 float64) float64 {
	return math.Acos(-cψEq(α, δ, α0, δ0))
}

// PhaseAngleEcl computes the phase angle of the Moon given ecliptic coordinates.
//
// Arguments λ, β, Δ are geocentric longitude, latitude, and distance
// to the Moon; λ0, R  are longitude and distance to the Sun.  Angles must be
// in radians.  Distances must be in the same units as each other.
//
// Result in radians.
func PhaseAngleEcl(λ, β, Δ, λ0, R float64) float64 {
	return pa(Δ, R, cψEcl(λ, β, λ0))
}

// cos elongation from ecliptic coordinates
func cψEcl(λ, β, λ0 float64) float64 {
	return math.Cos(β) * math.Cos(λ-λ0)
}

// PhaseAngleEcl2 computes the phase angle of the Moon given ecliptic coordinates.
//
// Less accurate than PhaseAngleEcl.
//
// Arguments λ, β are geocentric longitude and latitude of the Moon;
// λ0 is longitude of the Sun.  Angles must be in radians.
//
// Result in radians.
func PhaseAngleEcl2(λ, β, λ0 float64) float64 {
	return math.Acos(-cψEcl(λ, β, λ0))
}

// PhaseAngle3 computes the phase angle of the Moon given a julian day.
//
// Less accurate than PhaseAngle functions taking coordinates.
//
// Result in radians.
func PhaseAngle3(jde float64) float64 {
	T := base.J2000Century(jde)
	const p = math.Pi / 180
	D := base.Horner(T, 297.8501921*p, 445267.1114034*p,
		-.0018819*p, p/545868, -p/113065000)
	M := base.Horner(T, 357.5291092*p, 35999.0502909*p,
		-.0001535*p, p/24490000)
	Mʹ := base.Horner(T, 134.9633964*p, 477198.8675055*p,
		.0087414*p, p/69699, -p/14712000)
	return math.Pi - base.PMod(D, 2*math.Pi) +
		-6.289*p*math.Sin(Mʹ) +
		2.1*p*math.Sin(M) +
		-1.274*p*math.Sin(2*D-Mʹ) +
		-.658*p*math.Sin(2*D) +
		-.214*p*math.Sin(2*Mʹ) +
		-.11*p*math.Sin(D)
}
