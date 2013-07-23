// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Parabolic: Chapter 34, Parabolic Motion.
package parabolic

import (
	"math"

	"github.com/soniakeys/meeus/base"
)

// Elements holds parabolic elements needed for computing true anomaly and distance.
type Elements struct {
	TimeP float64 // time of perihelion, T
	PDis  float64 // perihelion distance, q
}

// AnomalyDistance returns true anomaly and distance of a body in a parabolic orbit of the Sun.
//
// True anomaly ν returned in radians.
// Distance r returned in AU.
func (e *Elements) AnomalyDistance(jde float64) (ν, r float64) {
	W := 3 * base.K / math.Sqrt2 * (jde - e.TimeP) / e.PDis / math.Sqrt(e.PDis)
	G := W * .5
	Y := math.Cbrt(G + math.Sqrt(G*G+1))
	s := Y - 1/Y
	ν = 2 * math.Atan(s)
	r = e.PDis * (1 + s*s)
	return
}
