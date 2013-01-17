// Copyright 2012 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package meeus

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
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

// UnitSymbols holds symbols for formatting FmtAngle, FmtHourAngle, FmtRA,
// and FmtTime types.
type UnitSymbols struct {
	First, M, S rune
}

// DMSRunes specifies symbols use when formatting angles.  You can change
// these, perhaps to ASCII 'd', 'm', and 's', as needed.
var DMSRunes = UnitSymbols{'°', '′', '″'}

// HMSRunes specifies symbols use when formatting hour angles, right
// ascensions, and times.
//
// You can change these, perhaps to ASCII 'h', 'm', and 's', as needed.
var HMSRunes = UnitSymbols{'ʰ', 'ᵐ', 'ˢ'}

// WidthError is an explanatory error set when a formatting operation outputs
// all stars, indicating a format overflow error.
type WidthError string

// Error implements the built in error interface.
func (e WidthError) Error() string {
	return string(e)
}

// Predefined WidthErrors.  The custom formatters for FmtAngle, FmtHourAngle,
// FmtRA, and FmtTime emit all asterisks, "*************", in these overflow
// cases.  The exact error is stored in the WidthError field of the type.
var (
	WidthErrorInvalidPrecision = WidthError("Invalid precision")
	WidthErrorLossOfPrecision  = WidthError("Possible loss of precision")
	WidthErrorDegreeOverflow   = WidthError("Degrees overflow width")
	WidthErrorHourOverflow     = WidthError("Hours overflow width")
	WidthErrorInfP             = WidthError("+Inf")
	WidthErrorInfN             = WidthError("-Inf")
	WidthErrorNaN              = WidthError("NaN")
)

// Split60 splits a decimal segment from a floating point number that
// will be formatted in some sexagesimal notation.
//
// It is a low-level function used internally but exported to support
// formatting cases not handled by the custom formatters of the FmtAngle,
// FmtHourAngle, FmtRA, and FmtTime types.
//
// Argument x is the number to split and prec specifies the number digits
// to place past the decimal point in the decimal segment.
//
// Return value neg will be true if x is < 0. x60 and seg are then returned as
// non-negative numbers.  Seg will be a formatted string in the range [0,60)
// and the relation
//	x60 * 60 + seg = abs(x)
// will hold.
//
// Seg is returned as a string because x is rounded specifically for the
// the specified precision.  Do not convert seg to a floating point number and
// do further operations on it or you risk seeing results like 23′60″.
// Seg will always have at least 1 digit to the left of the decimal point.
// Set argument pad to true to 0-pad seg to two digits to the left.
//
// Maximum allowed precision is 15, but that is only valid for angles smaller
// than a few arc seconds.  Larger angles will give a "Possible loss of
// precision" error.  The maximum precision before getting a loss of precision
// error decreases as the angle magnitude increases.  At one degree you can
// get 12 digits of precision.  At 360 degrees you get 9.
func Split60(x float64, prec int, pad bool) (neg bool, x60 int64, seg string, err error) {

	switch {
	case math.IsNaN(x):
		err = WidthErrorNaN
	case !math.IsInf(x, 0):
		goto P
	case math.IsInf(x, 1):
		err = WidthErrorInfP
	default:
		err = WidthErrorInfN
	}
	return
P:
	// limit of 15 set by max power of 10 that is exactly representable
	// as a float64
	if prec < 0 || prec > 15 {
		err = WidthErrorInvalidPrecision
		return
	}
	if x < 0 {
		x = -x
		neg = true
	}
	// precision factor, known to be exact
	pf := math.Pow(10, float64(prec))
	xs := x * pf // scale to precision

	// check that we can represent xs exactly
	i := int64(xs + .5) // round
	if i > 1<<52 {
		err = WidthErrorLossOfPrecision
		return
	}
	// compute final return values
	p60 := 60 * int64(pf)
	x60 = i / p60
	// digits of segment, scaled to precision
	digits := prec + 1
	if pad {
		digits++
	}
	seg = fmt.Sprintf("%0*d", digits, i%p60)
	if prec > 0 {
		split := len(seg) - prec
		seg = seg[:split] + "." + seg[split:]
	}
	return
}

// Angle represents a general purpose angle.
//
// Unit is radians.
type Angle float64

// NewAngle constructs a new Angle value from sign, degree, minute, and second
// components.
func NewAngle(neg bool, d, m int, s float64) Angle {
	return Angle(DMSToDeg(neg, d, m, s) * (math.Pi / 180))
}

// Rad returns the value of the angle as a float64.
func (a Angle) Rad() float64 { return float64(a) }

// FmtAngle is represents a formattable angle.
type FmtAngle struct {
	Angle            // embedded Angle with Rad method
	WidthError error // valid after format
}

// NewFmtAngle constructs a new FmtAngle from a float64 in radians.
func NewFmtAngle(rad float64) *FmtAngle {
	return &FmtAngle{Angle: Angle(rad)}
}

// SetDMS sets the value of an FAngle from sign, degree, minute, and second
// components.
//
// The receiver is returned as a convenience.
func (a *FmtAngle) SetDMS(neg bool, d, m int, s float64) *FmtAngle {
	a.Angle = NewAngle(neg, d, m, s)
	return a
}

// Format implements fmt.Formatter, formatting to degrees, minutes,
// and seconds.
func (a *FmtAngle) Format(f fmt.State, c rune) {
	a.WidthError = formatSex(a.Rad()*3600*180/math.Pi, fsAngle, nil, f, c)
}

// String implements fmt.Stringer
func (a *FmtAngle) String() string {
	return fmt.Sprintf("%s", a)
}

// HourAngle represents an angle corresponding to angular rotation of
// the Earth in a specified time.
//
// Unit is radians.
type HourAngle float64

// NewHourAngle constructs a new HourAngle value from sign, hour, minute,
// and second components.
func NewHourAngle(neg bool, h, m int, s float64) HourAngle {
	return HourAngle(DMSToDeg(neg, h, m, s) * 15 * math.Pi / 180)
}

// Rad returns the value of the hour angle as a float64.
//
// Unit is radians.
func (a HourAngle) Rad() float64 { return float64(a) }

// FmtHourAngle represents a formattable angle hour.
type FmtHourAngle struct {
	HourAngle        // embedded HourAngle with Rad method
	WidthError error // valid after format
}

// NewFmtHourAngle constructs a new FmtHourAngle from a float64 in radians.
func NewFmtHourAngle(rad float64) *FmtHourAngle {
	return &FmtHourAngle{HourAngle: HourAngle(rad)}
}

// SetHMS sets the value of the HourAngle from time components sign, hour,
// minute, and second.
//
// The receiver is returned as a convenience.
func (ha *FmtHourAngle) SetHMS(neg bool, h, m int, s float64) *FmtHourAngle {
	ha.HourAngle = NewHourAngle(neg, h, m, s)
	return ha
}

// Format implements fmt.Formatter, formatting to hours, minutes, and seconds.
func (ha *FmtHourAngle) Format(f fmt.State, c rune) {
	ha.WidthError =
		formatSex(ha.Rad()*3600*180/math.Pi/15, fsHourAngle, nil, f, c)
}

// String implements fmt.Stringer
func (ha *FmtHourAngle) String() string {
	return fmt.Sprintf("%s", ha)
}

// RA represents a value of right ascension.
//
// Unit is radians.
type RA float64

// NewRA constructs a new RA value from hour, minute, and second components.
//
// Negative values are not supported, and NewRA wraps values larger than 24
// to the range [0,24) hours.
func NewRA(h, m int, s float64) RA {
	hr := math.Mod(DMSToDeg(false, h, m, s), 24)
	return RA(hr * 15 * math.Pi / 180)
}

// Rad returns the RA value as a float64.
//
// Unit is radians.
func (ra RA) Rad() float64 { return float64(ra) }

// FmtRA represents a formattable right ascension.
type FmtRA struct {
	RA               // embedded RA with Rad method
	WidthError error // valid after format
}

// NewFmtRA constructs a new FmtRA from a float64 in radians.
//
// The value is wrapped to the range [0,24) hours.
func NewFmtRA(rad float64) *FmtRA {
	rad = math.Mod(rad, 2*math.Pi)
	if rad < 0 {
		rad += 2 * math.Pi
	}
	return &FmtRA{RA: RA(rad)}
}

// SetHMS sets the value of RA from components hour, minute, and second.
// Negative values are not supported, and SetHMS wraps values larger than 24
// to the range [0,24) hours.
//
// The receiver is returned as a convenience.
func (ra *FmtRA) SetHMS(h, m int, s float64) *FmtRA {
	ra.RA = NewRA(h, m, s)
	return ra
}

// Format implements fmt.Formatter, formatting to hours, minutes, and seconds.
func (ra *FmtRA) Format(f fmt.State, c rune) {
	// repeat mod in case Rad was directly set to something out of range
	decimalHours := math.Mod(ra.Rad()*180/math.Pi/15, 24)
	if decimalHours < 0 {
		decimalHours += 24
	}
	ra.WidthError = formatSex(decimalHours*3600, fsRA, nil, f, c)
}

// String implements fmt.Stringer
func (ra *FmtRA) String() string {
	return fmt.Sprintf("%s", ra)
}

// Time represents a duration or relative time.
//
// Unit is seconds.
type Time float64

// NewTime constructs a new Time value from sign, hour, minute, and
// second components.
func NewTime(neg bool, h, m int, s float64) Time {
	s += float64((h*60 + m) * 60)
	if neg {
		return Time(-s)
	}
	return Time(s)
}

// Sec returns the Time value as a float64.
//
// Unit is seconds.
func (t Time) Sec() float64 { return float64(t) }

// Min returns time in minutes.
func (t Time) Min() float64 { return float64(t) / 60 }

// Hour returns time in hours.
func (t Time) Hour() float64 { return float64(t) / 3600 }

// Day returns time in days.
func (t Time) Day() float64 { return float64(t) / 3600 / 24 }

// Rad returns time in radians, where 1 day = 2 Pi radians.
func (t Time) Rad() float64 { return float64(t) * math.Pi / 12 / 3600 }

// FmtTime represents a formattable duration or relative time.
type FmtTime struct {
	Time             // embedded Time with Sec method
	WidthError error // valid after format
}

// NewFmtTime constructs a new FmtTime from a float64 in seconds.
func NewFmtTime(sec float64) *FmtTime {
	return &FmtTime{Time: Time(sec)}
}

// SetHMS sets the value of FmtTime from time components sign, hour,
// minute, and second.
//
// The receiver is returned as a convenience.
func (t *FmtTime) SetHMS(neg bool, h, m int, s float64) *FmtTime {
	t.Time = NewTime(neg, h, m, s)
	return t
}

// Format implements fmt.Formatter, formatting to hours, minutes, and seconds.
func (t *FmtTime) Format(f fmt.State, c rune) {
	t.WidthError = formatSex(t.Sec(), fsTime, nil, f, c)
}

// String implements fmt.Stringer
func (t *FmtTime) String() string {
	return fmt.Sprintf("%s", t)
}

const (
	fsAngle = iota
	fsHourAngle
	fsRA
	fsTime
)

func formatSex(x float64, caller int, mock *string, f fmt.State, c rune) error {
	// valiate verb
	switch c {
	case 's', 'd', 'c', 'v', 'x':
	default:
		fmt.Fprintf(f, "Invalid verb: %%%c", c)
		return nil // not an overflow error
	}
	// declare some variables ahead of goto
	var (
		d, m     int64
		s1       string
		sexRune  UnitSymbols
		wid1     int
		wid1Spec bool
	)
	// get meaningful precision
	prec, ok := f.Precision()
	if !ok {
		prec = 0
	}
	neg, x60, r, err := Split60(x, prec, f.Flag('0'))
	if err != nil {
		goto Overflow
	}
	// add seconds unit symbol
	switch {
	case c == 'x':
		sexRune = UnitSymbols{' ', ' ', ' '}
	case caller == fsAngle:
		sexRune = DMSRunes
	default:
		sexRune = HMSRunes
	}
	switch c {
	case 's', 'v':
		r += string(sexRune.S)
	case 'd':
		r = DecSymAdd(r, sexRune.S)
	case 'c':
		r = DecSymCombine(r, sexRune.S)
	}
	// add degrees, minutes to partial result
	d = x60 / 60
	m = x60 % 60
	s1 = strconv.FormatInt(d, 10)
	wid1, wid1Spec = f.Width()
	if wid1Spec {
		// simple rule applies in all cases where width is specified:
		if len(s1) > wid1 {
			if caller == fsAngle {
				err = WidthErrorDegreeOverflow
			} else {
				err = WidthErrorHourOverflow
			}
			goto Overflow
		}
	}
	if f.Flag('#') || d > 0 {
		if f.Flag('0') {
			r = fmt.Sprintf("%0*s%c%02d%c%s",
				wid1, s1, sexRune.First, m, sexRune.M, r)
		} else {
			r = fmt.Sprintf("%s%c%d%c%s",
				s1, sexRune.First, m, sexRune.M, r)
		}
	} else if m > 0 {
		if f.Flag('0') {
			r = fmt.Sprintf("%02d%c%s", m, sexRune.M, r)
		} else {
			r = fmt.Sprintf("%d%c%s", m, sexRune.M, r)
		}
	}
	// add leading sign
	if caller != fsRA {
		switch {
		case neg:
			r = "-" + r
		case f.Flag('+'):
			r = "+" + r
		case f.Flag(' '):
			r = " " + r
		}
	}
	if mock == nil {
		f.Write([]byte(r))
	} else {
		*mock = r
	}
	return nil
Overflow:
	err1 := err
	var width int
	if mock != nil { // detect recursive loop
		width = 10
	} else {
		var valid string
		formatSex(0, caller, &valid, f, c)
		width = utf8.RuneCountInString(valid)
	}
	f.Write(bytes.Repeat([]byte{'*'}, width)) // emit overflow indicator
	return err1
}
