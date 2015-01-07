// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package base

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"strings"
	"unicode/utf8"
)

const (
	secAppend    = 's'
	secCombine   = 'c'
	secInsert    = 'd'
	minAppend    = 'm'
	minCombine   = 'n'
	minInsert    = 'o'
	hrDegAppend  = 'h'
	hrDegCombine = 'i'
	hrDegInsert  = 'j'
)

// InsertUnit inserts a unit indicator into a formatted decimal number.
//
// The indicator is inserted just before the decimal separator if one is
// present, or at the end of the number otherwise.
//
// The package variable DecSep is used to identify the decimal separator.
// If DecSep is non-empty and occurrs in d, unit is added just before the
// occurrence.  Otherwise unit is appended to the end of d.
//
// See also CombineUnit, StripUnit.
func InsertUnit(d, unit string) string {
	if DecSep == "" {
		return d + unit // DecSep empty, append unit
	}
	i := strings.Index(d, DecSep)
	if i < 0 {
		return d + unit // no DecSep found, append unit
	}
	// insert unit before DecSep
	return d[:i] + unit + d[i:]
}

// CombineUnit inserts a unit indicator into a formatted decimal number,
// combining it if possible with the decimal separator.
//
// The package variable DecSep is used to identify the decimal separator.
// If DecSep is non-empty and occurrs in d, the occurrence is replaced with
// argument 'unit' and package variable DecCombine.  Otherwise unit is
// appended to the end of d.
//
// See also InsertUnit, StripUnit.
func CombineUnit(d, unit string) string {
	if DecSep == "" {
		return d + unit // DecSep empty, append unit
	}
	i := strings.Index(d, DecSep)
	if i < 0 {
		return d + unit // no DecSep found, append unit
	}
	// insert unit, replace DecSep occurrence with DecCombine
	return d[:i] + unit + string(DecCombine) + d[i+len(DecSep):]
}

// StripUnit reverses the action of InsertUnit or CombineUnit,
// removing the specified unit indicator and restoring a following
// DecCombine to DecSep.
func StripUnit(d, unit string) string {
	xu := strings.Index(d, unit)
	if xu < 0 {
		return d
	}
	xd := xu + len(unit)
	if xd == len(d) {
		return d[:xu] // string ends with unit.  just remove the unit.
	}
	if strings.HasPrefix(d[xd:], DecSep) {
		return d[:xu] + d[xd:] // remove unit, retain DecSep
	}
	if r, sz := utf8.DecodeRuneInString(d[xd:]); r == DecCombine {
		// replace unit and DecCombine with DecSep
		return d[:xu] + DecSep + d[xd+sz:]
	}
	return d // otherwise don't mess with it
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
	HrDeg, Min, Sec string
}

// DMSUnits, HMSUnits, and DecSep specify unit and decimal indicators.
//
// You can change these as needed, for example to ASCII symbols.
// It is valid to use multiple character strings for DMSUnits and HMSUnits.
// It is valid to use empty strings with a fixed width format.
// DecCombine should be a rune of Unicode category "Mn" (mark, nonspacing).
var (
	DMSUnits   = UnitSymbols{"°", "′", "″"}
	HMSUnits   = UnitSymbols{"ʰ", "ᵐ", "ˢ"}
	DecSep     = "."
	DecCombine = '\u0323'
)

// Predefined errors indicate that a value could not be formatted.
// Custom formatters of FmtAngle, FmtHourAngle, FmtRA, and FmtTime types
// may store these in the Err field of the value being formatted.
var (
	ErrLossOfPrecision = errors.New("Loss of precision")
	ErrDegreeOverflow  = errors.New("Degrees overflow width")
	ErrHourOverflow    = errors.New("Hours overflow width")
	ErrPosInf          = errors.New("+Inf")
	ErrNegInf          = errors.New("-Inf")
	ErrNaN             = errors.New("NaN")
)

var (
	tenf = [16]float64{1e0, 1e1, 1e2, 1e3, 1e4, 1e5,
		1e6, 1e7, 1e8, 1e9, 1e10, 1e11, 1e12, 1e13, 1e14, 1e15}
	teni = [16]int64{1e0, 1e1, 1e2, 1e3, 1e4, 1e5,
		1e6, 1e7, 1e8, 1e9, 1e10, 1e11, 1e12, 1e13, 1e14, 1e15}
)

// sig verifies and returns significant digits of a number at a precision.
//
// x must be >= 0.  prec must be 0..15.
//
// the digits are returned as xs = int64(x * 10**prec + .5), as long as
// the result xs is small enough that all digits are significant given
// float64 representation.
// if xs does not represent a fully significant result -1 is returned.
func sig(x float64, prec int) int64 {
	xs := x*tenf[prec] + .5
	if !(xs <= 1<<52) { // 52 mantissa bits in float64
		return -1
	}
	return int64(xs)
}

// Angle represents a general purpose angle.
//
// Unit is radians.
type Angle float64

// NewAngle constructs a new Angle value from sign, degree, minute, and second
// components.
func NewAngle(neg bool, d, m int, s float64) Angle {
	return Angle(DMSToDeg(neg, d, m, s) * math.Pi / 180)
}

// Rad returns the angle in radians.
//
// This is the underlying representation and involves no scaling.
func (a Angle) Rad() float64 { return float64(a) }

// Deg returns the angle in degrees.
func (a Angle) Deg() float64 { return float64(a) * 180 / math.Pi }

// FmtAngle is represents a formattable angle.
type FmtAngle struct {
	Angle
	Err error // set each time the value is formatted.
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

// Format implements fmt.Formatter
func (a *FmtAngle) Format(f fmt.State, c rune) {
	s := state{
		State:  f,
		verb:   c,
		hrDeg:  a.Deg(),
		caller: fsAngle,
	}
	a.Err = s.writeFormatted()
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

// Rad returns the hour angle as an angle in radians.
//
// This is the underlying representation and involves no scaling.
func (a HourAngle) Rad() float64 { return float64(a) }

// Hour returns the hour angle as hours of time.
func (a HourAngle) Hour() float64 { return float64(a) * 12 / math.Pi }

// FmtHourAngle represents a formattable angle hour.
type FmtHourAngle struct {
	HourAngle
	Err error // set each time the value is formatted.
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

// Format implements fmt.Formatter
func (ha *FmtHourAngle) Format(f fmt.State, c rune) {
	s := &state{
		State:  f,
		verb:   c,
		hrDeg:  ha.Hour(),
		caller: fsHourAngle,
	}
	ha.Err = s.writeFormatted()
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

// Rad returns the right ascension as an angle in radians.
//
// This is the underlying representation and involves no scaling.
func (ra RA) Rad() float64 { return float64(ra) }

// Hour returns the right ascension as hours of time.
func (ra RA) Hour() float64 { return float64(ra) * 12 / math.Pi }

// FmtRA represents a formattable right ascension.
type FmtRA struct {
	RA
	Err error // set each time the value is formatted.
}

// NewFmtRA constructs a new FmtRA from a float64 in radians.
//
// The value is wrapped to the range [0,24) hours.
func NewFmtRA(rad float64) *FmtRA {
	return &FmtRA{RA: RA(PMod(rad, 2*math.Pi))}
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
	s := &state{
		State: f,
		verb:  c,
		// PMod in case ra.RA was directly set to something out of range
		hrDeg:  PMod(ra.Hour(), 24),
		caller: fsRA,
	}
	ra.Err = s.writeFormatted()
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

// Sec returns the time in seconds.
//
// This is the underlying representation and involves no scaling.
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
	Time
	Err error // set each time the value is formatted.
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
	s := &state{
		State:  f,
		verb:   c,
		hrDeg:  t.Hour(),
		caller: fsTime,
	}
	t.Err = s.writeFormatted()
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

type state struct {
	fmt.State         // 'f' in fmt.Formatter doc.  kind of handy to embed this.
	verb      rune    // 'c' in fmt.Formatter doc
	hrDeg     float64 // input, value to format
	prec      int     // f.Precision with a default of 0
	caller    int     // use fs constants
	units     UnitSymbols
}

func (s *state) writeFormatted() error {
	// valiate verb, pick formatting method in the process
	var f func() (string, error)
	switch s.verb {
	case 'v':
		fallthrough
	case secAppend, secCombine, secInsert:
		f = s.decimalSec // it's a method value! see the spec.
	case minAppend, minCombine, minInsert:
		f = s.decimalMin
	case hrDegAppend, hrDegCombine, hrDegInsert:
		f = s.decimalHrDeg
	default:
		fmt.Fprintf(s, "%%!%c(BADVERB)", s.verb)
		return nil // not a value error
	}

	// validate precision, storing it in the receiver.
	// 0 is our default if it's not specified.
	// (the docs don't define what prec is returned for the !ok case)
	var ok bool
	switch s.prec, ok = s.Precision(); {
	case !ok:
		s.prec = 0
	case s.prec > 15:
		// limit of 15 set by max power of 10 that is exactly representable
		// as a float64.  later code depends on prec being in this range.
		fmt.Fprintf(s, "%%!(BADPREC %d)", s.prec)
		return nil // not a value error
	}

	// format validated, now preliminary checks on value:
	var (
		r   string
		err error
	)
	switch {
	case math.IsNaN(s.hrDeg):
		err = ErrNaN
		goto valErr
	case !math.IsInf(s.hrDeg, 0): // normal path
	case math.IsInf(s.hrDeg, 1):
		err = ErrPosInf
		goto valErr
	default:
		err = ErrNegInf
		goto valErr
	}
	// okay so far.  a little more set up,
	switch {
	case s.caller == fsAngle:
		s.units = DMSUnits
	default:
		s.units = HMSUnits
	}
	// and then call the formatting method picked above
	if r, err = f(); err == nil {
		s.Write([]byte(r))
		return nil // normal return
	}

	// If there was a value error, we output all '*'s
	// but we need a length.  The strategy here is to replace the invalid
	// value with something valid and call format again to get a mock
	// result, then use len(mock) for the number of '*'s to output.
valErr:
	s.hrDeg = 0
	width := 10 // default, defensive in case f somehow fails on 0.
	if mock, err2 := f(); err2 == nil {
		width = utf8.RuneCountInString(mock)
		if strings.IndexRune(mock, DecCombine) >= 0 {
			width--
		}
	}
	s.Write(bytes.Repeat([]byte{'*'}, width))
	return err
}

func (s *state) decimalHrDeg() (string, error) {
	i := sig(math.Abs(s.hrDeg), s.prec)
	if i < 0 {
		return "", ErrLossOfPrecision
	}
	if s.hrDeg < 0 {
		i = -i
	}
	var r, f string
	if wid, widSpec := s.Width(); !widSpec {
		if s.Flag('+') {
			f = "%+0*d"
		} else if s.Flag(' ') { // sign space if requested
			f = "% 0*d"
		} else {
			f = "%0*d"
		}
		// +1 forces at least one place left of decimal point
		r = fmt.Sprintf(f, s.prec+1, i)
	} else {
		// fixed width a little more involved
		if s.Flag('+') {
			f = "%+"
		} else {
			f = "% " // sign space forced with fixed width
		}
		if s.Flag('0') {
			f += "0*d"
		} else {
			f += "*d"
		}
		wf := s.prec + wid + 1 // +1 here is required space for sign
		r := fmt.Sprintf(f, wf, i)
		if len(r) > wf {
			if s.caller == fsAngle {
				return "", ErrDegreeOverflow
			}
			return "", ErrHourOverflow
		}
	}
	if s.prec > 0 {
		split := len(r) - s.prec
		r = r[:split] + DecSep + r[split:]
	}
	switch s.verb {
	case hrDegAppend:
		r += string(s.units.HrDeg)
	case hrDegCombine:
		r = CombineUnit(r, s.units.HrDeg)
	case hrDegInsert:
		r = InsertUnit(r, s.units.HrDeg)
	}
	return r, nil
}

func (s *state) decimalMin() (string, error) {
	i := sig(math.Abs(s.hrDeg)*60, s.prec) // hrDeg*60 gets minutes
	if i < 0 {
		return "", ErrLossOfPrecision
	}
	p60 := 60 * teni[s.prec]
	min := i / p60
	sec := i % p60

	r, minEl, err := s.firstSeg(min)
	if err != nil {
		return "", err
	}
	return r + s.lastSeg(sec, s.units.Min, minEl), nil
}

func (s *state) firstSeg(x int64) (r string, elided bool, err error) {
	switch wid, widSpec := s.Width(); {
	case widSpec:
		f := "%*d"
		if s.Flag('0') {
			f = "%0*d"
		}
		r = fmt.Sprintf(f, wid, x)
		if len(r) > wid {
			if s.caller == fsAngle {
				return "", false, ErrDegreeOverflow
			}
			return "", false, ErrHourOverflow
		}
		r += s.units.HrDeg
	case x > 0 || s.Flag('#'):
		r = fmt.Sprintf("%d%s", x, s.units.HrDeg)
	default:
		elided = true
	}
	switch {
	case s.hrDeg < 0:
		r = "-" + r
	case s.Flag('+'):
		r = "+" + r
	case s.Flag(' '):
		r = " " + r
	}
	return r, elided, nil
}

func (s *state) lastSeg(sec int64, unit string, first bool) string {
	wid := s.prec + 1
	_, widSpec := s.Width()
	if s.Flag('0') && (widSpec || !first) {
		wid++
	}
	r := fmt.Sprintf("%0*d", wid, sec)
	if widSpec && len(r) < s.prec+2 {
		r = " " + r
	}
	if s.prec > 0 {
		split := len(r) - s.prec
		r = r[:split] + DecSep + r[split:]
	}
	switch s.verb {
	case secCombine, minCombine:
		return CombineUnit(r, unit)
	case secInsert, minInsert:
		return InsertUnit(r, unit)
	}
	return r + unit
}

func (s *state) decimalSec() (string, error) {
	i := sig(math.Abs(s.hrDeg)*3600, s.prec) // hrDeg*3600 gets seconds
	if i < 0 {
		return "", ErrLossOfPrecision
	}
	p60 := 60 * teni[s.prec]
	sec := i % p60
	i /= p60
	min := i % 60
	hrDeg := i / 60
	r, firstEl, err := s.firstSeg(hrDeg)
	if err != nil {
		return "", err
	}
	f := "%s%d%s"
	minEl := false
	if s.Flag('0') && !firstEl {
		f = "%s%02d%s"
	} else {
		switch _, widSpec := s.Width(); {
		case widSpec:
			f = "%s%2d%s"
		case firstEl && min == 0:
			minEl = true
			goto last
		}
	}
	r = fmt.Sprintf(f, r, min, s.units.Min)
last:
	return r + s.lastSeg(sec, s.units.Sec, minEl), nil
}
