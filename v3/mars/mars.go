// Copyright 2013 Sonia Keys
// License: MIT

// Mars: Chapter 42, Ephemeris for Physical Observations of Mars.
package mars

import (
	"math"

	"github.com/soniakeys/meeus/v3/base"
	"github.com/soniakeys/meeus/v3/coord"
	"github.com/soniakeys/meeus/v3/illum"
	"github.com/soniakeys/meeus/v3/nutation"
	pp "github.com/soniakeys/meeus/v3/planetposition"
	"github.com/soniakeys/unit"
)

// Physical computes quantities for physical observations of Mars.
//
// Results:
//	DE  planetocentric declination of the Earth.
//	DS  planetocentric declination of the Sun.
//	ω   Areographic longitude of the central meridian, as seen from Earth.
//	P   Geocentric position angle of Mars' northern rotation pole.
//	Q   Position angle of greatest defect of illumination.
//	d   Apparent diameter of Mars.
//	q   Greatest defect of illumination.
//	k   Illuminated fraction of the disk.
func Physical(jde float64, earth, mars *pp.V87Planet) (DE, DS, ω, P, Q, d, q unit.Angle, k float64) {
	// Step 1.
	T := base.J2000Century(jde)
	const p = math.Pi / 180
	// (42.1) p. 288
	λ0 := 352.9065*p + 1.1733*p*T
	β0 := 63.2818*p - .00394*p*T
	// Step 2.
	l0, b0, R := earth.Position(jde)
	l0, b0 = pp.ToFK5(l0, b0, jde)
	// Steps 3, 4.
	sl0, cl0 := l0.Sincos()
	sb0 := b0.Sin()
	Δ := .5 // surely better than 0.
	τ := base.LightTime(Δ)
	var l, b unit.Angle
	var r, x, y, z float64
	f := func() {
		l, b, r = mars.Position(jde - τ)
		l, b = pp.ToFK5(l, b, jde)
		sb, cb := b.Sincos()
		sl, cl := l.Sincos()
		// (42.2) p. 289
		x = r*cb*cl - R*cl0
		y = r*cb*sl - R*sl0
		z = r*sb - R*sb0
		// (42.3) p. 289
		Δ = math.Sqrt(x*x + y*y + z*z)
		τ = base.LightTime(Δ)
	}
	f()
	f()
	// Step 5.
	λ := math.Atan2(y, x)
	β := math.Atan(z / math.Hypot(x, y))
	// Step 6.
	sβ0, cβ0 := math.Sincos(β0)
	sβ, cβ := math.Sincos(β)
	DE = unit.Angle(math.Asin(-sβ0*sβ - cβ0*cβ*math.Cos(λ0-λ)))
	// Step 7.
	N := 49.5581*p + .7721*p*T
	lʹ := l.Rad() - .00697*p/r
	bʹ := b.Rad() - .000225*p*math.Cos(l.Rad()-N)/r
	// Step 8.
	sbʹ, cbʹ := math.Sincos(bʹ)
	DS = unit.Angle(math.Asin(-sβ0*sbʹ - cβ0*cbʹ*math.Cos(λ0-lʹ)))
	// Step 9.
	W := 11.504*p + 350.89200025*p*(jde-τ-2433282.5)
	// Step 10.
	ε0 := nutation.MeanObliquity(jde)
	sε0, cε0 := ε0.Sincos()
	α0, δ0 := coord.EclToEq(unit.Angle(λ0), unit.Angle(β0), sε0, cε0)
	// Step 11.
	u := y*cε0 - z*sε0
	v := y*sε0 + z*cε0
	α := math.Atan2(u, x)
	δ := math.Atan(v / math.Hypot(x, u))
	sδ, cδ := math.Sincos(δ)
	sδ0, cδ0 := δ0.Sincos()
	sα0α, cα0α := math.Sincos(α0.Rad() - α)
	ζ := math.Atan2(sδ0*cδ*cα0α-sδ*cδ0, cδ*sα0α)
	// Step 12.
	ω = unit.Angle(W - ζ).Mod1()
	// Step 13.
	Δψ, Δε := nutation.Nutation(jde)
	// Step 14.
	sl0λ, cl0λ := math.Sincos(l0.Rad() - λ)
	λ += .005693 * p * cl0λ / cβ
	β += .005693 * p * sl0λ * sβ
	// Step 15.
	λ0 += Δψ.Rad()
	λ += Δψ.Rad()
	ε := ε0 + Δε
	// Step 16.
	sε, cε := ε.Sincos()
	α0ʹ, δ0ʹ := coord.EclToEq(unit.Angle(λ0), unit.Angle(β0), sε, cε)
	αʹ, δʹ := coord.EclToEq(unit.Angle(λ), unit.Angle(β), sε, cε)
	// Step 17.
	sδ0ʹ, cδ0ʹ := δ0ʹ.Sincos()
	sδʹ, cδʹ := δʹ.Sincos()
	sα0ʹαʹ, cα0ʹαʹ := (α0ʹ - αʹ).Sincos()
	// (42.4) p. 290
	P = unit.Angle(math.Atan2(cδ0ʹ*sα0ʹαʹ, sδ0ʹ*cδʹ-cδ0ʹ*sδʹ*cα0ʹαʹ))
	if P < 0 {
		P += 2 * math.Pi
	}
	// Step 18.
	s := l0 + math.Pi
	ss, cs := s.Sincos()
	αs := math.Atan2(cε*ss, cs)
	δs := math.Asin(sε * ss)
	sδs, cδs := math.Sincos(δs)
	sαsα, cαsα := math.Sincos(αs - α)
	χ := math.Atan2(cδs*sαsα, sδs*cδ-cδs*sδ*cαsα)
	Q = unit.Angle(χ) + math.Pi
	// Step 19.
	d = unit.AngleFromSec(9.36) / unit.Angle(Δ)
	k = illum.Fraction(r, Δ, R)
	q = d.Mul(1 - k)
	return
}
