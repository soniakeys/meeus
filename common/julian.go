// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package common

// Julian and Besselian years described in chapter 21, Precession.

// JMod is the Julian date of the modified Julian date epoch.
const JMod = 2400000.5

// J2000 is the Julian date corresponding to January 1.5, year 2000.
const J2000 = 2451545.0
const J1900 = 2415020.0
const B1900 = 2415020.3135

const JulianYear = 365.25   // days
const JulianCentury = 36525 // days

const BesselianYear = 365.2421988 // days

func JulianYearToJDE(jy float64) float64 {
	return J2000 + JulianYear*(jy-2000)
}

func JDEToJulianYear(jde float64) float64 {
	return 2000 + (jde-J2000)/JulianYear
}

func BesselianYearToJDE(by float64) float64 {
	return B1900 + BesselianYear*(by-1900)
}

func JDEToBesselianYear(jde float64) float64 {
	return 1900 + (jde-B1900)/BesselianYear
}
