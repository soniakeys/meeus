// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Perihelion: Chapter 38, Planets in Perihelion and Aphelion.
package perihelion

import (
	"math"

	"github.com/soniakeys/meeus/base"
)

// Planet constants for first argument of Perihelion and Aphelion functions.
const (
	Mercury = iota
	Venus
	Earth
	Mars
	Jupiter
	Saturn
	EMBary
)

// Perihelion returns the periheion event nearest the given time.
//
// Argument p must be one of the planet constants above, y is a year number
// near the perihelion event.
func Perihelion(p int, y float64) (jde float64) {
	return ap(p, y, pf, false)
}

func pf(x float64) float64 {
	return math.Floor(x + .5)
}

// Aphelion returns the periheion event nearest the given time.
//
// Argument p must be one of the planet constants above, y is a year number
// near the aphelion event.
func Aphelion(p int, y float64) (jde float64) {
	return ap(p, y, af, true)
}

func af(x float64) float64 {
	return math.Floor(x) + .5
}

func ap(p int, y float64, f func(float64) float64, a bool) float64 {
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
}

var c = [][]float64{
	{2451590.257, 87.96934963},
	{2451738.233, 224.7008188, -.0000000327},
	{2451547.507, 365.2596358, .0000000156},
	{2452195.026, 686.9957857, -.0000001187},
	{2455636.936, 4332.897065, .0001367},
	{2452830.12, 10764.21676, .000827},
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
