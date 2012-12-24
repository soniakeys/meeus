// Copyright 2012 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Julian: Chapter 7, Julian day.
package julian

// JDMod is the Julian date of the modified Julian date epoch.
const JDMod = 2400000.5

// GregYMDToJD converts a Gregorian year, month, and day of month
// to Julian date.
func GregYMDToJD(y, m int, d float64) float64 {
	switch m {
	case 1, 2:
		y--
		m += 12
	}
	a := y / 100
	b := 2 - a + a/4
	return float64(36525*(int64(y)+4716)/100) +
		float64(306*(m+1)/10+b) + d - 1524.5
}

// JuliYMDToJD converts a Julian year, month, and day of month to Julian date.
func JuliYMDToJD(y, m int, d float64) float64 {
	switch m {
	case 1, 2:
		y--
		m += 12
	}
	return float64(36525*(int64(y)+4716)/100) +
		float64(306*(m+1)/10) + d - 1524.5
}
