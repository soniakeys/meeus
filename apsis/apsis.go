// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Apsis: Chapter 50, Perigee and apogee of the Moon

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

// Perigee returns the jde of perigee of the Moon nearest the given date.
//
// Year is a decimal year specifying a date.
func Perigee(year float64) float64 {
	l := newLa(year, 0)
	return mean(l.T) + l.pc()
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

// PerigeeParallax returns equatorial horizontal parallax of the Moon at the Perigee nearest the given date.
//
// Year is a decimal year specifying a date.
//
// Result in radians.
func PerigeeParallax(year float64) float64 {
	return newLa(year, 0).pp()
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

// perigee correction
func (l *la) pc() float64 {
	return -1.6769*math.Sin(2*l.D) +
		.4589*math.Sin(4*l.D) +
		-.1856*math.Sin(6*l.D) +
		.0883*math.Sin(8*l.D) +
		(-.0773+.00019*l.T)*math.Sin(2*l.D-l.M) +
		(.0502-.00013*l.T)*math.Sin(l.M) +
		-.046*math.Sin(10*l.D) +
		(.0422-.00011*l.T)*math.Sin(4*l.D-l.M) +
		-.0256*math.Sin(6*l.D-l.M) +
		.0253*math.Sin(12*l.D) +
		.0237*math.Sin(l.D) +
		.0162*math.Sin(8*l.D-l.M) +
		-.0145*math.Sin(14*l.D) +
		.0129*math.Sin(2*l.F) +
		-.0112*math.Sin(3*l.D) +
		-.0104*math.Sin(10*l.D-l.M) +
		.0086*math.Sin(16*l.D) +
		.0069*math.Sin(12*l.D-l.M) +
		.0066*math.Sin(5*l.D) +
		-.0053*math.Sin(2*(l.D+l.F)) +
		-.0052*math.Sin(18*l.D) +
		-.0046*math.Sin(14*l.D-l.M) +
		-.0041*math.Sin(7*l.D) +
		.004*math.Sin(2*l.D+l.M) +
		.0032*math.Sin(20*l.D) +
		-.0032*math.Sin(l.D+l.M) +
		.0031*math.Sin(16*l.D-l.M) +
		-.0029*math.Sin(4*l.D+l.M) +
		.0027*math.Sin(9*l.D) +
		.0027*math.Sin(4*l.D+2*l.F) +
		-.0027*math.Sin(2*(l.D-l.M)) +
		.0024*math.Sin(4*l.D-2*l.M) +
		-.0021*math.Sin(6*l.D-2*l.M) +
		-.0021*math.Sin(22*l.D) +
		-.0021*math.Sin(18*l.D-l.M) +
		.0019*math.Sin(6*l.D+l.M) +
		-.0018*math.Sin(11*l.D) +
		-.0014*math.Sin(8*l.D+l.M) +
		-.0014*math.Sin(4*l.D-2*l.F) +
		-.0014*math.Sin(6*l.D+2*l.F) +
		.0014*math.Sin(3*l.D+l.M) +
		-.0014*math.Sin(5*l.D+l.M) +
		.0013*math.Sin(13*l.D) +
		.0013*math.Sin(20*l.D-l.M) +
		.0011*math.Sin(3*l.D+2*l.M) +
		-.0011*math.Sin(2*(2*l.D+l.F-l.M)) +
		-.001*math.Sin(l.D+2*l.M) +
		-.0009*math.Sin(22*l.D-l.M) +
		-.0008*math.Sin(4*l.F) +
		.0008*math.Sin(6*l.D-2*l.F) +
		.0008*math.Sin(2*(l.D-l.F)+l.M) +
		.0007*math.Sin(2*l.M) +
		.0007*math.Sin(2*l.F-l.M) +
		.0007*math.Sin(2*l.D+4*l.F) +
		-.0006*math.Sin(2*(l.F-l.M)) +
		-.0006*math.Sin(2*(l.D-l.F+l.M)) +
		.0006*math.Sin(24*l.D) +
		.0005*math.Sin(4*(l.D-l.F)) +
		.0005*math.Sin(2*(l.D+l.M)) +
		-.0004*math.Sin(l.D-l.M)
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

// perigee parallax
func (l *la) pp() float64 {
	const s = math.Pi / 180 / 3600
	return 3629.215*s +
		63.224*s*math.Cos(2*l.D) +
		-6.990*s*math.Cos(4*l.D) +
		(2.834*s-0.0071*l.T*s)*math.Cos(2*l.D-l.M) +
		1.927*s*math.Cos(6*l.D) +
		-1.263*s*math.Cos(l.D) +
		-0.702*s*math.Cos(8*l.D) +
		(0.696*s-0.0017*l.T*s)*math.Cos(l.M) +
		-0.690*s*math.Cos(2*l.F) +
		(-0.629*s+0.0016*l.T*s)*math.Cos(4*l.D-l.M) +
		-0.392*s*math.Cos(2*(l.D-l.F)) +
		0.297*s*math.Cos(10*l.D) +
		0.260*s*math.Cos(6*l.D-l.M) +
		0.201*s*math.Cos(3*l.D) +
		-0.161*s*math.Cos(2*l.D+l.M) +
		0.157*s*math.Cos(l.D+l.M) +
		-0.138*s*math.Cos(12*l.D) +
		-0.127*s*math.Cos(8*l.D-l.M) +
		0.104*s*math.Cos(2*(l.D+l.F)) +
		0.104*s*math.Cos(2*(l.D-l.M)) +
		-0.079*s*math.Cos(5*l.D) +
		0.068*s*math.Cos(14*l.D) +
		0.067*s*math.Cos(10*l.D-l.M) +
		0.054*s*math.Cos(4*l.D+l.M) +
		-0.038*s*math.Cos(12*l.D-l.M) +
		-0.038*s*math.Cos(4*l.D-2*l.M) +
		0.037*s*math.Cos(7*l.D) +
		-0.037*s*math.Cos(4*l.D+2*l.F) +
		-0.035*s*math.Cos(16*l.D) +
		-0.030*s*math.Cos(3*l.D+l.M) +
		0.029*s*math.Cos(l.D-l.M) +
		-0.025*s*math.Cos(6*l.D+l.M) +
		0.023*s*math.Cos(2*l.M) +
		0.023*s*math.Cos(14*l.D-l.M) +
		-0.023*s*math.Cos(2*(l.D+l.M)) +
		0.022*s*math.Cos(6*l.D-2*l.M) +
		-0.021*s*math.Cos(2*l.D-2*l.F-l.M) +
		-0.020*s*math.Cos(9*l.D) +
		0.019*s*math.Cos(18*l.D) +
		0.017*s*math.Cos(6*l.D+2*l.F) +
		0.014*s*math.Cos(2*l.F-l.M) +
		-0.014*s*math.Cos(16*l.D-l.M) +
		0.013*s*math.Cos(4*l.D-2*l.F) +
		0.012*s*math.Cos(8*l.D+l.M) +
		0.011*s*math.Cos(11*l.D) +
		0.010*s*math.Cos(5*l.D+l.M) +
		-0.010*s*math.Cos(20*l.D)
}
