// Copyright 2013 Sonia Keys
// License: MIT

// Parallactic: Chapter 14, The Parallactic Angle, and three other Topics.
package parallactic

import (
	"math"

	"github.com/soniakeys/unit"
)

// ParallacticAngle returns parallactic angle of a celestial object.
//
//	φ is geographic latitude of observer.
//	δ is declination of observed object.
//	H is hour angle of observed object.
func ParallacticAngle(φ, δ unit.Angle, H unit.HourAngle) unit.Angle {
	sδ, cδ := δ.Sincos()
	sH, cH := H.Sincos()
	// (14.1) p. 98
	return unit.Angle(math.Atan2(sH, φ.Tan()*cδ-sδ*cH))
}

// ParallacticAngleOnHorizon is a special case of ParallacticAngle.
//
// The hour angle is not needed as an input and the math inside simplifies.
func ParallacticAngleOnHorizon(φ, δ unit.Angle) unit.Angle {
	return unit.Angle(math.Acos(φ.Sin() / δ.Cos()))
}

// EclipticAtHorizon computes how the plane of the ecliptic intersects
// the horizon at a given local sidereal time as observed from a given
// geographic latitude.
//
//	ε is obliquity of the ecliptic.
//	φ is geographic latitude of observer.
//	θ is local sidereal time.
//
//	λ1 and λ2 are ecliptic longitudes where the ecliptic intersects the horizon.
//	I is the angle at which the ecliptic intersects the horizon.
func EclipticAtHorizon(ε, φ unit.Angle, θ unit.Time) (λ1, λ2, I unit.Angle) {
	sε, cε := ε.Sincos()
	sφ, cφ := φ.Sincos()
	sθ, cθ := θ.Angle().Sincos()
	// (14.2) p. 99
	λ := unit.Angle(math.Atan2(-cθ, sε*(sφ/cφ)+cε*sθ))
	if λ < 0 {
		λ += math.Pi
	}
	// (14.3) p. 99
	return λ, λ + math.Pi, unit.Angle(math.Acos(cε*sφ - sε*cφ*sθ))
}

// EclipticAtEquator computes the angle between the ecliptic and the parallels
// of ecliptic latitude at a given ecliptic longitude.
//
// (The function name EclipticAtEquator is for consistency with the Meeus text,
// and works if you consider the equator a nominal parallel of latitude.)
//
//	λ is ecliptic longitude.
//	ε is obliquity of the ecliptic.
func EclipticAtEquator(λ, ε unit.Angle) unit.Angle {
	return unit.Angle(math.Atan(-λ.Cos() * ε.Tan()))
}

// DiurnalPathAtHorizon computes the angle of the path a celestial object
// relative to the horizon at the time of its rising or setting.
//
//	δ is declination of the object.
//	φ is geographic latitude of observer.
func DiurnalPathAtHorizon(δ, φ unit.Angle) (J unit.Angle) {
	tφ := φ.Tan()
	b := δ.Tan() * tφ
	c := math.Sqrt(1 - b*b)
	return unit.Angle(math.Atan(c * δ.Cos() / tφ))
}
