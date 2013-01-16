// Greenwich: Chapter 12, Sidereal Time at Greenwich
package greenwich

import (
	"math"

	"github.com/soniakeys/meeus"
	"github.com/soniakeys/meeus/julian"
)

// jdToCFrac returns values for use in computing sidereal time at Greenwich.
//
// Cen is centuries from J2000 of the JD at 0h UT of argument jd.  This is
// the value to use for evaluating the IAU sidereal time polynomial.
// DayFrac is the fraction of jd after 0h UT.  It is used to compute the
// final value of sidereal time.
func jdToCFrac(jd float64) (cen, dayFrac float64) {
	j0, f := math.Modf(jd + .5)
	return (j0 - .5 - julian.J2000) / 36525, f
}

// iau82 is a polynomial giving mean sidereal time at Greenwich at 0h UT.
//
// The polynomial is in centuries from J2000.0, as given by JDToCFrac.
// Coefficients are those adopted in 1982 by the International Astronomical
// Union.
var iau82 = []float64{24110.54841, 8640184.812866, 0.093104, 0.0000062}

// MeanSidereal returns mean sidereal time at Greenwich for the given JD.
//
// Computation is by IAU 1982 coefficients.  The results is in seconds of
// time and is in the range [0,86400).
func MeanSidereal(jd float64) float64 {
	cen, dayFrac := jdToCFrac(jd)
	s := math.Mod(meeus.Horner(cen, iau82)+dayFrac*1.00273790935*86400, 86400)
	if s < 0 {
		s += 86400
	}
	return s
}
