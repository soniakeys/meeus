// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package base

import "math"

// Angle represents a general purpose angle.
//
// Unit is radians.
type Angle float64

// NewAngle constructs a new Angle value from sign, degree, minute, and second
// components.
//
// For argument neg, pass '-' to negate the result.  Any other argument
// value, such as ' ', '+', or simply 0, leaves the result non-negated.
func NewAngle(neg byte, d, m int, s float64) Angle {
	return Angle(FromSexa(neg, d, m, s) / 180 * math.Pi)
}

func AngleFromDeg(d float64) Angle {
	// 180 deg or pi radians in a half-circle.
	return Angle(d / 180 * math.Pi)
}

func AngleFromMin(m float64) Angle {
	// 60 min in a degree, 180 deg or pi radians in a half-circle.
	return Angle(m / 60 / 180 * math.Pi)
}

func AngleFromSec(s float64) Angle {
	// 3600 sec in a degree, 180 deg or pi radians in a half-circle.
	return Angle(s / 3600 / 180 * math.Pi)
}

// Rad returns the angle in radians.
//
// This is the underlying representation and involves no scaling.
func (a Angle) Rad() float64 { return float64(a) }

// Deg returns the angle in degrees.
func (a Angle) Deg() float64 { return float64(a) * 180 / math.Pi }

// Min returns the angle in minutes.
func (a Angle) Min() float64 { return float64(a) * 60 * 180 / math.Pi }

// Sec returns the angle in seconds.
func (a Angle) Sec() float64 { return float64(a) * 3600 * 180 / math.Pi }

// FromSexa converts from parsed sexagesimal angle components to a single
// float64 value.
//
// Typically you pass non-negative values for d, m, and s, and to indicate
// a negative value, pass '-' for neg.  Any other value, such as ' ', '+',
// or simply 0, leaves the result non-negative.
//
// There are no limits on d, m, or s however.  Negative values or values
// > 60 for m and s are allowed for example.  The segment values are
// combined and then if neg is '-' that sum is negated.
//
// Also, the interpretation of d as degrees is arbitrary.  The function works
// as well on hours minutes and seconds.  Generally, m is a sexagesimal part
// of d and s is a sexagesimal part of m.
func FromSexa(neg byte, d, m int, s float64) float64 {
	s = (float64((d*60+m)*60) + s) / 3600
	if neg == '-' {
		return -s
	}
	return s
}

// HourAngle represents an angle corresponding to angular rotation of
// the Earth in a specified time.
//
// Unit is radians.
type HourAngle float64

// NewHourAngle constructs a new HourAngle value from sign, hour, minute,
// and second components.
//
// For argument neg, pass '-' to indicate a negative hour angle.  Any other
// argument value, such as ' ', '+', or simply 0, leaves the result
// non-negative.
func NewHourAngle(neg byte, h, m int, s float64) HourAngle {
	return HourAngle(FromSexa(neg, h, m, s) / 12 * math.Pi)
}

func HourAngleFromHour(h float64) HourAngle {
	// 12 hours or pi radians in a half-revolution
	return HourAngle(h / 12 * math.Pi)
}

func HourAngleFromMin(m float64) HourAngle {
	// 60 sec in an hour, 12 hours or pi radians in a half-revolution
	return HourAngle(m / 60 / 12 * math.Pi)
}

func HourAngleFromSec(s float64) HourAngle {
	// 3600 sec in an hour, 12 hours or pi radians in a half-revolution
	return HourAngle(s / 3600 / 12 * math.Pi)
}

// Rad returns the hour angle as an angle in radians.
//
// This is the underlying representation and involves no scaling.
func (h HourAngle) Rad() float64 { return float64(h) }

// Hour returns the hour angle as hours of time.
func (h HourAngle) Hour() float64 { return float64(h) * 12 / math.Pi }

func (h HourAngle) Min() float64 { return float64(h) * 60 * 12 / math.Pi }
func (h HourAngle) Sec() float64 { return float64(h) * 3600 * 12 / math.Pi }

// RA represents a value of right ascension.
//
// Unit is radians.
type RA float64

// NewRA constructs a new RA value from hour, minute, and second components.
//
// The result is wrapped to the range [0,2π), or [0,24) hours.
func NewRA(h, m int, s float64) RA {
	return RAFromRad(FromSexa(0, h, m, s) / 12 * math.Pi)
}

// RAFromRad constructs a new RA value from radians.
//
// The result is wrapped to the range [0,2π), or [0,24) hours.
func RAFromRad(rad float64) RA { return RA(PMod(rad, 2*math.Pi)) }

func RAFromDeg(d float64) RA  { return RAFromRad(d / 180 * math.Pi) }
func RAFromHour(h float64) RA { return RAFromRad(h / 12 * math.Pi) }
func RAFromMin(m float64) RA  { return RAFromRad(m / 60 / 12 * math.Pi) }
func RAFromSec(s float64) RA  { return RAFromRad(s / 3600 / 12 * math.Pi) }

func (ra RA) Add(h HourAngle) RA { return RAFromRad(ra.Rad() + h.Rad()) }

// Rad returns the RA as an angle in radians.
//
// This is the underlying representation and involves no scaling.
func (ra RA) Rad() float64 { return float64(ra) }

func (ra RA) Deg() float64 { return float64(ra) * 180 / math.Pi }

// Hour returns the RA as hours of RA.
func (ra RA) Hour() float64 { return float64(ra) * 12 / math.Pi }

func (ra RA) Min() float64 { return float64(ra) * 60 * 12 / math.Pi }
func (ra RA) Sec() float64 { return float64(ra) * 3600 * 12 / math.Pi }

// Time represents a duration or relative time.
//
// Unit is seconds.
type Time float64

// NewTime constructs a new Time value from sign, hour, minute, and
// second components.
//
// For argument neg, pass '-' to indicate a negative time delta.  Any other
// argument value, such as ' ', '+', or simply 0, leaves the result
// non-negative.
func NewTime(neg byte, h, m int, s float64) Time {
	s += float64((h*60 + m) * 60)
	if neg == '-' {
		return Time(-s)
	}
	return Time(s)
}

func TimeFromDay(d float64) Time {
	// 3600 sec in an hour, 24 hours in a day
	return Time(d * 3600 * 24)
}

func TimeFromHour(h float64) Time {
	// 3600 sec in an hour
	return Time(h * 3600)
}

func TimeFromMin(m float64) Time {
	// 60 sec in a min
	return Time(m * 60)
}

func TimeFromRad(rad float64) Time {
	// 3600 sec in an hour, 12 hours or pi radians in a half-day
	return Time(rad * 3600 * 12 / math.Pi)
}

// Day returns time in days.
func (t Time) Day() float64 { return float64(t) / 3600 / 24 }

// Hour returns time in hours.
func (t Time) Hour() float64 { return float64(t) / 3600 }

// Min returns time in minutes.
func (t Time) Min() float64 { return float64(t) / 60 }

// Rad returns time in radians, where 1 day = 2 Pi radians of rotation.
func (t Time) Rad() float64 { return float64(t) / 3600 / 12 * math.Pi }

// Sec returns the time in seconds.
//
// This is the underlying representation and involves no scaling.
func (t Time) Sec() float64 { return float64(t) }
