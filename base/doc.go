// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Base: Functions and other definitions useful with multiple packages.
//
// The book Astrononomical Algorithms (AA) begins with an unnumbered chapter
// titled "Some Symbols and Abbreviations."  In addition to a list of symbols
// and abbreviations are a few paragraphs introducing sexagesimal notation.
// Chapter 1, Hints and Tips contains additional information about sexagesimal
// numbers.  It made sense to combine these in one package.  Also here
// are various definitions and support functions useful in multiple chapters.
//
// Sexagesimal types
//
// Of great use are four types for quantities commonly expressed
// in sexagesimal format: Angle, HourAngle, RA, and Time.
// The underlying type of each is float64.  The unit for Angle, HourAngle,
// and RA is radians.  The unit for Time is seconds.
// Each type has a constructor that takes sexegesimal components.  Each type
// also has a method, Rad or Seconds, that simply returns the underlying type.
// Being based on float64s, these types are relatively efficient.
//
// Custom formatters
//
// Parallel to the four types just described are four types with custom
// formatters that produce sexagesimal formattng.  These types are FmtAngle,
// FmtHourAngle, FmtRA, and FmtTime.  These types are structs with methods
// that have pointer receivers.  There is more overhead with these types
// than with the basic Angle, HourAngle, RA, and Time types.
//
// Unit indicators
//
// The symbols used for degrees, minutes, and seconds for the FmtAngle type
// are taken from the package variable DMSUnits.  The symbols for
// hours, minutes, and seconds for the FmtHourAngle, FmtRA, and FmtTime
// types are taken from HMSUnits.
//
// Decimal unit indication
//
// The decimal separator, if it appears, is always in the last segment.
// Symbols used for decimal separators are taken from package variables
// DecSep and DecCombine.
//
// Three conventions are supported for unit indication on the decimal segment.
// By default (with %v, for example) the unit follows the segment.
//
// 1°23′45.6″
//
// Described on AA p.6 is a convention for placing a unit symbol directly above
// the decimal point of a decimal number. This sometimes can be approximated in
// Unicode with codes of the category "Mn", for example "combining dot below"
// u+0323.  Example (that may or may not look right*)
//
// 1°23′45″̣6
//
// For cases where software does not render this satisfactorily, an
// alternative convention is to simply insert the unit symbol ahead of the
// decimal separator as in
//
// 1°23′45″.6
//
//   * Footnote about combining dot.  The combining dot only looks right
//     to the extent that software (such as fonts and browsers) can render it.
//     See http://www.unicode.org/faq/char_combmark.html#12b for a description
//     of the issues.  It seems that monospace fonts are more problematic.
//     The examples above are aligned flush left to avoid godoc coding
//     them monospace in the HTML.  For example 1°23′45″̣6 is less likely to
//     look right.  Other contexts likely to use monospace fonts and so likely
//     to have trouble with the combining dot are operating system shells and
//     source code text editors.
//
// Format specifiers
//
// The syntax of a format specifier is
//
//    %[flags][width][.precision]verb
//
// The syntax is set by the Go fmt package, but this package customizes
// the meaning of all format specifier components.
//
// Verbs specify one of the above decimal unit conventions and also the unit
// of the decimal (right most) segment.  The decimal unit determines the
// the potential number of segments.  Full sexagesimal format has three
// segments with the decimal separator in seconds.  Decimal minutes format has
// an hour or degrees segment, a minutes segment with the decimal separator,
// and no seconds segment.  Decimal hour or degree format has only a single
// decimal segment.
//
// This table gives the verbs for the combinations of decimal unit indication
// and decimal segment:
//
//    decimal-unit indication:             following  combined  inserted
//
//    three segments, decimal in seconds:      %s        %c        %d
//    two segments, decimal in minutes:        %m        %n        %o
//    one segment, decimal in hr/degs:         %h        %i        %j
//
// Also %v is equivalent to %s.
//
// The following flags are supported:
//  +   always print leading sign
//  ' ' (space) leave space for elided + sign
//  #   display all segments, even if 0
//  0   pad displayed segments with leading zeros
//
// A + flag takes precedence over a ' ' (space) flag.
//
// The # flag forces output to have all segments, even if 0.  Without it,
// leading zero segments are elided.  (Consider formatting coordinates with #;
// distances and durations without.)
//
// The 0 flag pads with a leading zero on non-first (sexagesimal) segments.
// If a width is specfied, the 0 flag pads with leading zeros on the first
// (hr/deg) segment as well.
//
// For the RA type, sign formatting flags '+' and ' ' are ignored.
//
// Specifying width forces a fixed width format.  Flag '#' is implied, ' ' is
// implied unless '+' is given, and segments are space padded unless '0' is
// given.  The width number specifies the number of digits in the integer part
// of the most significant segment, hours or degrees — not the total width.
// For example you would typically use the number 2 for RA, 3 for longitude.
// Also with fixed width consider avoiding the combining dot verbs.
// (See note above on rendering of the combining dot.)  With fixed width
// sexagesimal formats, the sign indicator is always the left-most column;
// with fixed width space padded decimal hour or degree formats, the sign
// indicator is formatted immediately in front of the number within the
// space padded field.
//
// Precision specifies the number of places past the decimal separator
// of the decimal segment.  The default is 0.  There is no variable precision
// format.
//
// Errors
//
// A value that cannot be expressed the in the requested format represents
// an overflow condition.  In this case, the custom formatters emit all
// asterisks "*************" and leave a more descriptive error in the
// Err field of the value.
//
// If you specifiy width, digits of the integer part of the first segment must
// fit in the specified width.  Larger values cause overflow.
//
// Overflow also happens if more precision is requested than is represented
// in the underlying float64.  In the case of an angle formatted with the
// decimal separator in seconds, precision of 15 is possible only for angles
// less than a few arc seconds.  As angle values increase, fewer digits of
// precision are possible.  At one degree, you can get 12 digits of precision
// in the seconds segment of a full sexagesimal number, at 360 degrees,
// you can get 9.  For all formats, an angle too large for the specified
// precision causes overflow.
//
// +Inf, -Inf, and NaN always cause overflow.
//
// Only errors related to the value being formatted are handled as overflow.
// Errors of format specification are handled with the standard Printf
// convention of emitting the error in the formatted result.
//
// Bessellian and Julian Year
//
// Chapter 21, Precession actually contains these defintions.  They are moved
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
package base
