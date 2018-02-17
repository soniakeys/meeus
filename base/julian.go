// Copyright 2013 Sonia Keys
// License: MIT

package base

// Julian and Besselian years described in chapter 21, Precession.
// T, Julian centuries since J2000 described in chapter 22, Nutation.

// JMod is the Julian date of the modified Julian date epoch.
const JMod = 2400000.5

// J2000 is the Julian date corresponding to January 1.5, year 2000.
const J2000 = 2451545.0

// Julian days of common epochs.
const (
	J1900 = 2415020.0
	B1900 = 2415020.3135
	B1950 = 2433282.4235
)

// B1900, B1950 from p. 133

// JulianYear and other common periods.
const (
	JulianYear    = 365.25      // days
	JulianCentury = 36525       // days
	BesselianYear = 365.2421988 // days
)

// JulianYearToJDE returns the Julian ephemeris day for a Julian year.
func JulianYearToJDE(jy float64) float64 {
	return J2000 + JulianYear*(jy-2000)
}

// JDEToJulianYear returns a Julian year for a Julian ephemeris day.
func JDEToJulianYear(jde float64) float64 {
	return 2000 + (jde-J2000)/JulianYear
}

// BesselianYearToJDE returns the Julian ephemeris day for a Besselian year.
func BesselianYearToJDE(by float64) float64 {
	return B1900 + BesselianYear*(by-1900)
}

// JDEToBesselianYear returns the Besselian year for a Julian ephemeris day.
func JDEToBesselianYear(jde float64) float64 {
	return 1900 + (jde-B1900)/BesselianYear
}

// J2000Century returns the number of Julian centuries since J2000.
//
// The quantity appears as T in a number of time series.
func J2000Century(jde float64) float64 {
	// The formula is given in a number of places in the book, for example
	// (12.1) p. 87.
	// (22.1) p. 143.
	// (25.1) p. 163.
	return (jde - J2000) / JulianCentury
}
