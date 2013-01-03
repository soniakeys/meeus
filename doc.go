// Copyright 2012 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Package meeus and all of its subpackages implement algorithms from
// the book "Astronomical Algorithms" (AA) by Jean Meeus.
//
// They follow the second edition, copyright 1998, with corrections as of
// August 10, 2009.  Each package in a subdirectory implements algorithms
// of a chapter of the book.  Package meeus itself implements some introductory
// material.
//
// AA begins with an unnumbered chapter titled "Some Symbols and
// Abbreviations."  In addition to a list of symbols and abbreviations
// are a few paragraphs introducing sexagesimal notation.  Chapter 1,
// Hints and Tips contains additional information about sexagesimal
// numbers.  It made sense to combine these in one package.  Also here
// are support functions useful in multiple chapters.
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
//
// Custom formatters
//
// Not described in AA, but of great use, three types are defined with custom
// formatters that produce sexagesimal formattng.  These types are the Angle,
// HourAngle, and RA types.
// The syntax of a format specifier is
//  %[flags][width][.precision]verb
//
// Verbs are s, d, c, x, and v.  The meanings are different than for
// common Go types.  Given an Angle equivalent to 1.23 seconds,
//  %.2s formats as 1.23″   (s for standard formatting)
//  %.2d formats as 1″.23   (d for decimal symbol, as in DecSymAdd)
//  %.2c formats as 1″̣23    (c for combining dot, as in DecSymCombine)
//  %.2x formats as 123     (x for space, suppresses unit symbols and decimal point)
//  %v formats the same as %s
//
// The following flags are supported:
//  + always print leading sign
//  ' ' (space) leave space for elided sign
//  # display all three segments, even if 0
//  0 pad all segments with leading zeros
//
// A + flag takes precedence over a ' ' (space) flag.
// The # flag forces all formatted strings to have three numeric components,
// an hour or degree, a minute, and a second.  Without the # flag, small vaues
// will have zero values of hours, degrees, or minutes elided.
// The 0 flag pads with leading zeros on minutes and seconds, and if a
// width is specfied, leading zeros on the first segment as well.
// For the RA type, sign formatting flags '+' and ' ' are ignored.
//
// Width specifies the number of digits in the most significant segment,
// degrees or hours (not the total width of all three segments.)
// Precision specifies the number of places past the decimal point
// of the last (seconds) segment.
//
// Precision specifies the number of places to display past the decimal point.
//
// To ensure fixed width output, use one of the + or ' ' (space) flags,
// use the 0 flag, and use a width.
//
// The symbols used for degrees, minutes, and seconds for the Angle type
// are taken from the package variable DMSRunes.  The symbols for
// hours, minutes, and seconds for the HourAngle and RA types are taken
// from HMSRunes.
//
// Width Errors
//
// For various types of overflow, the custom formatters emit all asterisks
// "*************" and leave an exact error in the WidthError field of the
// type.
//
// Precision is limited to the range [0,15].  Values outside of that range
// will cause this overflow condition.
//
// Precision of 15 is possible only for angles less than a few arc seconds.
// As angle values increase, fewer digits of precision are possible.  At one
// degree, you can get 12 digits of precision, at 360 degrees, you can get 9.
// An angle too large for the specified precision causes overflow.
//
// If you specifiy width, the first segment, degrees or hours, must fit in the
// specified width.  Larger values cause overflow.
//
// +Inf, -Inf, and NaN always cause overflow.
package meeus
