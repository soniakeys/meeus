// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Elliptic: Chapter 33, Elliptic Motion
package elliptic

import (
	"math"

	"github.com/soniakeys/meeus/apparent"
	"github.com/soniakeys/meeus/coord"
	"github.com/soniakeys/meeus/kepler"
	"github.com/soniakeys/meeus/nutation"
	pp "github.com/soniakeys/meeus/planetposition"
	"github.com/soniakeys/meeus/solarxyz"
)

// Planet returns observed equatorial coordinates of a planet at a given time.
//
// Argument p must be a valid V87Planet object for the observed planet.
// Arguemnt earth must be a valed V87Planet object for Earth.
func Planet(p, earth *pp.V87Planet, jde float64) (α, δ float64) {
	L0, B0, R0 := earth.Position(jde)
	L, B, R := p.Position(jde)
	sB0, cB0 := math.Sincos(B0)
	sL0, cL0 := math.Sincos(L0)
	sB, cB := math.Sincos(B)
	sL, cL := math.Sincos(L)
	x := R*cB*cL - R0*cB0*cL0
	y := R*cB*sL - R0*cB0*sL0
	z := R*sB - R0*sB0
	{
		Δ := math.Sqrt(x*x + y*y + z*z) // (33.4) p. 224
		τ := lightTime(Δ)
		// repeating with jde-τ
		L, B, R = p.Position(jde - τ)
		sB, cB = math.Sincos(B)
		sL, cL = math.Sincos(L)
		x = R*cB*cL - R0*cB0*cL0
		y = R*cB*sL - R0*cB0*sL0
		z = R*sB - R0*sB0
	}
	λ := math.Atan2(y, x)                // (33.1) p. 223
	β := math.Atan2(z, math.Hypot(x, y)) // (33.2) p. 223
	Δλ, Δβ := apparent.EclipticAberration(λ, β, jde)
	λ, β = pp.ToFK5(λ+Δλ, β+Δβ, jde)
	Δψ, Δε := nutation.Nutation(jde)
	λ += Δψ
	sε, cε := math.Sincos(nutation.MeanObliquity(jde) + Δε)
	return coord.EclToEq(λ, β, sε, cε)
}

func lightTime(Δ float64) float64 {
	return .0057755183 * Δ // (33.3) p. 224
}

type Keplerian struct {
	Axis float64 // Semimajor axis, a, in AU
	Ecc  float64 // Eccentricity, e
	Inc  float64 // Inclination, i, in radians
	Peri float64 // Argument of perihelion, ω, in radians
	Node float64 // Longitude of ascending node, Ω, in radians
	Time float64 // Time of perihelion, T, as jde
}

func Elements(k *Keplerian, jde float64, e *pp.V87Planet) (α, δ float64) {
	X, Y, Z := solarxyz.PositionJ2000(e, jde)
	// (33.6) p. 227
	n := .9856076686 * math.Pi / 180 / k.Axis / math.Sqrt(k.Axis)
	// J2000 values, given on p. 228
	const sε = .397777156
	const cε = .917482062
	sΩ, cΩ := math.Sincos(k.Node)
	si, ci := math.Sincos(k.Inc)
	// (33.7) p. 228
	F := cΩ
	G := sΩ * cε
	H := sΩ * sε
	P := -sΩ * ci
	Q := cΩ*ci*cε - si*sε
	R := cΩ*ci*sε + si*cε
	// (33.8) p. 229
	A := math.Atan2(F, P)
	B := math.Atan2(G, Q)
	C := math.Atan2(H, R)
	a := math.Hypot(F, P)
	b := math.Hypot(G, Q)
	c := math.Hypot(H, R)

	M := n * (jde - k.Time)
	E, err := kepler.Kepler2b(k.Ecc, M, 15)
	if err != nil {
		E = kepler.Kepler3(k.Ecc, M)
	}
	ν := kepler.True(E, k.Ecc)
	r := kepler.Radius(E, k.Ecc, k.Axis)
	// (33.9) p. 229
	x := r * a * math.Sin(A+k.Peri+ν)
	y := r * b * math.Sin(B+k.Peri+ν)
	z := r * c * math.Sin(C+k.Peri+ν)
	// (33.10) p. 229
	ξ := X + x
	η := Y + y
	ζ := Z + z
	Δ := math.Sqrt(ξ*ξ + η*η + ζ*ζ)
	{
		τ := lightTime(Δ)
		// repeating with jde-τ
		M = n * (jde - τ - k.Time)
		E, err = kepler.Kepler2b(k.Ecc, M, 15)
		if err != nil {
			E = kepler.Kepler3(k.Ecc, M)
		}
		ν = kepler.True(E, k.Ecc)
		r = kepler.Radius(E, k.Ecc, k.Axis)
		x = r * a * math.Sin(A+k.Peri+ν)
		y = r * b * math.Sin(B+k.Peri+ν)
		z = r * c * math.Sin(C+k.Peri+ν)
		ξ = X + x
		η = Y + y
		ζ = Z + z
		Δ = math.Sqrt(ξ*ξ + η*η + ζ*ζ)
	}
	α = math.Atan2(η, ξ)
	δ = math.Asin(ζ / Δ)
	return
}
