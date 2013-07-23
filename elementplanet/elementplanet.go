// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Elementplanet: Chapter 31, Elements of Planetary Orbits.
//
// Partial: Only Mercury is implemented, and only mean elements.  The tables
// present too much chance of typographic errors.
package elementplanet

import (
	"math"

	"github.com/soniakeys/meeus/base"
)

const (
	Mercury = iota
	Venus
	Earth
	Mars
	Jupiter
	Saturn
	Uranus
	Neptune
	nPlanets // :(
)

// Elements contains orbital elements as returned by functions in this package.
//
// Some other elements easily derived from these are
//
//	Mean Anomolay, M = Lon - Peri
//	Argument of Perihelion, ω = Peri - Node
type Elements struct {
	Lon  float64 // mean longitude, L
	Axis float64 // semimajor axis, a
	Ecc  float64 // eccentricity, e
	Inc  float64 // inclination, i
	Node float64 // longitude of ascending node, Ω
	Peri float64 // longitude of perihelion, ϖ (Meeus likes π better)
}

type c6 struct {
	L, a, e, i, Ω, ϖ []float64
}

var cMean = []c6{
	{ // Mercury
		[]float64{252.250906, 149474.0722491, .0003035, .000000018},
		[]float64{.38709831},
		[]float64{.20563175, .000020407, -.0000000283, -.00000000018},
		[]float64{7.004986, .0018215, -.00001810, .000000056},
		[]float64{48.330893, 1.1861883, .00017542, .00000215},
		[]float64{77.456119, 1.5564776, .00029544, .000000009},
	},
}

// Mean returns mean orbital elements for a planet
//
// Argument p must be a planet const as defined above, argument e is
// a result parameter.  A valid non-nil pointer to an Elements struct
// must be passed in.
//
// Results are referenced to mean dynamical ecliptic and equinox of date.
//
// Semimajor axis is in AU, angular elements are in radians.
func Mean(p int, jde float64, e *Elements) {
	T := base.J2000Century(jde)
	c := &cMean[p]
	e.Lon = base.PMod(base.Horner(T, c.L...)*math.Pi/180, 2*math.Pi)
	e.Axis = base.Horner(T, c.a...)
	e.Ecc = base.Horner(T, c.e...)
	e.Inc = base.Horner(T, c.i...) * math.Pi / 180
	e.Node = base.Horner(T, c.Ω...) * math.Pi / 180
	e.Peri = base.Horner(T, c.ϖ...) * math.Pi / 180
}
