// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Apparent: Chapter 23, Apparent Place of a Star
package apparent

import (
	"math"

	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/meeus/coord"
	"github.com/soniakeys/meeus/nutation"
	"github.com/soniakeys/meeus/precess"
	"github.com/soniakeys/meeus/solar"
	"github.com/soniakeys/unit"
)

// Nutation returns corrections due to nutation for equatorial coordinates
// of an object.
//
// Results are invalid for objects very near the celestial poles.
func Nutation(α unit.RA, δ unit.Angle, jd float64) (Δα1 unit.HourAngle, Δδ1 unit.Angle) {
	ε := nutation.MeanObliquity(jd)
	sε, cε := ε.Sincos()
	Δψ, Δε := nutation.Nutation(jd)
	sα, cα := α.Sincos()
	tδ := δ.Tan()
	// (23.1) p. 151
	Δα1 = unit.HourAngle((cε+sε*sα*tδ)*Δψ.Rad() - cα*tδ*Δε.Rad())
	Δδ1 = Δψ.Mul(sε*cα) + Δε.Mul(sα)
	return
}

// κ is the constnt of aberration in radians.
var κ = unit.AngleFromSec(20.49552)

// longitude of perihelian of Earth's orbit.
func perihelion(T float64) unit.Angle {
	return unit.AngleFromDeg(base.Horner(T, 102.93735, 1.71946, .00046))
}

// EclipticAberration returns corrections due to aberration for ecliptic
// coordinates of an object.
func EclipticAberration(λ, β unit.Angle, jd float64) (Δλ, Δβ unit.Angle) {
	T := base.J2000Century(jd)
	s, _ := solar.True(T)
	e := solar.Eccentricity(T)
	π := perihelion(T)
	sβ, cβ := β.Sincos()
	ssλ, csλ := (s - λ).Sincos()
	sπλ, cπλ := (π - λ).Sincos()
	// (23.2) p. 151
	Δλ = κ.Mul((e*cπλ - csλ) / cβ)
	Δβ = -κ.Mul(sβ * (ssλ - e*sπλ))
	return
}

// Aberration returns corrections due to aberration for equatorial
// coordinates of an object.
func Aberration(α unit.RA, δ unit.Angle, jd float64) (Δα2 unit.HourAngle, Δδ2 unit.Angle) {
	ε := nutation.MeanObliquity(jd)
	T := base.J2000Century(jd)
	s, _ := solar.True(T)
	e := solar.Eccentricity(T)
	π := perihelion(T)
	sα, cα := α.Sincos()
	sδ, cδ := δ.Sincos()
	ss, cs := s.Sincos()
	sπ, cπ := π.Sincos()
	cε := ε.Cos()
	tε := ε.Tan()
	q1 := cα * cε
	// (23.3) p. 152
	Δα2 = unit.HourAngle(κ.Rad() * (e*(q1*cπ+sα*sπ) - (q1*cs + sα*ss)) / cδ)
	q2 := cε * (tε*cδ - sα*sδ)
	q3 := cα * sδ
	Δδ2 = κ.Mul(e*(cπ*q2+sπ*q3) - (cs*q2 + ss*q3))
	return
}

// Position computes the apparent position of an object.
//
// Position is computed for equatorial coordinates in eqFrom, considering
// proper motion, precession, nutation, and aberration.  Result is in
// eqTo.  EqFrom and eqTo must be non-nil, but may point to the same struct.
func Position(eqFrom, eqTo *coord.Equatorial, epochFrom, epochTo float64, mα unit.HourAngle, mδ unit.Angle) *coord.Equatorial {
	precess.Position(eqFrom, eqTo, epochFrom, epochTo, mα, mδ)
	jd := base.JulianYearToJDE(epochTo)
	Δα1, Δδ1 := Nutation(eqTo.RA, eqTo.Dec, jd)
	Δα2, Δδ2 := Aberration(eqTo.RA, eqTo.Dec, jd)
	eqTo.RA = eqTo.RA.Add(Δα1 + Δα2)
	eqTo.Dec += Δδ1 + Δδ2
	return eqTo
}

// AberrationRonVondrak uses the Ron-Vondrák expression to compute corrections
// due to aberration for equatorial coordinates of an object.
func AberrationRonVondrak(α unit.RA, δ unit.Angle, jd float64) (Δα unit.HourAngle, Δδ unit.Angle) {
	T := base.J2000Century(jd)
	r := &rv{
		T:  T,
		L2: 3.1761467 + 1021.3285546*T,
		L3: 1.7534703 + 628.3075849*T,
		L4: 6.2034809 + 334.0612431*T,
		L5: 0.5995465 + 52.9690965*T,
		L6: 0.8740168 + 21.3299095*T,
		L7: 5.4812939 + 7.4781599*T,
		L8: 5.3118863 + 3.8133036*T,
		Lp: 3.8103444 + 8399.6847337*T,
		D:  5.1984667 + 7771.3771486*T,
		Mp: 2.3555559 + 8328.6914289*T,
		F:  1.6279052 + 8433.4661601*T,
	}
	var Xp, Yp, Zp float64
	// sum smaller terms first
	for i := 35; i >= 0; i-- {
		x, y, z := rvTerm[i](r)
		Xp += x
		Yp += y
		Zp += z
	}
	sα, cα := α.Sincos()
	sδ, cδ := δ.Sincos()
	// (23.4) p. 156
	Δα = unit.HourAngle((Yp*cα - Xp*sα) / (c * cδ))
	Δδ = unit.Angle(-((Xp*cα+Yp*sα)*sδ - Zp*cδ) / c)
	return
}

const c = 17314463350 // unit is 1e-8 AU / day

type rv struct {
	T, L2, L3, L4, L5, L6, L7, L8, Lp, D, Mp, F float64
}

type rvFunc func(*rv) (x, y, z float64)

var rvTerm = [36]rvFunc{
	func(r *rv) (x, y, z float64) { // 1
		sA, cA := math.Sincos(r.L3)
		return (-1719914-2*r.T)*sA - 25*cA,
			(25-13*r.T)*sA + (1578089+156*r.T)*cA,
			(10+32*r.T)*sA + (684185-358*r.T)*cA
	},
	func(r *rv) (x, y, z float64) { // 2
		sA, cA := math.Sincos(2 * r.L3)
		return (6434+141*r.T)*sA + (28007-107*r.T)*cA,
			(25697-95*r.T)*sA + (-5904-130*r.T)*cA,
			(11141-48*r.T)*sA + (-2559-55*r.T)*cA
	},
	func(r *rv) (x, y, z float64) { // 3
		sA, cA := math.Sincos(r.L5)
		return 715 * sA, 6*sA - 657*cA, -15*sA - 282*cA
	},
	func(r *rv) (x, y, z float64) { // 4
		sA, cA := math.Sincos(r.Lp)
		return 715 * sA, -656 * cA, -285 * cA
	},
	func(r *rv) (x, y, z float64) { // 5
		sA, cA := math.Sincos(3 * r.L3)
		return (486-5*r.T)*sA + (-236-4*r.T)*cA,
			(-216-4*r.T)*sA + (-446+5*r.T)*cA,
			-94*sA - 193*cA
	},
	func(r *rv) (x, y, z float64) { // 6
		sA, cA := math.Sincos(r.L6)
		return 159 * sA, 2*sA - 147*cA, -6*sA - 61*cA
	},
	func(r *rv) (x, y, z float64) { // 7
		cA := math.Cos(r.F)
		return 0, 26 * cA, -59 * cA
	},
	func(r *rv) (x, y, z float64) { // 8
		sA, cA := math.Sincos(r.Lp + r.Mp)
		return 39 * sA, -36 * cA, -16 * cA
	},
	func(r *rv) (x, y, z float64) { // 9
		sA, cA := math.Sincos(2 * r.L5)
		return 33*sA - 10*cA, -9*sA - 30*cA, -5*sA - 13*cA
	},
	func(r *rv) (x, y, z float64) { // 10
		sA, cA := math.Sincos(2*r.L3 - r.L5)
		return 31*sA + cA, sA - 28*cA, -12 * cA
	},
	func(r *rv) (x, y, z float64) { // 11
		sA, cA := math.Sincos(3*r.L3 - 8*r.L4 + 3*r.L5)
		return 8*sA - 28*cA, 25*sA + 8*cA, 11*sA + 3*cA
	},
	func(r *rv) (x, y, z float64) { // 12
		sA, cA := math.Sincos(5*r.L3 - 8*r.L4 + 3*r.L5)
		return 8*sA - 28*cA, -25*sA - 8*cA, -11*sA + -3*cA
	},
	func(r *rv) (x, y, z float64) { // 13
		sA, cA := math.Sincos(2*r.L2 - r.L3)
		return 21 * sA, -19 * cA, -8 * cA
	},
	func(r *rv) (x, y, z float64) { // 14
		sA, cA := math.Sincos(r.L2)
		return -19 * sA, 17 * cA, 8 * cA
	},
	func(r *rv) (x, y, z float64) { // 15
		sA, cA := math.Sincos(r.L7)
		return 17 * sA, -16 * cA, -7 * cA
	},
	func(r *rv) (x, y, z float64) { // 16
		sA, cA := math.Sincos(r.L3 - 2*r.L5)
		return 16 * sA, 15 * cA, sA + 7*cA
	},
	func(r *rv) (x, y, z float64) { // 17
		sA, cA := math.Sincos(r.L8)
		return 16 * sA, sA - 15*cA, -3*sA - 6*cA
	},
	func(r *rv) (x, y, z float64) { // 18
		sA, cA := math.Sincos(r.L3 + r.L5)
		return 11*sA - cA, -sA - 10*cA, -sA - 5*cA
	},
	func(r *rv) (x, y, z float64) { // 19
		sA, cA := math.Sincos(2*r.L2 - 2*r.L3)
		return -11 * cA, -10 * sA, -4 * sA
	},
	func(r *rv) (x, y, z float64) { // 20
		sA, cA := math.Sincos(r.L3 - r.L5)
		return -11*sA - 2*cA, -2*sA + 9*cA, -sA + 4*cA
	},
	func(r *rv) (x, y, z float64) { // 21
		sA, cA := math.Sincos(4 * r.L3)
		return -7*sA - 8*cA, -8*sA + 6*cA, -3*sA + 3*cA
	},
	func(r *rv) (x, y, z float64) { // 22
		sA, cA := math.Sincos(3*r.L3 - 2*r.L5)
		return -10 * sA, 9 * cA, 4 * cA
	},
	func(r *rv) (x, y, z float64) { // 23
		sA, cA := math.Sincos(r.L2 - 2*r.L3)
		return -9 * sA, -9 * cA, -4 * cA
	},
	func(r *rv) (x, y, z float64) { // 24
		sA, cA := math.Sincos(2*r.L2 - 3*r.L3)
		return -9 * sA, -8 * cA, -4 * cA
	},
	func(r *rv) (x, y, z float64) { // 25
		sA, cA := math.Sincos(2 * r.L6)
		return -9 * cA, -8 * sA, -3 * sA
	},
	func(r *rv) (x, y, z float64) { // 26
		sA, cA := math.Sincos(2*r.L2 - 4*r.L3)
		return -9 * cA, 8 * sA, 3 * sA
	},
	func(r *rv) (x, y, z float64) { // 27
		sA, cA := math.Sincos(3*r.L3 - 2*r.L4)
		return 8 * sA, -8 * cA, -3 * cA
	},
	func(r *rv) (x, y, z float64) { // 28
		sA, cA := math.Sincos(r.Lp + 2*r.D - r.Mp)
		return 8 * sA, -7 * cA, -3 * cA
	},
	func(r *rv) (x, y, z float64) { // 29
		sA, cA := math.Sincos(8*r.L2 - 12*r.L3)
		return -4*sA - 7*cA, -6*sA + 4*cA, -3*sA + 2*cA
	},
	func(r *rv) (x, y, z float64) { // 30
		sA, cA := math.Sincos(8*r.L2 - 14*r.L3)
		return -4*sA - 7*cA, 6*sA - 4*cA, 3*sA - 2*cA
	},
	func(r *rv) (x, y, z float64) { // 31
		sA, cA := math.Sincos(2 * r.L4)
		return -6*sA - 5*cA, -4*sA + 5*cA, -2*sA + 2*cA
	},
	func(r *rv) (x, y, z float64) { // 32
		sA, cA := math.Sincos(3*r.L2 - 4*r.L3)
		return -sA - cA, -2*sA - 7*cA, sA - 4*cA
	},
	func(r *rv) (x, y, z float64) { // 33
		sA, cA := math.Sincos(2*r.L3 - 2*r.L5)
		return 4*sA - 6*cA, -5*sA - 4*cA, -2*sA - 2*cA
	},
	func(r *rv) (x, y, z float64) { // 34
		sA, cA := math.Sincos(3*r.L2 - 3*r.L3)
		return -7 * cA, -6 * sA, -3 * sA
	},
	func(r *rv) (x, y, z float64) { // 35
		sA, cA := math.Sincos(2*r.L3 - 2*r.L4)
		return 5*sA - 5*cA, -4*sA - 5*cA, -2*sA - 2*cA
	},
	func(r *rv) (x, y, z float64) { // 36
		sA, cA := math.Sincos(r.Lp - 2*r.D)
		return 5 * sA, -5 * cA, -2 * cA
	},
}

// PositionRonVondrak computes the apparent position of an object using
// the Ron-Vondrák expression for aberration.
//
// Position is computed for equatorial coordinates in eqFrom, considering
// proper motion, aberration, precession, and nutation.  Result is in
// eqTo.  EqFrom and eqTo must be non-nil, but may point to the same struct.
//
// Note the Ron-Vondrák expression is only valid for the epoch J2000.
// EqFrom must be coordinates at epoch J2000.
func PositionRonVondrak(eqFrom, eqTo *coord.Equatorial, epochTo float64, mα unit.HourAngle, mδ unit.Angle) *coord.Equatorial {
	t := epochTo - 2000
	eqTo.RA = eqFrom.RA.Add(mα.Mul(t))
	eqTo.Dec = eqFrom.Dec + mδ.Mul(t)
	jd := base.JulianYearToJDE(epochTo)
	Δα, Δδ := AberrationRonVondrak(eqTo.RA, eqTo.Dec, jd)
	eqTo.RA = eqTo.RA.Add(Δα)
	eqTo.Dec += Δδ
	precess.Position(eqTo, eqTo, 2000, epochTo, 0, 0)
	Δα1, Δδ1 := Nutation(eqTo.RA, eqTo.Dec, jd)
	eqTo.RA = eqTo.RA.Add(Δα1)
	eqTo.Dec += Δδ1
	return eqTo
}
