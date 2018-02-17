// Copyright 2013 Sonia Keys
// License: MIT

// Refraction: Chapter 16: Atmospheric Refraction.
//
// Functions here assume atmospheric pressure of 1010 mb, temperature of
// 10°C, and yellow light.
package refraction

import (
	"math"

	"github.com/soniakeys/unit"
)

var (
	gt15true1 = unit.AngleFromSec(58.294)
	gt15true2 = unit.AngleFromSec(.0668)
	gt15app1  = unit.AngleFromSec(58.276)
	gt15app2  = unit.AngleFromSec(.0824)
)

// Gt15True returns refraction for obtaining true altitude when altitude
// is greater than 15 degrees (about .26 radians.)
//
// h0 must be a measured apparent altitude of a celestial body.
//
// Result is refraction to be subtracted from h0 to obtain the true altitude
// of the body.
func Gt15True(h0 unit.Angle) unit.Angle {
	// (16.1) p. 105
	t := (math.Pi/2 - h0).Tan()
	return gt15true1.Mul(t) - gt15true2.Mul(t*t*t)
}

// Gt15Apparent returns refraction for obtaining apparent altitude when
// altitude is greater than 15 degrees (about .26 radians.)
//
// h must be a computed true "airless" altitude of a celestial body.
//
// Result is refraction to be added to h to obtain the apparent altitude
// of the body.
func Gt15Apparent(h unit.Angle) unit.Angle {
	// (16.2) p. 105
	t := (math.Pi/2 - h).Tan()
	return gt15app1.Mul(t) - gt15app2.Mul(t*t*t)
}

// Bennett returns refraction for obtaining true altitude.
//
// h0 must be a measured apparent altitude of a celestial body in radians.
//
// Results are accurate to .07 arc min from horizon to zenith.
//
// Result is refraction to be subtracted from h0 to obtain the true altitude
// of the body.
func Bennett(h0 unit.Angle) unit.Angle {
	// (16.3) p. 106
	hd := h0.Deg()
	return unit.AngleFromMin(1 / math.Tan((hd+7.31/(hd+4.4))*math.Pi/180))
}

// Bennett2 returns refraction for obtaining true altitude.
//
// Similar to Bennett, but a correction is applied to give a more accurate
// result.
//
// Results are accurate to .015 arc min.  Result unit is radians.
func Bennett2(h0 unit.Angle) unit.Angle {
	R := Bennett(h0).Min()
	return unit.AngleFromMin(R - .06*math.Sin((14.7*R+13)*math.Pi/180))
}

// Saemundsson returns refraction for obtaining apparent altitude.
//
// h must be a computed true "airless" altitude of a celestial body in radians.
//
// Result is refraction to be added to h to obtain the apparent altitude
// of the body.
//
// Results are consistent with Bennett to within 4 arc sec.
func Saemundsson(h unit.Angle) unit.Angle {
	// (16.4) p. 106
	hd := h.Deg()
	return unit.AngleFromMin(1.02 / math.Tan((hd+10.3/(hd+5.11))*math.Pi/180))
}
