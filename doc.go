// Meeus implements algorithms from the book "Astronomical Algorithms" by Jean Meeus.
//
// It follows the second edition, copyright 1998, with corrections as of
// August 10, 2009.
//
// Library Goals
//
// Jean Meeus's book has long been respected as a broad-reaching source of
// astronomical algorithms, and many code libraries have been based on it.
// This library will be distinct in several respects, I hope.
//
// First of all it is in the Go language, a programming language new enough
// that it well postdates the book itself.  Go has many advantages for a
// large and diverse library such as this and it is a fine language for
// for scientific computations.  I hope that a Go implementation will prove
// relevant for some time in the future.
//
// Next, this library attempts fairly comprehensive coverage of the book.
// Each chapter of the book is addressed, and in the very few cases where
// there seems no code from the chapter that is applicable in Go, similar
// and more appropriate techniques are at least discussed in documentation.
// If meaningful, examples are given as well.
//
// While this library attempts fairly comprehensive coverage of the book,
// it does not attempt to present a complete, well rounded, and polished
// astronomy library.  Such a production-quality astronomy library would
// likely include some updated routines and data, routines and data from
// other sources, and would fill in various holes of functionality which
// Meeus elects to gloss over.  Such a library could certainly be derived
// from this one, but it is beyond the scope of what is attempted here.
//
// Thus, this library should represent a solid foundation for the development
// of a broad range of astronomy software.  Much software should be able to
// use this library directly.  Some software will need routines from additional
// sources.  When the API of this library begins to present friction with that
// of other code, it may be time to fork or otherwise derive a new library
// from this one.  Please feel free to do this, respecting of course the MIT
// license under which this software is offered.
//
// Package Contents
//
// By Go convention, each package is in its own subdirectory.  The
// "subdirectories" list of this documentation page lists all packages of
// of the library.  Each package also corresponds to exacly one chapter of
// the book.  The package documenation heading references the chapter number
// and a cross reference is given below of chapter numbers and package names.
//
// Within a chapter of the book, Meeus presents explanatory text, numbered
// formulas, numbered examples, and other exercises.  Within a package of this
// library, there are library functions and other codified definitions; there
// are Go examples which appear in documentation and which are also evaluated
// and verified to produce correct output by the go test feature; and there is
// test code which is neither part of the API nor the documentation but which
// verified by the go test feature.
//
// The "API", or choice of functions to implement in Go, covers nearly all of
// Meeus's numbered formulas.  The correspondence is not one-to-one, but often
// "refactored" into functions that seem more idiomatic to Go.  Also if an
// example or exercise from the book illustrates an algorithm that seems of
// general utility, it will be included in the API as well.  This is set as
// the limit of the API however, and thus the limit of the functionality
// offered by this library.
//
// Each numbered example in the book is also translated to a Go example
// function.  This typically shows how to use the implemented API to compute
// the results of the example.  As the go test feature validates these
// results, the examples also serve as baseline tests of the correctness
// of the API code.  Relevant "exercises" from the book are also often
// implemented as Go examples.
//
// In addition to the chapter packages, there is a package called "common".
// This contains a few definitions that are provided by Meeus but are of such
// general use that they really don't belong with any one chapter.  The much
// greater bulk of common however, are functions which Meeus does not explicity
// provide, but again are of general use.  The nature of these functions is
// as helper subroutines or IO subroutines.  The functions do not offer
// additional astronomy algorithms beyond those provided by Meeus.
//
// Chapter Cross-reference
//
// .
//
//	Chapter                                              Package
//
//	1. Hints and Tips                                    hints
//	3. Interpolation                                     interp
//	4. Curve Fitting                                     fit
//	7. Julian Day                                        julian
//	8. Date of Easter                                    easter
//	10. Dynamical Time and Universal Time                deltat
//	11. The Earth's Globe                                globe
//	12. Sidereal Time at Greenwich                       sidereal
//	13. Transformation of Coordinates                    coord
//	14. The Parallactic Angle, and three other Topics    parallactic
//	15. Rising, Transit, and Setting                     rise
//	16. Atmospheric Refraction                           refraction
//	17. Angular Separation                               angle
//	18. Planetary Conjunctions                           conjunction
//	21. Precession                                       precess
//	22. Nutation and the Obliquity of the Ecliptic       nutation
package meeus
