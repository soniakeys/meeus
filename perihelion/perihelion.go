// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Perihelion: Chapter 38, Planets in Perihelion and Aphelion.
package perihelion

import (
	"math"

	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/meeus/interp"
	pp "github.com/soniakeys/meeus/planetposition"
)

// Planet constants for first argument of Perihelion and Aphelion functions.
const (
	Mercury = iota
	Venus
	Earth
	Mars
	Jupiter
	Saturn
	Uranus
	Neptune
	EMBary
)

// Perihelion returns an approximate jde of the perihelion event nearest the given time.
//
// Argument p must be one of the planet constants above, y is a year number
// indicating a time near the perihelion event.
func Perihelion(p int, y float64) (jde float64) {
	return ap(p, y, false, pf)
}

func pf(x float64) float64 {
	return math.Floor(x + .5)
}

// Aphelion returns an approximate jde of the aphelion event nearest the given time.
//
// Argument p must be one of the planet constants above, y is a year number
// indicating a time near the aphelion event.
func Aphelion(p int, y float64) (jde float64) {
	return ap(p, y, true, af)
}

func af(x float64) float64 {
	return math.Floor(x) + .5
}

func ap(p int, y float64, a bool, f func(float64) float64) float64 {
	i := p
	if i == EMBary {
		i = Earth
	}
	k := f(ka[i].a * (y - ka[i].b))
	j := base.Horner(k, c[i]...)
	if p == Earth {
		c := ep
		if a {
			c = ea
		}
		for i := 0; i < 5; i++ {
			j += c[i] * math.Sin((ec[i].a+ec[i].b*k)*math.Pi/180)
		}
	}
	return j
}

type ab struct {
	a, b float64
}

var ka = []ab{
	{4.15201, 2000.12},
	{1.62549, 2000.53},
	{.99997, 2000.01},
	{.53166, 2001.78},
	{.0843, 2011.2},
	{.03393, 2003.52},
	{.0119, 2051.1},
	{.00607, 2047.5},
}

var c = [][]float64{
	{2451590.257, 87.96934963},
	{2451738.233, 224.7008188, -.0000000327},
	{2451547.507, 365.2596358, .0000000156},
	{2452195.026, 686.9957857, -.0000001187},
	{2455636.936, 4332.897065, .0001367},
	{2452830.12, 10764.21676, .000827},
	{2470213.5, 30694.8767, -.00541},
	{2468895.1, 60190.33, .03429},
}

var ec = []ab{
	{328.41, 132.788585},
	{316.13, 584.903153},
	{346.2, 450.380738},
	{136.95, 659.306737},
	{249.52, 329.653368},
}

var ep = []float64{1.278, -.055, -.091, -.056, -.045}
var ea = []float64{-1.352, .061, .062, .029, .031}

// Perihelion2 returns the perihelion event nearest the given time.
//
// Argument p must be one of the planet constants Mercury through Neptune;
// EMBary is not allowed.  Y is a year number near the perihelion event.
// D is the desired precision of the time result, in days. V must be a
// planetposition.V87Planet object consistent with argument p.
//
// Result jde is the time of the event, r is the distance of the planet
// from the Sun in AU.
func Perihelion2(p int, y, d float64, v *pp.V87Planet) (jde, r float64) {
	return ap2(p, y, d, v, false, pf)
}

// Aphelion2 returns the aphelion event nearest the given time.
//
// Argument p must be one of the planet constants Mercury through Neptune;
// EMBary is not allowed.  Y is a year number near the perihelion event.
// D is the desired precision of the time result, in days. V must be a
// planetposition.V87Planet object consistent with argument p.
//
// Result jde is the time of the event, r is the distance of the planet
// from the Sun in AU.
func Aphelion2(p int, y, d float64, v *pp.V87Planet) (jde, r float64) {
	return ap2(p, y, d, v, true, af)
}

func ap2(p int, y, d float64, v *pp.V87Planet, a bool, f func(float64) float64) (jde, r float64) {
	j1 := ap(p, y, a, f)
	if p != Neptune {
		return ap2a(j1, d, a, v)
	}
	// Meeus doesn't give an algorithm to handle the double extrema of
	// Neptune.  The algorithm here is to pick starting points several years
	// either side of the approximate date and let ap2a follow the slopes
	// from there.  It's rather slow, but seems to find correct answers.
	j0, r0 := ap2a(j1-5000, d, a, v)
	j2, r2 := ap2a(j1+5000, d, a, v)
	if r0 > r2 == a {
		return j0, r0
	}
	return j2, r2
}

func ap2a(j1, d float64, a bool, v *pp.V87Planet) (jde, r float64) {
	// Meeus doesn't give a complete algorithm for finding accurate answers.
	// The algorithm here starts with the approximate result algorithm
	// then crawls along the curve at resultion d until three points
	// are found containing the desired extremum.  It's slow if the starting
	// point is far away (the case of Neptune) or if high accuracy is
	// demanded.  1 day accuracy is generally quick.
	j0 := j1 - d
	j2 := j1 + d
	rr := make([]float64, 3)
	_, _, rr[0] = v.Position2000(j0)
	_, _, rr[1] = v.Position2000(j1)
	_, _, rr[2] = v.Position2000(j2)
	for {
		if a {
			if rr[1] > rr[0] && rr[1] > rr[2] {
				break
			}
		} else {
			if rr[1] < rr[0] && rr[1] < rr[2] {
				break
			}
		}
		if rr[0] < rr[2] == a {
			j0 = j1
			j1 = j2
			j2 += d
			rr[0] = rr[1]
			rr[1] = rr[2]
			_, _, rr[2] = v.Position2000(j2)
		} else {
			j2 = j1
			j1 = j0
			j0 -= d
			rr[2] = rr[1]
			rr[1] = rr[0]
			_, _, rr[0] = v.Position2000(j0)
		}
	}
	l, err := interp.NewLen3(j0, j2, rr)
	if err != nil {
		panic(err) // unexpected.
	}
	jde, r, err = l.Extremum()
	if err != nil {
		panic(err) // unexpected.
	}
	return
}
