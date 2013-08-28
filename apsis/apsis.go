// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Apsis: Chapter 50, Perigee and apogee of the Moon
//
// Incomplete:  Perigee and PerigeeParallax not implemented for lack of test
// cases.  Implementation involves copying tables of coefficients, involving
// risk of typographical error.
package apsis

import (
	"math"

	"github.com/soniakeys/meeus/base"
)

// conversion factor from k to T, given in (50.3) p. 356
const ck = 1 / 1325.55

// (50.1) p. 355
func mean(T float64) float64 {
	return base.Horner(T, 2451534.6698, 27.55454989/ck,
		-.0006691, -.000001098, .0000000052)
}

// snap returns k at half h nearest year y.
func snap(y, h float64) float64 {
	k := (y - 1999.97) * 13.2555 // (50.2) p. 355
	return math.Floor(k-h+.5) + h
}

// MeanPerigee returns the jde of the mean perigee of the Moon nearest the given date.
//
// Year is a decimal year specifying a date.
func MeanPerigee(year float64) float64 {
	return mean(snap(year, 0) * ck)
}

// MeanApogee returns the jde of the mean apogee of the Moon nearest the given date.
//
// Year is a decimal year specifying a date.
func MeanApogee(year float64) float64 {
	return mean(snap(year, .5) * ck)
}

// Apogee returns the jde of apogee of the Moon nearest the given date.
//
// Year is a decimal year specifying a date.
func Apogee(year float64) float64 {
	l := newLa(year, .5)
	return mean(l.T) + l.ac()
}

// ApogeeParallax returns equatorial horizontal parallax of the Moon at the Apogee nearest the given date.
//
// Year is a decimal year specifying a date.
//
// Result in radians.
func ApogeeParallax(year float64) float64 {
	return newLa(year, .5).ap()
}

type la struct {
	k, T    float64
	D, M, F float64
}

const p = math.Pi / 180

func newLa(y, h float64) *la {
	l := &la{k: snap(y, h)}
	l.T = l.k * ck // (50.3) p. 350
	l.D = base.Horner(l.T, 171.9179*p, 335.9106046*p/ck,
		-.0100383*p, -.00001156*p, .000000055*p)
	l.M = base.Horner(l.T, 347.3477*p, 27.1577721*p/ck,
		-.000813*p, -.000001*p)
	l.F = base.Horner(l.T, 316.6109*p, 364.5287911*p/ck,
		-.0125053*p, -.0000148*p)
	return l
}

// apogee correction
func (l *la) ac() float64 {
	return .4392*math.Sin(2*l.D) +
		.0684*math.Sin(4*l.D) +
		(.0456-.00011*l.T)*math.Sin(l.M) +
		(.0426-.00011*l.T)*math.Sin(2*l.D-l.M) +
		.0212*math.Sin(2*l.F) +
		-.0189*math.Sin(l.D) +
		.0144*math.Sin(6*l.D) +
		.0113*math.Sin(4*l.D-l.M) +
		.0047*math.Sin(2*(l.D+l.F)) +
		.0036*math.Sin(l.D+l.M) +
		.0035*math.Sin(8*l.D) +
		.0034*math.Sin(6*l.D-l.M) +
		-.0034*math.Sin(2*(l.D-l.F)) +
		.0022*math.Sin(2*(l.D-l.M)) +
		-.0017*math.Sin(3*l.D) +
		.0013*math.Sin(4*l.D+2*l.F) +
		.0011*math.Sin(8*l.D-l.M) +
		.001*math.Sin(4*l.D-2*l.M) +
		.0009*math.Sin(10*l.D) +
		.0007*math.Sin(3*l.D+l.M) +
		.0006*math.Sin(2*l.M) +
		.0005*math.Sin(2*l.D+l.M) +
		.0005*math.Sin(2*(l.D+l.M)) +
		.0004*math.Sin(6*l.D+2*l.F) +
		.0004*math.Sin(6*l.D-2*l.M) +
		.0004*math.Sin(10*l.D-l.M) +
		-.0004*math.Sin(5*l.D) +
		-.0004*math.Sin(4*l.D-2*l.F) +
		.0003*math.Sin(2*l.F+l.M) +
		.0003*math.Sin(12*l.D) +
		.0003*math.Sin(2*l.D+2*l.F-l.M) +
		-.0003*math.Sin(l.D-l.M)
}

// apogee parallax
func (l *la) ap() float64 {
	const s = math.Pi / 180 / 3600
	return 3245.251*s +
		-9.147*s*math.Cos(2*l.D) +
		-.841*s*math.Cos(l.D) +
		.697*s*math.Cos(2*l.F) +
		(-.656*s+.0016*s*l.T)*math.Cos(l.M) +
		.355*s*math.Cos(4*l.D) +
		.159*s*math.Cos(2*l.D-l.M) +
		.127*s*math.Cos(l.D+l.M) +
		.065*s*math.Cos(4*l.D-l.M) +
		.052*s*math.Cos(6*l.D) +
		.043*s*math.Cos(2*l.D+l.M) +
		.031*s*math.Cos(2*(l.D+l.F)) +
		-.023*s*math.Cos(2*(l.D-l.F)) +
		.022*s*math.Cos(2*(l.D-l.M)) +
		.019*s*math.Cos(2*(l.D+l.M)) +
		-.016*s*math.Cos(2*l.M) +
		.014*s*math.Cos(6*l.D-l.M) +
		.01*s*math.Cos(8*l.D)
}
