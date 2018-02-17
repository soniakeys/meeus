// Copyright 2013 Sonia Keys
// License: MIT

// Eclipse: Chapter 54, Eclipses.
package eclipse

import (
	"math"

	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/meeus/moonphase"
	"github.com/soniakeys/unit"
)

func g(k, jm, c1, c2 float64) (eclipse bool, jmax, γ, u, Mʹ float64) {
	const ck = 1 / 1236.85
	const p = math.Pi / 180
	T := k * ck
	F := base.Horner(T, 160.7108*p, 390.67050284*p/ck,
		-.0016118*p, -.00000227*p, .000000011*p)
	if math.Abs(math.Sin(F)) > .36 {
		return // no eclipse
	}
	eclipse = true
	E := base.Horner(T, 1, -.002516, -.0000074)
	M := base.Horner(T, 2.5534*p, 29.1053567*p/ck,
		-.0000014*p, -.00000011*p)
	Mʹ = base.Horner(T, 201.5643*p, 385.81693528*p/ck,
		.0107582*p, .00001238*p, -.000000058*p)
	Ω := base.Horner(T, 124.7746*p, -1.56375588*p/ck,
		.0020672*p, .00000215*p)
	sΩ := math.Sin(Ω)
	F1 := F - .02665*p*sΩ
	A1 := base.Horner(T, 299.77*p, .107408*p/ck, -.009173*p)
	// (54.1) p. 380
	jmax = jm +
		c1*math.Sin(Mʹ) +
		c2*math.Sin(M)*E +
		.0161*math.Sin(2*Mʹ) +
		-.0097*math.Sin(2*F1) +
		.0073*math.Sin(Mʹ-M)*E +
		-.005*math.Sin(Mʹ+M)*E +
		-.0023*math.Sin(Mʹ-2*F1) +
		.0021*math.Sin(2*M)*E +
		.0012*math.Sin(Mʹ+2*F1) +
		.0006*math.Sin(2*Mʹ+M)*E +
		-.0004*math.Sin(3*Mʹ) +
		-.0003*math.Sin(M+2*F1)*E +
		.0003*math.Sin(A1) +
		-.0002*math.Sin(M-2*F1)*E +
		-.0002*math.Sin(2*Mʹ-M)*E +
		-.0002*sΩ
	P := .207*math.Sin(M)*E +
		.0024*math.Sin(2*M)*E +
		-.0392*math.Sin(Mʹ) +
		.0116*math.Sin(2*Mʹ) +
		-.0073*math.Sin(Mʹ+M)*E +
		.0067*math.Sin(Mʹ-M)*E +
		.0118*math.Sin(2*F1)
	Q := 5.2207 +
		-.0048*math.Cos(M)*E +
		.002*math.Cos(2*M)*E +
		-.3299*math.Cos(Mʹ) +
		-.006*math.Cos(Mʹ+M)*E +
		.0041*math.Cos(Mʹ-M)*E
	sF1, cF1 := math.Sincos(F1)
	W := math.Abs(cF1)
	γ = (P*cF1 + Q*sF1) * (1 - .0048*W)
	u = .0059 +
		.0046*math.Cos(M)*E +
		-.0182*math.Cos(Mʹ) +
		.0004*math.Cos(2*Mʹ) +
		-.0005*math.Cos(M+Mʹ)
	return
}

// Eclipse type identifiers returned from Solar and Lunar.
const (
	None         = iota
	Partial      // for solar eclipses
	Annular      // solar
	AnnularTotal // solar
	Penumbral    // for lunar eclipses
	Umbral       // lunar
	Total        // solar or lunar
)

// Snap returns k at specified quarter q nearest year y.
// Cut and paste from moonphase.  Time corresponding to k needed in these
// algorithms but otherwise not meaningful enough to export from moonphase.
func snap(y, q float64) float64 {
	k := (y - 2000) * 12.3685 // (49.2) p. 350
	return math.Floor(k-q+.5) + q
}

// Solar computes quantities related to solar eclipses.
//
// Argument year is a decimal year specifying a date.
//
// eclipseType will be None, Partial, Annular, AnnularTotal, or Total.
// If None, none of the other return values may be meaningful.
//
// central is true if the center of the eclipse shadow touches the Earth.
//
// jmax is the jde when the center of the eclipse shadow is closest to the
// Earth center, in a plane through the center of the Earth.
//
// γ is the distance from the eclipse shadow center to the Earth center
// at time jmax.
//
// u is the radius of the Moon's umbral cone in the plane of the Earth.
//
// p is the radius of the penumbral cone.
//
// mag is eclipse magnitude for partial eclipses.  It is not valid for other
// eclipse types.
//
// γ, u, and p are in units of equatorial Earth radii.
func Solar(year float64) (eclipseType int, central bool, jmax, γ, u, p, mag float64) {
	var e bool
	e, jmax, γ, u, _ = g(snap(year, 0), moonphase.MeanNew(year), -.4075, .1721)
	p = u + .5461
	if !e {
		return // no eclipse
	}
	aγ := math.Abs(γ)
	if aγ > 1.5433+u {
		return // no eclipse
	}
	central = aγ < .9972 // eclipse center touches Earth
	switch {
	case !central:
		eclipseType = Partial // most common case
		if aγ < 1.026 {       // umbral cone may touch earth
			if aγ < .9972+math.Abs(u) { // total or annular
				eclipseType = Total // report total in both cases
			}
		}
	case u < 0:
		eclipseType = Total
	case u > .0047:
		eclipseType = Annular
	default:
		ω := .00464 * math.Sqrt(1-γ*γ)
		if u < ω {
			eclipseType = AnnularTotal
		} else {
			eclipseType = Annular
		}
	}
	if eclipseType == Partial {
		// (54.2) p. 382
		mag = (1.5433 + u - aγ) / (.5461 + 2*u)
	}
	return
}

// Lunar computes quantities related to lunar eclipses.
//
// Argument year is a decimal year specifying a date.
//
// eclipseType will be None, Penumbral, Umbral, or Total.
// If None, none of the other return values may be meaningful.
//
// jmax is the jde when the center of the eclipse shadow is closest to the
// Moon center, in a plane through the center of the Moon.
//
// γ is the distance from the eclipse shadow center to the moon center
// at time jmax.
//
// σ is the radius of the umbral cone in the plane of the Moon.
//
// ρ is the radius of the penumbral cone.
//
// mag is eclipse magnitude.
//
// sd- return values are semidurations of the phases of the eclipse.
//
// γ, σ, and ρ are in units of equatorial Earth radii.
func Lunar(year float64) (eclipseType int, jmax, γ, ρ, σ, mag float64, sdTotal, sdPartial, sdPenumbral unit.Time) {
	var e bool
	var u, Mʹ float64
	e, jmax, γ, u, Mʹ = g(snap(year, .5),
		moonphase.MeanFull(year), -.4065, .1727)
	if !e {
		return // no eclipse
	}
	ρ = 1.2848 + u
	σ = .7403 - u
	aγ := math.Abs(γ)
	mag = (1.0128 - u - aγ) / .545 // (54.3) p. 382
	switch {
	case mag > 1:
		eclipseType = Total
	case mag > 0:
		eclipseType = Umbral
	default:
		mag = (1.5573 + u - aγ) / .545 // (54.4) p. 382
		if mag < 0 {
			return // no eclipse
		}
		eclipseType = Penumbral
	}
	p := 1.0128 - u
	t := .4678 - u
	n := .5458 + .04*math.Cos(Mʹ)
	γ2 := γ * γ
	switch eclipseType {
	case Total:
		sdTotal = unit.TimeFromHour(math.Sqrt(t*t-γ2) / n)
		fallthrough
	case Umbral:
		sdPartial = unit.TimeFromHour(math.Sqrt(p*p-γ2) / n)
		fallthrough
	default:
		h := 1.5573 + u
		sdPenumbral = unit.TimeFromHour(math.Sqrt(h*h-γ2) / n)
	}
	return
}
