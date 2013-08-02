// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Mars: Chapter 42, Ephemeris for Physical Observations of Mars.
package mars

import (
	//	"fmt"
	"math"

	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/meeus/coord"
	"github.com/soniakeys/meeus/illum"
	"github.com/soniakeys/meeus/nutation"
	pp "github.com/soniakeys/meeus/planetposition"
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
//	k   Illuminated fraction of the disk.
//	q   Greatest defect of illumination.
//
// All angular results (all results except k) are in radians.
func Physical(jde float64, earth, mars *pp.V87Planet) (DE, DS, ω, P, Q, d, k, q float64) {
	// Step 1.
	T := base.J2000Century(jde)
	//	fmt.Printf("T: %.10f\n", T)
	const p = math.Pi / 180
	λ0 := 352.9065*p + 1.1733*p*T
	β0 := 63.2818*p - .00394*p*T
	//	fmt.Printf("λ0: %.5f deg\n", λ0/p)
	//	fmt.Printf("β0: %+.5f deg\n", β0/p)
	// Step 2.
	l0, b0, R := earth.Position(jde)
	l0, b0 = pp.ToFK5(l0, b0, jde)
	//	fmt.Printf("l0: %.6f deg\n", l0/p)
	//	fmt.Printf("b0: %.6f deg\n", b0/p)
	//	fmt.Printf("R: %.8f\n", R)
	// Steps 3, 4.
	sl0, cl0 := math.Sincos(l0)
	sb0 := math.Sin(b0)
	Δ := .5 // surely better than 0.
	τ := base.LightTime(Δ)
	var l, b, r, x, y, z float64
	f := func() {
		l, b, r = mars.Position(jde - τ)
		l, b = pp.ToFK5(l, b, jde)
		sb, cb := math.Sincos(b)
		sl, cl := math.Sincos(l)
		x = r*cb*cl - R*cl0
		y = r*cb*sl - R*sl0
		z = r*sb - R*sb0
		Δ = math.Sqrt(x*x + y*y + z*z)
		τ = base.LightTime(Δ)
	}
	f()
	//	fmt.Printf("l: %.6f deg\n", l/p)
	//	fmt.Printf("b: %.6f deg\n", b/p)
	//	fmt.Printf("r: %.8f\n", r)
	//	fmt.Printf("Δ: %.9f\n", Δ)
	//	fmt.Printf("τ: %.6f\n", τ)
	f()
	//	fmt.Printf("l: %.6f deg\n", l/p)
	//	fmt.Printf("b: %.6f deg\n", b/p)
	//	fmt.Printf("r: %.8f\n", r)
	//	fmt.Printf("Δ: %.9f\n", Δ)
	//	fmt.Printf("τ: %.6f\n", τ)
	// Step 5.
	λ := math.Atan2(y, x)
	β := math.Atan(z / math.Hypot(x, y))
	//	fmt.Printf("λ: %.6f\n", λ/p)
	//	fmt.Printf("β: %.6f\n", β/p)
	// Step 6.
	sβ0, cβ0 := math.Sincos(β0)
	sβ, cβ := math.Sincos(β)
	DE = math.Asin(-sβ0*sβ - cβ0*cβ*math.Cos(λ0-λ))
	// Step 7.
	N := 49.5581*p + .7721*p*T
	lʹ := l - .00697*p/r
	bʹ := b - .000225*p*math.Cos(l-N)/r
	//	fmt.Printf("N: %.4f deg\n", N/p)
	//	fmt.Printf("lʹ: %.6f deg\n", lʹ/p)
	//	fmt.Printf("bʹ: %+.6f deg\n", bʹ/p)
	// Step 8.
	sbʹ, cbʹ := math.Sincos(bʹ)
	DS = math.Asin(-sβ0*sbʹ - cβ0*cbʹ*math.Cos(λ0-lʹ))
	// Step 9.
	W := 11.504*p + 350.89200025*p*(jde-τ-2433282.5)
	//	fmt.Printf("W: %.4f deg\n", base.PMod(W/p, 360))
	// Step 10.
	ε0 := nutation.MeanObliquity(jde)
	sε0, cε0 := math.Sincos(ε0)
	α0, δ0 := coord.EclToEq(λ0, β0, sε0, cε0)
	//	fmt.Printf("ε0: %.6f deg\n", ε0/p)
	//	fmt.Printf("α0: %.6f deg\n", α0/p)
	//	fmt.Printf("δ0: %+.6f deg\n", δ0/p)
	// Step 11.
	u := y*cε0 - z*sε0
	v := y*sε0 + z*cε0
	//	fmt.Printf("u: %.7f\n", u)
	//	fmt.Printf("v: %.7f\n", v)
	α := math.Atan2(u, x)
	δ := math.Atan(v / math.Hypot(x, u))
	sδ, cδ := math.Sincos(δ)
	sδ0, cδ0 := math.Sincos(δ0)
	sα0α, cα0α := math.Sincos(α0 - α)
	ζ := math.Atan2(sδ0*cδ*cα0α-sδ*cδ0, cδ*sα0α)
	//	fmt.Printf("α: %.6f deg\n", α/p)
	//	fmt.Printf("δ: %+.6f deg\n", δ/p)
	//	fmt.Printf("ζ: %.4f deg\n", ζ/p+360)
	// Step 12.
	ω = base.PMod(W-ζ, 2*math.Pi)
	// Step 13.
	Δψ, Δε := nutation.Nutation(jde)
	//	fmt.Printf("Δψ: %.2f sec\n", Δψ*180/math.Pi*60*60)
	//	fmt.Printf("Δε: %.2f sec\n", Δε*180/math.Pi*60*60)
	// Step 14.
	sl0λ, cl0λ := math.Sincos(l0 - λ)
	λ += .005693 * p * cl0λ / cβ
	β += .005693 * p * sl0λ * sβ
	//	fmt.Printf("λ: %.6f\n", λ/p)
	//	fmt.Printf("β: %.6f\n", β/p)
	// Step 15.
	λ0 += Δψ
	λ += Δψ
	ε := ε0 + Δε
	//	fmt.Printf("λ0: %.6f deg\n", λ0/p)
	//	fmt.Printf("λ: %.6f deg\n", λ/p)
	//	fmt.Printf("ε: %.6f deg\n", ε/p)
	// Step 16.
	sε, cε := math.Sincos(ε)
	α0ʹ, δ0ʹ := coord.EclToEq(λ0, β0, sε, cε)
	αʹ, δʹ := coord.EclToEq(λ, β, sε, cε)
	//	fmt.Printf("α0ʹ: %.6f deg\n", α0ʹ/p)
	//	fmt.Printf("δ0ʹ: %+.6f deg\n", δ0ʹ/p)
	//	fmt.Printf("αʹ: %.6f deg\n", αʹ/p)
	//	fmt.Printf("δʹ: %+.6f deg\n", δʹ/p)
	// Step 17.
	sδ0ʹ, cδ0ʹ := math.Sincos(δ0ʹ)
	sδʹ, cδʹ := math.Sincos(δʹ)
	sα0ʹαʹ, cα0ʹαʹ := math.Sincos(α0ʹ - αʹ)
	P = math.Atan2(cδ0ʹ*sα0ʹαʹ, sδ0ʹ*cδʹ-cδ0ʹ*sδʹ*cα0ʹαʹ)
	if P < 0 {
		P += 2 * math.Pi
	}
	// Step 18.
	s := l0 + math.Pi
	ss, cs := math.Sincos(s)
	αs := math.Atan2(cε*ss, cs)
	δs := math.Asin(sε * ss)
	//	fmt.Printf("αs: %.3f deg\n", αs/p+360)
	//	fmt.Printf("δs: %+.3f deg\n", δs/p)
	sδs, cδs := math.Sincos(δs)
	sαsα, cαsα := math.Sincos(αs - α)
	χ := math.Atan2(cδs*sαsα, sδs*cδ-cδs*sδ*cαsα)
	Q = χ + math.Pi
	//	fmt.Printf("χ: %.2f deg\n", χ/p)
	// Step 19.
	d = 9.36 / 60 / 60 * math.Pi / 180 / Δ
	k = illum.Fraction2(r, Δ, R)
	q = (1 - k) * d
	return
}
