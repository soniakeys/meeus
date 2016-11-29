// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Planetelements: Chapter 31, Elements of Planetary Orbits.
//
// Partial:  Only implemented for mean equinox of date.
package planetelements

import (
	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/unit"
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
//	Mean Anomaly, M = Lon - Peri
//	Argument of Perihelion, ω = Peri - Node
type Elements struct {
	Lon  unit.Angle // mean longitude, L
	Axis float64    // semimajor axis, a
	Ecc  float64    // eccentricity, e
	Inc  unit.Angle // inclination, i
	Node unit.Angle // longitude of ascending node, Ω
	Peri unit.Angle // longitude of perihelion, ϖ (Meeus likes π better)
}

type c6 struct {
	L, a, e, i, Ω, ϖ []float64
}

// Table 31.A, p. 212
var cMean = []c6{
	{ // Mercury
		[]float64{252.250906, 149474.0722491, .0003035, .000000018},
		[]float64{.38709831},
		[]float64{.20563175, .000020407, -.0000000283, -.00000000018},
		[]float64{7.004986, .0018215, -.0000181, .000000056},
		[]float64{48.330893, 1.1861883, .00017542, .000000215},
		[]float64{77.456119, 1.5564776, .00029544, .000000009},
	},
	{ // Venus
		[]float64{181.979801, 58519.2130302, .00031014, .000000015},
		[]float64{.72332982},
		[]float64{.00677192, -.000047765, .0000000981, .00000000046},
		[]float64{3.394662, .0010037, -.00000088, -.000000007},
		[]float64{76.67992, .9011206, .00040618, -.000000093},
		[]float64{131.563703, 1.4022288, -.00107618, -.000005678},
	},
	{ // Earth
		[]float64{100.466457, 36000.7698278, .00030322, .00000002},
		[]float64{1.000001018},
		[]float64{.01670863, -.000042037, -.0000001267, .00000000014},
		[]float64{0},
		nil,
		[]float64{102.937348, 1.7195366, .00045688, -.000000018},
	},
	{ // Mars
		[]float64{355.433, 19141.6964471, .00031052, .000000016},
		[]float64{1.523679342},
		[]float64{.09340065, .000090484, -.0000000806, -.00000000025},
		[]float64{1.849726, -.0006011, .00001276, -.000000007},
		[]float64{49.558093, .7720959, .00001557, .000002267},
		[]float64{336.060234, 1.8410449, .00013477, .000000536},
	},
	{ // Jupiter
		[]float64{34.351519, 3036.3027748, .0002233, .000000037},
		[]float64{5.202603209, .0000001913},
		[]float64{.04849793, .000163225, -.0000004714, -.00000000201},
		[]float64{1.303267, -.0054965, .00000466, -.000000002},
		[]float64{100.464407, 1.0209774, .00040315, .000000404},
		[]float64{14.331207, 1.6126352, .00103042, -.000004464},
	},
	{ // Saturn
		[]float64{50.077444, 1223.5110686, .00051908, -.00000003},
		[]float64{9.554909192, -.0000021390, .000000004},
		[]float64{.05554814, -.000346641, -.0000006436, .0000000034},
		[]float64{2.488879, -.0037362, -.00001519, .000000087},
		[]float64{113.665503, .877088, -.00012176, -.000002249},
		[]float64{93.057237, 1.9637613, .00083753, .000004928},
	},
	{ // Uranus
		[]float64{314.055005, 429.8640561, .0003039, .000000026},
		[]float64{19.218446062, -.0000000372, .00000000098},
		[]float64{.04638122, -.000027293, .0000000789, .00000000024},
		[]float64{.773197, .0007744, .00003749, -.000000092},
		[]float64{74.005957, .5211278, .00133947, .000018484},
		[]float64{173.005291, 1.486379, .00021406, .000000434},
	},
	{ // Neptune
		[]float64{304.348665, 219.8833092, .00030882, .000000018},
		[]float64{30.110386869, -.0000001663, .00000000069},
		[]float64{.00945575, .000006033, 0, -.00000000005},
		[]float64{1.769953, -.0093082, -.00000708, .000000027},
		[]float64{131.784057, 1.1022039, .00025952, -.000000637},
		[]float64{48.120276, 1.4262957, .00038434, .00000002},
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
	e.Lon = unit.AngleFromDeg(base.Horner(T, c.L...)).Mod1()
	e.Axis = base.Horner(T, c.a...)
	e.Ecc = base.Horner(T, c.e...)
	e.Inc = unit.AngleFromDeg(base.Horner(T, c.i...))
	e.Node = unit.AngleFromDeg(base.Horner(T, c.Ω...))
	e.Peri = unit.AngleFromDeg(base.Horner(T, c.ϖ...))
}

// Inc returns mean inclination for a planet at a date.
//
// Result is the same as the Inc field returned by function Mean.  That is,
// referenced to mean dynamical ecliptic and equinox of date.
func Inc(p int, jde float64) unit.Angle {
	return unit.AngleFromDeg(base.Horner(base.J2000Century(jde), cMean[p].i...))
}

// Node returns mean longitude of ascending node for a planet at a date.
//
// Result is the same as the Node field returned by function Mean.  That is,
// radians, referenced to mean dynamical ecliptic and equinox of date.
func Node(p int, jde float64) unit.Angle {
	return unit.AngleFromDeg(base.Horner(base.J2000Century(jde), cMean[p].Ω...))
}
