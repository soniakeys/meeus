// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Precession: Chapter 21, Precession.
//
// Functions in this package take Julian epoch argurments rather than Julian
// days.  Use base.JDEToJulianYear() to convert.
//
// Also in package base are some definitions related to the Besselian and
// Julian Year.
//
// Partial:  Precession from FK4 not implemented.  Meeus gives no test cases.
// It's a fair amount of code and data, representing significant chances for
// errors.  And precession from FK4 would seem to be of little interest today.
//
// Proper motion units
//
// Meeus gives some example annual proper motions in units of seconds of
// right ascension and seconds of declination.  To make units clear,
// functions in this package take proper motions with argument types of
// unit.HourAngle and unit.Angle respectively.  Error-prone conversions
// can be avoided by using the constructors for these unit types.
//
// For example, given an annual proper motion in right ascension of -0ˢ.03847,
// rather than
//
//	mra := -0.03847 / 13751 // as Meeus suggests on p. 141
//
// or
//
//	mra := -0.03847 * (15/3600) * (pi/180) // less magic
//
// use
//
//	mra := unit.HourAngleFromSec(-0.03847)
//
// Unless otherwise indicated, functions in this library expect proper motions
// to be annual proper motions, so the unit denominator is years.
// (The code, following Meeus's example, technically treats it as Julian years.)
package precess

import (
	"math"

	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/meeus/coord"
	"github.com/soniakeys/meeus/elementequinox"
	"github.com/soniakeys/meeus/nutation"
	"github.com/soniakeys/unit"
)

// ApproxAnnualPrecession returns approximate annual precision in right
// ascension and declination.
//
// The two epochs should be within a few hundred years.
// The declinations should not be too close to the poles.
func ApproxAnnualPrecession(eq *coord.Equatorial, epochFrom, epochTo float64) (Δα unit.HourAngle, Δδ unit.Angle) {
	m, nα, nδ := mn(epochFrom, epochTo)
	sα, cα := eq.RA.Sincos()
	// (21.1) p. 132
	Δα = m + nα.Mul(sα*eq.Dec.Tan())
	Δδ = nδ.Mul(cα)
	return
}

// mn as separate function for testing purposes
func mn(epochFrom, epochTo float64) (m, nα unit.HourAngle, nδ unit.Angle) {
	T := (epochTo - epochFrom) * .01
	m = unit.HourAngleFromSec(3.07496 + 0.00186*T)
	nα = unit.HourAngleFromSec(1.33621 - 0.00057*T)
	nδ = unit.AngleFromSec(20.0431 - 0.0085*T)
	return
}

// ApproxPosition uses ApproxAnnualPrecession to compute a simple and quick
// precession while still considering proper motion.
//
// Both eqFrom and eqTo must be non-nil, although they may point to the same
// struct.  EqTo is returned for convenience.
func ApproxPosition(eqFrom, eqTo *coord.Equatorial, epochFrom, epochTo float64, mα unit.HourAngle, mδ unit.Angle) *coord.Equatorial {
	Δα, Δδ := ApproxAnnualPrecession(eqFrom, epochFrom, epochTo)
	dy := epochTo - epochFrom
	eqTo.RA = eqFrom.RA.Add((Δα + mα).Mul(dy))
	eqTo.Dec = eqFrom.Dec + (Δδ + mδ).Mul(dy)
	return eqTo
}

// Precessor represents precession from one epoch to another.
//
// Construct with NewPrecessor, then call method Precess.
// After construction, Precess may be called multiple times to precess
// different coordinates with the same initial and final epochs.
type Precessor struct {
	ζ      unit.RA
	z      unit.Angle
	sθ, cθ float64
}

const d = math.Pi / 180
const s = d / 3600

// Package variables allow these slices to be reused.  (As composite
// literals inside of NewPrecessor they would be reallocated on every
// function call.)
var (
	// coefficients from (21.2) p. 134
	ζT = []float64{2306.2181 * s, 1.39656 * s, -0.000139 * s}
	zT = []float64{2306.2181 * s, 1.39656 * s, -0.000139 * s}
	θT = []float64{2004.3109 * s, -0.8533 * s, -0.000217 * s}

	// coefficients from (21.3) p. 134
	ζt = []float64{2306.2181 * s, 0.30188 * s, 0.017998 * s}
	zt = []float64{2306.2181 * s, 1.09468 * s, 0.018203 * s}
	θt = []float64{2004.3109 * s, -0.42665 * s, -0.041833 * s}
)

// NewPrecessor constructs a Precessor object and initializes it to precess
// coordinates from epochFrom to epochTo.
func NewPrecessor(epochFrom, epochTo float64) *Precessor {
	// (21.2) p. 134
	ζCoeff := ζt
	zCoeff := zt
	θCoeff := θt
	if epochFrom != 2000 {
		T := (epochFrom - 2000) * .01
		ζCoeff = []float64{
			base.Horner(T, ζT...),
			0.30188*s - 0.000344*s*T,
			0.017998 * s}
		zCoeff = []float64{
			base.Horner(T, zT...),
			1.09468*s + 0.000066*s*T,
			0.018203 * s}
		θCoeff = []float64{
			base.Horner(T, θT...),
			-0.42665*s - 0.000217*s*T,
			-0.041833 * s}
	}
	t := (epochTo - epochFrom) * .01
	p := &Precessor{
		ζ: unit.RA(base.Horner(t, ζCoeff...) * t),
		z: unit.Angle(base.Horner(t, zCoeff...) * t),
	}
	θ := base.Horner(t, θCoeff...) * t
	p.sθ, p.cθ = math.Sincos(θ)
	return p
}

// Precess precesses coordinates eqFrom, leaving result in eqTo.
//
// The same struct may be used for eqFrom and eqTo.
// EqTo is returned for convenience.
func (p *Precessor) Precess(eqFrom, eqTo *coord.Equatorial) *coord.Equatorial {
	// (21.4) p. 134
	sδ, cδ := eqFrom.Dec.Sincos()
	sαζ, cαζ := (eqFrom.RA + p.ζ).Sincos()
	A := cδ * sαζ
	B := p.cθ*cδ*cαζ - p.sθ*sδ
	C := p.sθ*cδ*cαζ + p.cθ*sδ
	eqTo.RA = unit.RAFromRad(math.Atan2(A, B) + p.z.Rad())
	if C < base.CosSmallAngle {
		eqTo.Dec = unit.Angle(math.Asin(C))
	} else {
		eqTo.Dec = unit.Angle(math.Acos(math.Hypot(A, B))) // near pole
	}
	return eqTo
}

// Position precesses equatorial coordinates from one epoch to another,
// including proper motions.
//
// If proper motions are not to be considered or are not applicable, pass 0, 0
// for mα, mδ
//
// Both eqFrom and eqTo must be non-nil, although they may point to the same
// struct.  EqTo is returned for convenience.
func Position(eqFrom, eqTo *coord.Equatorial, epochFrom, epochTo float64, mα unit.HourAngle, mδ unit.Angle) *coord.Equatorial {
	p := NewPrecessor(epochFrom, epochTo)
	t := epochTo - epochFrom
	eqTo.RA = unit.RAFromRad(eqFrom.RA.Rad() + mα.Rad()*t)
	eqTo.Dec = eqFrom.Dec + mδ*unit.Angle(t)
	return p.Precess(eqTo, eqTo)
}

// EclipticPrecessor represents precession from one epoch to another.
//
// Construct with NewEclipticPrecessor, then call method Precess.
// After construction, Precess may be called multiple times to precess
// different coordinates with the same initial and final epochs.
type EclipticPrecessor struct {
	sη, cη float64
	π, p   unit.Angle
}

var (
	// coefficients from (21.5) p. 136, scaled to radians
	ηT = []float64{47.0029 * s, -0.06603 * s, 0.000598 * s}
	πT = []float64{174.876384 * d, 3289.4789 * s, 0.60622 * s}
	pT = []float64{5029.0966 * s, 2.22226 * s, -0.000042 * s}

	// coefficients from (21.6) p. 136, scaled to radians
	ηt = []float64{47.0029 * s, -0.03302 * s, 0.000060 * s}
	πt = []float64{174.876384 * d, -869.8089 * s, 0.03536 * s}
	pt = []float64{5029.0966 * s, 1.11113 * s, -0.000006 * s}
)

// NewEclipticPrecessor constructs an EclipticPrecessor object and initializes
// it to precess coordinates from epochFrom to epochTo.
func NewEclipticPrecessor(epochFrom, epochTo float64) *EclipticPrecessor {
	// (21.5) p. 136
	ηCoeff := ηt
	πCoeff := πt
	pCoeff := pt
	if epochFrom != 2000 {
		T := (epochFrom - 2000) * .01
		ηCoeff = []float64{
			base.Horner(T, ηT...),
			-0.03302*s + 0.000598*s*T,
			0.000060 * s}
		πCoeff = []float64{
			base.Horner(T, πT...),
			-869.8089*s - 0.50491*s*T,
			0.03536 * s}
		pCoeff = []float64{
			base.Horner(T, pT...),
			1.11113*s - 0.000042*s*T,
			-0.000006 * s}
	}
	t := (epochTo - epochFrom) * .01
	p := &EclipticPrecessor{
		π: unit.Angle(base.Horner(t, πCoeff...)),
		p: unit.Angle(base.Horner(t, pCoeff...) * t),
	}
	η := unit.Angle(base.Horner(t, ηCoeff...) * t)
	p.sη, p.cη = η.Sincos()
	return p
}

// EclipticPrecess precesses coordinates eclFrom, leaving result in eclTo.
//
// The same struct may be used for eclFrom and eclTo.
// EclTo is returned for convenience.
func (p *EclipticPrecessor) Precess(eclFrom, eclTo *coord.Ecliptic) *coord.Ecliptic {
	// (21.7) p. 137
	sβ, cβ := eclFrom.Lat.Sincos()
	sd, cd := (p.π - eclFrom.Lon).Sincos()
	A := p.cη*cβ*sd - p.sη*sβ
	B := cβ * cd
	C := p.cη*sβ + p.sη*cβ*sd
	eclTo.Lon = p.p + p.π - unit.Angle(math.Atan2(A, B))
	if C < base.CosSmallAngle {
		eclTo.Lat = unit.Angle(math.Asin(C))
	} else {
		eclTo.Lat = unit.Angle(math.Acos(math.Hypot(A, B))) // near pole
	}
	return eclTo
}

// ReduceElements reduces orbital elements of a solar system body from one
// equinox to another.
//
// This function is described in chapter 24, but is located in this
// package so it can be a method of EclipticPrecessor.
func (p *EclipticPrecessor) ReduceElements(eFrom, eTo *elementequinox.Elements) *elementequinox.Elements {
	ψ := p.π + p.p
	si, ci := eFrom.Inc.Sincos()
	snp, cnp := (eFrom.Node - p.π).Sincos()
	// (24.1) p. 159
	eTo.Inc = unit.Angle(math.Acos(ci*p.cη + si*p.sη*cnp))
	// (24.2) p. 159
	eTo.Node = ψ +
		unit.Angle(math.Atan2(si*snp, p.cη*si*cnp-p.sη*ci))
	// (24.3) p. 160
	eTo.Peri = eFrom.Peri +
		unit.Angle(math.Atan2(-p.sη*snp, si*p.cη-ci*p.sη*cnp))
	return eTo
}

// EclipticPosition precesses ecliptic coordinates from one epoch to another,
// including proper motions.
//
// While eclFrom is given as ecliptic coordinates, proper motions mα, mδ are
// still expected to be equatorial.  If proper motions are not to be considered
// or are not applicable, pass 0, 0.
//
// Both eclFrom and eclTo must be non-nil, although they may point to the same
// struct.  EclTo is returned for convenience.
func EclipticPosition(eclFrom, eclTo *coord.Ecliptic, epochFrom, epochTo float64, mα unit.HourAngle, mδ unit.Angle) *coord.Ecliptic {
	p := NewEclipticPrecessor(epochFrom, epochTo)
	*eclTo = *eclFrom
	if mα != 0 || mδ != 0 {
		mλ, mβ := eqProperMotionToEcl(mα, mδ, epochFrom, eclFrom)
		t := epochTo - epochFrom
		eclTo.Lon += mλ.Mul(t)
		eclTo.Lat += mβ.Mul(t)
	}
	return p.Precess(eclTo, eclTo)
}

func eqProperMotionToEcl(mα unit.HourAngle, mδ unit.Angle, epoch float64, pos *coord.Ecliptic) (mλ, mβ unit.Angle) {
	ε := nutation.MeanObliquity(base.JulianYearToJDE(epoch))
	sε, cε := ε.Sincos()
	α, δ := coord.EclToEq(pos.Lon, pos.Lat, sε, cε)
	sα, cα := α.Sincos()
	sδ, cδ := δ.Sincos()
	cβ := pos.Lat.Cos()
	mλ = (mδ.Mul(sε*cα) + unit.Angle(mα).Mul(cδ*(cε*cδ+sε*sδ*sα))).Div(cβ * cβ)
	mβ = (mδ.Mul(cε*cδ+sε*sδ*sα) - unit.Angle(mα).Mul(sε*cα*cδ)).Div(cβ)
	return
}

// ProperMotion3D takes the 3D equatorial coordinates of an object
// at one epoch and computes its coordinates at a new epoch, considering
// proper motion and radial velocity.
//
// Radial distance (r) must be in parsecs, radial velocitiy (mr) in
// parsecs per year.
//
// Both eqFrom and eqTo must be non-nil, although they may point to the same
// struct.  EqTo is returned for convenience.
func ProperMotion3D(eqFrom, eqTo *coord.Equatorial, epochFrom, epochTo, r, mr float64, mα unit.HourAngle, mδ unit.Angle) *coord.Equatorial {
	sα, cα := eqFrom.RA.Sincos()
	sδ, cδ := eqFrom.Dec.Sincos()
	x := r * cδ * cα
	y := r * cδ * sα
	z := r * sδ
	mrr := mr / r
	zmδ := z * mδ.Rad()
	mx := x*mrr - zmδ*cα - y*mα.Rad()
	my := y*mrr - zmδ*sα + x*mα.Rad()
	mz := z*mrr + r*mδ.Rad()*cδ
	t := epochTo - epochFrom
	xp := x + t*mx
	yp := y + t*my
	zp := z + t*mz
	eqTo.RA = unit.RAFromRad(math.Atan2(yp, xp))
	eqTo.Dec = unit.Angle(math.Atan2(zp, math.Hypot(xp, yp)))
	return eqTo
}
