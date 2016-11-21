// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Refraction: Chapter 16: Atmospheric Refraction.
//
// Functions here assume atmospheric pressure of 1010 mb, temperature of
// 10°C, and yellow light.
package refraction

import (
	"math"

	"github.com/soniakeys/meeus/base"
)

var (
	gt15true1 = base.AngleFromSec(58.294).Rad()
	gt15true2 = base.AngleFromSec(.0668).Rad()
	gt15app1  = base.AngleFromSec(58.276).Rad()
	gt15app2  = base.AngleFromSec(.0824).Rad()
)

// Gt15True returns refraction for obtaining true altitude when altitude
// is greater than 15 degrees (about .26 radians.)
//
// h0 must be a measured apparent altitude of a celestial body in radians.
//
// Result is refraction to be subtracted from h0 to obtain the true altitude
// of the body.  Unit is radians.
func Gt15True(h0 float64) float64 {
	// (16.1) p. 105
	t := math.Tan(math.Pi/2 - h0)
	return gt15true1*t - gt15true2*t*t*t
}

// Gt15Apparent returns refraction for obtaining apparent altitude when
// altitude is greater than 15 degrees (about .26 radians.)
//
// h must be a computed true "airless" altitude of a celestial body in radians.
//
// Result is refraction to be added to h to obtain the apparent altitude
// of the body.  Unit is radians.
func Gt15Apparent(h float64) float64 {
	// (16.2) p. 105
	t := math.Tan(math.Pi/2 - h)
	return gt15app1*t - gt15app2*t*t*t
}

// Bennett returns refraction for obtaining true altitude.
//
// h0 must be a measured apparent altitude of a celestial body in radians.
//
// Results are accurate to .07 arc min from horizon to zenith.
//
// Result is refraction to be subtracted from h0 to obtain the true altitude
// of the body.  Unit is radians.
func Bennett(h0 float64) float64 {
	// (16.3) p. 106
	const cRad = math.Pi / 180
	const c1 = cRad / 60
	const c731 = 7.31 * cRad * cRad
	const c44 = 4.4 * cRad
	return c1 / math.Tan(h0+c731/(h0+c44))
}

// Bennett2 returns refraction for obtaining true altitude.
//
// Similar to Bennett, but a correction is applied to give a more accurate
// result.
//
// Results are accurate to .015 arc min.  Result unit is radians.
func Bennett2(h0 float64) float64 {
	const cRad = math.Pi / 180
	const cMin = 60 / cRad
	const c06 = .06 / cMin
	const c147 = 14.7 * cMin * cRad
	const c13 = 13 * cRad
	R := Bennett(h0)
	return R - c06*math.Sin(c147*R+c13)
}

// Saemundsson returns refraction for obtaining apparent altitude.
//
// h must be a computed true "airless" altitude of a celestial body in radians.
//
// Result is refraction to be added to h to obtain the apparent altitude
// of the body.
//
// Results are consistent with Bennett to within 4 arc sec.
// Result unit is radians.
func Saemundsson(h float64) float64 {
	// (16.4) p. 106
	const cRad = math.Pi / 180
	const c102 = 1.02 * cRad / 60
	const c103 = 10.3 * cRad * cRad
	const c511 = 5.11 * cRad
	return c102 / math.Tan(h+c103/(h+c511))
}
