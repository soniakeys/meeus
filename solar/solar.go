// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Solar: Chapter 25, Solar Coordinates.
package solar

import (
	"math"

	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/meeus/coord"
	"github.com/soniakeys/meeus/nutation"
	pp "github.com/soniakeys/meeus/planetposition"
)

// True returns true geometric longitude and anomaly of the sun.
//
// Argument T is the number of Julian centuries since J2000.
// See base.J2000Century.
//
// Results:
//	s = true geometric longitude, ☉, in radians
//	ν = true anomaly in radians
func True(T float64) (s, ν float64) {
	L0 := base.Horner(T, 280.46646, 36000.76983, 0.0003032) *
		math.Pi / 180
	M := base.Horner(T, 357.52911, 35999.05029, -0.0001537) *
		math.Pi / 180
	C := (base.Horner(T, 1.914602, -0.004817, -.000014)*
		math.Sin(M) +
		(0.019993-.000101*T)*math.Sin(2*M) +
		0.000289*math.Sin(3*M)) * math.Pi / 180
	return base.PMod(L0+C, 2*math.Pi), base.PMod(M+C, 2*math.Pi)
}

// Eccentricity returns eccentricity of the Earth's orbit around the sun.
//
// Argument T is the number of Julian centuries since J2000.
// See base.J2000Century.
func Eccentricity(T float64) float64 {
	return base.Horner(T, 0.016708634, -0.000042037, -0.0000001267)
}

// Radius returns the Sun-Earth distance in AU.
//
// Argument T is the number of Julian centuries since J2000.
// See base.J2000Century.
func Radius(T float64) float64 {
	_, ν := True(T)
	e := Eccentricity(T)
	return 1.000001018 * (1 - e*e) / (1 + e*math.Cos(ν))
}

// ApparentLongitude returns apparent longitude of the Sun referenced to the true equinox of date.
//
// Argument T is the number of Julian centuries since J2000.
// See base.J2000Century.
//
// Result includes correction for nutation and aberration.  Unit is radians.
func ApparentLongitude(T float64) float64 {
	Ω := 125.04*math.Pi/180 - 1934.136*math.Pi/180*T
	s, _ := True(T)
	return s - .00569*math.Pi/180 - .00478*math.Pi/180*math.Sin(Ω)
}

// ApparentLongitudeJ2000 returns apparent longitude of the Sun referenced to equinox J2000.
//
// Argument T is the number of Julian centuries since J2000.
// See base.J2000Century.
//
// Result includes correction for nutation and aberration.  Unit is radians.
func ApparentLongitudeJ2000(T float64) float64 {
	Ω := node(T)
	s, _ := True(T)
	return s - .00569*math.Pi/180 - .00478*math.Pi/180*math.Sin(Ω)
}

func node(T float64) float64 {
	return 125.04*math.Pi/180 - 1934.136*math.Pi/180*T
}

// TrueEquatorial returns the true geometric position of the Sun as equatorial coordinates.
func TrueEquatorial(jde float64) (α, δ float64) {
	s, _ := True(base.J2000Century(jde))
	ε := nutation.MeanObliquity(jde)
	ss, cs := math.Sincos(s)
	sε, cε := math.Sincos(ε)
	return math.Atan2(cε*ss, cs), sε * ss
}

// ApparentEquatorial returns the apparent position of the Sun as equatorial coordinates.
//
//	α: right ascension in radians
//	δ: declination in radians
func ApparentEquatorial(jde float64) (α, δ float64) {
	T := base.J2000Century(jde)
	λ := ApparentLongitude(T)
	ε := nutation.MeanObliquity(jde)
	sλ, cλ := math.Sincos(λ)
	sε, cε := math.Sincos(ε + .00256*math.Pi/180*math.Cos(node(T)))
	return math.Atan2(cε*sλ, cλ), math.Asin(sε * sλ)
}

// TrueVSOP87 returns the true geometric position of the sun as ecliptic coordinates.
//
// Result computed by truncated VSOP87.
//
//	s: ecliptic longitude in radians
//	β: ecliptic latitude in radians
//	R: range in AU
func TrueVSOP87(jde float64) (s, β, R float64) {
	l, b, r := pp.VSOP87(pp.Earth, jde)
	s = l + math.Pi
	// FK5 correction.
	λp := base.Horner(base.J2000Century(jde),
		s, -1.397*math.Pi/180, -.00031*math.Pi/180)
	sλp, cλp := math.Sincos(λp)
	Δβ := .03916 / 3600 * math.Pi / 180 * (cλp - sλp)
	return base.PMod(s-.09033/3600*math.Pi/180, 2*math.Pi), Δβ - b, r
}

// ApparentVSOP87 returns the apparent position of the sun as ecliptic coordinates.
//
// Result computed by truncated VSOP87 and includes effects of nutation and
// aberration.
//
//  λ: ecliptic longitude in radians
//  β: ecliptic latitude in radians
//  R: range in AU
func ApparentVSOP87(jde float64) (λ, β, R float64) {
	s, β, R := TrueVSOP87(jde)
	Δψ, _ := nutation.Nutation(jde)
	Δλ := -20.4898 / 3600 * math.Pi / 180 / R // aberration
	return s + Δψ + Δλ, β, R
}

// ApparentEquatorialVSOP87 returns the apparent position of the sun as equatorial coordinates.
//
// Result computed by truncated VSOP87 and includes effects of nutation and
// aberration.
//
//	α: right ascension in radians
//	δ: declination in radians
//	R: range in AU
func ApparentEquatorialVSOP87(jde float64) (α, δ, R float64) {
	// duplicate code from ApparentVSOP87 so we can keep Δε
	s, β, R := TrueVSOP87(jde)
	Δψ, Δε := nutation.Nutation(jde)
	Δλ := -20.4898 / 3600 * math.Pi / 180 / R // aberration
	ecl := &coord.Ecliptic{
		Lon: s + Δψ + Δλ,
		Lat: β,
	}
	ε := nutation.MeanObliquity(jde) + Δε
	sε, cε := math.Sincos(ε)
	eq := &coord.Equatorial{}
	eq.EclToEq(ecl, sε, cε)
	return eq.RA, eq.Dec, R
}
