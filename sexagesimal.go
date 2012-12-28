// Copyright 2012 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package meeus

import (
	"bytes"
	"fmt"
	"math"
	"strings"
	"unicode/utf8"
)

// DecSymAdd adds a symbol representing units to a formatted decimal number.
// The symbol is added just before the decimal point.
func DecSymAdd(d string, sym rune) string {
	i := strings.IndexRune(d, '.')
	if i < 0 {
		return d + string(sym) // no decimal point, append symbol
	}
	// insert c before decimal point
	return d[:i] + string(sym) + d[i:]
}

// DecSymCombine adds a symbol like DecSymAdd but replaces the decimal point
// with the Unicode combining dot below (u+0323) so that it will combine
// with the added symbol.
func DecSymCombine(d string, sym rune) string {
	i := strings.IndexRune(d, '.')
	if i < 0 {
		return d + string(sym) // no decimal point, append symbol
	}
	// insert c, replace decimal point with combining dot below
	return d[:i] + string(sym) + "̣" + d[i+1:]
}

// DecSymStrip reverses the action of DecSymAdd or DecSymCombine,
// removing the specified unit symbol and restoring a combining dot
// to an ordinary decimal point.
func DecSymStrip(d string, sym rune) string {
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
		if i+sl == len(d) {
			// no decimal point or combining dot found, but string ends
			// with sym.  just remove the symbol.
			return d[:i]
		}
	}
	// otherwise don't mess with it
	return d
}

// DMSToDeg converts from parsed sexagesimal angle components to decimal
// degrees.
func DMSToDeg(neg bool, d, m int, s float64) float64 {
	s = (float64((d*60+m)*60) + s) / 3600
	if neg {
		return -s
	}
	return s
}

// Angle, Time, and RA are defined with custom formatters.
//
// Given a value equivalent to 1.23 seconds,
//	%s formats as 1.23″   (s for standard formatting)
//	%d formats as 1″.23   (d for decimal symbol, as in DecSymAdd)
//	%c formats as 1″̣23    (c for combining dot, as in DecSymCombine)
//	%x formats as 1       (x for Fortran space, suppresses unit symbols)
//	%0.2x formats as 00 00 0123
//	%v formats the same as %s
//
// Width and precision are supported.  Precision must be <= 15.
// The following flags are supported:
//	+ always print leading sign
//	- left justify within width
//	' ' (space) leave space for elided sign
//	0 pad all segments with leading zeros
//
// The 0 flag forces all formatted strings to have three numeric components,
// An hour or degree, a minute, and a second.  Without the 0 flag, small vaues
// will have zero values of hours, degrees, or minutes elided or space padded.
// In most cases where you use the %x verb, you should also use the 0 flag.
//
// Note: for default floating point formatting, simply convert to float64.
type Angle struct {
	Rad           float64 // Angle in radians.
	WidthOverflow string  // Message valid after format.
}

func NewAngle(rad float64) *Angle {
	return &Angle{Rad: rad}
}

func (a *Angle) SetDMS(neg bool, d, m int, s float64) {
	a.Rad = DMSToDeg(neg, d, m, s) * (math.Pi / 180)
}

// implement fmt.Formatter
func (a Angle) Format(f fmt.State, c rune) {
	a.WidthOverflow = ""
	// valiate verb
	switch c {
	case 's', 'd', 'c', 'v', 'x':
	default:
		fmt.Fprintf(f, "Invalid verb for Angle: %%%c", c)
		return
	}
	// get meaningful precision
	prec, ok := f.Precision()
	if !ok {
		prec = 0
	}
	// compute a width suitable for overflow result. the default is the width
	// specified in the format spec, otherwise a reasonable minimum width.
	wid, widSpec := f.Width()
	if !widSpec {
		if f.Flag('0') {
			wid = 9 // 000 00 00
		} else {
			wid = 1 // at least one digit to left of decimal point
		}
		if prec > 0 {
			wid = prec + 1
		}
		if f.Flag(' ') || f.Flag('+') {
			wid++ // leading sign
		}
	}
	// quick sanity checks.
	// prec <= 13 keeps 60*pi exact in calculation a little bit below.
	if prec > 13 {
		a.WidthOverflow = "max allowed prec is 13"
		f.Write(bytes.Repeat([]byte{'*'}, wid)) // overflow result
		return
	}
	if math.IsNaN(a.Rad) || math.IsInf(a.Rad, 0) {
		a.WidthOverflow = "Nan or Inf"
		f.Write(bytes.Repeat([]byte{'*'}, wid)) // overflow result
		return
	}
	// work with positive value
	neg := false
	x := a.Rad
	if x < 0 {
		neg = true
		x = -x
	}
	pf := math.Pow(10, float64(prec)) // precision factor
	// xs = x in seconds, scaled to precision
	xs := x * pf * 3600 * 180 / math.Pi
	// check that we can work with as as int64 without overflow
	i := int64(xs + .5) // round
	if i >= 1<<52 || math.Abs(xs-float64(i)) >= 1 {
		a.WidthOverflow = "loss of precision"
		f.Write(bytes.Repeat([]byte{'*'}, wid)) // overflow result
		return
	}
	// compute integer values of segments
	pi := int64(pf)
	s := i % (60 * pi) // second segment, scaled to precision
	i /= 60 * pi       // i now == minutes
	m := i % 60        // minute segment
	d := i / 60        // degree segment

	// format seconds into partial result r
	sw := prec + 1
	if f.Flag('0') {
		sw++
	}
	r := fmt.Sprintf("%0*d", sw, s) // format with leading zeros
	if prec > 0 && c != 'x' {
		r = r[:len(r)-prec] + "." + r[len(r)-prec:] // insert decimal point
	}
	// add seconds unit symbol
	switch c {
	case 's', 'v':
		r += "″"
	case 'd':
		r = DecSymAdd(r, '″')
	case 'c':
		r = DecSymCombine(r, '″')
	}
	// add degrees, minutes to partial result
	if c == 'x' {
		switch {
		case f.Flag('0'):
			r = fmt.Sprintf("%03d %02d %s", d, m, r)
		case d > 0:
			r = fmt.Sprintf("%d %d %s", d, m, r)
		case m > 0:
			r = fmt.Sprintf("%d %s", m, r)
		}
	} else {
		switch {
		case f.Flag('0'):
			r = fmt.Sprintf("%03d°%02d′%s", d, m, r)
		case d > 0:
			r = fmt.Sprintf("%d°%d′%s", d, m, r)
		case m > 0:
			r = fmt.Sprintf("%d′%s", m, r)
		}
	}
	// add leading sign
	switch {
	case neg:
		r = "-" + r
	case f.Flag('+'):
		r = "+" + r
	case f.Flag(' '):
		r = " " + r
	}
	// format complete, check for width overflow
	if widSpec {
		pad := wid - utf8.RuneCountInString(r)
		if pad < 0 {
			a.WidthOverflow = "result overflows width"
			f.Write(bytes.Repeat([]byte{'*'}, wid)) // overflow result
			return
		}
		if pad > 0 {
			if f.Flag('-') {
				r += strings.Repeat(" ", pad)
			} else {
				r = strings.Repeat(" ", pad) + r
			}
		}
	}
	f.Write([]byte(r))
}

func (a Angle) String() string {
	return fmt.Sprintf("%s", a)
}

// RA formats to hours, minutes, seconds like Time, but wrapped to the range
// 0 to 24 hours.  Sign formatting flags '+' and ' ' are ignored.
// Otherwise see formatting notes under Angle.
type RA struct {
	Rad           float64 // Right ascension in radians.
	WidthOverflow string  // Message valid after format.
}

func (ra *RA) SetHMS(neg bool, h, m int, s float64) {
	hr := math.Mod(DMSToDeg(neg, h, m, s), 24)
	if hr < 0 {
		hr += 24
	}
	ra.Rad = hr * 15 * math.Pi / 180
}

// HourAngle formats to hours, minutes, seconds.  See formatting notes under Angle.
type HourAngle struct {
	Rad           float64 // Hour angle in radians.
	WidthOverflow string  // Message valid after format.
}

func (ha *HourAngle) SetHMS(neg bool, h, m int, s float64) {
	ha.Rad = DMSToDeg(neg, h, m, s) * 15 * math.Pi / 180
}

type WidthError string

func (e WidthError) Error() string {
	return string(e)
}

var WidthErrorInvalidPrecision = WidthError("Invalid precision")
var WidthErrorLossOfPrecision = WidthError("Possible loss of precision")

func Split60(x float64, prec int) (neg bool, quo int64, rem string, err error) {
	// limit of 15 set by max power of 10 is exactly representable as a float64
	if prec < 0 || prec > 15 {
		err = WidthErrorInvalidPrecision
		return
	}
	if x < 0 {
		x = -x
		neg = true
	}
	pf := math.Pow(10, float64(prec)) // precision factor, exact
	xs := x * pf                      // scale to precision
	// check that we can work with xs as uint64 without overflow
	i := int64(xs + .5) // round
	if i > 1<<52 {
		err = WidthErrorLossOfPrecision
		return
	}
	// compute integer values of segments
	p60 := 60 * int64(pf)
	quo = i / p60
	// digits of second segment, scaled to precision
	rem = fmt.Sprintf("%0*d", prec+1, i%p60)
	if prec > 0 {
		split := len(rem) - prec
		rem = rem[:split] + "." + rem[split:]
	}
	return
}
