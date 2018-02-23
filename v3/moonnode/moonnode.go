// Copyright 2013 Sonia Keys
// License: MIT

// Moonnode: Chapter 51, Passages of the Moon through the Nodes.
package moonnode

import (
	"math"

	"github.com/soniakeys/meeus/v3/base"
)

// Ascending returns the date of passage of the Moon through an ascending node.
//
// Argument year is a decimal year specifying a date near the event.
//
// Returned is the jde of the event nearest the given date.
func Ascending(year float64) float64 {
	return node(year, 0)
}

// Descending returns the date of passage of the Moon through a descending node.
//
// Argument year is a decimal year specifying a date near the event.
//
// Returned is the jde of the event nearest the given date.
func Descending(year float64) float64 {
	return node(year, .5)
}

func node(y, h float64) float64 {
	k := (y - 2000.05) * 13.4223 // (50.1) p. 355
	k = math.Floor(k-h+.5) + h   // snap to half orbit
	const p = math.Pi / 180
	const ck = 1 / 1342.23
	T := k * ck
	D := base.Horner(T, 183.638*p, 331.73735682*p/ck,
		.0014852*p, .00000209*p, -.00000001*p)
	M := base.Horner(T, 17.4006*p, 26.8203725*p/ck,
		.0001186*p, .00000006*p)
	Mʹ := base.Horner(T, 38.3776*p, 355.52747313*p/ck,
		.0123499*p, .000014627*p, -.000000069*p)
	Ω := base.Horner(T, 123.9767*p, -1.44098956*p/ck,
		.0020608*p, .00000214*p, -.000000016*p)
	V := base.Horner(T, 299.75*p, 132.85*p, -.009173*p)
	P := Ω + 272.75*p - 2.3*p*T
	E := base.Horner(T, 1, -.002516, -.0000074)
	return base.Horner(T, 2451565.1619, 27.212220817/ck,
		.0002762, .000000021, -.000000000088) +
		-.4721*math.Sin(Mʹ) +
		-.1649*math.Sin(2*D) +
		-.0868*math.Sin(2*D-Mʹ) +
		.0084*math.Sin(2*D+Mʹ) +
		-.0083*math.Sin(2*D-M)*E +
		-.0039*math.Sin(2*D-M-Mʹ)*E +
		.0034*math.Sin(2*Mʹ) +
		-.0031*math.Sin(2*(D-Mʹ)) +
		.003*math.Sin(2*D+M)*E +
		.0028*math.Sin(M-Mʹ)*E +
		.0026*math.Sin(M)*E +
		.0025*math.Sin(4*D) +
		.0024*math.Sin(D) +
		.0022*math.Sin(M+Mʹ)*E +
		.0017*math.Sin(Ω) +
		.0014*math.Sin(4*D-Mʹ) +
		.0005*math.Sin(2*D+M-Mʹ)*E +
		.0004*math.Sin(2*D-M+Mʹ)*E +
		-.0003*math.Sin(2*(D-M))*E +
		.0003*math.Sin(4*D-M)*E +
		.0003*math.Sin(V) +
		.0003*math.Sin(P)
}
