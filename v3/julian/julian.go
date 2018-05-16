// Copyright 2012 Sonia Keys
// License: MIT

// Julian: Chapter 7, Julian day.
//
// Under "General remarks", Meeus describes the INT function as used in the
// book.  In some contexts, math.Floor might be suitable, but I think more
// often, the functions base.FloorDiv or base.FloorDiv64 will be more
// appropriate.  See documentation in package base.
//
// On p. 63, Modified Julian Day is defined.  See constant JMod in package
// base.
//
// See also related functions JulianToGregorian and GregorianToJulian in
// package jm.
package julian

import (
	"math"
	"time"

	"github.com/soniakeys/meeus/v3/base"
)

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
	a := base.FloorDiv(y, 100)
	b := 2 - a + base.FloorDiv(a, 4)
	// (7.1) p. 61
	return float64(base.FloorDiv64(36525*(int64(y+4716)), 100)) +
		float64(base.FloorDiv(306*(m+1), 10)+b) + d - 1524.5
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
	return float64(base.FloorDiv64(36525*(int64(y+4716)), 100)) +
		float64(base.FloorDiv(306*(m+1), 10)) + d - 1524.5
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
//
// Note that this function returns a date in either the Julian or Gregorian
// Calendar, as appropriate.
func JDToCalendar(jd float64) (year, month int, day float64) {
	zf, f := math.Modf(jd + .5)
	z := int64(zf)
	a := z
	if z >= 2299161 {
		α := base.FloorDiv64(z*100-186721625, 3652425)
		a = z + 1 + α - base.FloorDiv64(α, 4)
	}
	b := a + 1524
	c := base.FloorDiv64(b*100-12210, 36525)
	d := base.FloorDiv64(36525*c, 100)
	e := int(base.FloorDiv64((b-d)*1e4, 306001))
	// compute return values
	day = float64(int(b-d)-base.FloorDiv(306001*e, 1e4)) + f
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

// jdToCalendarGregorian returns the Gregorian calendar date for the given jd.
//
// Note that it returns a Gregorian date even for dates before the start of
// the Gregorian calendar.  The function is useful when working with Go
// time.Time values because they are always based on the Gregorian calendar.
func jdToCalendarGregorian(jd float64) (year, month int, day float64) {
	zf, f := math.Modf(jd + .5)
	z := int64(zf)
	α := base.FloorDiv64(z*100-186721625, 3652425)
	a := z + 1 + α - base.FloorDiv64(α, 4)
	b := a + 1524
	c := base.FloorDiv64(b*100-12210, 36525)
	d := base.FloorDiv64(36525*c, 100)
	e := int(base.FloorDiv64((b-d)*1e4, 306001))
	// compute return values
	day = float64(int(b-d)-base.FloorDiv(306001*e, 1e4)) + f
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

// JDToTime takes a JD and returns a Go time.Time value.
func JDToTime(jd float64) time.Time {
	// time.Time is always Gregorian
	y, m, d := jdToCalendarGregorian(jd)
	t := time.Date(y, time.Month(m), 0, 0, 0, 0, 0, time.UTC)
	return t.Add(time.Duration(d * 24 * float64(time.Hour)))
}

// TimeToJD takes a Go time.Time and returns a JD as float64.
//
// Any time zone offset in the time.Time is ignored and the time is
// treated as UTC.
func TimeToJD(t time.Time) float64 {
	ut := t.UTC()
	y, m, _ := ut.Date()
	d := ut.Sub(time.Date(y, m, 0, 0, 0, 0, 0, time.UTC))
	// time.Time is always Gregorian
	return CalendarGregorianToJD(y, int(m), float64(d)/float64(24*time.Hour))
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
