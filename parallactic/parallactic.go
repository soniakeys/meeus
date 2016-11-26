// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Parallactic: Chapter 14, The Parallactic Angle, and three other Topics.
package parallactic

import (
	"math"

	"github.com/soniakeys/meeus/base"
)

// ParallacticAngle returns parallactic angle of a celestial object.
//
//	φ is geographic latitude of observer.
//	δ is declination of observed object.
//	H is hour angle of observed object.
func ParallacticAngle(φ, δ base.Angle, H base.HourAngle) base.Angle {
	sδ, cδ := math.Sincos(δ.Rad())
	sH, cH := math.Sincos(H.Rad())
	// (14.1) p. 98
	return base.Angle(math.Atan2(sH, math.Tan(φ.Rad())*cδ-sδ*cH))
}

// ParallacticAngleOnHorizon is a special case of ParallacticAngle.
//
// The hour angle is not needed as an input and the math inside simplifies.
func ParallacticAngleOnHorizon(φ, δ base.Angle) base.Angle {
	return base.Angle(math.Acos(math.Sin(φ.Rad()) / math.Cos(δ.Rad())))
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
func EclipticAtHorizon(ε, φ base.Angle, θ base.Time) (λ1, λ2, I base.Angle) {
	sε, cε := math.Sincos(ε.Rad())
	sφ, cφ := math.Sincos(φ.Rad())
	sθ, cθ := math.Sincos(θ.Rad())
	// (14.2) p. 99
	λ := base.Angle(math.Atan2(-cθ, sε*(sφ/cφ)+cε*sθ))
	if λ < 0 {
		λ += math.Pi
	}
	// (14.3) p. 99
	return λ, λ + math.Pi, base.Angle(math.Acos(cε*sφ - sε*cφ*sθ))
}

// EclipticAtEquator computes the angle between the ecliptic and the parallels
// of ecliptic latitude at a given ecliptic longitude.
//
// (The function name EclipticAtEquator is for consistency with the Meeus text,
// and works if you consider the equator a nominal parallel of latitude.)
//
//	λ is ecliptic longitude.
//	ε is obliquity of the ecliptic.
func EclipticAtEquator(λ, ε base.Angle) base.Angle {
	return base.Angle(math.Atan(-math.Cos(λ.Rad()) * math.Tan(ε.Rad())))
}

// DiurnalPathAtHorizon computes the angle of the path a celestial object
// relative to the horizon at the time of its rising or setting.
//
//	δ is declination of the object.
//	φ is geographic latitude of observer.
func DiurnalPathAtHorizon(δ, φ base.Angle) (J base.Angle) {
	tφ := math.Tan(φ.Rad())
	b := math.Tan(δ.Rad()) * tφ
	c := math.Sqrt(1 - b*b)
	return base.Angle(math.Atan(c * math.Cos(δ.Rad()) / tφ))
}
