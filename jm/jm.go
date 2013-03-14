// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// JM: Chapter 9, Jewish and Moslem Calendars.
//
// The Jewish calendar routines are implemented as a monolithic function,
// because computations of the various results build off of common
// intermediate results.
//
// The Moslem calendar routines break down nicely into some separate functions.
//
// Included in these are two functions that convert between Gregorian and
// Julian calendar days without going through Julian day (JD).  As such,
// I suppose, these or similar routines are not in chapter 7, Julian Day.
// Package base might also be a suitable place for these, but I'm not sure
// they are used anywhere else in the book.  Anyway, they have the quirk
// that they are not direct inverses:  JulianToGregorian returns the day number
// of the day of the Gregorian year, but GregorianToJulian wants the Gregorian
// month and day of month as input.
package jm

import (
	"math"

	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/meeus/julian"
)

// JewishCalendar returns interesting dates and facts about a given year.
//
// Input is a Julian or Gregorian year.
//
// Outputs:
//	A:      Year number in the Jewish Calendar
//	mP:     Month number of Pesach.
//	dP:     Day number of Pesach.
//	mNY:    Month number of the Jewish new year.
//	dNY:    Day number of the Jewish new year.
//	months: Number of months in this year.
//	days:   Number of days in this year.
func JewishCalendar(y int) (A, mP, dP, mNY, dNY, months, days int) {
	A = y + 3760
	D := bigD(y)
	mP = 3
	dP = D
	if dP > 31 {
		mP++
		dP -= 31
	}
	// A simplification of Meeus's rule to add 163 days.  Months of Pesach
	// are either March or April with D based off of March.  Months of New
	// year are either September or August so D+163-(days from March to
	// September == 184) = D-21 must be based off of September.
	mNY = 9
	dNY = D - 21
	if dNY > 30 {
		mNY++
		dNY -= 30
	}
	months = 12
	switch A % 19 {
	case 0, 3, 6, 8, 11, 14, 17:
		months++
	}
	// Similarly, A simplification of Meeus's rule to take the difference
	// in calendar days from NY of one year to NY of the next.  NY is based
	// on D, so difference in D is difference in day numbers of year.  Result
	// is sum of this number and the number of days in the Western calandar
	// year.
	y1 := y + 1
	lf := julian.LeapYearGregorian
	if y1 < 1583 {
		lf = julian.LeapYearJulian
	}
	days = 365
	if lf(y1) {
		days++
	}
	days += bigD(y1) - D
	return
}

func bigD(y int) int {
	C := base.FloorDiv(y, 100)
	var S int
	if y >= 1583 {
		S = base.FloorDiv(3*C-5, 4)
	}
	a := (12*y + 12) % 19
	b := y % 4
	Q := -1.904412361576 + 1.554241796621*float64(a) + .25*float64(b) -
		.003177794022*float64(y) + float64(S)
	fq := math.Floor(Q)
	iq := int(fq)
	j := (iq + 3*y + 5*b + 2 - S) % 7
	r := Q - fq
	var D int
	switch {
	case j == 2 || j == 4 || j == 6:
		D = iq + 23
	case j == 1 && a > 6 && r >= .63287037:
		D = iq + 24
	case j == 0 && a > 11 && r >= .897723765:
		D = iq + 23
	default:
		D = iq + 22
	}
	return D
}

// MoslemToJulian converts a Moslem calandar date to a Julian year and day number.
func MoslemToJulian(y, m, d int) (jY, jDN int) {
	N := d + base.FloorDiv(295001*(m-1)+9900, 10000)
	Q := base.FloorDiv(y, 30)
	R := y % 30
	A := base.FloorDiv(11*R+3, 30)
	W := 404*Q + 354*R + 208 + A
	Q1 := base.FloorDiv(W, 1461)
	Q2 := W % 1461
	G := 621 + 28*Q + 4*Q1
	K := base.FloorDiv(Q2*10000, 3652422)
	E := base.FloorDiv(3652422*K, 10000)
	J := Q2 - E + N - 1
	X := G + K
	switch {
	case J > 366 && X%4 == 0:
		J -= 366
		X++
	case J > 365 && X%4 > 0:
		J -= 365
		X++
	}
	return X, J
}

// JulianToGregorian converts a Julian calendar year and day number to a year, month, and day in the Gregorian calendar.
func JulianToGregorian(y, dn int) (gY, gM, gD int) {
	JD := base.FloorDiv(36525*(y-1), 100) + 1721423 + dn
	α := base.FloorDiv(JD*100-186721625, 3652425)
	β := JD
	if JD >= 2299161 {
		β += 1 + α - base.FloorDiv(α, 4)
	}
	b := β + 1524
	return ymd(b)
}

func ymd(b int) (y, m, d int) {
	c := base.FloorDiv(b*100-12210, 36525)
	d = base.FloorDiv(36525*c, 100) // borrow the variable
	e := base.FloorDiv((b-d)*10000, 306001)
	// compute named return values
	d = b - d - base.FloorDiv(306001*e, 10000)
	if e < 14 {
		m = e - 1
	} else {
		m = e - 13
	}
	if m > 2 {
		y = c - 4716
	} else {
		y = c - 4715
	}
	return
}

// MoslemLeapYear returns true if year y of the Moslem calendar is a leap year.
func MoslemLeapYear(y int) bool {
	R := y % 30
	return (11*R+3)%30 > 18
}

// GregorianToJulian takes a year, month, and day of the Gregorian calendar and returns the equivalent year, month, and day of the Julian calendar.
func GregorianToJulian(y, m, d int) (jy, jm, jd int) {
	if m < 3 {
		y--
		m += 12
	}
	α := base.FloorDiv(y, 100)
	β := 2 - α + base.FloorDiv(α, 4)
	b := base.FloorDiv(36525*y, 100) +
		base.FloorDiv(306001*(m+1), 10000) +
		d + 1722519 + β
	return ymd(b)
}

// JulianToMoslem takes a year, month, and day of the Julian calendar and returns the equivalent year, month, and day of the Moslem calendar.
func JulianToMoslem(y, m, d int) (my, mm, md int) {
	W := 2
	if y%4 == 0 {
		W = 1
	}
	N := base.FloorDiv(275*m, 9) - W*base.FloorDiv(m+9, 12) + d - 30
	A := y - 623
	B := base.FloorDiv(A, 4)
	C2 := func(A int) int {
		C := A % 4
		C1 := 365.25001 * float64(C)
		C2 := math.Floor(C1)
		if C1-C2 > .5 {
			return int(C2) + 1
		}
		return int(C2)
	}(A)
	Dp := 1461*B + 170 + C2
	Q := base.FloorDiv(Dp, 10631)
	R := Dp % 10631
	J := base.FloorDiv(R, 354)
	K := R % 354
	O := base.FloorDiv(11*J+14, 30)
	H := 30*Q + J + 1
	JJ := K - O + N - 1
	days := 354
	if MoslemLeapYear(y) {
		days++
	}
	if JJ > days {
		JJ -= days
		H++
	}
	if JJ == 355 {
		mm = 12
		md = 30
	} else {
		S := base.FloorDiv((JJ-1)*10, 295)
		mm = 1 + S
		md = base.FloorDiv(10*JJ-295*S, 10)
	}
	return H, mm, md
}

// An MMonth specifies a month of the Moslum Calendar (Muharram = 1, ...).
//
// This type is modeled after the Month type of the time package in the
// Go standard library.
type MMonth int

// Upgraded to Unicode from the spellings given by Meeus.
// Source: http://en.wikipedia.org/wiki/Islamic_calendar.
var mmonths = [12]string{
	"Muḥarram",
	"Ṣafar",
	"Rabīʿ I",
	"Rabīʿ II",
	"Jumādā I",
	"Jumādā II",
	"Rajab",
	"Shaʿbān",
	"Ramaḍān",
	"Shawwāl",
	"Dhū al-Qaʿda",
	"Dhū al-Ḥijja",
}

// String returns the Romanization of the month ("Muḥarram", "Ṣafar", ...).
func (m MMonth) String() string { return mmonths[m-1] }
