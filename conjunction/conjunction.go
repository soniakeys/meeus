// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Conjunction: Chapter 18: Planetary Conjunctions.
package conjunction

import (
	"errors"

	"github.com/soniakeys/meeus/interp"
)

// Planetary computes a conjunction between two moving objects, such as planets.
//
// Conjunction is found with interpolation against length 5 ephemerides.
//
// T1, t5 are times of first and last rows of ephemerides.  The scale is
// arbitrary.
//
// R1, d1 is the ephemeris of the first object.  The columns may be celestial
// coordinates in right ascension and declination or ecliptic coordinates in
// longitude and latitude.
//
// R2, d2 is the ephemeris of the second object, in the same frame as the first.
//
// Return value t is time of conjunction in the scale of t1, t5.
// Δd is the amount that object 2 was "above" object 1 at the time of
// conjunction.
func Planetary(t1, t5 float64, r1, d1, r2, d2 []float64) (t, Δd float64, err error) {
	if len(r1) != 5 || len(d1) != 5 || len(r2) != 5 || len(d2) != 5 {
		err = errors.New("Five rows required in ephemerides")
		return
	}
	dr := make([]float64, 5, 10)
	dd := dr[5:10]
	for i, r := range r1 {
		dr[i] = r2[i] - r
		dd[i] = d2[i] - d1[i]
	}
	return conj(t1, t5, dr, dd)
}

// Stellar computes a conjunction between a moving and non-moving object.
//
// Arguments and return values same as with Planetary, except the non-moving
// object is r1, d1.  The ephemeris of the moving object is r2, d2.
func Stellar(t1, t5, r1, d1 float64, r2, d2 []float64) (t, Δd float64, err error) {
	if len(r2) != 5 || len(d2) != 5 {
		err = errors.New("Five rows required in ephemeris")
		return
	}
	dr := make([]float64, 5, 10)
	dd := dr[5:10]
	for i, r := range r2 {
		dr[i] = r - r1
		dd[i] = d2[i] - d1
	}
	return conj(t1, t5, dr, dd)
}

func conj(t1, t5 float64, dr, dd []float64) (t, Δd float64, err error) {
	var l5 *interp.Len5
	if l5, err = interp.NewLen5(t1, t5, dr); err != nil {
		return
	}
	if t, err = l5.Zero(true); err != nil {
		return
	}
	if l5, err = interp.NewLen5(t1, t5, dd); err != nil {
		return
	}
	Δd, err = l5.InterpolateXStrict(t)
	return
}
