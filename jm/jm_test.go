// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package jm_test

import (
	"fmt"
	"time"

	"github.com/soniakeys/meeus/jm"
)

func ExampleJewishCalendar() {
	// Example 9.a, p. 73.
	A, mP, dP, mNY, dNY, months, days := jm.JewishCalendar(1990)
	fmt.Println("Jewish Year:", A)
	fmt.Println("Pesach:     ", time.Month(mP), dP)
	fmt.Println("New Year:   ", time.Month(mNY), dNY)
	fmt.Println("Months:     ", months)
	fmt.Println("Days:       ", days)
	// Output:
	// Jewish Year: 5750
	// Pesach:      April 10
	// New Year:    September 20
	// Months:      12
	// Days:        354
}

func ExampleMoslemToJulian() {
	// Example 9.b, p. 75, conversion to Julian.
	y, dn := jm.MoslemToJulian(1421, 1, 1)
	fmt.Println(y, dn)
	// Output:
	// 2000 84
}

func ExampleJulianToGregorian() {
	// Example 9.b, p. 75, conversion to Gregorian.
	y, m, d := jm.JulianToGregorian(2000, 84)
	fmt.Println(d, time.Month(m), y)
	// Output:
	// 6 April 2000
}

func ExampleMoslemLeapYear() {
	// Example 9.b, p. 75, indication of leap year.
	if jm.MoslemLeapYear(1421) {
		fmt.Println("Moslem year 1421 is a leap year of 355 days.")
	} else {
		fmt.Println("Moslem year 1421 is a common year of 354 days.")
	}
	// Output:
	// Moslem year 1421 is a common year of 354 days.
}

func ExampleGregorianToJulian() {
	// Example 9.c, p. 76, conversion to Julian Calendar.
	y, m, d := jm.GregorianToJulian(1991, 8, 13)
	fmt.Println(y, time.Month(m), d, "Julian")
	// Output:
	// 1991 July 31 Julian
}

func ExampleJulianToMoslem() {
	// Example 9.c, p. 76, final output.
	y, m, d := jm.JulianToMoslem(1991, 7, 31)
	fmt.Println(d, jm.MMonth(m), "of A.H.", y)
	// Output:
	// 2 Ṣafar of A.H. 1412
}
