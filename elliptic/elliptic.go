// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Elliptic: Chapter 33, Elliptic Motion.
//
// Partial: Various formulas and algorithms are unimplemented for lack of
// examples or test cases.
package elliptic

import (
	"math"

	"github.com/soniakeys/meeus/apparent"
	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/meeus/coord"
	"github.com/soniakeys/meeus/kepler"
	"github.com/soniakeys/meeus/nutation"
	pp "github.com/soniakeys/meeus/planetposition"
	"github.com/soniakeys/meeus/solarxyz"
	"github.com/soniakeys/unit"
)

// Position returns observed equatorial coordinates of a planet at a given time.
//
// Argument p must be a valid V87Planet object for the observed planet.
// Argument earth must be a valid V87Planet object for Earth.
//
// Results are right ascension and declination, α and δ in radians.
func Position(p, earth *pp.V87Planet, jde float64) (α unit.RA, δ unit.Angle) {
	L0, B0, R0 := earth.Position(jde)
	L, B, R := p.Position(jde)
	sB0, cB0 := B0.Sincos()
	sL0, cL0 := L0.Sincos()
	sB, cB := B.Sincos()
	sL, cL := L.Sincos()
	x := R*cB*cL - R0*cB0*cL0
	y := R*cB*sL - R0*cB0*sL0
	z := R*sB - R0*sB0
	{
		Δ := math.Sqrt(x*x + y*y + z*z) // (33.4) p. 224
		τ := base.LightTime(Δ)
		// repeating with jde-τ
		L, B, R = p.Position(jde - τ)
		sB, cB = B.Sincos()
		sL, cL = L.Sincos()
		x = R*cB*cL - R0*cB0*cL0
		y = R*cB*sL - R0*cB0*sL0
		z = R*sB - R0*sB0
	}
	λ := unit.Angle(math.Atan2(y, x))                // (33.1) p. 223
	β := unit.Angle(math.Atan2(z, math.Hypot(x, y))) // (33.2) p. 223
	Δλ, Δβ := apparent.EclipticAberration(λ, β, jde)
	λ, β = pp.ToFK5(λ+Δλ, β+Δβ, jde)
	Δψ, Δε := nutation.Nutation(jde)
	λ += Δψ
	sε, cε := (nutation.MeanObliquity(jde) + Δε).Sincos()
	return coord.EclToEq(λ, β, sε, cε)
	// Meeus gives a formula for elongation but doesn't spell out how to
	// obtain term λ0 and doesn't give an example solution.
}

// Elements holds keplerian elements.
type Elements struct {
	Axis  float64    // Semimajor axis, a, in AU
	Ecc   float64    // Eccentricity, e
	Inc   unit.Angle // Inclination, i
	ArgP  unit.Angle // Argument of perihelion, ω
	Node  unit.Angle // Longitude of ascending node, Ω
	TimeP float64    // Time of perihelion, T, as jde
}

// Position returns observed equatorial coordinates of a body with Keplerian elements.
//
// Argument e must be a valid V87Planet object for Earth.
//
// Results are right ascension and declination α and δ, and elongation ψ,
// all in radians.
func (k *Elements) Position(jde float64, e *pp.V87Planet) (α unit.RA, δ, ψ unit.Angle) {
	// (33.6) p. 227
	n := base.K / k.Axis / math.Sqrt(k.Axis)
	const sε = base.SOblJ2000
	const cε = base.COblJ2000
	sΩ, cΩ := k.Node.Sincos()
	si, ci := k.Inc.Sincos()
	// (33.7) p. 228
	F := cΩ
	G := sΩ * cε
	H := sΩ * sε
	P := -sΩ * ci
	Q := cΩ*ci*cε - si*sε
	R := cΩ*ci*sε + si*cε
	// (33.8) p. 229
	A := unit.Angle(math.Atan2(F, P))
	B := unit.Angle(math.Atan2(G, Q))
	C := unit.Angle(math.Atan2(H, R))
	a := math.Hypot(F, P)
	b := math.Hypot(G, Q)
	c := math.Hypot(H, R)

	f := func(jde float64) (x, y, z float64) {
		M := unit.Angle(n * (jde - k.TimeP))
		E, err := kepler.Kepler2b(k.Ecc, M, 15)
		if err != nil {
			E = kepler.Kepler3(k.Ecc, M)
		}
		ν := kepler.True(E, k.Ecc)
		r := kepler.Radius(E, k.Ecc, k.Axis)
		// (33.9) p. 229
		x = r * a * (A + k.ArgP + ν).Sin()
		y = r * b * (B + k.ArgP + ν).Sin()
		z = r * c * (C + k.ArgP + ν).Sin()
		return
	}
	return AstrometricJ2000(f, jde, e)
}

// AstrometricJ2000 is a utility function for computing astrometric coordinates.
//
// It is used internally and only exported so that it can be used from
// multiple packages.  It is not otherwise expected to be used.
//
// Argument f is a function that returns J2000 equatorial rectangular
// coodinates of a body.
//
// Results are J2000 right ascention, declination, and elongation.
func AstrometricJ2000(f func(float64) (x, y, z float64), jde float64, e *pp.V87Planet) (α unit.RA, δ, ψ unit.Angle) {
	X, Y, Z := solarxyz.PositionJ2000(e, jde)
	x, y, z := f(jde)
	// (33.10) p. 229
	ξ := X + x
	η := Y + y
	ζ := Z + z
	Δ := math.Sqrt(ξ*ξ + η*η + ζ*ζ)
	{
		τ := base.LightTime(Δ)
		x, y, z = f(jde - τ)
		ξ = X + x
		η = Y + y
		ζ = Z + z
		Δ = math.Sqrt(ξ*ξ + η*η + ζ*ζ)
	}
	α = unit.RAFromRad(math.Atan2(η, ξ))
	δ = unit.Angle(math.Asin(ζ / Δ))
	R0 := math.Sqrt(X*X + Y*Y + Z*Z)
	ψ = unit.Angle(math.Acos((ξ*X + η*Y + ζ*Z) / R0 / Δ))
	return
}

// Velocity returns instantaneous velocity of a body in elliptical orbit around the Sun.
//
// Argument a is the semimajor axis of the body, r is the instaneous distance
// to the Sun, both in AU.
//
// Result is in Km/sec.
func Velocity(a, r float64) float64 {
	return 42.1219 * math.Sqrt(1/r-.5/a)
}

// Velocity returns the velocity of a body at aphelion.
//
// Argument a is the semimajor axis of the body in AU, e is eccentricity.
//
// Result is in Km/sec.
func VAphelion(a, e float64) float64 {
	return 29.7847 * math.Sqrt((1-e)/(1+e)/a)
}

// Velocity returns the velocity of a body at perihelion.
//
// Argument a is the semimajor axis of the body in AU, e is eccentricity.
//
// Result is in Km/sec.
func VPerihelion(a, e float64) float64 {
	return 29.7847 * math.Sqrt((1+e)/(1-e)/a)
}

// Length1 returns Ramanujan's approximation for the length of an elliptical
// orbit.
//
// Argument a is semimajor axis, e is eccentricity.
//
// Result is in units used for semimajor axis, typically AU.
func Length1(a, e float64) float64 {
	b := a * math.Sqrt(1-e*e)
	return math.Pi * (3*(a+b) - math.Sqrt((a+3*b)*(3*a+b)))
}

// Length2 returns an alternate approximation for the length of an elliptical
// orbit.
//
// Argument a is semimajor axis, e is eccentricity.
//
// Result is in units used for semimajor axis, typically AU.
func Length2(a, e float64) float64 {
	b := a * math.Sqrt(1-e*e)
	s := a + b
	p := a * b
	A := s * .5
	G := math.Sqrt(p)
	H := 2 * p / s
	return math.Pi * (21*A - 2*G - 3*H) * .125
}

// Length3 returns the length of an elliptical orbit.
//
// Argument a is semimajor axis, e is eccentricity.
//
// Result is exact, and in units used for semimajor axis, typically AU.
/* As Meeus notes, Length4 converges faster.  There is no reason to use
this function
func Length3(a, e float64) float64 {
	sum0 := 1.
	e2 := e * e
	term := e2 * .25
	sum1 := 1. - term
	nf := 1.
	df := 2.
	for sum1 != sum0 {
		term *= nf
		nf += 2
		df += 2
		term *= nf * e2 / (df * df)
		sum0 = sum1
		sum1 -= term
	}
	return 2 * math.Pi * a * sum0
}*/

// Length4 returns the length of an elliptical orbit.
//
// Argument a is semimajor axis, e is eccentricity.
//
// Result is exact, and in units used for semimajor axis, typically AU.
func Length4(a, e float64) float64 {
	b := a * math.Sqrt(1-e*e)
	m := (a - b) / (a + b)
	m2 := m * m
	sum0 := 1.
	term := m2 * .25
	sum1 := 1. + term
	nf := -1.
	df := 2.
	for sum1 != sum0 {
		nf += 2
		df += 2
		term *= nf * nf * m2 / (df * df)
		sum0 = sum1
		sum1 += term
	}
	return 2 * math.Pi * a * sum0 / (1 + m)
}
