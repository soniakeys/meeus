// Copyright 2013 Sonia Keys
// License: MIT

// Moonmaxdec: Chapter 52, Maximum Declinations of the Moon
package moonmaxdec

import (
	"math"

	"github.com/soniakeys/meeus/v3/base"
	"github.com/soniakeys/unit"
)

// North computes the maximum northern declination of the Moon near a given date.
//
// Argument year is a decimal year specifying a date near the event.
//
// Returned is the jde of the event nearest the given date and the declination
// of the Moon at that time.
func North(y float64) (jde float64, δ unit.Angle) {
	return max(y, &nc)
}

// South computes the maximum southern declination of the Moon near a given date.
//
// Argument year is a decimal year specifying a date near the event.
//
// Returned is the jde of the event nearest the given date and the declination
// of the Moon at that time.
func South(y float64) (jde float64, δ unit.Angle) {
	return max(y, &sc)
}

const p = math.Pi / 180

func max(y float64, c *mc) (jde float64, δ unit.Angle) {
	k := (y - 2000.03) * 13.3686 // (52.1) p. 367
	k = math.Floor(k + .5)
	const ck = 1 / 1336.86
	T := k * ck
	D := base.Horner(T, c.D, 333.0705546*p/ck, -.0004214*p, .00000011*p)
	M := base.Horner(T, c.M, 26.9281592*p/ck, -.0000355*p, -.0000001*p)
	Mʹ := base.Horner(T, c.Mʹ, 356.9562794*p/ck, .0103066*p, .00001251*p)
	F := base.Horner(T, c.F, 1.4467807*p/ck, -.002069*p, -.00000215*p)
	E := base.Horner(T, 1, -.002516, -.0000074)
	jde = base.Horner(T, c.JDE, 27.321582247/ck, .000119804, -.000000141) +
		c.tc[0]*math.Cos(F) +
		c.tc[1]*math.Sin(Mʹ) +
		c.tc[2]*math.Sin(2*F) +
		c.tc[3]*math.Sin(2*D-Mʹ) +
		c.tc[4]*math.Cos(Mʹ-F) +
		c.tc[5]*math.Cos(Mʹ+F) +
		c.tc[6]*math.Sin(2*D) +
		c.tc[7]*math.Sin(M)*E +
		c.tc[8]*math.Cos(3*F) +
		c.tc[9]*math.Sin(Mʹ+2*F) +
		c.tc[10]*math.Cos(2*D-F) +
		c.tc[11]*math.Cos(2*D-Mʹ-F) +
		c.tc[12]*math.Cos(2*D-Mʹ+F) +
		c.tc[13]*math.Cos(2*D+F) +
		c.tc[14]*math.Sin(2*Mʹ) +
		c.tc[15]*math.Sin(Mʹ-2*F) +
		c.tc[16]*math.Cos(2*Mʹ-F) +
		c.tc[17]*math.Sin(Mʹ+3*F) +
		c.tc[18]*math.Sin(2*D-M-Mʹ)*E +
		c.tc[19]*math.Cos(Mʹ-2*F) +
		c.tc[20]*math.Sin(2*(D-Mʹ)) +
		c.tc[21]*math.Sin(F) +
		c.tc[22]*math.Sin(2*D+Mʹ) +
		c.tc[23]*math.Cos(Mʹ+2*F) +
		c.tc[24]*math.Sin(2*D-M)*E +
		c.tc[25]*math.Sin(Mʹ+F) +
		c.tc[26]*math.Sin(M-Mʹ)*E +
		c.tc[27]*math.Sin(Mʹ-3*F) +
		c.tc[28]*math.Sin(2*Mʹ+F) +
		c.tc[29]*math.Cos(2*(D-Mʹ)-F) +
		c.tc[30]*math.Sin(3*F) +
		c.tc[31]*math.Cos(Mʹ+3*F) +
		c.tc[32]*math.Cos(2*Mʹ) +
		c.tc[33]*math.Cos(2*D-Mʹ) +
		c.tc[34]*math.Cos(2*D+Mʹ+F) +
		c.tc[35]*math.Cos(Mʹ) +
		c.tc[36]*math.Sin(3*Mʹ+F) +
		c.tc[37]*math.Sin(2*D-Mʹ+F) +
		c.tc[38]*math.Cos(2*(D-Mʹ)) +
		c.tc[39]*math.Cos(D+F) +
		c.tc[40]*math.Sin(M+Mʹ)*E +
		c.tc[41]*math.Sin(2*(D-F)) +
		c.tc[42]*math.Cos(2*Mʹ+F) +
		c.tc[43]*math.Cos(3*Mʹ+F)
	δ = unit.Angle(23.6961*p - .013004*p*T +
		c.dc[0]*math.Sin(F) +
		c.dc[1]*math.Cos(2*F) +
		c.dc[2]*math.Sin(2*D-F) +
		c.dc[3]*math.Sin(3*F) +
		c.dc[4]*math.Cos(2*(D-F)) +
		c.dc[5]*math.Cos(2*D) +
		c.dc[6]*math.Sin(Mʹ-F) +
		c.dc[7]*math.Sin(Mʹ+2*F) +
		c.dc[8]*math.Cos(F) +
		c.dc[9]*math.Sin(2*D+M-F)*E +
		c.dc[10]*math.Sin(Mʹ+3*F) +
		c.dc[11]*math.Sin(D+F) +
		c.dc[12]*math.Sin(Mʹ-2*F) +
		c.dc[13]*math.Sin(2*D-M-F)*E +
		c.dc[14]*math.Sin(2*D-Mʹ-F) +
		c.dc[15]*math.Cos(Mʹ+F) +
		c.dc[16]*math.Cos(Mʹ+2*F) +
		c.dc[17]*math.Cos(2*Mʹ+F) +
		c.dc[18]*math.Cos(Mʹ-3*F) +
		c.dc[19]*math.Cos(2*Mʹ-F) +
		c.dc[20]*math.Cos(Mʹ-2*F) +
		c.dc[21]*math.Sin(2*Mʹ) +
		c.dc[22]*math.Sin(3*Mʹ+F) +
		c.dc[23]*math.Cos(2*D+M-F)*E +
		c.dc[24]*math.Cos(Mʹ-F) +
		c.dc[25]*math.Cos(3*F) +
		c.dc[26]*math.Sin(2*D+F) +
		c.dc[27]*math.Cos(Mʹ+3*F) +
		c.dc[28]*math.Cos(D+F) +
		c.dc[29]*math.Sin(2*Mʹ-F) +
		c.dc[30]*math.Cos(3*Mʹ+F) +
		c.dc[31]*math.Cos(2*(D+Mʹ)+F) +
		c.dc[32]*math.Sin(2*(D-Mʹ)-F) +
		c.dc[33]*math.Cos(2*Mʹ) +
		c.dc[34]*math.Cos(Mʹ) +
		c.dc[35]*math.Sin(2*F) +
		c.dc[36]*math.Sin(Mʹ+F))
	return jde, δ.Mul(c.s)
}

type mc struct {
	D, M, Mʹ, F, JDE, s float64
	tc                  [44]float64
	dc                  [37]float64
}

// north coefficients
var nc = mc{
	D:   152.2029 * p,
	M:   14.8591 * p,
	Mʹ:  4.6881 * p,
	F:   325.8867 * p,
	JDE: 2451562.5897,
	s:   1,
	tc: [44]float64{
		.8975,
		-.4726,
		-.1030,
		-.0976,
		-.0462,
		-.0461,
		-.0438,
		.0162,
		-.0157,
		.0145,
		.0136,
		-.0095,
		-.0091,
		-.0089,
		.0075,
		-.0068,
		.0061,
		-.0047,
		-.0043,
		-.004,
		-.0037,
		.0031,
		.0030,
		-.0029,
		-.0029,
		-.0027,
		.0024,
		-.0021,
		.0019,
		.0018,
		.0018,
		.0017,
		.0017,
		-.0014,
		.0013,
		.0013,
		.0012,
		.0011,
		-.0011,
		.001,
		.001,
		-.0009,
		.0007,
		-.0007,
	},
	dc: [37]float64{
		5.1093 * p,
		.2658 * p,
		.1448 * p,
		-.0322 * p,
		.0133 * p,
		.0125 * p,
		-.0124 * p,
		-.0101 * p,
		.0097 * p,
		-.0087 * p,
		.0074 * p,
		.0067 * p,
		.0063 * p,
		.0060 * p,
		-.0057 * p,
		-.0056 * p,
		.0052 * p,
		.0041 * p,
		-.004 * p,
		.0038 * p,
		-.0034 * p,
		-.0029 * p,
		.0029 * p,
		-.0028 * p,
		-.0028 * p,
		-.0023 * p,
		-.0021 * p,
		.0019 * p,
		.0018 * p,
		.0017 * p,
		.0015 * p,
		.0014 * p,
		-.0012 * p,
		-.0012 * p,
		-.001 * p,
		-.001 * p,
		.0006 * p,
	},
}

// south coefficients
var sc = mc{
	D:   345.6676 * p,
	M:   1.3951 * p,
	Mʹ:  186.21 * p,
	F:   145.1633 * p,
	JDE: 2451548.9289,
	s:   -1,
	tc: [44]float64{
		-.8975,
		-.4726,
		-.1030,
		-.0976,
		.0541,
		.0516,
		-.0438,
		.0112,
		.0157,
		.0023,
		-.0136,
		.011,
		.0091,
		.0089,
		.0075,
		-.003,
		-.0061,
		-.0047,
		-.0043,
		.004,
		-.0037,
		-.0031,
		.0030,
		.0029,
		-.0029,
		-.0027,
		.0024,
		-.0021,
		-.0019,
		-.0006,
		-.0018,
		-.0017,
		.0017,
		.0014,
		-.0013,
		-.0013,
		.0012,
		.0011,
		.0011,
		.001,
		.001,
		-.0009,
		-.0007,
		-.0007,
	},
	dc: [37]float64{
		-5.1093 * p,
		.2658 * p,
		-.1448 * p,
		.0322 * p,
		.0133 * p,
		.0125 * p,
		-.0015 * p,
		.0101 * p,
		-.0097 * p,
		.0087 * p,
		.0074 * p,
		.0067 * p,
		-.0063 * p,
		-.0060 * p,
		.0057 * p,
		-.0056 * p,
		-.0052 * p,
		-.0041 * p,
		-.004 * p,
		-.0038 * p,
		.0034 * p,
		-.0029 * p,
		.0029 * p,
		.0028 * p,
		-.0028 * p,
		.0023 * p,
		.0021 * p,
		.0019 * p,
		.0018 * p,
		-.0017 * p,
		.0015 * p,
		.0014 * p,
		.0012 * p,
		-.0012 * p,
		.001 * p,
		-.001 * p,
		.0037 * p,
	},
}
