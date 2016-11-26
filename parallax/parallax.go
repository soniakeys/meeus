// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Parallax: Chapter 40, Correction for Parallax.
package parallax

import (
	"math"

	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/meeus/globe"
	"github.com/soniakeys/meeus/sidereal"
)

// constant for Horizontal.  p. 279.
var hp = base.AngleFromSec(8.794)

// Horizontal returns equatorial horizontal parallax of a body.
//
// Argument Δ is distance in AU.
//
// Meeus mentions use of this function for the Moon, Sun, planet, or comet.
// That is, for relatively distant objects.  For parallax of the Moon (or
// other relatively close object) see moonposition.Parallax.
func Horizontal(Δ float64) (π base.Angle) {
	return hp.Div(Δ) // (40.1) p. 279
}

// Topocentric returns topocentric positions including parallax.
//
// Arguments α, δ are geocentric right ascension and declination in radians.
// Δ is distance to the observed object in AU.  ρsφʹ, ρcφʹ are parallax
// constants (see package globe.) L is geographic longitude of the observer,
// jde is time of observation.
//
// Results are observed topocentric ra and dec in radians.
func Topocentric(α base.RA, δ base.Angle, Δ, ρsφʹ, ρcφʹ float64, L base.Angle, jde float64) (αʹ base.RA, δʹ base.Angle) {
	π := Horizontal(Δ)
	θ0 := sidereal.Apparent(jde)
	H := (θ0.Angle() - L - base.Angle(α)).Mod1()
	sπ := math.Sin(π.Rad())
	sH, cH := math.Sincos(H.Rad())
	sδ, cδ := math.Sincos(δ.Rad())
	// (40.2) p. 279
	Δα := base.HourAngle(math.Atan2(-ρcφʹ*sπ*sH, cδ-ρcφʹ*sπ*cH))
	αʹ = α.Add(Δα)
	// (40.3) p. 279
	δʹ = base.Angle(math.Atan2((sδ-ρsφʹ*sπ)*math.Cos(Δα.Rad()), cδ-ρcφʹ*sπ*cH))
	return
}

// Topocentric2 returns topocentric corrections including parallax.
//
// This function implements the "non-rigorous" method descripted in the text.
//
// Note that results are corrections, not corrected coordinates.
func Topocentric2(α base.RA, δ base.Angle, Δ, ρsφʹ, ρcφʹ float64, L base.Angle, jde float64) (Δα base.HourAngle, Δδ base.Angle) {
	π := Horizontal(Δ)
	θ0 := sidereal.Apparent(jde)
	H := (θ0.Angle() - L - base.Angle(α)).Mod1()
	sH, cH := math.Sincos(H.Rad())
	sδ, cδ := math.Sincos(δ.Rad())
	Δα = base.HourAngle(-π.Mul(ρcφʹ * sH / cδ)) // (40.4) p. 280
	Δδ = -π.Mul(ρsφʹ*cδ - ρcφʹ*cH*sδ)           // (40.5) p. 280
	return
}

// Topocentric3 returns topocentric hour angle and declination including parallax.
//
// This function implements the "alternative" method described in the text.
// The method should be similarly rigorous to that of Topocentric() and results
// should be virtually consistent.
func Topocentric3(α base.RA, δ base.Angle, Δ, ρsφʹ, ρcφʹ float64, L base.Angle, jde float64) (Hʹ base.HourAngle, δʹ base.Angle) {
	π := Horizontal(Δ)
	θ0 := sidereal.Apparent(jde)
	H := (θ0.Angle() - L - base.Angle(α)).Mod1()
	sπ := math.Sin(π.Rad())
	sH, cH := math.Sincos(H.Rad())
	sδ, cδ := math.Sincos(δ.Rad())
	A := cδ * sH
	B := cδ*cH - ρcφʹ*sπ
	C := sδ - ρsφʹ*sπ
	q := math.Sqrt(A*A + B*B + C*C)
	Hʹ = base.HourAngle(math.Atan2(A, B))
	δʹ = base.Angle(math.Asin(C / q))
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
// Results are observed topocentric coordinates and semidiameter.
func TopocentricEcliptical(λ, β, s, φ base.Angle, h float64, ε base.Angle, θ base.Time, π base.Angle) (λʹ, βʹ, sʹ base.Angle) {
	S, C := globe.Earth76.ParallaxConstants(φ, h)
	sλ, cλ := math.Sincos(λ.Rad())
	sβ, cβ := math.Sincos(β.Rad())
	sε, cε := math.Sincos(ε.Rad())
	sθ, cθ := math.Sincos(θ.Rad())
	sπ := math.Sin(π.Rad())
	N := cλ*cβ - C*sπ*cθ
	λʹ = base.Angle(math.Atan2(sλ*cβ-sπ*(S*sε+C*cε*sθ), N))
	if λʹ < 0 {
		λʹ += 2 * math.Pi
	}
	cλʹ := math.Cos(λʹ.Rad())
	βʹ = base.Angle(math.Atan(cλʹ * (sβ - sπ*(S*cε-C*sε*sθ)) / N))
	sʹ = base.Angle(math.Asin(cλʹ * math.Cos(βʹ.Rad()) * math.Sin(s.Rad()) / N))
	return
}
