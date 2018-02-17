// Copyright 2012 Sonia Keys
// License: MIT

// Hints: Chapter 1, Hints and Tips.
//
// This is a documentation only package.  Code suggested by this chapter
// is in the package base.  The sections of the chapter are briefly
// discussed here.
//
// Trigonometric functions of large angles:  The function base.PMod reduces
// angles (or any floating point quantity) to a range from 0 to a given
// positive number.  This satisfies the suggestion of this section, but see
// the Go examples for this function.  Reducing the range of a number may
// or may not offer accuracy advantages.
//
// Angle modes:  Functions in the standard Go math library work in radians.
// To avoid inefficiencies of repeated conversions, all packages of this
// library work with angle and angle-like quantities in radians.  Formulas
// in the book typically take degrees.  These formulas are generally translated
// to radians so that as much work as possible can be done in radians.  It
// is true that IO will often be done in degrees, but IO is not speed limiting.
// For computational reasons, interal formats, function arguments, and return
// values are almost always radians.
//
// Right ascensions:  The base package has the function FromSexa for ingesting
// sexagesimal quantities such as right ascensions, and it defines types for
// angles, hour angles, right ascensions, and times.  Example 1.a of this
// chapter is implemented as a package example below.
//
// The correct quadrant:  Go has a complete set of inverse trigonometric
// functions, including math.Atan2.
//
// The input of negative angles:  The function base.FromSexa has a "neg"
// parameter to negate the overall quantity.  The numeric components are
// generally passed as positive numbers.  This avoids the problems described
// in the text.
//
// Powers of time:  Meeus offers some cautionary anecdotes here, but provides
// no concrete rules to follow, no algorithms for determining when periodic
// terms might be neglected.  There is no code from this section.
//
// Avoiding powers:  The function base.Horner implements Horner's method,
// and is used heavily by other packages.
//
// To shorten a program:  Shortening a program is rarely a goal these days.
// Computer storage is huge and program or data size is almost never a limit.
//
// Go's features for included data are not much like
// the DATA statement of BASIC.  Generally packages of this library which
// include lists or tables of data simply encode them as statically typed
// literals.  See deltat.table10A for an example of a table encoded as a
// literal.  (Click on the function name Interp10A when viewing the
// documentation in a browser, and a source window opens.  Scroll up
// slightly to see the table.)
//
// Another possibility in Go is to encode data as "raw strings"
// (see http://golang.org/ref/spec#String_literals) and them parse them
// as needed.  This technique can be useful if you have data in a text file
// which is provided to you from some external source.  Pasting the text
// file as a large raw string can avoid transcription errors that might occur
// if you manually reformat data as typed literals.  Consider though, the
// run-time cost of parsing the data.  Even if it is done at program startup,
// it can lead to significant program startup times.  Ultimately you might
// write a standalone utility program that reads the text file and emits a Go
// source file containing the data reformatted as literals.
//
// Safety tests: There seems nothing unique to astronomy here.  Write correct
// code, check all errors, handle all cases, code defensively.  For those new
// to Go, read up on error handling philosophy of Go.  We don't use assertions,
// we don't use exceptions as error handlers.
//
// Debugging.  Go has numerous features to avoid errors and help programmers
// write error free code.  Exception conditions terminate the program and
// identify the site of the failure.  If the bug is not obvious from examining
// the code at the site of failure, it then rarely takes more than a test-print
// or two to find a bug.  For more traditional debugging, the compiled
// executables include data to enable debugging with popular debuggers.
// You can swap numbers with x, y = y, x.
//
// Checking the results:  See the go test feature.
package hints
