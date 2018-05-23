// Copyright 2013 Sonia Keys
// License: MIT

// Nutation: Chapter 22, Nutation and the Obliquity of the Ecliptic.
package nutation

import (
	"math"

	"github.com/soniakeys/meeus/v3/base"
	"github.com/soniakeys/unit"
)

// Nutation returns nutation in longitude (Δψ) and nutation in obliquity (Δε)
// for a given JDE.
//
// JDE = UT + ΔT, see package deltat.
//
// Computation is by 1980 IAU theory, with terms < .0003″ neglected.
func Nutation(jde float64) (Δψ, Δε unit.Angle) {
	T := base.J2000Century(jde)
	D := base.Horner(T,
		297.85036, 445267.11148, -0.0019142, 1./189474) * math.Pi / 180
	M := base.Horner(T,
		357.52772, 35999.050340, -0.0001603, -1./300000) * math.Pi / 180
	N := base.Horner(T,
		134.96298, 477198.867398, 0.0086972, 1./56250) * math.Pi / 180
	F := base.Horner(T,
		93.27191, 483202.017538, -0.0036825, 1./327270) * math.Pi / 180
	Ω := base.Horner(T,
		125.04452, -1934.136261, 0.0020708, 1./450000) * math.Pi / 180
	// sum in reverse order to accumulate smaller terms first
	var Δψs, Δεs float64
	for i := len(table22A) - 1; i >= 0; i-- {
		row := table22A[i]
		arg := row.d*D + row.m*M + row.n*N + row.f*F + row.ω*Ω
		s, c := math.Sincos(arg)
		Δψs += s * (row.s0 + row.s1*T)
		Δεs += c * (row.c0 + row.c1*T)
	}
	Δψ = unit.AngleFromSec(Δψs * .0001)
	Δε = unit.AngleFromSec(Δεs * .0001)
	return
}

// ApproxNutation returns a fast approximation of nutation in longitude (Δψ)
// and nutation in obliquity (Δε) for a given JDE.
//
// Accuracy is 0.5″ in Δψ, 0.1″ in Δε.
func ApproxNutation(jde float64) (Δψ, Δε unit.Angle) {
	T := (jde - base.J2000) / 36525
	Ω := (125.04452 - 1934.136261*T) * math.Pi / 180
	L := (280.4665 + 36000.7698*T) * math.Pi / 180
	N := (218.3165 + 481267.8813*T) * math.Pi / 180
	sΩ, cΩ := math.Sincos(Ω)
	s2L, c2L := math.Sincos(2 * L)
	s2N, c2N := math.Sincos(2 * N)
	s2Ω, c2Ω := math.Sincos(2 * Ω)
	Δψ = unit.AngleFromSec(-17.2*sΩ - 1.32*s2L - 0.23*s2N + 0.21*s2Ω)
	Δε = unit.AngleFromSec(9.2*cΩ + 0.57*c2L + 0.1*c2N - 0.09*c2Ω)
	return
}

// MeanObliquity returns mean obliquity (ε₀) following the IAU 1980
// polynomial.
//
// Accuracy is 1″ over the range 1000 to 3000 years and 10″ over the range
// 0 to 4000 years.
func MeanObliquity(jde float64) unit.Angle {
	// (22.2) p. 147
	return unit.AngleFromSec(base.Horner(base.J2000Century(jde),
		unit.FromSexaSec(' ', 23, 26, 21.448),
		-46.815,
		-0.00059,
		0.001813))
}

// MeanObliquityLaskar returns mean obliquity (ε₀) following the Laskar
// 1986 polynomial.
//
// Accuracy over the range 1000 to 3000 years is .01″.
//
// Accuracy over the valid date range of -8000 to +12000 years is
// "a few seconds."
func MeanObliquityLaskar(jde float64) unit.Angle {
	// (22.3) p. 147
	return unit.AngleFromSec(base.Horner(base.J2000Century(jde)*.01,
		unit.FromSexaSec(' ', 23, 26, 21.448),
		-4680.93,
		-1.55,
		1999.25,
		-51.38,
		-249.67,
		-39.05,
		7.12,
		27.87,
		5.79,
		2.45))
}

// NutationInRA returns "nutation in right ascension" or "equation of the
// equinoxes."
func NutationInRA(jde float64) unit.HourAngle {
	// ch 12, p.88
	Δψ, Δε := Nutation(jde)
	ε0 := MeanObliquity(jde)
	return unit.HourAngle(Δψ.Rad() * math.Cos((ε0 + Δε).Rad()))
}

var table22A = []struct {
	d, m, n, f, ω  float64
	s0, s1, c0, c1 float64
}{
	{0, 0, 0, 0, 1, -171996, -174.2, 92025, 8.9},
	{-2, 0, 0, 2, 2, -13187, -1.6, 5736, -3.1},
	{0, 0, 0, 2, 2, -2274, -0.2, 977, -0.5},
	{0, 0, 0, 0, 2, 2062, 0.2, -895, 0.5},
	{0, 1, 0, 0, 0, 1426, -3.4, 54, -0.1},
	{0, 0, 1, 0, 0, 712, 0.1, -7, 0},
	{-2, 1, 0, 2, 2, -517, 1.2, 224, -0.6},
	{0, 0, 0, 2, 1, -386, -0.4, 200, 0},
	{0, 0, 1, 2, 2, -301, 0, 129, -0.1},
	{-2, -1, 0, 2, 2, 217, -0.5, -95, 0.3},
	{-2, 0, 1, 0, 0, -158, 0, 0, 0},
	{-2, 0, 0, 2, 1, 129, 0.1, -70, 0},
	{0, 0, -1, 2, 2, 123, 0, -53, 0},
	{2, 0, 0, 0, 0, 63, 0, 0, 0},
	{0, 0, 1, 0, 1, 63, 0.1, -33, 0},
	{2, 0, -1, 2, 2, -59, 0, 26, 0},
	{0, 0, -1, 0, 1, -58, -0.1, 32, 0},
	{0, 0, 1, 2, 1, -51, 0, 27, 0},
	{-2, 0, 2, 0, 0, 48, 0, 0, 0},
	{0, 0, -2, 2, 1, 46, 0, -24, 0},
	{2, 0, 0, 2, 2, -38, 0, 16, 0},
	{0, 0, 2, 2, 2, -31, 0, 13, 0},
	{0, 0, 2, 0, 0, 29, 0, 0, 0},
	{-2, 0, 1, 2, 2, 29, 0, -12, 0},
	{0, 0, 0, 2, 0, 26, 0, 0, 0},
	{-2, 0, 0, 2, 0, -22, 0, 0, 0},
	{0, 0, -1, 2, 1, 21, 0, -10, 0},
	{0, 2, 0, 0, 0, 17, -0.1, 0, 0},
	{2, 0, -1, 0, 1, 16, 0, -8, 0},
	{-2, 2, 0, 2, 2, -16, 0.1, 7, 0},
	{0, 1, 0, 0, 1, -15, 0, 9, 0},
	{-2, 0, 1, 0, 1, -13, 0, 7, 0},
	{0, -1, 0, 0, 1, -12, 0, 6, 0},
	{0, 0, 2, -2, 0, 11, 0, 0, 0},
	{2, 0, -1, 2, 1, -10, 0, 5, 0},
	{2, 0, 1, 2, 2, -8, 0, 3, 0},
	{0, 1, 0, 2, 2, 7, 0, -3, 0},
	{-2, 1, 1, 0, 0, -7, 0, 0, 0},
	{0, -1, 0, 2, 2, -7, 0, 3, 0},
	{2, 0, 0, 2, 1, -7, 0, 3, 0},
	{2, 0, 1, 0, 0, 6, 0, 0, 0},
	{-2, 0, 2, 2, 2, 6, 0, -3, 0},
	{-2, 0, 1, 2, 1, 6, 0, -3, 0},
	{2, 0, -2, 0, 1, -6, 0, 3, 0},
	{2, 0, 0, 0, 1, -6, 0, 3, 0},
	{0, -1, 1, 0, 0, 5, 0, 0, 0},
	{-2, -1, 0, 2, 1, -5, 0, 3, 0},
	{-2, 0, 0, 0, 1, -5, 0, 3, 0},
	{0, 0, 2, 2, 1, -5, 0, 3, 0},
	{-2, 0, 2, 0, 1, 4, 0, 0, 0},
	{-2, 1, 0, 2, 1, 4, 0, 0, 0},
	{0, 0, 1, -2, 0, 4, 0, 0, 0},
	{-1, 0, 1, 0, 0, -4, 0, 0, 0},
	{-2, 1, 0, 0, 0, -4, 0, 0, 0},
	{1, 0, 0, 0, 0, -4, 0, 0, 0},
	{0, 0, 1, 2, 0, 3, 0, 0, 0},
	{0, 0, -2, 2, 2, -3, 0, 0, 0},
	{-1, -1, 1, 0, 0, -3, 0, 0, 0},
	{0, 1, 1, 0, 0, -3, 0, 0, 0},
	{0, -1, 1, 2, 2, -3, 0, 0, 0},
	{2, -1, -1, 2, 2, -3, 0, 0, 0},
	{0, 0, 3, 2, 2, -3, 0, 0, 0},
	{2, -1, 0, 2, 2, -3, 0, 0, 0},
}
