// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Apparent: Chapter 23, Apparent Place of a Star
package apparent

import (
	"math"

	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/meeus/nutation"
	"github.com/soniakeys/meeus/solar"
)

func Nutation(α, δ, jd float64) (Δα1, Δδ1 float64) {
	ε := nutation.MeanObliquity(jd)
	sε, cε := math.Sincos(ε)
	Δψ, Δε := nutation.Nutation(jd)
	sα, cα := math.Sincos(α)
	tδ := math.Tan(δ)
	Δα1 = (cε+sε*sα*tδ)*Δψ - cα*tδ*Δε
	Δδ1 = sε*cα*Δψ + sα*Δε
	return
}

// κ is the constnt of abberation in radians.
const κ = 20.49552 * math.Pi / 180 / 3600

func perihelion(T float64) float64 {
	return base.Horner(T, 102.93735, 1.71946, .00046) * math.Pi / 180
}

func EclipticAbberation(λ, β, jd float64) (Δλ, Δβ float64) {
	T := base.J2000Century(jd)
	s, _ := solar.True(T)
	e := solar.Eccentricity(T)
	π := perihelion(T)
	sβ, cβ := math.Sincos(β)
	ssλ, csλ := math.Sincos(s - λ)
	sπλ, cπλ := math.Sincos(π - λ)
	Δλ = κ * (e*cπλ - csλ) / cβ
	Δβ = -κ * sβ * (ssλ - e*sπλ)
	return
}

func EquatorialAbberation(α, δ, jd float64) (Δα2, Δδ2 float64) {
	ε := nutation.MeanObliquity(jd)
	T := base.J2000Century(jd)
	s, _ := solar.True(T)
	e := solar.Eccentricity(T)
	π := perihelion(T)
	sα, cα := math.Sincos(α)
	sδ, cδ := math.Sincos(δ)
	ss, cs := math.Sincos(s)
	sπ, cπ := math.Sincos(π)
	cε := math.Cos(ε)
	tε := math.Tan(ε)
	q1 := cα * cε
	Δα2 = κ * (e*(q1*cπ+sα*sπ) - (q1*cs + sα*ss)) / cδ
	q2 := cε * (tε*cδ - sα*sδ)
	q3 := cα * sδ
	Δδ2 = κ * (e*(cπ*q2+sπ*q3) - (cs*q2 + ss*q3))
	return
}
