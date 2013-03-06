// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Parallactic: Chapter 14, The Parallactic Angle, and three other Topics.
package parallactic

import "math"

// ParallacticAngle returns parallactic angle of a celestial object.
//
//	φ is geographic latitude of observer.
//	δ is declination of observed object.
//	H is hour angle of observed object.
//
// All angles including result are in radians.
func ParallacticAngle(φ, δ, H float64) float64 {
	sδ, cδ := math.Sincos(δ)
	sH, cH := math.Sincos(H)
	return math.Atan2(sH, math.Tan(φ)*cδ-sδ*cH)
}

// ParallacticAngleOnHorizon is a special case of ParallacticAngle.
//
// The hour angle is not needed as an input and the math inside simplifies.
func ParallacticAngleOnHorizon(φ, δ float64) float64 {
	return math.Acos(math.Sin(φ) / math.Cos(δ))
}

// EclipticAtHorizon computes how the plane of the ecliptic intersects
// the horizon at a given local sidereal time as observed from a given
// geographic latitude.
//
//	ε is obliquity of the ecliptic.
//	φ is geographic latitude of observer.
//	θ is local sidereal time expressed as an hour angle.
//
//	λ1 and λ2 are ecliptic longitudes where the ecliptic intersects the horizon.
//	I is the angle at which the ecliptic intersects the horizon.
//
// All angles, arguments and results, are in radians.
func EclipticAtHorizon(ε, φ, θ float64) (λ1, λ2, I float64) {
	sε, cε := math.Sincos(ε)
	sφ, cφ := math.Sincos(φ)
	sθ, cθ := math.Sincos(θ)
	λ := math.Atan2(-cθ, sε*(sφ/cφ)+cε*sθ)
	if λ < 0 {
		λ += math.Pi
	}
	return λ, λ + math.Pi, math.Acos(cε*sφ - sε*cφ*sθ)
}

// EclipticAtEquator computes the angle between the ecliptic and the parallels
// of ecliptic latitude at a given ecliptic longitude.
//
// (The function name EclipticAtEquator is for consistency with the Meeus text,
// and works if you consider the equator a nominal parallel of latitude.)
//
//	λ is ecliptic longitude.
//	ε is obliquity of the ecliptic.
//
// All angles in radians.
func EclipticAtEquator(λ, ε float64) float64 {
	return math.Atan(-math.Cos(λ) * math.Tan(ε))
}

// DiurnalPathAtHorizon computes the angle of the path a celestial object
// relative to the horizon at the time of its rising or setting.
//
//	δ is declination of the object.
//	φ is geographic latitude of observer.
//
// All angles in radians.
func DiurnalPathAtHorizon(δ, φ float64) (J float64) {
	tφ := math.Tan(φ)
	b := math.Tan(δ) * tφ
	c := math.Sqrt(1 - b*b)
	return math.Atan(c * math.Cos(δ) / tφ)
}
