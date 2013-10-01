// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Moon: Chapter 53, Ephemeris for Physical Observations of the Moon.
//
// Incomplete.  Topocentric functions are commented out for lack of test data.
package moon

import (
	"math"

	"github.com/soniakeys/meeus/base"
	//	"github.com/soniakeys/meeus/parallax"
	"github.com/soniakeys/meeus/coord"
	"github.com/soniakeys/meeus/moonposition"
	"github.com/soniakeys/meeus/nutation"

	pp "github.com/soniakeys/meeus/planetposition"
	"github.com/soniakeys/meeus/solar"
)

const p = math.Pi / 180
const _I = 1.54242 * p // IAU value of inclination of mean lunar equator
var sI, cI = math.Sincos(_I)

// Physical returns quantities useful for physical observation of the Moon.
//
// Returned l, b are librations in selenographic longitude and latitude.
// They represent combined optical and physical librations.  Topocentric
// librations are not considered.
//
// Returned P is the the position angle of the Moon's axis of rotation.
//
// Returned l0, b0 are the selenographic coordinates of the Sun.
//
// Returned values all in radians.
func Physical(jde float64, earth *pp.V87Planet) (l, b, P, l0, b0 float64) {
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
	Δψ      float64 // nutation in longitude
	F       float64 // argument of latitude of Moon
	Ω       float64 // mean longitude of the ascending node of lunar orbit
	sε, cε  float64 // true obliquity of the ecliptic
	ρ, σ, τ float64
}

func newMoon(jde float64) *moon {
	m := &moon{jde: jde}
	// Δψ, F, Ω, p. 372.
	var Δε float64
	m.Δψ, Δε = nutation.Nutation(jde)
	T := base.J2000Century(jde)
	F := base.Horner(T, 93.272095*p, 483202.0175233*p,
		-.0036539*p, -p/3526000, p/863310000)
	m.F = F
	m.Ω = base.Horner(T, 125.0445479*p, -1934.1362891*p, .0020754*p,
		p/467441, -p/60616000)
	// true ecliptic
	m.sε, m.cε = math.Sincos(nutation.MeanObliquity(jde) + Δε)
	// ρ, σ, τ, p. 372,373
	D := base.Horner(T, 297.8501921*p, 445267.1114034*p,
		-.0018819*p, p/545868, -p/113065000)
	M := base.Horner(T, 357.5291092*p, 35999.0502909*p,
		-.0001535*p, p/24490000)
	Mʹ := base.Horner(T, 134.9633964*p, 477198.8675055*p,
		.0087414*p, p/69699, -p/14712000)
	E := base.Horner(T, 1, -.002516, -.0000074)
	K1 := 119.75*p + 131.849*p*T
	K2 := 72.56*p + 20.186*p*T
	m.ρ = -.02752*p*math.Cos(Mʹ) +
		-.02245*p*math.Sin(F) +
		.00684*p*math.Cos(Mʹ-2*F) +
		-.00293*p*math.Cos(2*F) +
		-.00085*p*math.Cos(2*(F-D)) +
		-.00054*p*math.Cos(Mʹ-2*D) +
		-.0002*p*math.Sin(Mʹ+F) +
		-.0002*p*math.Cos(Mʹ+2*F) +
		-.0002*p*math.Cos(Mʹ-F) +
		.00014*p*math.Cos(Mʹ+2*(F-D))
	m.σ = -.02816*p*math.Sin(Mʹ) +
		.02244*p*math.Cos(F) +
		-.00682*p*math.Sin(Mʹ-2*F) +
		-.00279*p*math.Sin(2*F) +
		-.00083*p*math.Sin(2*(F-D)) +
		.00069*p*math.Sin(Mʹ-2*D) +
		.0004*p*math.Cos(Mʹ+F) +
		-.00025*p*math.Sin(2*Mʹ) +
		-.00023*p*math.Sin(Mʹ+2*F) +
		.0002*p*math.Cos(Mʹ-F) +
		.00019*p*math.Sin(Mʹ-F) +
		.00013*p*math.Sin(Mʹ+2*(F-D)) +
		-.0001*p*math.Cos(Mʹ-3*F)
	m.τ = .0252*p*math.Sin(M)*E +
		.00473*p*math.Sin(2*(Mʹ-F)) +
		-.00467*p*math.Sin(Mʹ) +
		.00396*p*math.Sin(K1) +
		.00276*p*math.Sin(2*(Mʹ-D)) +
		.00196*p*math.Sin(m.Ω) +
		-.00183*p*math.Cos(Mʹ-F) +
		.00115*p*math.Sin(Mʹ-2*D) +
		-.00096*p*math.Sin(Mʹ-D) +
		.00046*p*math.Sin(2*(F-D)) +
		-.00039*p*math.Sin(Mʹ-F) +
		-.00032*p*math.Sin(Mʹ-M-D) +
		.00027*p*math.Sin(2*(Mʹ-D)-M) +
		.00023*p*math.Sin(K2) +
		-.00014*p*math.Sin(2*D) +
		.00014*p*math.Cos(2*(Mʹ-F)) +
		-.00012*p*math.Sin(Mʹ-2*F) +
		-.00012*p*math.Sin(2*Mʹ) +
		.00011*p*math.Sin(2*(Mʹ-M-D))
	return m
}

// lib() curiously serves for computing both librations and solar coordinates,
// depending on the coordinates λ, β passed in.  Quantity A not described in
// the book, but clearly depends on the λ, β of the current context and so
// does not belong in the moon struct.  Instead just return it from optical
// and pass it along to physical.
func (m *moon) lib(λ, β float64) (l, b float64) {
	lʹ, bʹ, A := m.optical(λ, β)
	lʺ, bʺ := m.physical(A, bʹ)
	l = lʹ + lʺ
	if l > math.Pi {
		l -= 2 * math.Pi
	}
	b = bʹ + bʺ
	return
}

func (m *moon) optical(λ, β float64) (lʹ, bʹ, A float64) {
	// (53.1) p. 372
	W := λ - m.Ω // (λ without nutation)
	sW, cW := math.Sincos(W)
	sβ, cβ := math.Sincos(β)
	A = math.Atan2(sW*cβ*cI-sβ*sI, cW*cβ)
	lʹ = base.PMod(A-m.F, 2*math.Pi)
	bʹ = math.Asin(-sW*cβ*sI - sβ*cI)
	return
}

func (m *moon) physical(A, bʹ float64) (lʺ, bʺ float64) {
	// (53.2) p. 373
	sA, cA := math.Sincos(A)
	lʺ = -m.τ + (m.ρ*cA+m.σ*sA)*math.Tan(bʹ)
	bʺ = m.σ*cA - m.ρ*sA
	return
}

func (m *moon) pa(λ, β, b float64) float64 {
	V := m.Ω + m.Δψ + m.σ/sI
	sV, cV := math.Sincos(V)
	sIρ, cIρ := math.Sincos(_I + m.ρ)
	X := sIρ * sV
	Y := sIρ*cV*m.cε - cIρ*m.sε
	ω := math.Atan2(X, Y)
	α, _ := coord.EclToEq(λ+m.Δψ, β, m.sε, m.cε)
	P := math.Asin(math.Hypot(X, Y) * math.Cos(α-ω) / math.Cos(b))
	if P < 0 {
		P += 2 * math.Pi
	}
	return P
}

func (m *moon) sun(λ, β, Δ float64, earth *pp.V87Planet) (l0, b0 float64) {
	λ0, _, R := solar.ApparentVSOP87(earth, m.jde)
	ΔR := Δ / (R * base.AU)
	λH := λ0 + math.Pi + ΔR*math.Cos(β)*math.Sin(λ0-λ)
	βH := ΔR * β
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
//
// Result is altitude in radians.
func SunAltitude(η, θ, l0, b0 float64) float64 {
	c0 := math.Pi/2 - l0
	sb0, cb0 := math.Sincos(b0)
	sθ, cθ := math.Sincos(θ)
	return math.Asin(sb0*sθ + cb0*cθ*math.Sin(c0+η))
}

// Sunrise returns time of sunrise for a point on the Moon near the given date.
//
// Arguments η, θ are selenographic longitude and latitude of a site on the
// Moon, jde can be any date.
//
// Returned is the time of sunrise as a jde nearest the given jde.
func Sunrise(η, θ, jde float64, earth *pp.V87Planet) float64 {
	jde -= srCorr(η, θ, jde, earth)
	return jde - srCorr(η, θ, jde, earth)
}

// Sunset returns time of sunset for a point on the Moon near the given date.
//
// Arguments η, θ are selenographic longitude and latitude of a site on the
// Moon, jde can be any date.
//
// Returned is the time of sunset as a jde nearest the given jde.
func Sunset(η, θ, jde float64, earth *pp.V87Planet) float64 {
	jde += srCorr(η, θ, jde, earth)
	return jde + srCorr(η, θ, jde, earth)
}

func srCorr(η, θ, jde float64, earth *pp.V87Planet) float64 {
	_, _, _, l0, b0 := Physical(jde, earth)
	h := SunAltitude(η, θ, l0, b0)
	return h / (12.19075 * p * math.Cos(θ))
}
