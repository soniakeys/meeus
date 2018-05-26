// Copyright 2013 Sonia Keys
// License: MIT

// Solstice: Chapter 27: Equinoxes and Solstices.
package solstice

import (
	"math"

	"github.com/soniakeys/meeus/v3/base"
	pp "github.com/soniakeys/meeus/v3/planetposition"
	"github.com/soniakeys/meeus/v3/solar"
	"github.com/soniakeys/unit"
)

var (
	mc0 = []float64{1721139.29189, 365242.13740, .06134, .00111, -.00071}
	jc0 = []float64{1721233.25401, 365241.72562, -.05323, .00907, .00025}
	sc0 = []float64{1721325.70455, 365242.49558, -.11677, -.00297, .00074}
	dc0 = []float64{1721414.39987, 365242.88257, -.00769, -.00933, -.00006}

	mc2 = []float64{2451623.80984, 365242.37404, .05169, -.00411, -.00057}
	jc2 = []float64{2451716.56767, 365241.62603, .00325, .00888, -.00030}
	sc2 = []float64{2451810.21715, 365242.01767, -.11575, .00337, .00078}
	dc2 = []float64{2451900.05952, 365242.74049, -.06223, -.00823, .00032}
)

type term struct {
	a, b, c float64
}

var terms = []term{
	{485, 324.96, 1934.136},
	{203, 337.23, 32964.467},
	{199, 342.08, 20.186},
	{182, 27.85, 445267.112},
	{156, 73.14, 45036.886},
	{136, 171.52, 22518.443},
	{77, 222.54, 65928.934},
	{74, 296.72, 3034.906},
	{70, 243.58, 9037.513},
	{58, 119.81, 33718.147},
	{52, 297.17, 150.678},
	{50, 21.02, 2281.226},

	{45, 247.54, 29929.562},
	{44, 325.15, 31555.956},
	{29, 60.93, 4443.417},
	{18, 155.12, 67555.328},
	{17, 288.79, 4562.452},
	{16, 198.04, 62894.029},
	{14, 199.76, 31436.921},
	{12, 95.39, 14577.848},
	{12, 287.11, 31931.756},
	{12, 320.81, 34777.259},
	{9, 227.73, 1222.114},
	{8, 15.45, 16859.074},
}

// March returns the JDE of the March equinox for the given year.
//
// Results are valid for the years -1000 to +3000.
//
// Accuracy is within one minute of time for the years 1951-2050.
func March(y int) float64 {
	if y < 1000 {
		return eq(y, mc0)
	}
	return eq(y-2000, mc2)
}

// June returns the JDE of the June solstice for the given year.
//
// Results are valid for the years -1000 to +3000.
//
// Accuracy is within one minute of time for the years 1951-2050.
func June(y int) float64 {
	if y < 1000 {
		return eq(y, jc0)
	}
	return eq(y-2000, jc2)
}

// September returns the JDE of the September equinox for the given year.
//
// Results are valid for the years -1000 to +3000.
//
// Accuracy is within one minute of time for the years 1951-2050.
func September(y int) float64 {
	if y < 1000 {
		return eq(y, sc0)
	}
	return eq(y-2000, sc2)
}

// December returns the JDE of the December solstice for a given year.
//
// Results are valid for the years -1000 to +3000.
//
// Accuracy is within one minute of time for the years 1951-2050.
func December(y int) float64 {
	if y < 1000 {
		return eq(y, dc0)
	}
	return eq(y-2000, dc2)
}

func eq(y int, c []float64) float64 {
	J0 := base.Horner(float64(y)*.001, c...)
	T := base.J2000Century(J0)
	W := 35999.373*math.Pi/180*T - 2.47*math.Pi/180
	Δλ := 1 + .0334*math.Cos(W) + .0007*math.Cos(2*W)
	S := 0.
	for i := len(terms) - 1; i >= 0; i-- {
		t := &terms[i]
		S += t.a * math.Cos((t.b+t.c*T)*math.Pi/180)
	}
	return J0 + .00001*S/Δλ
}

// March2 returns a more accurate JDE of the March equinox.
//
// Result is accurate to one second of time.
//
// Parameter e must be a V87Planet object representing Earth, obtained with
// the package planetposition and code similar to
//
//	e, err := planetposition.LoadPlanet(planetposition.Earth, "")
//	    if err != nil {
//	        ....
//
// See example under June2.
func March2(y int, e *pp.V87Planet) float64 {
	if y < 1000 {
		return eq2(y, e, 0, mc0)
	}
	return eq2(y-2000, e, 0, mc2)
}

// June2 returns a more accurate JDE of the June solstice.
//
// Result is accurate to one second of time.
//
// Parameter e must be a V87Planet object representing Earth, obtained with
// the package planetposition.
func June2(y int, e *pp.V87Planet) float64 {
	if y < 1000 {
		return eq2(y, e, math.Pi/2, jc0)
	}
	return eq2(y-2000, e, math.Pi/2, jc2)
}

// September2 returns a more accurate JDE of the September equinox.
//
// Result is accurate to one second of time.
//
// Parameter e must be a V87Planet object representing Earth, obtained with
// the package planetposition and code similar to
//
//	e, err := planetposition.LoadPlanet(planetposition.Earth, "")
//	    if err != nil {
//	        ....
//
// See example under June2.
func September2(y int, e *pp.V87Planet) float64 {
	if y < 1000 {
		return eq2(y, e, math.Pi, sc0)
	}
	return eq2(y-2000, e, math.Pi, sc2)
}

// December2 returns a more accurate JDE of the December solstice.
//
// Result is accurate to one second of time.
//
// Parameter e must be a V87Planet object representing Earth, obtained with
// the package planetposition and code similar to
//
//	e, err := planetposition.LoadPlanet(planetposition.Earth, "")
//	    if err != nil {
//	        ....
//
// See example under June2.
func December2(y int, e *pp.V87Planet) float64 {
	if y < 1000 {
		return eq2(y, e, math.Pi*3/2, dc0)
	}
	return eq2(y-2000, e, math.Pi*3/2, dc2)
}

func eq2(y int, e *pp.V87Planet, q unit.Angle, c []float64) float64 {
	J0 := base.Horner(float64(y)*.001, c...)
	for {
		λ, _, _ := solar.ApparentVSOP87(e, J0)
		c := 58 * (q - λ).Sin() // (27.1) p. 180
		J0 += c
		if math.Abs(c) < .000005 {
			break
		}
	}
	return J0
}
