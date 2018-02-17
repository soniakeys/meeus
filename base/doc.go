// Copyright 2013 Sonia Keys
// License: MIT

// Base: Functions and other definitions useful with multiple packages.
//
// Base contains various definitions and support functions useful in multiple
// chapters.
//
// Bessellian and Julian Year
//
// Chapter 21, Precession actually contains these definitions.  They are moved
// here because of their general utility.
//
// Chapter 22, Nutation contains the function for Julian centuries since J2000.
//
// Phase angle functions
//
// Two functions, Illuminated and Limb, concern the illumnated phase of a body
// and are given in two chapters, 41 an 48.  They are collected here because
// the identical functions apply in both chapters.
//
// General purpose math functions
//
// SmallAngle is recommended in chapter 17, p. 109.
//
// PMod addresses the issue on p. 7, chapter 1, in the section "Trigonometric
// functions of large angles", but the function is not written to be specific
// to angles and so has more general utility.
//
// Horner is described on p. 10, chapter 1.
//
// FloorDiv and FloorDiv64 are optimizations for the INT function described
// on p. 60, chapter 7.
//
// Unit types and conversions
//
// While the library uses float64s for many parameter and return values, it
// sometimes uses a defined type to clarify interpretation of a value.
// Types Angle, RA, HourAngle, and Time are defined here for this purpose.
// These defined types have a number of constructors and methods that perform
// useful conversions.
//
// Also the function FromSexa takes sexagesimal components such as degrees
// minutes and seconds and returns a single value.
package base
