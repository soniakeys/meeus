// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Jupiter: Chapter 43, Ephemeris for Physical Observations of Jupiter.
package jupiter

import (
	"math"

	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/meeus/nutation"
	pp "github.com/soniakeys/meeus/planetposition"
	"github.com/soniakeys/unit"
)

// Physical computes quantities for physical observations of Jupiter.
//
// Results:
//	DS  Planetocentric declination of the Sun.
//	DE  Planetocentric declination of the Earth.
//	ω1  Longitude of the System I central meridian of the illuminated disk,
//	    as seen from Earth.
//	ω2  Longitude of the System II central meridian of the illuminated disk,
//	    as seen from Earth.
//	P   Geocentric position angle of Jupiter's northern rotation pole.
func Physical(jde float64, earth, jupiter *pp.V87Planet) (DS, DE, ω1, ω2, P unit.Angle) {
	// Step 1.
	d := jde - 2433282.5
	T1 := d / base.JulianCentury
	const p = math.Pi / 180
	α0 := 268*p + .1061*p*T1
	δ0 := 64.5*p - .0164*p*T1
	// Step 2.
	W1 := 17.71*p + 877.90003539*p*d
	W2 := 16.838*p + 870.27003539*p*d
	// Step 3.
	l0, b0, R := earth.Position(jde)
	l0, b0 = pp.ToFK5(l0, b0, jde)
	// Steps 4-7.
	sl0, cl0 := l0.Sincos()
	sb0 := b0.Sin()
	Δ := 4. // surely better than 0.
	var l, b unit.Angle
	var r, x, y, z float64
	f := func() {
		τ := base.LightTime(Δ)
		l, b, r = jupiter.Position(jde - τ)
		l, b = pp.ToFK5(l, b, jde)
		sb, cb := b.Sincos()
		sl, cl := l.Sincos()
		// (42.2) p. 289
		x = r*cb*cl - R*cl0
		y = r*cb*sl - R*sl0
		z = r*sb - R*sb0
		// (42.3) p. 289
		Δ = math.Sqrt(x*x + y*y + z*z)
	}
	f()
	f()
	// Step 8.
	ε0 := nutation.MeanObliquity(jde)
	// Step 9.
	sε0, cε0 := ε0.Sincos()
	sl, cl := l.Sincos()
	sb, cb := b.Sincos()
	αs := math.Atan2(cε0*sl-sε0*sb/cb, cl)
	δs := math.Asin(cε0*sb + sε0*cb*sl)
	// Step 10.
	sδs, cδs := math.Sincos(δs)
	sδ0, cδ0 := math.Sincos(δ0)
	DS = unit.Angle(math.Asin(-sδ0*sδs - cδ0*cδs*math.Cos(α0-αs)))
	// Step 11.
	u := y*cε0 - z*sε0
	v := y*sε0 + z*cε0
	α := math.Atan2(u, x)
	δ := math.Atan(v / math.Hypot(x, u))
	sδ, cδ := math.Sincos(δ)
	sα0α, cα0α := math.Sincos(α0 - α)
	ζ := math.Atan2(sδ0*cδ*cα0α-sδ*cδ0, cδ*sα0α)
	// Step 12.
	DE = unit.Angle(math.Asin(-sδ0*sδ - cδ0*cδ*math.Cos(α0-α)))
	// Step 13.
	ω1 = unit.Angle(W1 - ζ - 5.07033*p*Δ)
	ω2 = unit.Angle(W2 - ζ - 5.02626*p*Δ)
	// Step 14.
	C := unit.Angle((2*r*Δ + R*R - r*r - Δ*Δ) / (4 * r * Δ))
	if (l - l0).Sin() < 0 {
		C = -C
	}
	ω1 = (ω1 + C).Mod1()
	ω2 = (ω2 + C).Mod1()
	// Step 15.
	Δψ, Δε := nutation.Nutation(jde)
	ε := ε0 + Δε
	// Step 16.
	sε, cε := ε.Sincos()
	sα, cα := math.Sincos(α)
	α += .005693 * p * (cα*cl0*cε + sα*sl0) / cδ
	δ += .005693 * p * (cl0*cε*(sε/cε*cδ-sα*sδ) + cα*sδ*sl0)
	// Step 17.
	tδ := sδ / cδ
	Δα := (cε+sε*sα*tδ)*Δψ.Rad() - cα*tδ*Δε.Rad()
	Δδ := sε*cα*Δψ.Rad() + sα*Δε.Rad()
	αʹ := α + Δα
	δʹ := δ + Δδ
	sα0, cα0 := math.Sincos(α0)
	tδ0 := sδ0 / cδ0
	Δα0 := (cε+sε*sα0*tδ0)*Δψ.Rad() - cα0*tδ0*Δε.Rad()
	Δδ0 := sε*cα0*Δψ.Rad() + sα0*Δε.Rad()
	α0ʹ := α0 + Δα0
	δ0ʹ := δ0 + Δδ0
	// Step 18.
	sδʹ, cδʹ := math.Sincos(δʹ)
	sδ0ʹ, cδ0ʹ := math.Sincos(δ0ʹ)
	sα0ʹαʹ, cα0ʹαʹ := math.Sincos(α0ʹ - αʹ)
	// (42.4) p. 290
	P = unit.Angle(math.Atan2(cδ0ʹ*sα0ʹαʹ, sδ0ʹ*cδʹ-cδ0ʹ*sδʹ*cα0ʹαʹ))
	if P < 0 {
		P += 2 * math.Pi
	}
	return
}

// Physical2 computes quantities for physical observations of Jupiter.
//
// Results are less accurate than with Physical().
//
// Results:
//	DS  Planetocentric declination of the Sun.
//	DE  Planetocentric declination of the Earth.
//	ω1  Longitude of the System I central meridian of the illuminated disk,
//	    as seen from Earth.
//	ω2  Longitude of the System II central meridian of the illuminated disk,
//	    as seen from Earth.
//
// All angular results in radians.
func Physical2(jde float64) (DS, DE, ω1, ω2 unit.Angle) {
	d := jde - base.J2000
	const p = math.Pi / 180
	V := 172.74*p + .00111588*p*d
	M := 357.529*p + .9856003*p*d
	sV := math.Sin(V)
	N := 20.02*p + .0830853*p*d + .329*p*sV
	J := 66.115*p + .9025179*p*d - .329*p*sV
	sM, cM := math.Sincos(M)
	sN, cN := math.Sincos(N)
	s2M, c2M := math.Sincos(2 * M)
	s2N, c2N := math.Sincos(2 * N)
	A := 1.915*p*sM + .02*p*s2M
	B := 5.555*p*sN + .168*p*s2N
	K := J + A - B
	R := 1.00014 - .01671*cM - .00014*c2M
	r := 5.20872 - .25208*cN - .00611*c2N
	sK, cK := math.Sincos(K)
	Δ := math.Sqrt(r*r + R*R - 2*r*R*cK)
	ψ := math.Asin(R / Δ * sK)
	dd := d - Δ/173
	ω1 = unit.Angle(210.98*p + 877.8169088*p*dd + ψ - B)
	ω2 = unit.Angle(187.23*p + 870.1869088*p*dd + ψ - B)
	C := unit.Angle(math.Sin(ψ / 2))
	C *= C
	if sK > 0 {
		C = -C
	}
	ω1 = (ω1 + C).Mod1()
	ω2 = (ω2 + C).Mod1()
	λ := 34.35*p + .083091*p*d + .329*p*sV + B
	DS = unit.Angle(3.12 * p * math.Sin(λ+42.8*p))
	DE = DS - unit.Angle(2.22*p*math.Sin(ψ)*math.Cos(λ+22*p)) -
		unit.Angle(1.3*p*(r-Δ)/Δ*math.Sin(λ-100.5*p))
	return
}
