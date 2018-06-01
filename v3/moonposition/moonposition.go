// Copyright 2013 Sonia Keys
// License: MIT

// Moonposition: Chapter 47, Position of the Moon.
package moonposition

import (
	"math"

	"github.com/soniakeys/meeus/v3/base"
	"github.com/soniakeys/unit"
)

// Parallax returns equatorial horizontal parallax of the Moon.
//
// Argument Δ is distance between centers of the Earth and Moon, in km.
func Parallax(Δ float64) unit.Angle {
	// p. 337
	return unit.Angle(math.Asin(6378.14 / Δ))
}

const p = math.Pi / 180

func dmf(T float64) (D, M, Mʹ, F float64) {
	D = base.Horner(T, 297.8501921*p, 445267.1114034*p,
		-.0018819*p, p/545868, -p/113065000)
	M = base.Horner(T, 357.5291092*p, 35999.0502909*p,
		-.0001536*p, p/24490000)
	Mʹ = base.Horner(T, 134.9633964*p, 477198.8675055*p,
		.0087414*p, p/69699, -p/14712000)
	F = base.Horner(T, 93.272095*p, 483202.0175233*p,
		-.0036539*p, -p/3526000, p/863310000)
	return
}

// Position returns geocentric location of the Moon.
//
// Results are referenced to mean equinox of date and do not include
// the effect of nutation.
//
//	λ  Geocentric longitude.
//	β  Geocentric latidude.
//	Δ  Distance between centers of the Earth and Moon, in km.
func Position(jde float64) (λ, β unit.Angle, Δ float64) {
	T := base.J2000Century(jde)
	Lʹ := base.Horner(T, 218.3164477*p, 481267.88123421*p,
		-.0015786*p, p/538841, -p/65194000)
	D, M, Mʹ, F := dmf(T)
	A1 := 119.75*p + 131.849*p*T
	A2 := 53.09*p + 479264.29*p*T
	A3 := 313.45*p + 481266.484*p*T
	E := base.Horner(T, 1, -.002516, -.0000074)
	E2 := E * E
	Σl := 3958*math.Sin(A1) + 1962*math.Sin(Lʹ-F) + 318*math.Sin(A2)
	Σr := 0.
	Σb := -2235*math.Sin(Lʹ) + 382*math.Sin(A3) + 175*math.Sin(A1-F) +
		175*math.Sin(A1+F) + 127*math.Sin(Lʹ-Mʹ) - 115*math.Sin(Lʹ+Mʹ)
	for i := range ta {
		r := &ta[i]
		sa, ca := math.Sincos(D*r.D + M*r.M + Mʹ*r.Mʹ + F*r.F)
		switch r.M {
		case 0:
			Σl += r.Σl * sa
			Σr += r.Σr * ca
		case 1, -1:
			Σl += r.Σl * sa * E
			Σr += r.Σr * ca * E
		case 2, -2:
			Σl += r.Σl * sa * E2
			Σr += r.Σr * ca * E2
		}
	}
	for i := range tb {
		r := &tb[i]
		sb := math.Sin(D*r.D + M*r.M + Mʹ*r.Mʹ + F*r.F)
		switch r.M {
		case 0:
			Σb += r.Σb * sb
		case 1, -1:
			Σb += r.Σb * sb * E
		case 2, -2:
			Σb += r.Σb * sb * E2
		}
	}
	λ = unit.Angle(Lʹ).Mod1() + unit.AngleFromDeg(Σl*1e-6)
	β = unit.AngleFromDeg(Σb * 1e-6)
	Δ = 385000.56 + Σr*1e-3
	return
}

type tas struct{ D, M, Mʹ, F, Σl, Σr float64 }

var ta = [...]tas{
	{0, 0, 1, 0, 6288774, -20905355},
	{2, 0, -1, 0, 1274027, -3699111},
	{2, 0, 0, 0, 658314, -2955968},
	{0, 0, 2, 0, 213618, -569925},

	{0, 1, 0, 0, -185116, 48888},
	{0, 0, 0, 2, -114332, -3149},
	{2, 0, -2, 0, 58793, 246158},
	{2, -1, -1, 0, 57066, -152138},

	{2, 0, 1, 0, 53322, -170733},
	{2, -1, 0, 0, 45758, -204586},
	{0, 1, -1, 0, -40923, -129620},
	{1, 0, 0, 0, -34720, 108743},

	{0, 1, 1, 0, -30383, 104755},
	{2, 0, 0, -2, 15327, 10321},
	{0, 0, 1, 2, -12528, 0},
	{0, 0, 1, -2, 10980, 79661},

	{4, 0, -1, 0, 10675, -34782},
	{0, 0, 3, 0, 10034, -23210},
	{4, 0, -2, 0, 8548, -21636},
	{2, 1, -1, 0, -7888, 24208},

	{2, 1, 0, 0, -6766, 30824},
	{1, 0, -1, 0, -5163, -8379},
	{1, 1, 0, 0, 4987, -16675},
	{2, -1, 1, 0, 4036, -12831},

	{2, 0, 2, 0, 3994, -10445},
	{4, 0, 0, 0, 3861, -11650},
	{2, 0, -3, 0, 3665, 14403},
	{0, 1, -2, 0, -2689, -7003},

	{2, 0, -1, 2, -2602, 0},
	{2, -1, -2, 0, 2390, 10056},
	{1, 0, 1, 0, -2348, 6322},
	{2, -2, 0, 0, 2236, -9884},

	{0, 1, 2, 0, -2120, 5751},
	{0, 2, 0, 0, -2069, 0},
	{2, -2, -1, 0, 2048, -4950},
	{2, 0, 1, -2, -1773, 4130},

	{2, 0, 0, 2, -1595, 0},
	{4, -1, -1, 0, 1215, -3958},
	{0, 0, 2, 2, -1110, 0},
	{3, 0, -1, 0, -892, 3258},

	{2, 1, 1, 0, -810, 2616},
	{4, -1, -2, 0, 759, -1897},
	{0, 2, -1, 0, -713, -2117},
	{2, 2, -1, 0, -700, 2354},

	{2, 1, -2, 0, 691, 0},
	{2, -1, 0, -2, 596, 0},
	{4, 0, 1, 0, 549, -1423},
	{0, 0, 4, 0, 537, -1117},

	{4, -1, 0, 0, 520, -1571},
	{1, 0, -2, 0, -487, -1739},
	{2, 1, 0, -2, -399, 0},
	{0, 0, 2, -2, -381, -4421},

	{1, 1, 1, 0, 351, 0},
	{3, 0, -2, 0, -340, 0},
	{4, 0, -3, 0, 330, 0},
	{2, -1, 2, 0, 327, 0},

	{0, 2, 1, 0, -323, 1165},
	{1, 1, -1, 0, 299, 0},
	{2, 0, 3, 0, 294, 0},
	{2, 0, -1, -2, 0, 8752},
}

type tbs struct{ D, M, Mʹ, F, Σb float64 }

var tb = [...]tbs{
	{0, 0, 0, 1, 5128122},
	{0, 0, 1, 1, 280602},
	{0, 0, 1, -1, 277693},
	{2, 0, 0, -1, 173237},

	{2, 0, -1, 1, 55413},
	{2, 0, -1, -1, 46271},
	{2, 0, 0, 1, 32573},
	{0, 0, 2, 1, 17198},

	{2, 0, 1, -1, 9266},
	{0, 0, 2, -1, 8822},
	{2, -1, 0, -1, 8216},
	{2, 0, -2, -1, 4324},

	{2, 0, 1, 1, 4200},
	{2, 1, 0, -1, -3359},
	{2, -1, -1, 1, 2463},
	{2, -1, 0, 1, 2211},

	{2, -1, -1, -1, 2065},
	{0, 1, -1, -1, -1870},
	{4, 0, -1, -1, 1828},
	{0, 1, 0, 1, -1794},

	{0, 0, 0, 3, -1749},
	{0, 1, -1, 1, -1565},
	{1, 0, 0, 1, -1491},
	{0, 1, 1, 1, -1475},

	{0, 1, 1, -1, -1410},
	{0, 1, 0, -1, -1344},
	{1, 0, 0, -1, -1335},
	{0, 0, 3, 1, 1107},

	{4, 0, 0, -1, 1021},
	{4, 0, -1, 1, 833},

	{0, 0, 1, -3, 777},
	{4, 0, -2, 1, 671},
	{2, 0, 0, -3, 607},
	{2, 0, 2, -1, 596},

	{2, -1, 1, -1, 491},
	{2, 0, -2, 1, -451},
	{0, 0, 3, -1, 439},
	{2, 0, 2, 1, 422},

	{2, 0, -3, -1, 421},
	{2, 1, -1, 1, -366},
	{2, 1, 0, 1, -351},
	{4, 0, 0, 1, 331},

	{2, -1, 1, 1, 315},
	{2, -2, 0, -1, 302},
	{0, 0, 1, 3, -283},
	{2, 1, 1, -1, -229},

	{1, 1, 0, -1, 223},
	{1, 1, 0, 1, 223},
	{0, 1, -2, -1, -220},
	{2, 1, -1, -1, -220},

	{1, 0, 1, 1, -185},
	{2, -1, -2, -1, 181},
	{0, 1, 2, 1, -177},
	{4, 0, -2, -1, 176},

	{4, -1, -1, -1, 166},
	{1, 0, 1, -1, -164},
	{4, 0, 1, -1, 132},
	{1, 0, -1, -1, -119},

	{4, -1, 0, -1, 115},
	{2, -2, 0, 1, 107},
}

// Node returns longitude of the mean ascending node of the lunar orbit.
func Node(jde float64) unit.Angle {
	return unit.AngleFromDeg(base.Horner(base.J2000Century(jde),
		125.0445479, -1934.1362891, .0020754, 1/467441, -1/60616000)).Mod1()
}

// Perigee returns longitude of perigee of the lunar orbit.
func Perigee(jde float64) unit.Angle {
	return unit.AngleFromDeg(base.Horner(base.J2000Century(jde),
		83.3532465, 4069.0137287, -.01032, -1/80053, 1/18999000)).Mod1()
}

// TrueNode returns longitude of the true ascending node.
//
// That is, the node of the instantaneous lunar orbit.
func TrueNode(jde float64) unit.Angle {
	D, M, Mʹ, F := dmf(base.J2000Century(jde))
	return Node(jde) + unit.AngleFromDeg(
		-1.4979*math.Sin(2*(D-F))+
			-.15*math.Sin(M)+
			-.1226*math.Sin(2*D)+
			.1176*math.Sin(2*F)+
			-.0801*math.Sin(2*(Mʹ-F)))
}
