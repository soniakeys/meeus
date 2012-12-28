// Copyright 2012 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Meeus implements algorithms from the book "Astronomical Algorithms" (AA)
// by Jean Meeus.
//
// It follows the second edition, copyright 1998, with corrections as of
// August 10, 2009.  Each package in a subdirectory implements algorithms
// of a chapter of the book.
//
// AA begins with an unnumbered chapter titled "Some Symbols and
// Abbreviations."  In addition to a list of symbols and abbreviations
// are a few paragraphs introducing sexagesimal notation.  Chapter 1,
// Hints and Tips contains additional information about sexagesimal
// numbers.  It made sense to combine these in one package.
//
// Package meeus contains:
//	Routines inspired by the initial unnamed chapter.
//	Routines from Chapter1, Hints and Tips.
//	Additional routines that are applicable to multiple chapters.
//
// Decimal Symbols
//
// Described on p.6 is a convention for placing a
// unit symbol directly above the decimal point of a decimal number.
// This can be done with Unicode by replacing the decimal point with
// the unit symbol and "combining dot below," u+0323.  The function
// DecSymCombine here performs this substitution.  Of course this only
// works to the extent that software can render the combining character.
// For cases where rendering software fails badly, DecSymAdd is provided
// as a compromise.  It does not use the combining dot but simply places
// the unit symbol ahead of the decimal point.  Numbers modified with either
// function can be returned to their original form with DecSymStrip.
package meeus

import ()

// FloorDivInt returns the floor of x / y.
//
// It uses integer math only, so is more efficient than using floating point
// intermediate values.  This function can be used in many places where INT()
// appears in AA.  As with built in integer division, it panics with y == 0.
func FloorDivInt(x, y int) int {
	if (x < 0) == (y < 0) {
		return x / y
	}
	return x/y - 1
}

// FloorDivInt64 returns the floor of x / y.
//
// It uses integer math only, so is more efficient than using floating point
// intermediate values.  This function can be used in many places where INT()
// appears in AA.  As with built in integer division, it panics with y == 0.
func FloorDivInt64(x, y int64) int64 {
	if (x < 0) == (y < 0) {
		return x / y
	}
	return x/y - 1
}

// Cmp compares two float64s and returns -1, 0, or 1 if a is <, >, or == b,
// respectively.
//
// The name and semantics are chosen to match big.Cmp in the Go standard
// library.
func Cmp(a, b float64) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	}
	return 0
}
