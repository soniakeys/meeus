// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Solar: Chapter 25, Solar Coordinates.
package solar

import (
	"math"

	"github.com/soniakeys/meeus/base"
)

// True returns true geometric longitude and anomaly of the sun.
//
// Argument T is the number of Julian centuries since J2000.
// See base.J2000Century.
//
//	s = true geometric longitude, ☉
//	ν = true anomaly
func True(T float64) (s, ν float64) {
	L0 := base.Horner(T, []float64{280.46646, 36000.76983, 0.0003032}) *
		math.Pi / 180
	M := base.Horner(T, []float64{357.52911, 35999.05029, -0.0001537}) *
		math.Pi / 180
	C := (base.Horner(T, []float64{1.914602, -0.004817, -.000014})*
		math.Sin(M) +
		(0.019993-.000101*T)*math.Sin(2*M) +
		0.000289*math.Sin(3*M)) * math.Pi / 180
	return base.PMod(L0+C, 2*math.Pi), base.PMod(M+C, 2*math.Pi)
}

// Eccentricity returns eccentricity of the Earth's orbit.
//
// Argument T is the number of Julian centuries since J2000.
// See base.J2000Century.
func Eccentricity(T float64) float64 {
	return base.Horner(T, []float64{0.016708634, -0.000042037, -0.0000001267})
}
