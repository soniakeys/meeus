// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Elementequinox: Chapter 24, Reduction of Ecliptical Elements
// from one Equinox to another one.
//
// See package precess for the method EclipticPrecessor.ReduceElements and
// associated example.  The method is described in this chapter but located
// in package precess so that it can be a method of EclipticPrecessor.
package elementequinox

import (
	"math"

	"github.com/soniakeys/meeus/base"
)

// Elements are the orbital elements of a solar system object which change
// from one equinox to another.
type Elements struct {
	Inc  float64 // inclination
	Peri float64 // argument of perihelion (ω)
	Node float64 // longitude of ascending node (Ω)
}

// ReduceB1950ToJ2000 reduces orbital elements of a solar system body from
// equinox B1950 to J2000.
func ReduceB1950ToJ2000(eFrom, eTo *Elements) *Elements {
	const S = .0001139788
	const C = .9999999935
	W := eFrom.Node - 174.298782*math.Pi/180
	si, ci := math.Sincos(eFrom.Inc)
	sW, cW := math.Sincos(W)
	A := si * sW
	B := C*si*cW - S*ci
	eTo.Inc = math.Asin(math.Hypot(A, B))
	eTo.Node = base.PMod(174.997194*math.Pi/180+math.Atan2(A, B),
		2*math.Pi)
	eTo.Peri = base.PMod(eFrom.Peri+math.Atan2(-S*sW, C*si-S*ci*cW),
		2*math.Pi)
	return eTo
}

// ReduceB1950ToJ2000 reduces orbital elements of a solar system body from
// equinox B1950 in the FK4 system to equinox J2000 in the FK5 system.
func ReduceB1950FK4ToJ2000FK5(eFrom, eTo *Elements) *Elements {
	const (
		Lp = 4.50001688 * math.Pi / 180
		L  = 5.19856209 * math.Pi / 180
		J  = .00651966 * math.Pi / 180
	)
	W := L + eFrom.Node
	si, ci := math.Sincos(eFrom.Inc)
	sJ, cJ := math.Sincos(J)
	sW, cW := math.Sincos(W)
	eTo.Inc = math.Acos(ci*cJ - si*sJ*cW)
	eTo.Node = base.PMod(math.Atan2(si*sW, ci*sJ+si*cJ*cW)-Lp,
		2*math.Pi)
	eTo.Peri = base.PMod(eFrom.Peri+math.Atan2(sJ*sW, si*cJ+ci*sJ*cW),
		2*math.Pi)
	return eTo
}
