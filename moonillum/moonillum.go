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
func PhaseAngleEq(α base.RA, δ base.Angle, Δ float64, α0 base.RA, δ0 base.Angle, R float64) base.Angle {
	return pa(Δ, R, cψEq(α, α0, δ, δ0))
}

// cos elongation from equatorial coordinates
func cψEq(α, α0 base.RA, δ, δ0 base.Angle) float64 {
	sδ, cδ := math.Sincos(δ.Rad())
	sδ0, cδ0 := math.Sincos(δ0.Rad())
	return sδ0*sδ + cδ0*cδ*math.Cos((α0-α).Rad())
}

// phase angle from cos elongation and distances
func pa(Δ, R, cψ float64) base.Angle {
	sψ := math.Sin(math.Acos(cψ))
	i := base.Angle(math.Atan(R * sψ / (Δ - R*cψ)))
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
// α0, δ0  are coordinates of the Sun.
func PhaseAngleEq2(α base.RA, δ base.Angle, α0 base.RA, δ0 base.Angle) base.Angle {
	return base.Angle(math.Acos(-cψEq(α, α0, δ, δ0)))
}

// PhaseAngleEcl computes the phase angle of the Moon given ecliptic coordinates.
//
// Arguments λ, β, Δ are geocentric longitude, latitude, and distance
// to the Moon; λ0, R  are longitude and distance to the Sun.
// Distances must be in the same units as each other.
func PhaseAngleEcl(λ, β base.Angle, Δ float64, λ0 base.Angle, R float64) base.Angle {
	return pa(Δ, R, cψEcl(λ, β, λ0))
}

// cos elongation from ecliptic coordinates
func cψEcl(λ, β, λ0 base.Angle) float64 {
	return math.Cos(β.Rad()) * math.Cos((λ - λ0).Rad())
}

// PhaseAngleEcl2 computes the phase angle of the Moon given ecliptic coordinates.
//
// Less accurate than PhaseAngleEcl.
//
// Arguments λ, β are geocentric longitude and latitude of the Moon;
// λ0 is longitude of the Sun.
func PhaseAngleEcl2(λ, β, λ0 base.Angle) base.Angle {
	return base.Angle(math.Acos(-cψEcl(λ, β, λ0)))
}

// PhaseAngle3 computes the phase angle of the Moon given a julian day.
//
// Less accurate than PhaseAngle functions taking coordinates.
func PhaseAngle3(jde float64) base.Angle {
	T := base.J2000Century(jde)
	D := base.AngleFromDeg(base.Horner(T, 297.8501921,
		445267.1114034, -.0018819, 1/545868, -1/113065000)).Mod1().Rad()
	M := base.AngleFromDeg(base.Horner(T,
		357.5291092, 35999.0502909, -.0001535, 1/24490000)).Mod1().Rad()
	Mʹ := base.AngleFromDeg(base.Horner(T, 134.9633964,
		477198.8675055, .0087414, 1/69699, -1/14712000)).Mod1().Rad()
	return math.Pi - base.Angle(D) + base.AngleFromDeg(
		-6.289*math.Sin(Mʹ)+
			2.1*math.Sin(M)+
			-1.274*math.Sin(2*D-Mʹ)+
			-.658*math.Sin(2*D)+
			-.214*math.Sin(2*Mʹ)+
			-.11*math.Sin(D))
}
