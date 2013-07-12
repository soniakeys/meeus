// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Sidereal: Chapter 12, Sidereal Time at Greenwich.
package sidereal

import (
	"math"

	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/meeus/nutation"
)

// jdToCFrac returns values for use in computing sidereal time at Greenwich.
//
// Cen is centuries from J2000 of the JD at 0h UT of argument jd.  This is
// the value to use for evaluating the IAU sidereal time polynomial.
// DayFrac is the fraction of jd after 0h UT.  It is used to compute the
// final value of sidereal time.
func jdToCFrac(jd float64) (cen, dayFrac float64) {
	j0, f := math.Modf(jd + .5)
	return base.J2000Century(j0 - .5), f
}

// iau82 is a polynomial giving mean sidereal time at Greenwich at 0h UT.
//
// The polynomial is in centuries from J2000.0, as given by JDToCFrac.
// Coefficients are those adopted in 1982 by the International Astronomical
// Union and are given in (12.2) p. 87.
var iau82 = []float64{24110.54841, 8640184.812866, 0.093104, 0.0000062}

// Mean returns mean sidereal time at Greenwich for a given JD.
//
// Computation is by IAU 1982 coefficients.  The result is in seconds of
// time and is in the range [0,86400).
func Mean(jd float64) float64 {
	return base.PMod(mean(jd), 86400)
}

func mean(jd float64) float64 {
	s, f := mean0UT(jd)
	return s + f*1.00273790935*86400
}

// Mean0UT returns mean sidereal time at Greenwich at 0h UT on the given JD.
//
// The result is in seconds of time and is in the range [0,86400).
func Mean0UT(jd float64) float64 {
	s, _ := mean0UT(jd)
	return base.PMod(s, 86400)
}

func mean0UT(jd float64) (sidereal, dayFrac float64) {
	cen, f := jdToCFrac(jd)
	// (12.2) p. 87
	return base.Horner(cen, iau82...), f
}

// Apparent returns apparent sidereal time at Greenwich for the given JD.
//
// Apparent is mean plus the nutation in right ascension.
//
// The result is in seconds of time and is in the range [0,86400).
func Apparent(jd float64) float64 {
	s := mean(jd)                       // seconds of time
	n := nutation.NutationInRA(jd)      // angle (radians) of RA
	ns := n * 3600 * 180 / math.Pi / 15 // convert RA to time in seconds
	return base.PMod(s+ns, 86400)
}

// Apparent0UT returns apparent sidereal time at Greenwich at 0h UT
// on the given JD.
//
// The result is in seconds of time and is in the range [0,86400).
func Apparent0UT(jd float64) float64 {
	j0, f := math.Modf(jd + .5)
	cen := (j0 - .5 - base.J2000) / 36525
	s := base.Horner(cen, iau82...) + f*1.00273790935*86400
	n := nutation.NutationInRA(j0)      // angle (radians) of RA
	ns := n * 3600 * 180 / math.Pi / 15 // convert RA to time in seconds
	return base.PMod(s+ns, 86400)
}
