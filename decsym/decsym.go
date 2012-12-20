// Copyright 2012 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Decsym contains functions for combining unit symbols with decimal points.
//
// AA begins with an unnumbered chapter titled "Some Symbols and
// Abbreviations."  Described on p.6 is a convention for placing a
// unit symbol directly above the decimal point of a decimal number.
// This can be done with Unicode by replacing the decimal point with
// the unit symbol and "combining dot below," u+0323.  The function
// Combine here performs this substitution.  Of course this only
// works to the extent that software can render the combining character.
// For cases where rendering software fails badly, Add is provided
// as a compromise.  It does not use the combining dot but simply places
// the unit symbol ahead of the decimal point.  Numbers modified with either
// function can be returned to their original form with Strip.
package decsym

import (
	"strings"
	"unicode/utf8"
)

// Add adds a symbol representing units to a formatted decimal number.
// The symbol is added just before the decimal point.
func Add(d string, sym rune) string {
	i := strings.IndexRune(d, '.')
	if i < 0 {
		return d // no decimal point
	}
	// insert c before decimal point
	return d[:i] + string(sym) + d[i:]
}

// Combine is like Add but replaces the decimal point with the Unicode
// combining dot below (u+0323) so that it will combine with the added symbol.
func Combine(d string, sym rune) string {
	i := strings.IndexRune(d, '.')
	if i < 0 {
		return d // no decimal point
	}
	// insert c, replace decimal point with combining dot below
	return d[:i] + string(sym) + "̣" + d[i+1:]
}

// Strip reverses the action of Add or Combine,
// removing the specified unit symbol and restoring a combining dot
// to an ordinary decimal point.
func Strip(d string, sym rune) string {
	sl := utf8.RuneLen(sym)
	if i := strings.IndexRune(d, sym); i >= 0 {
		if i < len(d)-sl && d[i+sl] == '.' {
			// ordinary decimal point following unit
			return d[:i] + d[i+sl:]
		}
		if i < len(d)-sl-1 && d[i+sl:i+sl+2] == "̣" {
			// combining dot below following unit
			return d[:i] + "." + d[i+sl+2:]
		}
	}
	// otherwise don't mess with it
	return d
}
