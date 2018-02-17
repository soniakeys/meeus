// Copyright 2013 Sonia Keys
// License: MIT

// Meeus implements algorithms from the book "Astronomical Algorithms"
// by Jean Meeus.
//
// It follows the second edition, copyright 1998, with corrections as of
// August 10, 2009.
//
// It requires Go 1.1 or later.
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
// the book.  The package documentation heading references the chapter number
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
// The "API", or choice of functions to implement in Go, covers many of
// Meeus's numbered formulas and covers the algorithms needed to work most
// of the numbered examples.  The correspondence is not one-to-one, but often
// "refactored" into functions that seem more idiomatic to Go.  This is set
// as the limit of the API however, and thus the limit of the functionality
// offered by this library.
//
// Each numbered example in the book is also translated to a Go example
// function.  This typically shows how to use the implemented API to compute
// the results of the example.  As the go test feature validates these
// results, the examples also serve as baseline tests of the correctness
// of the API code.  Relevant "exercises" from the book are also often
// implemented as Go examples.
//
// A few packages remain incomplete.  A package is considered complete if
// it implements all major formulas and algorithms and if it implements all
// numbered examples.  For incomplete packages, the package documentation
// will describe the ways in which it is incomplete and typically give reasons
// for the incompleteness.
//
// In addition to the chapter packages, there is a package called "base".
// This contains a few definitions that are provided by Meeus but are of such
// general use that they really don't belong with any one chapter.  The much
// greater bulk of base however, are functions which Meeus does not explicitly
// provide, but again are of general use.  The nature of these functions is
// as helper subroutines or IO subroutines.  The functions do not offer
// additional astronomy algorithms beyond those provided by Meeus.
//
// Identifiers
//
// To more closely follow the book's use of Greek letters and other symbols,
// Unicode is freely used in the source code.  Recognizing that these symbols
// are awkard to enter in many environments however, they are avoided for
// exported symbols that comprise the library API.  The function Coord.EclToEq
// for example, returns (α, δ float64) but of course you can assign these
// return values to whatever variables you like.  The struct Coord.Equatorial
// on the other hand, has exported fields RA and Dec.  ASCII is used in this
// case to simplify using these symbols in your code.
//
// Some identifiers use the prime symbol (ʹ).  That's Unicode U+02B9,
// not the ASCII '.  Go uses ASCII ' for raw strings and does not allow it
// in identifiers.  U+02B9 on the other hand is Unicode category Lm,
// and is perfectly valid in Go identifiers.
//
// Unit types
//
// An earler version of this library used the Go type float64 for most
// parameters and return values.  This allowed terse, efficient code but
// required careful attention to the scaling or units used.  Go defined types
// are now used for Time, RA, HourAngle, and general Angle quantities in the
// interest of making units and coversions more clear.  These types are
// defined in the external package github.com/soniakeys/unit.
//
// Sexagesimal formatting
//
// An earlier version of this library included routines for formatting
// sexagesimal quantities.  These have been moved to the external package
// github.com/soniakeys/sexagesimal and use of this package is now restricted
// to examples and tests.
//
// Meeus packages and the sexagesimal package both depend on the unit package.
// Meeus packages do not depend on sexagesimal, although the Meeus tests do.
//
// Chapter Cross-reference
//
// .
//
//	Chapter                                                 Package
//
//	1.  Hints and Tips                                      hints
//	2.  About Accuracy                                      accuracy
//	3.  Interpolation                                       interp
//	4.  Curve Fitting                                       fit
//	5.  Iteration                                           iterate
//	6.  Sorting Numbers                                     sort
//	7.  Julian Day                                          julian
//	8.  Date of Easter                                      easter
//	9.  Jewish and Moslem Calendars                         jm
//	10. Dynamical Time and Universal Time                   deltat
//	11. The Earth's Globe                                   globe
//	12. Sidereal Time at Greenwich                          sidereal
//	13. Transformation of Coordinates                       coord
//	14. The Parallactic Angle, and three other Topics       parallactic
//	15. Rising, Transit, and Setting                        rise
//	16. Atmospheric Refraction                              refraction
//	17. Angular Separation                                  angle
//	18. Planetary Conjunctions                              conjunction
//	19. Bodies in Straight Line                             line
//	20. Smallest Circle containing three Celestial Bodies   circle
//	21. Precession                                          precess
//	22. Nutation and the Obliquity of the Ecliptic          nutation
//	23. Apparent Place of a Star                            apparent
//	24. Reduction of Ecliptical Elements from one Equinox   elementequinox
//	    to another one
//	25. Solar Coordinates                                   solar
//	26. Rectangular Coordinates of the Sun                  solarxyz
//	27. Equinoxes and Solstices                             solstice
//	28. Equation of Time                                    eqtime
//	29. Ephemeris for Physical Observations of the Sun      solardisk
//	30. Equation of Kepler                                  kepler
//	31. Elements of Planetary Orbits                        planetelements
//	32. Positions of the Planets                            planetposition
//	33. Elliptic Motion                                     elliptic
//	34. Parabolic Motion                                    parabolic
//	35. Near-parabolic Motion                               nearparabolic
//	36. The Calculation of some Planetary Phenomena         planetary
//	37. Pluto                                               pluto
//	38. Planets in Perihelion and in Aphelion               perihelion
//	39. Passages through the Nodes                          node
//	40. Correction for Parallax                             parallax
//	41. Illuminated Fraction of the Disk and Magnitude      illum
//	    of a Planet
//	42. Ephemeris for Physical Observations of Mars         mars
//	43. Ephemeris for Physical Observations of Jupiter      jupiter
//	44. Positions of the Satellites of Jupiter              jupitermoons
//	45. The Ring of Saturn                                  saturnring
//	46. Positions of the Satellites of Saturn               saturnmoons
//	47. Position of the Moon                                moonposition
//	48. Illuminated Fraction of the Moon's Disk             moonillum
//	49. Phases of the Moon                                  moonphase
//	50. Perigee and apogee of the Moon                      apsis
//	51. Passages of the Moon through the Nodes              moonnode
//	52. Maximum declinations of the Moon                    moonmaxdec
//	53. Ephemeris for Physical Observations of the Moon     moon
//	54. Eclipses                                            eclipse
//	55. Semidiameters of the Sun, Moon, and Planets         semidiameter
//	56. Stellar Magnitudes                                  stellar
//	57. Binary Stars                                        binary
//	58. Calculation of a Planar Sundial                     sundial
package meeus
