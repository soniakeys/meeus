// Copyright 2013 Sonia Keys
// License: MIT

// Moon: Chapter 53, Ephemeris for Physical Observations of the Moon.
//
// Incomplete.  Topocentric functions are commented out for lack of test data.
package moon

import (
	"math"

	"github.com/soniakeys/meeus/v3/base"
	"github.com/soniakeys/meeus/v3/coord"
	"github.com/soniakeys/meeus/v3/moonposition"
	"github.com/soniakeys/meeus/v3/nutation"
	//	"github.com/soniakeys/meeus/v3/parallax"
	pp "github.com/soniakeys/meeus/v3/planetposition"
	"github.com/soniakeys/meeus/v3/solar"
	"github.com/soniakeys/unit"
)

// IAU value of inclination of mean lunar equator, p. 372
var _I = unit.AngleFromDeg(1.54242)
var sI, cI = math.Sincos(_I.Rad())

// Physical returns quantities useful for physical observation of the Moon.
//
// Returned l, b are librations in selenographic longitude and latitude.
// They represent combined optical and physical librations.  Topocentric
// librations are not considered.
//
// Returned P is the the position angle of the Moon's axis of rotation.
//
// Returned l0, b0 are the selenographic coordinates of the Sun.
func Physical(jde float64, earth *pp.V87Planet) (l, b, P, l0, b0 unit.Angle) {
	λ, β, Δ := moonposition.Position(jde) // (λ without nutation)
	m := newMoon(jde)
	l, b = m.lib(λ, β)
	P = m.pa(λ, β, b)
	l0, b0 = m.sun(λ, β, Δ, earth)
	return
}

// Quantities computed for a jde and used in computing return values of
// Physical().  Computations are broken into several methods to organize
// the code.
type moon struct {
	jde     float64
	Δψ      unit.Angle // nutation in longitude
	F       unit.Angle // argument of latitude of Moon
	Ω       unit.Angle // mean longitude of the ascending node of lunar orbit
	sε, cε  float64    // true obliquity of the ecliptic
	ρ, σ, τ unit.Angle
}

func newMoon(jde float64) *moon {
	m := &moon{jde: jde}
	// Δψ, F, Ω, p. 372.
	var Δε unit.Angle
	m.Δψ, Δε = nutation.Nutation(jde)
	T := base.J2000Century(jde)
	m.F = unit.AngleFromDeg(base.Horner(T,
		93.272095, 483202.0175233, -.0036539, -1/3526000, 1/863310000))
	F := m.F.Rad()
	m.Ω = unit.AngleFromDeg(base.Horner(T,
		125.0445479, -1934.1362891, .0020754, 1/467441, -1/60616000))
	// true ecliptic
	m.sε, m.cε = math.Sincos((nutation.MeanObliquity(jde) + Δε).Rad())
	// ρ, σ, τ, p. 372,373
	D := unit.AngleFromDeg(base.Horner(T,
		297.8501921, 445267.1114034, -.0018819, 1/545868, -1/113065000)).Rad()
	M := unit.AngleFromDeg(base.Horner(T,
		357.5291092, 35999.0502909, -.0001536, 1/24490000)).Rad()
	Mʹ := unit.AngleFromDeg(base.Horner(T,
		134.9633964, 477198.8675055, .0087414, 1/69699, -1/14712000)).Rad()
	E := base.Horner(T, 1, -.002516, -.0000074)
	K1 := unit.AngleFromDeg(119.75 + 131.849*T).Rad()
	K2 := unit.AngleFromDeg(72.56 + 20.186*T).Rad()
	m.ρ = unit.AngleFromDeg(-.02752*math.Cos(Mʹ) +
		-.02245*math.Sin(F) +
		.00684*math.Cos(Mʹ-2*F) +
		-.00293*math.Cos(2*F) +
		-.00085*math.Cos(2*(F-D)) +
		-.00054*math.Cos(Mʹ-2*D) +
		-.0002*math.Sin(Mʹ+F) +
		-.0002*math.Cos(Mʹ+2*F) +
		-.0002*math.Cos(Mʹ-F) +
		.00014*math.Cos(Mʹ+2*(F-D)))
	m.σ = unit.AngleFromDeg(-.02816*math.Sin(Mʹ) +
		.02244*math.Cos(F) +
		-.00682*math.Sin(Mʹ-2*F) +
		-.00279*math.Sin(2*F) +
		-.00083*math.Sin(2*(F-D)) +
		.00069*math.Sin(Mʹ-2*D) +
		.0004*math.Cos(Mʹ+F) +
		-.00025*math.Sin(2*Mʹ) +
		-.00023*math.Sin(Mʹ+2*F) +
		.0002*math.Cos(Mʹ-F) +
		.00019*math.Sin(Mʹ-F) +
		.00013*math.Sin(Mʹ+2*(F-D)) +
		-.0001*math.Cos(Mʹ-3*F))
	m.τ = unit.AngleFromDeg(.0252*math.Sin(M)*E +
		.00473*math.Sin(2*(Mʹ-F)) +
		-.00467*math.Sin(Mʹ) +
		.00396*math.Sin(K1) +
		.00276*math.Sin(2*(Mʹ-D)) +
		.00196*math.Sin(m.Ω.Rad()) +
		-.00183*math.Cos(Mʹ-F) +
		.00115*math.Sin(Mʹ-2*D) +
		-.00096*math.Sin(Mʹ-D) +
		.00046*math.Sin(2*(F-D)) +
		-.00039*math.Sin(Mʹ-F) +
		-.00032*math.Sin(Mʹ-M-D) +
		.00027*math.Sin(2*(Mʹ-D)-M) +
		.00023*math.Sin(K2) +
		-.00014*math.Sin(2*D) +
		.00014*math.Cos(2*(Mʹ-F)) +
		-.00012*math.Sin(Mʹ-2*F) +
		-.00012*math.Sin(2*Mʹ) +
		.00011*math.Sin(2*(Mʹ-M-D)))
	return m
}

// lib() curiously serves for computing both librations and solar coordinates,
// depending on the coordinates λ, β passed in.  Quantity A not described in
// the book, but clearly depends on the λ, β of the current context and so
// does not belong in the moon struct.  Instead just return it from optical
// and pass it along to physical.
func (m *moon) lib(λ, β unit.Angle) (l, b unit.Angle) {
	lʹ, bʹ, A := m.optical(λ, β)
	lʺ, bʺ := m.physical(A, bʹ)
	l = lʹ + lʺ
	if l > math.Pi {
		l -= 2 * math.Pi
	}
	b = bʹ + bʺ
	return
}

func (m *moon) optical(λ, β unit.Angle) (lʹ, bʹ, A unit.Angle) {
	// (53.1) p. 372
	W := λ - m.Ω // (λ without nutation)
	sW, cW := W.Sincos()
	sβ, cβ := β.Sincos()
	A = unit.Angle(math.Atan2(sW*cβ*cI-sβ*sI, cW*cβ))
	lʹ = (A - m.F).Mod1()
	bʹ = unit.Angle(math.Asin(-sW*cβ*sI - sβ*cI))
	return
}

func (m *moon) physical(A, bʹ unit.Angle) (lʺ, bʺ unit.Angle) {
	// (53.2) p. 373
	sA, cA := A.Sincos()
	lʺ = -m.τ + (m.ρ.Mul(cA) + m.σ.Mul(sA)).Mul(bʹ.Tan())
	bʺ = m.σ.Mul(cA) - m.ρ.Mul(sA)
	return
}

func (m *moon) pa(λ, β, b unit.Angle) unit.Angle {
	V := m.Ω + m.Δψ + m.σ.Div(sI)
	sV, cV := V.Sincos()
	sIρ, cIρ := (_I + m.ρ).Sincos()
	X := sIρ * sV
	Y := sIρ*cV*m.cε - cIρ*m.sε
	ω := math.Atan2(X, Y)
	α, _ := coord.EclToEq(λ+m.Δψ, β, m.sε, m.cε)
	P := unit.Angle(math.Asin(math.Hypot(X, Y) * math.Cos(α.Rad()-ω) / b.Cos()))
	if P < 0 {
		P += 2 * math.Pi
	}
	return P
}

func (m *moon) sun(λ, β unit.Angle, Δ float64, earth *pp.V87Planet) (l0, b0 unit.Angle) {
	λ0, _, R := solar.ApparentVSOP87(earth, m.jde)
	ΔR := Δ / (R * base.AU)
	λH := λ0 + math.Pi + unit.AngleFromDeg(57.296).Mul(ΔR*(β.Cos()*(λ0-λ).Sin()))
	βH := β.Mul(ΔR)
	return m.lib(λH, βH)
}

/* commented out for lack of test data
func Topocentric(jde, ρsφʹ, ρcφʹ, L float64) (l, b, P float64) {
	λ, β, Δ := moonposition.Position(jde) // (λ without nutation)
	Δψ, Δε := nutation.Nutation(jde)
	sε, cε := math.Sincos(nutation.MeanObliquity(jde) + Δε)
	α, δ := coord.EclToEq(λ+Δψ, β, sε, cε)
	α, δ = parallax.Topocentric(α, δ, Δ/base.AU, ρsφʹ, ρcφʹ, L, jde)
	λ, β = coord.EqToEcl(α, δ, sε, cε)
	m := newMoon(jde)
	l, b = m.lib(λ, β)
	P = m.pa(λ, β, b)
	return
}

func TopocentricCorrections(jde, b, P, φ, δ, H, π float64) (Δl, Δb, ΔP float64) {
	sφ, cφ := math.Sincos(φ)
	sH, cH := math.Sincos(H)
	sδ, cδ := math.Sincos(δ)
	Q := math.Atan(cφ * sH / (cδ*sφ - sδ*cφ*cH))
	z := math.Acos(sδ*sφ + cδ*cφ*cH)
	πʹ := π * (math.Sin(z) + .0084*math.Sin(2*z))
	sQP, cQP := math.Sincos(Q - P)
	Δl = -πʹ * sQP / math.Cos(b)
	Δb = πʹ * cQP
	ΔP = Δl*math.Sin(b+Δb) - πʹ*math.Sin(Q)*math.Tan(δ)
	return
}
*/

// SunAltitude returns altitude of the Sun above the lunar horizon.
//
// Arguments η, θ are selenographic longitude and latitude of a site on the
// Moon, l0, b0 are selenographic coordinates of the Sun, as returned by
// Physical(), for example.
func SunAltitude(η, θ, l0, b0 unit.Angle) unit.Angle {
	c0 := math.Pi/2 - l0
	sb0, cb0 := b0.Sincos()
	sθ, cθ := θ.Sincos()
	return unit.Angle(math.Asin(sb0*sθ + cb0*cθ*(c0+η).Sin()))
}

// Sunrise returns time of sunrise for a point on the Moon near the given date.
//
// Arguments η, θ are selenographic longitude and latitude of a site on the
// Moon, jde can be any date.
//
// Returned is the time of sunrise as a jde nearest the given jde.
func Sunrise(η, θ unit.Angle, jde float64, earth *pp.V87Planet) float64 {
	jde -= srCorr(η, θ, jde, earth)
	return jde - srCorr(η, θ, jde, earth)
}

// Sunset returns time of sunset for a point on the Moon near the given date.
//
// Arguments η, θ are selenographic longitude and latitude of a site on the
// Moon, jde can be any date.
//
// Returned is the time of sunset as a jde nearest the given jde.
func Sunset(η, θ unit.Angle, jde float64, earth *pp.V87Planet) float64 {
	jde += srCorr(η, θ, jde, earth)
	return jde + srCorr(η, θ, jde, earth)
}

func srCorr(η, θ unit.Angle, jde float64, earth *pp.V87Planet) float64 {
	_, _, _, l0, b0 := Physical(jde, earth)
	h := SunAltitude(η, θ, l0, b0)
	return h.Deg() / 12.19075 / θ.Cos()
}
