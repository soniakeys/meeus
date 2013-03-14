// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Accuracy: Chapter 2, About accuracy.
//
// This is a documentation only package.  The book has good advice for
// program authors.  Within this library, the sexagesimal formatting routines
// of the package common have some features for managing accuracy on output.
//
// Precision can specified for these functions, but as both the specified
// precision and the numeric quantity increase, at some point the underlying
// float64 does not have significant bits to support the requested precision.
// At this point all asterisks are output as a loss of precision indicator.
// See package documentation for more details.
//
// Also for the case where seconds or even minutes might not be significant
// in a sexagesimal quantity, there is a feature by which output can be
// rounded and display of the insignificant sexagesimal components is
// suppressed.  Again, see package documentation for more details.
//
// In general, packages use the Go float64 type throughout.  Float64 implements
// IEEE-754 64 bit floating point numbers fairly closely.  See the Go
// programming language reference and the IEEE-754 reference for more details.
package accuracy
