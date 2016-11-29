// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Parabolic: Chapter 34, Parabolic Motion.
package parabolic

import (
	"math"

	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/unit"
)

// Elements holds parabolic elements needed for computing true anomaly and distance.
type Elements struct {
	TimeP float64 // time of perihelion, T, as JD
	PDis  float64 // perihelion distance, q, in AU
}

// AnomalyDistance returns true anomaly and distance of a body in a parabolic orbit of the Sun.
//
// Distance r returned in AU.
func (e *Elements) AnomalyDistance(jde float64) (ν unit.Angle, r float64) {
	W := 3 * base.K / math.Sqrt2 * (jde - e.TimeP) / e.PDis / math.Sqrt(e.PDis)
	G := W * .5
	Y := math.Cbrt(G + math.Sqrt(G*G+1))
	s := Y - 1/Y
	ν = unit.Angle(2 * math.Atan(s))
	r = e.PDis * (1 + s*s)
	return
}
