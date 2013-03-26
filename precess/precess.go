// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Precession: Chapter 21, Precession.
//
// Functions in this package take Julian epoch argurments rather than Julian
// days.  Use base.JDEToJulianYear() to convert.
//
// Also in package base are some definitions related to the Besselian and
// Julian Year.
package precess

import (
	"math"

	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/meeus/coord"
	"github.com/soniakeys/meeus/nutation"
)

// ApproxAnnualPrecession returns approximate annual precision in right
// ascension and declination.
//
// The two epochs should be within a few hundred years.
// The declinations should not be too close to the poles.
//
// Δα, Δδ both returned in radians of arc.
func ApproxAnnualPrecession(eq *coord.Equatorial, epochFrom, epochTo float64) (Δα, Δδ float64) {
	m, na, nd := mn(epochFrom, epochTo)
	sa, ca := math.Sincos(eq.RA)
	Δαs := m + na*sa*math.Tan(eq.Dec) // seconds of time
	Δδs := nd * ca                    // seconds of arc
	neg := false
	if Δδs < 0 {
		neg = true
		Δδs = -Δδs
	}
	return base.Time(Δαs).Rad(), base.NewAngle(neg, 0, 0, Δδs).Rad()
}

// mn as separate function for testing purposes
func mn(epochFrom, epochTo float64) (m, na, nd float64) {
	T := (epochTo - epochFrom) * .01
	m = 3.07496 + 0.00186*T
	na = 1.33621 - 0.00057*T
	nd = 20.0431 - 0.0085*T
	return
}

// PrecessApprox uses ApproxAnnualPrecession to perform a simple and quick
// precession while still considering proper motion.
//
// Proper motions mα and mδ must be in arc radians per year.
//
// Both eqFrom and eqTo must be non-nil, although they may point to the same
// struct.  EqTo is returned for convenience.
func PrecessApprox(eqFrom, eqTo *coord.Equatorial, epochFrom, epochTo, mα, mδ float64) *coord.Equatorial {
	Δα, Δδ := ApproxAnnualPrecession(eqFrom, epochFrom, epochTo)
	dy := epochTo - epochFrom
	eqTo.RA = eqFrom.RA + (Δα+mα)*dy
	eqTo.Dec = eqFrom.Dec + (Δδ+mδ)*dy
	return eqTo
}

var ζt = []float64{0, 2306.2181, 0.30188, 0.017998}
var zt = []float64{0, 2306.2181, 1.09468, 0.018203}
var θt = []float64{0, 2004.3109, -0.42665, -0.041833}

var ζT = []float64{2306.2181, 1.39656, -0.000139}
var zT = []float64{2306.2181, 1.39656, -0.000139}
var θT = []float64{2004.3109, -0.8533, -0.000217}

// Precess precesses equatorial coordinates from one epoch to another,
// including proper motions.
//
// Proper motions must be in arc radians per year.  Pass 0 if proper motions
// are not to be considered or are not applicable.
//
// Both eqFrom and eqTo must be non-nil, although they may point to the same
// struct.  EqTo is returned for convenience.
func Precess(eqFrom, eqTo *coord.Equatorial, epochFrom, epochTo, mα, mδ float64) *coord.Equatorial {
	ζCoeff := ζt
	zCoeff := zt
	θCoeff := θt
	if epochFrom != 2000 {
		T := (epochFrom - 2000) * .01
		ζCoeff = []float64{0,
			base.Horner(T, ζT...),
			0.30188 - 0.000344*T,
			0.017998}
		zCoeff = []float64{0,
			base.Horner(T, zT...),
			1.09468 - 0.000066*T,
			0.018203}
		θCoeff = []float64{0,
			base.Horner(T, θT...),
			-0.42665 - 0.000217*T,
			-0.041833}
	}
	t := (epochTo - epochFrom) * .01
	ζ := base.NewAngle(false, 0, 0, base.Horner(t, ζCoeff...)).Rad()
	z := base.NewAngle(false, 0, 0, base.Horner(t, zCoeff...)).Rad()
	θ := base.NewAngle(false, 0, 0, base.Horner(t, θCoeff...)).Rad()

	α := eqFrom.RA + mα*t*100
	δ := eqFrom.Dec + mδ*t*100

	sδ, cδ := math.Sincos(δ)
	sαζ, cαζ := math.Sincos(α + ζ)
	sθ, cθ := math.Sincos(θ)
	A := cδ * sαζ
	B := cθ*cδ*cαζ - sθ*sδ
	eqTo.RA = math.Atan2(A, B) + z
	if math.Pi-math.Abs(δ) > .001 {
		eqTo.Dec = math.Asin(sθ*cδ*cαζ + cθ*sδ)
	} else {
		eqTo.Dec = math.Acos(math.Sqrt(A*A + B*B)) // near pole
	}
	return eqTo
}

var ηt = []float64{0, 47.0029, -0.03302, 0.000060}
var πt = []float64{3600 * 174.876384, -869.8089, 0.03536}
var pt = []float64{0, 5029.0966, 1.11113, -0.000006}

var ηT = []float64{74.0029, -0.06603, 0.000598}
var πT = []float64{3600 * 174.876384, 3289.4789, 0.60622}
var pT = []float64{5029.0966, 2.22226, -0.000042}

// PrecessEcl precesses ecliptic coordinates from one epoch to another,
// including proper motions.
//
// Proper motions must be in arc radians per year.  Pass 0 if proper motions
// are not to be considered or are not applicable.
//
// Both eqFrom and eqTo must be non-nil, although they may point to the same
// struct.  EqTo is returned for convenience.
func PrecessEcl(eclFrom, eclTo *coord.Ecliptic, epochFrom, epochTo, mα, mδ float64) *coord.Ecliptic {
	ηCoeff := ηt
	πCoeff := πt
	pCoeff := pt
	if epochFrom != 2000 {
		T := (epochFrom - 2000) * .01
		ηCoeff = []float64{0,
			base.Horner(T, ηT...),
			-0.03302 + 0.000598*T,
			0.000060}
		πCoeff = []float64{
			base.Horner(T, πT...),
			-869.8089 - 0.50491*T,
			0.03536}
		pCoeff = []float64{0,
			base.Horner(T, pT...),
			1.11113 - 0.000042*T,
			-0.000006}
	}
	t := (epochTo - epochFrom) * .01
	η := base.NewAngle(false, 0, 0, base.Horner(t, ηCoeff...)).Rad()
	π := base.NewAngle(false, 0, 0, base.Horner(t, πCoeff...)).Rad()
	p := base.NewAngle(false, 0, 0, base.Horner(t, pCoeff...)).Rad()

	β := eclFrom.Lat
	λ := eclFrom.Lon
	if mα != 0 || mδ != 0 {
		mλ, mβ := eqProperMotionToEcl(mα, mδ, epochFrom, eclFrom)
		λ += mλ * t * 100
		β += mβ * t * 100
	}

	sη, cη := math.Sincos(η)
	sβ, cβ := math.Sincos(β)
	sd, cd := math.Sincos(π - λ)
	A := cη*cβ*sd - sη*sβ
	B := cβ * cd
	eclTo.Lon = p + π - math.Atan2(A, B)
	if math.Pi-math.Abs(β) > .001 {
		eclTo.Lat = math.Asin(cη*sβ + sη*cβ*sd)
	} else {
		eclTo.Lat = math.Acos(math.Sqrt(A*A + B*B)) // near pole
	}
	return eclTo
}

func eqProperMotionToEcl(mα, mδ, epoch float64, pos *coord.Ecliptic) (mλ, mβ float64) {
	ε := nutation.MeanObliquity(base.JulianYearToJDE(epoch))
	sε, cε := math.Sincos(ε)
	eqPos := new(coord.Equatorial).EclToEq(pos, sε, cε)
	sα, cα := math.Sincos(eqPos.RA)
	sδ, cδ := math.Sincos(eqPos.Dec)
	cβ := math.Cos(pos.Lat)
	mλ = (mδ*sε*cα + mα*cδ*(cε*cδ+sε*sδ*sα)) / (cβ * cβ)
	mβ = (mδ*(cε*cδ+sε*sδ*sα) - mα*sε*cα*cδ) / cβ
	return
}

// ProperMotion3D takes the 3D equatorial coordinates of an object
// at one epoch and computes its coordinates at a new epoch, considering
// proper motion and radial velocity.
//
// Proper motions (mα, mδ) must be in radians per year, radial distance (r)
// in parsecs, radial velocitiy (mr) in parsecs per year.
//
// Both eqFrom and eqTo must be non-nil, although they may point to the same
// struct.  EqTo is returned for convenience.
func ProperMotion3D(eqFrom, eqTo *coord.Equatorial, epochFrom, epochTo, r, mr, mα, mδ float64) *coord.Equatorial {
	sα, cα := math.Sincos(mα)
	sδ, cδ := math.Sincos(mδ)
	x := r * cδ * cα
	y := r * cδ * sα
	z := r * sδ
	mrr := mr / r
	zmδ := z * mδ
	mx := x*mrr - zmδ*cα - y*mα
	my := y*mrr - zmδ*sα + x*mα
	mz := z*mrr + r*mδ*cδ
	t := epochTo - epochFrom
	xp := x + t*mx
	yp := y + t*my
	zp := z + t*mz
	eqTo.RA = eqFrom.RA + math.Atan2(yp, xp)
	eqTo.Dec = eqFrom.Dec + math.Atan2(zp, math.Hypot(xp, yp))
	return eqTo
}
