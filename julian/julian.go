// Copyright 2012 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Julian: Chapter 7, Julian day.
package julian

import (
	"github.com/soniakeys/meeus"
	"math"
)

// JDMod is the Julian date of the modified Julian date epoch.
const JDMod = 2400000.5

// CalendarGregorianToJD converts a Gregorian year, month, and day of month
// to Julian day.
//
// Negative years are valid, back to JD 0.  The result is not valid for
// dates before JD 0.
func CalendarGregorianToJD(y, m int, d float64) float64 {
	switch m {
	case 1, 2:
		y--
		m += 12
	}
	a := meeus.FloorDivInt(y, 100)
	b := 2 - a + meeus.FloorDivInt(a, 4)
	return float64(meeus.FloorDivInt64(36525*(int64(y+4716)), 100)) +
		float64(meeus.FloorDivInt(306*(m+1), 10)+b) + d - 1524.5
}

// CalendarJulianToJD converts a Julian year, month, and day of month to Julian day.
//
// Negative years are valid, back to JD 0.  The result is not valid for
// dates before JD 0.
func CalendarJulianToJD(y, m int, d float64) float64 {
	switch m {
	case 1, 2:
		y--
		m += 12
	}
	return float64(meeus.FloorDivInt64(36525*(int64(y+4716)), 100)) +
		float64(meeus.FloorDivInt(306*(m+1), 10)) + d - 1524.5
}

// LeapYearJulian returns true if year y in the Julian calendar is a leap year.
func LeapYearJulian(y int) bool {
	return y%4 == 0
}

// LeapYearGregorian returns true if year y in the Gregorian calendar is a leap year.
func LeapYearGregorian(y int) bool {
	return (y%4 == 0 && y%100 != 0) || y%400 == 0
}

// JDToCalendar returns the calendar date for the given jd.
func JDToCalendar(jd float64) (year, month int, day float64) {
	zf, f := math.Modf(jd + .5)
	z := int64(zf)
	a := z
	if z >= 2299151 {
		α := meeus.FloorDivInt64(z*100-186721625, 3652425)
		a = z + 1 + α - meeus.FloorDivInt64(α, 4)
	}
	b := a + 1524
	c := meeus.FloorDivInt64(b*100-12210, 36525)
	d := meeus.FloorDivInt64(36525*c, 100)
	e := int(meeus.FloorDivInt64((b-d)*1e4, 306001))
	// compute return values
	day = float64(int(b-d)-meeus.FloorDivInt(306001*e, 1e4)) + f
	switch e {
	default:
		month = e - 1
	case 14, 15:
		month = e - 13
	}
	switch month {
	default:
		year = int(c) - 4716
	case 1, 2:
		year = int(c) - 4715
	}
	return
}

// DayOfWeek determines the day of the week for a given JD.
//
// The value returned is an integer in the range 0 to 6, where 0 represents
// Sunday.  This is the same convention followed in the time package of the
// Go standard library.
func DayOfWeek(jd float64) int {
	return int(jd+1.5) % 7
}

// DayOfYearGregorian computes the day number within the year of the Gregorian
// calendar.
func DayOfYearGregorian(y, m, d int) int {
	return DayOfYear(y, m, d, LeapYearGregorian(y))
}

// DayOfYearJulian computes the day number within the year of the Julian
// calendar.
func DayOfYearJulian(y, m, d int) int {
	return DayOfYear(y, m, d, LeapYearJulian(y))
}

// DayOfYear computes the day number within the year.
//
// This form of the function is not specific to the Julian or Gregorian
// calendar, but you must tell it whether the year is a leap year.
func DayOfYear(y, m, d int, leap bool) int {
	k := 2
	if leap {
		k--
	}
	return wholeMonths(m, k) + d
}

// DayOfYearToCalendar returns the calendar month and day for a given
// day of year and leap year status.
func DayOfYearToCalendar(n int, leap bool) (m, d int) {
	k := 2
	if leap {
		k--
	}
	if n < 32 {
		m = 1
	} else {
		m = (900*(k+n) + 98*275) / 27500
	}
	return m, n - wholeMonths(m, k)
}

func wholeMonths(m, k int) int {
	return 275*m/9 - k*((m+9)/12) - 30
}
