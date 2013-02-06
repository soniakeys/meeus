// Sidereal: Chapter 12, Sidereal Time at Greenwich
package sidereal

import (
	"math"

	"github.com/soniakeys/meeus/common"
	"github.com/soniakeys/meeus/hints"
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
	return (j0 - .5 - common.J2000) / 36525, f
}

// iau82 is a polynomial giving mean sidereal time at Greenwich at 0h UT.
//
// The polynomial is in centuries from J2000.0, as given by JDToCFrac.
// Coefficients are those adopted in 1982 by the International Astronomical
// Union.
var iau82 = []float64{24110.54841, 8640184.812866, 0.093104, 0.0000062}

// Mean returns mean sidereal time at Greenwich for a given JD.
//
// Computation is by IAU 1982 coefficients.  The result is in seconds of
// time and is in the range [0,86400).
func Mean(jd float64) float64 {
	cen, dayFrac := jdToCFrac(jd)
	s := math.Mod(hints.Horner(cen, iau82)+dayFrac*1.00273790935*86400, 86400)
	if s < 0 {
		s += 86400
	}
	return s
}

// Apparent returns apparent sidereal time at Greenwich for the given JD.
//
// Apparent is mean plus the nutation in right ascension.
//
// The result is in seconds of time and is in the range [0,86400).
func Apparent(jd float64) float64 {
	s := Mean(jd)                       // seconds of time
	n := nutation.NutationInRA(jd)      // angle (radians) of RA
	ns := n * 3600 * 180 / math.Pi / 15 // convert RA to time in seconds
	s = math.Mod(s+ns, 86400)
	if s < 0 {
		s += 86400
	}
	return s
}
