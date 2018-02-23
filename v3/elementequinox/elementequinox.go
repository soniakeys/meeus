// Copyright 2013 Sonia Keys
// License: MIT

// Elementequinox: Chapter 24, Reduction of Ecliptical Elements
// from one Equinox to another one.
//
// See package precess for the method EclipticPrecessor.ReduceElements and
// associated example.  The method is described in this chapter but located
// in package precess so that it can be a method of EclipticPrecessor.
package elementequinox

import (
	"math"

	"github.com/soniakeys/unit"
)

// Elements are the orbital elements of a solar system object which change
// from one equinox to another.
type Elements struct {
	Inc  unit.Angle // inclination
	Peri unit.Angle // argument of perihelion (ω)
	Node unit.Angle // longitude of ascending node (Ω)
}

// ReduceB1950ToJ2000 reduces orbital elements of a solar system body from
// equinox B1950 to J2000.
func ReduceB1950ToJ2000(eFrom, eTo *Elements) *Elements {
	// (24.4) p. 161
	const S = .0001139788
	const C = .9999999935
	W := eFrom.Node - unit.AngleFromDeg(174.298782)
	si, ci := eFrom.Inc.Sincos()
	sW, cW := W.Sincos()
	A := si * sW
	B := C*si*cW - S*ci
	eTo.Inc = unit.Angle(math.Asin(math.Hypot(A, B)))
	eTo.Node = (unit.AngleFromDeg(174.997194) +
		unit.Angle(math.Atan2(A, B))).Mod1()
	eTo.Peri = (eFrom.Peri +
		unit.Angle(math.Atan2(-S*sW, C*si-S*ci*cW))).Mod1()
	return eTo
}

var (
	_Lp = unit.AngleFromDeg(4.50001688)
	_L  = unit.AngleFromDeg(5.19856209)
	_J  = unit.AngleFromDeg(.00651966)
)

// ReduceB1950ToJ2000 reduces orbital elements of a solar system body from
// equinox B1950 in the FK4 system to equinox J2000 in the FK5 system.
func ReduceB1950FK4ToJ2000FK5(eFrom, eTo *Elements) *Elements {
	W := _L + eFrom.Node
	si, ci := eFrom.Inc.Sincos()
	sJ, cJ := _J.Sincos()
	sW, cW := W.Sincos()
	eTo.Inc = unit.Angle(math.Acos(ci*cJ - si*sJ*cW))
	eTo.Node = (unit.Angle(math.Atan2(si*sW, ci*sJ+si*cJ*cW)) -
		_Lp).Mod1()
	eTo.Peri = (eFrom.Peri +
		unit.Angle(math.Atan2(sJ*sW, si*cJ+ci*sJ*cW))).Mod1()
	return eTo
}
