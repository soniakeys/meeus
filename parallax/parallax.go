// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Parallax: Chapter 40, Correction for Parallax.
package parallax

import (
	"math"

	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/meeus/globe"
	"github.com/soniakeys/meeus/sidereal"
	"github.com/soniakeys/sexagesimal"
)

// constant for Horizontal.  p. 279.
var hp = sexa.NewAngle(' ', 0, 0, 8.794).Rad()

// Horizontal returns equatorial horizontal parallax of a body.
//
// Argument Δ is distance in AU.
//
// Result is parallax in radians.
//
// Meeus mentions use of this function for the Moon, Sun, planet, or comet.
// That is, for relatively distant objects.  For parallax of the Moon (or
// other relatively close object) see moonposition.Parallax.
func Horizontal(Δ float64) (π float64) {
	return hp / Δ // (40.1) p. 279
}

// Topocentric returns topocentric positions including parallax.
//
// Arguments α, δ are geocentric right ascension and declination in radians.
// Δ is distance to the observed object in AU.  ρsφʹ, ρcφʹ are parallax
// constants (see package globe.) L is geographic longitude of the observer,
// jde is time of observation.
//
// Results are observed topocentric ra and dec in radians.
func Topocentric(α, δ, Δ, ρsφʹ, ρcφʹ, L, jde float64) (αʹ, δʹ float64) {
	π := Horizontal(Δ)
	θ0 := sexa.Time(sidereal.Apparent(jde)).Rad()
	H := base.PMod(θ0-L-α, 2*math.Pi)
	sπ := math.Sin(π)
	sH, cH := math.Sincos(H)
	sδ, cδ := math.Sincos(δ)
	Δα := math.Atan2(-ρcφʹ*sπ*sH, cδ-ρcφʹ*sπ*cH) // (40.2) p. 279
	αʹ = α + Δα
	δʹ = math.Atan2((sδ-ρsφʹ*sπ)*math.Cos(Δα), cδ-ρcφʹ*sπ*cH) // (40.3) p. 279
	return
}

// Topocentric2 returns topocentric corrections including parallax.
//
// This function implements the "non-rigorous" method descripted in the text.
//
// Note that results are corrections, not corrected coordinates.
func Topocentric2(α, δ, Δ, ρsφʹ, ρcφʹ, L, jde float64) (Δα, Δδ float64) {
	π := Horizontal(Δ)
	θ0 := sexa.Time(sidereal.Apparent(jde)).Rad()
	H := base.PMod(θ0-L-α, 2*math.Pi)
	sH, cH := math.Sincos(H)
	sδ, cδ := math.Sincos(δ)
	Δα = -π * ρcφʹ * sH / cδ         // (40.4) p. 280
	Δδ = -π * (ρsφʹ*cδ - ρcφʹ*cH*sδ) // (40.5) p. 280
	return
}

// Topocentric3 returns topocentric hour angle and declination including parallax.
//
// This function implements the "alternative" method described in the text.
// The method should be similarly rigorous to that of Topocentric() and results
// should be virtually consistent.
func Topocentric3(α, δ, Δ, ρsφʹ, ρcφʹ, L, jde float64) (Hʹ, δʹ float64) {
	π := Horizontal(Δ)
	θ0 := sexa.Time(sidereal.Apparent(jde)).Rad()
	H := base.PMod(θ0-L-α, 2*math.Pi)
	sπ := math.Sin(π)
	sH, cH := math.Sincos(H)
	sδ, cδ := math.Sincos(δ)
	A := cδ * sH
	B := cδ*cH - ρcφʹ*sπ
	C := sδ - ρsφʹ*sπ
	q := math.Sqrt(A*A + B*B + C*C)
	Hʹ = math.Atan2(A, B)
	δʹ = math.Asin(C / q)
	return
}

// TopocentricEcliptical returns topocentric ecliptical coordinates including parallax.
//
// Arguments λ, β are geocentric ecliptical longitude and latitude of a body,
// s is its geocentric semidiameter. φ, h are the observer's latitude and
// and height above the ellipsoid in meters.  ε is the obliquity of the
// ecliptic, θ is local sidereal time, π is equatorial horizontal parallax
// of the body (see Horizonal()).
//
// All angular parameters and results are in radians.
//
// Results are observed topocentric coordinates and semidiameter.
func TopocentricEcliptical(λ, β, s, φ, h, ε, θ, π float64) (λʹ, βʹ, sʹ float64) {
	S, C := globe.Earth76.ParallaxConstants(φ, h)
	sλ, cλ := math.Sincos(λ)
	sβ, cβ := math.Sincos(β)
	sε, cε := math.Sincos(ε)
	sθ, cθ := math.Sincos(θ)
	sπ := math.Sin(π)
	N := cλ*cβ - C*sπ*cθ
	λʹ = math.Atan2(sλ*cβ-sπ*(S*sε+C*cε*sθ), N)
	if λʹ < 0 {
		λʹ += 2 * math.Pi
	}
	cλʹ := math.Cos(λʹ)
	βʹ = math.Atan(cλʹ * (sβ - sπ*(S*cε-C*sε*sθ)) / N)
	sʹ = math.Asin(cλʹ * math.Cos(βʹ) * math.Sin(s) / N)
	return
}
