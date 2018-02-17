// Copyright 2013 Sonia Keys
// License: MIT

// Nearparabolic: Chapter 35, Near-parabolic Motion.
package nearparabolic

import (
	"errors"
	"math"

	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/unit"
)

// Elements holds orbital elements for near-parabolic orbits.
type Elements struct {
	TimeP float64 // Time of Perihelion, T
	PDis  float64 // Perihelion distance, q
	Ecc   float64 // eccentricity, e
}

// AnomalyDistance returns true anomaly and distance for near-parabolic orbits.
//
// Distance r returned in AU.
// An error is returned if the algorithm fails to converge.
func (e *Elements) AnomalyDistance(jde float64) (ν unit.Angle, r float64, err error) {
	// fairly literal translation of code on p. 246
	q1 := base.K * math.Sqrt((1+e.Ecc)/e.PDis) / (2 * e.PDis) // line 20
	g := (1 - e.Ecc) / (1 + e.Ecc)                            // line 20

	t := jde - e.TimeP // line 22
	if t == 0 {        // line 24
		return 0, e.PDis, nil
	}
	d1, d := 10000., 1e-9        // line 14
	q2 := q1 * t                 // line 28
	s := 2. / (3 * math.Abs(q2)) // line 30
	s = 2 / math.Tan(2*math.Atan(math.Cbrt(math.Tan(math.Atan(s)/2))))
	if t < 0 { // line 34
		s = -s
	}
	if e.Ecc != 1 { // line 36
		l := 0 // line 38
		for {
			s0 := s // line 40
			z := 1.
			y := s * s
			g1 := -y * s
			q3 := q2 + 2*g*s*y/3 // line 42
			for {
				z += 1                          // line 44
				g1 = -g1 * g * y                // line 46
				z1 := (z - (z+1)*g) / (2*z + 1) // line 48
				f := z1 * g1                    // line 50
				q3 += f                         // line 52
				if z > 50 || math.Abs(f) > d1 { // line 54
					return 0, 0, errors.New("No convergence")
				}
				if math.Abs(f) <= d { // line 56
					break
				}
			}
			l++ // line 58
			if l > 50 {
				return 0, 0, errors.New("No convergence")
			}
			for {
				s1 := s // line 60
				s = (2*s*s*s/3 + q3) / (s*s + 1)
				if math.Abs(s-s1) <= d { // line 62
					break
				}
			}
			if math.Abs(s-s0) <= d { // line 64
				break
			}
		}
	}
	ν = unit.Angle(2 * math.Atan(s))               // line 66
	r = e.PDis * (1 + e.Ecc) / (1 + e.Ecc*ν.Cos()) // line 68
	if ν < 0 {                                     // line 70
		ν += 2 * math.Pi
	}
	return
}
