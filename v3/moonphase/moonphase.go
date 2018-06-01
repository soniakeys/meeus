// Copyright 2013 Sonia Keys
// License: MIT

// Moonphase: Chapter 49, Phases of the Moon
package moonphase

import (
	"math"

	"github.com/soniakeys/meeus/v3/base"
)

const ck = 1 / 1236.85

// (49.1) p. 349
func mean(T float64) float64 {
	return base.Horner(T, 2451550.09766, 29.530588861/ck,
		.00015437, -.00000015, .00000000073)
}

// snap returns k at specified quarter q nearest year y.
func snap(y, q float64) float64 {
	k := (y - 2000) * 12.3685 // (49.2) p. 350
	return math.Floor(k-q+.5) + q
}

// MeanNew returns the jde of the mean New Moon nearest the given date.
//
// Year is a decimal year specifying a date.
//
// The mean date is within .5 day of the true date of New Moon.
func MeanNew(year float64) float64 {
	return mean(snap(year, 0) * ck)
}

// MeanFirst returns the jde of the mean First Quarter Moon nearest the given date.
//
// Year is a decimal year specifying a date.
//
// The mean date is within .5 day of the true date of First Quarter Moon.
func MeanFirst(year float64) float64 {
	return mean(snap(year, .25) * ck)
}

// MeanFull returns the jde of the mean Full Moon nearest the given date.
//
// Year is a decimal year specifying a date.
//
// The mean date is within .5 day of the true date of New Moon.
func MeanFull(year float64) float64 {
	return mean(snap(year, .5) * ck)
}

// MeanLast returns the jde of the mean Last Quarter Moon nearest the given date.
//
// Year is a decimal year specifying a date.
//
// The mean date is within .5 day of the true date of Last Quarter Moon.
func MeanLast(year float64) float64 {
	return mean(snap(year, .75) * ck)
}

// New returns the jde of New Moon nearest the given date.
//
// Year is a decimal year specifying a date.
func New(year float64) float64 {
	m := newMp(year, 0)
	return mean(m.T) + m.nfc(&nc) + m.a()
}

// First returns the jde of First Quarter Moon nearest the given date.
//
// Year is a decimal year specifying a date.
func First(year float64) float64 {
	m := newMp(year, .25)
	return mean(m.T) + m.flc() + m.w() + m.a()
}

// Full returns the jde of Full Moon nearest the given date.
//
// Year is a decimal year specifying a date.
func Full(year float64) float64 {
	m := newMp(year, .5)
	return mean(m.T) + m.nfc(&fc) + m.a()
}

// Last returns the jde of Last Quarter Moon nearest the given date.
//
// Year is a decimal year specifying a date.
func Last(year float64) float64 {
	m := newMp(year, .75)
	return mean(m.T) + m.flc() - m.w() + m.a()
}

type mp struct {
	k, T           float64
	E, M, Mʹ, F, Ω float64
	A              [14]float64
}

const p = math.Pi / 180

func newMp(y, q float64) *mp {
	m := &mp{k: snap(y, q)}
	m.T = m.k * ck // (49.3) p. 350
	m.E = base.Horner(m.T, 1, -.002516, -.0000074)
	m.M = base.Horner(m.T, 2.5534*p, 29.1053567*p/ck,
		-.0000014*p, -.00000011*p)
	m.Mʹ = base.Horner(m.T, 201.5643*p, 385.81693528*p/ck,
		.0107582*p, .00001238*p, -.000000058*p)
	m.F = base.Horner(m.T, 160.7108*p, 390.67050284*p/ck,
		-.0016118*p, -.00000227*p, .000000011*p)
	m.Ω = base.Horner(m.T, 124.7746*p, -1.56375588*p/ck,
		.0020672*p, .00000215*p)
	m.A[0] = 299.7*p + .107408*p*m.k - .009173*m.T*m.T
	m.A[1] = 251.88*p + .016321*p*m.k
	m.A[2] = 251.83*p + 26.651886*p*m.k
	m.A[3] = 349.42*p + 36.412478*p*m.k
	m.A[4] = 84.66*p + 18.206239*p*m.k
	m.A[5] = 141.74*p + 53.303771*p*m.k
	m.A[6] = 207.17*p + 2.453732*p*m.k
	m.A[7] = 154.84*p + 7.30686*p*m.k
	m.A[8] = 34.52*p + 27.261239*p*m.k
	m.A[9] = 207.19*p + .121824*p*m.k
	m.A[10] = 291.34*p + 1.844379*p*m.k
	m.A[11] = 161.72*p + 24.198154*p*m.k
	m.A[12] = 239.56*p + 25.513099*p*m.k
	m.A[13] = 331.55*p + 3.592518*p*m.k
	return m
}

// new or full corrections
func (e *mp) nfc(c *[25]float64) float64 {
	return c[0]*math.Sin(e.Mʹ) +
		c[1]*math.Sin(e.M)*e.E +
		c[2]*math.Sin(2*e.Mʹ) +
		c[3]*math.Sin(2*e.F) +
		c[4]*math.Sin(e.Mʹ-e.M)*e.E +
		c[5]*math.Sin(e.Mʹ+e.M)*e.E +
		c[6]*math.Sin(2*e.M)*e.E*e.E +
		c[7]*math.Sin(e.Mʹ-2*e.F) +
		c[8]*math.Sin(e.Mʹ+2*e.F) +
		c[9]*math.Sin(2*e.Mʹ+e.M)*e.E +
		c[10]*math.Sin(3*e.Mʹ) +
		c[11]*math.Sin(e.M+2*e.F)*e.E +
		c[12]*math.Sin(e.M-2*e.F)*e.E +
		c[13]*math.Sin(2*e.Mʹ-e.M)*e.E +
		c[14]*math.Sin(e.Ω) +
		c[15]*math.Sin(e.Mʹ+2*e.M) +
		c[16]*math.Sin(2*(e.Mʹ-e.F)) +
		c[17]*math.Sin(3*e.M) +
		c[18]*math.Sin(e.Mʹ+e.M-2*e.F) +
		c[19]*math.Sin(2*(e.Mʹ+e.F)) +
		c[20]*math.Sin(e.Mʹ+e.M+2*e.F) +
		c[21]*math.Sin(e.Mʹ-e.M+2*e.F) +
		c[22]*math.Sin(e.Mʹ-e.M-2*e.F) +
		c[23]*math.Sin(3*e.Mʹ+e.M) +
		c[24]*math.Sin(4*e.Mʹ)
}

// new coefficients
var nc = [25]float64{
	-.4072,
	.17241,
	.01608,
	.01039,
	.00739,
	-.00514,
	.00208,
	-.00111,
	-.00057,
	.00056,
	-.00042,
	.00042,
	.00038,
	-.00024,
	-.00017,
	-.00007,
	.00004,
	.00004,
	.00003,
	.00003,
	-.00003,
	.00003,
	-.00002,
	-.00002,
	.00002,
}

// full coefficients
var fc = [25]float64{
	-.40614,
	.17302,
	.01614,
	.01043,
	.00734,
	-.00515,
	.00209,
	-.00111,
	-.00057,
	.00056,
	-.00042,
	.00042,
	.00038,
	-.00024,
	-.00017,
	-.00007,
	.00004,
	.00004,
	.00003,
	.00003,
	-.00003,
	.00003,
	-.00002,
	-.00002,
	.00002,
}

// first or last corrections
func (m *mp) flc() float64 {
	return -.62801*math.Sin(m.Mʹ) +
		.17172*math.Sin(m.M)*m.E +
		-.01183*math.Sin(m.Mʹ+m.M)*m.E +
		.00862*math.Sin(2*m.Mʹ) +
		.00804*math.Sin(2*m.F) +
		.00454*math.Sin(m.Mʹ-m.M)*m.E +
		.00204*math.Sin(2*m.M)*m.E*m.E +
		-.0018*math.Sin(m.Mʹ-2*m.F) +
		-.0007*math.Sin(m.Mʹ+2*m.F) +
		-.0004*math.Sin(3*m.Mʹ) +
		-.00034*math.Sin(2*m.Mʹ-m.M)*m.E +
		.00032*math.Sin(m.M+2*m.F)*m.E +
		.00032*math.Sin(m.M-2*m.F)*m.E +
		-.00028*math.Sin(m.Mʹ+2*m.M)*m.E*m.E +
		.00027*math.Sin(2*m.Mʹ+m.M)*m.E +
		-.00017*math.Sin(m.Ω) +
		-.00005*math.Sin(m.Mʹ-m.M-2*m.F) +
		.00004*math.Sin(2*m.Mʹ+2*m.F) +
		-.00004*math.Sin(m.Mʹ+m.M+2*m.F) +
		.00004*math.Sin(m.Mʹ-2*m.M) +
		.00003*math.Sin(m.Mʹ+m.M-2*m.F) +
		.00003*math.Sin(3*m.M) +
		.00002*math.Sin(2*m.Mʹ-2*m.F) +
		.00002*math.Sin(m.Mʹ-m.M+2*m.F) +
		-.00002*math.Sin(3*m.Mʹ+m.M)
}

func (m *mp) w() float64 {
	return .00306 - .00038*m.E*math.Cos(m.M) + .00026*math.Cos(m.Mʹ) -
		.00002*(math.Cos(m.Mʹ-m.M)-math.Cos(m.Mʹ+m.M)-math.Cos(2*m.F))
}

// additional corrections
func (m *mp) a() float64 {
	var a float64
	for i, c := range ac {
		a += c * math.Sin(m.A[i])
	}
	return a
}

var ac = [14]float64{
	.000325,
	.000165,
	.000164,
	.000126,
	.00011,
	.000062,
	.00006,
	.000056,
	.000047,
	.000042,
	.000040,
	.000037,
	.000035,
	.000023,
}
