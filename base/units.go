// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package base

import "math"

type Angle float64

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

func (a Angle) Rad() float64 { return float64(a) }
func (a Angle) Deg() float64 { return float64(a) * 180 / math.Pi }
func (a Angle) Min() float64 { return float64(a) * 60 * 180 / math.Pi }
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

type HourAngle float64

func HourAngleFromHours(h float64) HourAngle {
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

func (h HourAngle) Rad() float64   { return float64(h) }
func (h HourAngle) Hours() float64 { return float64(h) * 12 / math.Pi }
func (h HourAngle) Min() float64   { return float64(h) * 60 * 12 / math.Pi }
func (h HourAngle) Sec() float64   { return float64(h) * 3600 * 12 / math.Pi }

type RA float64

func RAFromRad(rad float64) RA { return RA(PMod(rad, 2*math.Pi)) }
func RAFromDeg(d float64) RA   { return RAFromRad(d / 180 * math.Pi) }
func RAFromHours(h float64) RA { return RAFromRad(h / 12 * math.Pi) }
func RAFromMin(m float64) RA   { return RAFromRad(m / 60 / 12 * math.Pi) }
func RAFromSec(s float64) RA   { return RAFromRad(s / 3600 / 12 * math.Pi) }

func (ra RA) Add(h HourAngle) RA { return RAFromRad(ra.Rad() + h.Rad()) }

func (ra RA) Rad() float64  { return float64(ra) }
func (ra RA) Deg() float64  { return float64(ra) * 180 / math.Pi }
func (ra RA) Hour() float64 { return float64(ra) * 12 / math.Pi }
func (ra RA) Min() float64  { return float64(ra) * 60 * 12 / math.Pi }
func (ra RA) Sec() float64  { return float64(ra) * 3600 * 12 / math.Pi }

type Time float64

func TimeFromDays(d float64) Time {
	// 3600 sec in an hour, 24 hours in a day
	return Time(d * 3600 * 24)
}

func TimeFromHours(h float64) Time {
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

func (t Time) Days() float64  { return float64(t) / 3600 / 24 }
func (t Time) Hours() float64 { return float64(t) / 3600 }
func (t Time) Min() float64   { return float64(t) / 60 }
func (t Time) Rad() float64   { return float64(t) / 3600 / 12 * math.Pi }
func (t Time) Sec() float64   { return float64(t) }
