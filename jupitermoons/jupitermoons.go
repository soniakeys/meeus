// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Jupitermoons: Chapter 42, Positions of the Satellites of Jupiter.
package jupitermoons

import (
	"math"

	"github.com/soniakeys/meeus/base"
	pe "github.com/soniakeys/meeus/planetelements"
	pp "github.com/soniakeys/meeus/planetposition"
	"github.com/soniakeys/meeus/solar"
)

// XY used for returning coordinates of moons.
type XY struct {
	X, Y float64
}

// Positions computes positions of moons of Jupiter.
//
// Returned coordinates are in units of Jupiter radii.
func Positions(jde float64) (pI, pII, pIII, pIV XY) {
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
	λ := 34.35*p + .083091*p*d + .329*p*sV + B
	DS := 3.12 * p * math.Sin(λ+42.8*p)
	DE := DS - 2.22*p*math.Sin(ψ)*math.Cos(λ+22*p) -
		1.3*p*(r-Δ)/Δ*math.Sin(λ-100.5*p)
	dd := d - Δ/173
	u1 := 163.8069*p + 203.4058646*p*dd + ψ - B
	u2 := 358.414*p + 101.2916335*p*dd + ψ - B
	u3 := 5.7176*p + 50.234518*p*dd + ψ - B
	u4 := 224.8092*p + 21.48798*p*dd + ψ - B
	G := 331.18*p + 50.310482*p*dd
	H := 87.45*p + 21.569231*p*dd
	s212, c212 := math.Sincos(2 * (u1 - u2))
	s223, c223 := math.Sincos(2 * (u2 - u3))
	sG, cG := math.Sincos(G)
	sH, cH := math.Sincos(H)
	c1 := .473 * p * s212
	c2 := 1.065 * p * s223
	c3 := .165 * p * sG
	c4 := .843 * p * sH
	r1 := 5.9057 - .0244*c212
	r2 := 9.3966 - .0882*c223
	r3 := 14.9883 - .0216*cG
	r4 := 26.3627 - .1939*cH
	sDE := math.Sin(DE)
	xy := func(u, r float64) XY {
		su, cu := math.Sincos(u)
		return XY{r * su, -r * cu * sDE}
	}
	return xy(u1+c1, r1), xy(u2+c2, r2), xy(u3+c3, r3), xy(u4+c4, r4)
}

// Positions computes positions of moons of Jupiter.
//
// High accuracy method based on theory "E5."  Results returned in
// argument pos, which must not be nil.  Returned coordinates in units
// of Jupiter radii.
func E5(jde float64, earth, jupiter *pp.V87Planet, pos *[4]XY) {
	// I'll interject that I don't trust the results of this function.
	// There is obviously a great chance of typographic errors.
	// My Y results for the test case of the example don't agree with
	// Meeus's well at all, but do agree with the results from the less
	// accurate method.  This would seem to indicate a typo in Meeus's
	// computer implementation.  On the other hand, while my X results
	// agree reasonably well with his, our X results for satellite III
	// don't agree well with the result from the less accurate method,
	// perhaps indicating a typo in the presented algorithm.

	// variables assigned in following block
	var λ0, β0, t float64
	Δ := 5.
	{
		s, β, R := solar.TrueVSOP87(earth, jde)
		ss, cs := math.Sincos(s.Rad())
		sβ := math.Sin(β.Rad())
		τ := base.LightTime(Δ)
		var x, y, z float64
		f := func() {
			l, b, r := jupiter.Position(jde - τ)
			sl, cl := math.Sincos(l.Rad())
			sb, cb := math.Sincos(b.Rad())
			x = r*cb*cl + R*cs
			y = r*cb*sl + R*ss
			z = r*sb + R*sβ
			Δ = math.Sqrt(x*x + y*y + z*z)
			τ = base.LightTime(Δ)
		}
		f()
		f()
		λ0 = math.Atan2(y, x)
		β0 = math.Atan(z / math.Hypot(x, y))
		t = jde - 2443000.5 - τ
	}
	const p = math.Pi / 180
	l1 := 106.07719*p + 203.48895579*p*t
	l2 := 175.73161*p + 101.374724735*p*t
	l3 := 120.55883*p + 50.317609207*p*t
	l4 := 84.44459*p + 21.571071177*p*t
	π1 := 97.0881*p + .16138586*p*t
	π2 := 154.8663*p + .04726307*p*t
	π3 := 188.184*p + .00712734*p*t
	π4 := 335.2868*p + .00184*p*t
	ω1 := 312.3346*p - .13279386*p*t
	ω2 := 100.4411*p - .03263064*p*t
	ω3 := 119.1942*p - .00717703*p*t
	ω4 := 322.6186*p - .00175934*p*t
	Γ := .33033*p*math.Sin(163.679*p+.0010512*p*t) +
		.03439*p*math.Sin(34.486*p-.0161731*p*t)
	Φλ := 199.6766*p + .1737919*p*t
	ψ := 316.5182*p - .00000208*p*t
	G := 30.23756*p + .0830925701*p*t + Γ
	Gʹ := 31.97853*p + .0334597339*p*t
	const Π = 13.469942 * p

	Σ1 := .47259*p*math.Sin(2*(l1-l2)) +
		-.03478*p*math.Sin(π3-π4) +
		.01081*p*math.Sin(l2-2*l3+π3) +
		.00738*p*math.Sin(Φλ) +
		.00713*p*math.Sin(l2-2*l3+π2) +
		-.00674*p*math.Sin(π1+π3-2*Π-2*G) +
		.00666*p*math.Sin(l2-2*l3+π4) +
		.00445*p*math.Sin(l1-π3) +
		-.00354*p*math.Sin(l1-l2) +
		-.00317*p*math.Sin(2*ψ-2*Π) +
		.00265*p*math.Sin(l1-π4) +
		-.00186*p*math.Sin(G) +
		.00162*p*math.Sin(π2-π3) +
		.00158*p*math.Sin(4*(l1-l2)) +
		-.00155*p*math.Sin(l1-l3) +
		-.00138*p*math.Sin(ψ+ω3-2*Π-2*G) +
		-.00115*p*math.Sin(2*(l1-2*l2+ω2)) +
		.00089*p*math.Sin(π2-π4) +
		.00085*p*math.Sin(l1+π3-2*Π-2*G) +
		.00083*p*math.Sin(ω2-ω3) +
		.00053*p*math.Sin(ψ-ω2)
	Σ2 := 1.06476*p*math.Sin(2*(l2-l3)) +
		.04256*p*math.Sin(l1-2*l2+π3) +
		.03581*p*math.Sin(l2-π3) +
		.02395*p*math.Sin(l1-2*l2+π4) +
		.01984*p*math.Sin(l2-π4) +
		-.01778*p*math.Sin(Φλ) +
		.01654*p*math.Sin(l2-π2) +
		.01334*p*math.Sin(l2-2*l3+π2) +
		.01294*p*math.Sin(π3-π4) +
		-.01142*p*math.Sin(l2-l3) +
		-.01057*p*math.Sin(G) +
		-.00775*p*math.Sin(2*(ψ-Π)) +
		.00524*p*math.Sin(2*(l1-l2)) +
		-.0046*p*math.Sin(l1-l3) +
		.00316*p*math.Sin(ψ-2*G+ω3-2*Π) +
		-.00203*p*math.Sin(π1+π3-2*Π-2*G) +
		.00146*p*math.Sin(ψ-ω3) +
		-.00145*p*math.Sin(2*G) +
		.00125*p*math.Sin(ψ-ω4) +
		-.00115*p*math.Sin(l1-2*l3+π3) +
		-.00094*p*math.Sin(2*(l2-ω2)) +
		.00086*p*math.Sin(2*(l1-2*l2+ω2)) +
		-.00086*p*math.Sin(5*Gʹ-2*G+52.225*p) +
		-.00078*p*math.Sin(l2-l4) +
		-.00064*p*math.Sin(3*l3-7*l4+4*π4) +
		.00064*p*math.Sin(π1-π4) +
		-.00063*p*math.Sin(l1-2*l3+π4) +
		.00058*p*math.Sin(ω3-ω4) +
		.00056*p*math.Sin(2*(ψ-Π-G)) +
		.00056*p*math.Sin(2*(l2-l4)) +
		.00055*p*math.Sin(2*(l1-l3)) +
		.00052*p*math.Sin(3*l3-7*l4+π3+3*π4) +
		-.00043*p*math.Sin(l1-π3) +
		.00041*p*math.Sin(5*(l2-l3)) +
		.00041*p*math.Sin(π4-Π) +
		.00032*p*math.Sin(ω2-ω3) +
		.00032*p*math.Sin(2*(l3-G-Π))
	Σ3 := .1649*p*math.Sin(l3-π3) +
		.09081*p*math.Sin(l3-π4) +
		-.06907*p*math.Sin(l2-l3) +
		.03784*p*math.Sin(π3-π4) +
		.01846*p*math.Sin(2*(l3-l4)) +
		-.0134*p*math.Sin(G) +
		-.01014*p*math.Sin(2*(ψ-Π)) +
		.00704*p*math.Sin(l2-2*l3+π3) +
		-.0062*p*math.Sin(l2-2*l3+π2) +
		-.00541*p*math.Sin(l3-l4) +
		.00381*p*math.Sin(l2-2*l3+π4) +
		.00235*p*math.Sin(ψ-ω3) +
		.00198*p*math.Sin(ψ-ω4) +
		.00176*p*math.Sin(Φλ) +
		.0013*p*math.Sin(3*(l3-l4)) +
		.00125*p*math.Sin(l1-l3) +
		-.00119*p*math.Sin(5*Gʹ-2*G+52.225*p) +
		.00109*p*math.Sin(l1-l2) +
		-.001*p*math.Sin(3*l3-7*l4+4*π4) +
		.00091*p*math.Sin(ω3-ω4) +
		.0008*p*math.Sin(3*l3-7*l4+π3+3*π4) +
		-.00075*p*math.Sin(2*l2-3*l3+π3) +
		.00072*p*math.Sin(π1+π3-2*Π-2*G) +
		.00069*p*math.Sin(π4-Π) +
		-.00058*p*math.Sin(2*l3-3*l4+π4) +
		-.00057*p*math.Sin(l3-2*l4+π4) +
		.00056*p*math.Sin(l3+π3-2*Π-2*G) +
		-.00052*p*math.Sin(l2-2*l3+π1) +
		-.00050*p*math.Sin(π2-π3) +
		.00048*p*math.Sin(l3-2*l4+π3) +
		-.00045*p*math.Sin(2*l2-3*l3+π4) +
		-.00041*p*math.Sin(π2-π4) +
		-.00038*p*math.Sin(2*G) +
		-.00037*p*math.Sin(π3-π4+ω3-ω4) +
		-.00032*p*math.Sin(3*l3-7*l4+2*π3+2*π4) +
		.0003*p*math.Sin(4*(l3-l4)) +
		.00029*p*math.Sin(l3+π4-2*Π-2*G) +
		-.00028*p*math.Sin(ω3+ψ-2*Π-2*G) +
		.00026*p*math.Sin(l3-Π-G) +
		.00024*p*math.Sin(l2-3*l3+2*l4) +
		.00021*p*math.Sin(2*(l3-Π-G)) +
		-.00021*p*math.Sin(l3-π2) +
		.00017*p*math.Sin(2*(l3-π3))
	Σ4 := .84287*p*math.Sin(l4-π4) +
		.03431*p*math.Sin(π4-π3) +
		-.03305*p*math.Sin(2*(ψ-Π)) +
		-.03211*p*math.Sin(G) +
		-.01862*p*math.Sin(l4-π3) +
		.01186*p*math.Sin(ψ-ω4) +
		.00623*p*math.Sin(l4+π4-2*G-2*Π) +
		.00387*p*math.Sin(2*(l4-π4)) +
		-.00284*p*math.Sin(5*Gʹ-2*G+52.225*p) +
		-.00234*p*math.Sin(2*(ψ-π4)) +
		-.00223*p*math.Sin(l3-l4) +
		-.00208*p*math.Sin(l4-Π) +
		.00178*p*math.Sin(ψ+ω4-2*π4) +
		.00134*p*math.Sin(π4-Π) +
		.00125*p*math.Sin(2*(l4-G-Π)) +
		-.00117*p*math.Sin(2*G) +
		-.00112*p*math.Sin(2*(l3-l4)) +
		.00107*p*math.Sin(3*l3-7*l4+4*π4) +
		.00102*p*math.Sin(l4-G-Π) +
		.00096*p*math.Sin(2*l4-ψ-ω4) +
		.00087*p*math.Sin(2*(ψ-ω4)) +
		-.00085*p*math.Sin(3*l3-7*l4+π3+3*π4) +
		.00085*p*math.Sin(l3-2*l4+π4) +
		-.00081*p*math.Sin(2*(l4-ψ)) +
		.00071*p*math.Sin(l4+π4-2*Π-3*G) +
		.00061*p*math.Sin(l1-l4) +
		-.00056*p*math.Sin(ψ-ω3) +
		-.00054*p*math.Sin(l3-2*l4+π3) +
		.00051*p*math.Sin(l2-l4) +
		.00042*p*math.Sin(2*(ψ-G-Π)) +
		.00039*p*math.Sin(2*(π4-ω4)) +
		.00036*p*math.Sin(ψ+Π-π4-ω4) +
		.00035*p*math.Sin(2*Gʹ-G+188.37*p) +
		-.00035*p*math.Sin(l4-π4+2*Π-2*ψ) +
		-.00032*p*math.Sin(l4+π4-2*Π-G) +
		.0003*p*math.Sin(2*Gʹ-2*G+149.15*p) +
		.00029*p*math.Sin(3*l3-7*l4+2*π3+2*π4) +
		.00028*p*math.Sin(l4-π4+2*ψ-2*Π) +
		-.00028*p*math.Sin(2*(l4-ω4)) +
		-.00027*p*math.Sin(π3-π4+ω3-ω4) +
		-.00026*p*math.Sin(5*Gʹ-3*G+188.37*p) +
		.00025*p*math.Sin(ω4-ω3) +
		-.00025*p*math.Sin(l2-3*l3+2*l4) +
		-.00023*p*math.Sin(3*(l3-l4)) +
		.00021*p*math.Sin(2*l4-2*Π-3*G) +
		-.00021*p*math.Sin(2*l3-3*l4+π4) +
		.00019*p*math.Sin(l4-π4-G) +
		-.00019*p*math.Sin(2*l4-π3-π4) +
		-.00018*p*math.Sin(l4-π4+G) +
		-.00016*p*math.Sin(l4+π3-2*Π-2*G)
	L1 := l1 + Σ1
	L2 := l2 + Σ2
	L3 := l3 + Σ3
	L4 := l4 + Σ4
	// variables assigned in following block
	var I float64
	X := make([]float64, 5)
	Y := make([]float64, 5)
	Z := make([]float64, 5)
	var R [4]float64
	{
		L := [...]float64{L1, L2, L3, L4}
		B := [...]float64{
			math.Atan(.0006393*p*math.Sin(L1-ω1) +
				.0001825*p*math.Sin(L1-ω2) +
				.0000329*p*math.Sin(L1-ω3) +
				-.0000311*p*math.Sin(L1-ψ) +
				.0000093*p*math.Sin(L1-ω4) +
				.0000075*p*math.Sin(3*L1-4*l2-1.9927*Σ1+ω2) +
				.0000046*p*math.Sin(L1+ψ-2*Π-2*G)),
			math.Atan(.0081004*p*math.Sin(L2-ω2) +
				.0004512*p*math.Sin(L2-ω3) +
				-.0003284*p*math.Sin(L2-ψ) +
				.0001160*p*math.Sin(L2-ω4) +
				.0000272*p*math.Sin(l1-2*l3+1.0146*Σ2+ω2) +
				-.0000144*p*math.Sin(L2-ω1) +
				.0000143*p*math.Sin(L2+ψ-2*Π-2*G) +
				.0000035*p*math.Sin(L2-ψ+G) +
				-.0000028*p*math.Sin(l1-2*l3+1.0146*Σ2+ω3)),
			math.Atan(.0032402*p*math.Sin(L3-ω3) +
				-.0016911*p*math.Sin(L3-ψ) +
				.0006847*p*math.Sin(L3-ω4) +
				-.0002797*p*math.Sin(L3-ω2) +
				.0000321*p*math.Sin(L3+ψ-2*Π-2*G) +
				.0000051*p*math.Sin(L3-ψ+G) +
				-.0000045*p*math.Sin(L3-ψ-G) +
				-.0000045*p*math.Sin(L3+ψ-2*Π) +
				.0000037*p*math.Sin(L3+ψ-2*Π-3*G) +
				.000003*p*math.Sin(2*l2-3*L3+4.03*Σ3+ω2) +
				-.0000021*p*math.Sin(2*l2-3*L3+4.03*Σ3+ω3)),
			math.Atan(-.0076579*p*math.Sin(L4-ψ) +
				.0044134*p*math.Sin(L4-ω4) +
				-.0005112*p*math.Sin(L4-ω3) +
				.0000773*p*math.Sin(L4+ψ-2*Π-2*G) +
				.0000104*p*math.Sin(L4-ψ+G) +
				-.0000102*p*math.Sin(L4-ψ-G) +
				.0000088*p*math.Sin(L4+ψ-2*Π-3*G) +
				-.0000038*p*math.Sin(L4+ψ-2*Π-G)),
		}
		R = [...]float64{
			5.90569 * (1 +
				-.0041339*math.Cos(2*(l1-l2)) +
				-.0000387*math.Cos(l1-π3) +
				-.0000214*math.Cos(l1-π4) +
				.000017*math.Cos(l1-l2) +
				-.0000131*math.Cos(4*(l1-l2)) +
				.0000106*math.Cos(l1-l3) +
				-.0000066*math.Cos(l1+π3-2*Π-2*G)),
			9.39657 * (1 +
				.0093848*math.Cos(l1-l2) +
				-.0003116*math.Cos(l2-π3) +
				-.0001744*math.Cos(l2-π4) +
				-.0001442*math.Cos(l2-π2) +
				.0000553*math.Cos(l2-l3) +
				.0000523*math.Cos(l1-l3) +
				-.0000290*math.Cos(2*(l1-l2)) +
				.0000164*math.Cos(2*(l2-ω2)) +
				.0000107*math.Cos(l1-2*l3+π3) +
				-.0000102*math.Cos(l2-π1) +
				-.0000091*math.Cos(2*(l1-l3))),
			14.98832 * (1 +
				-.0014388*math.Cos(l3-π3) +
				-.0007917*math.Cos(l3-π4) +
				.0006342*math.Cos(l2-l3) +
				-.0001761*math.Cos(2*(l3-l4)) +
				.0000294*math.Cos(l3-l4) +
				-.0000156*math.Cos(3*(l3-l4)) +
				.0000156*math.Cos(l1-l3) +
				-.0000153*math.Cos(l1-l2) +
				.000007*math.Cos(2*l2-3*l3+π3) +
				-.0000051*math.Cos(l3+π3-2*Π-2*G)),
			26.36273 * (1 +
				-.0073546*math.Cos(l4-π4) +
				.0001621*math.Cos(l4-π3) +
				.0000974*math.Cos(l3-l4) +
				-.0000543*math.Cos(l4+π4-2*Π-2*G) +
				-.0000271*math.Cos(2*(l4-π4)) +
				.0000182*math.Cos(l4-Π) +
				.0000177*math.Cos(2*(l3-l4)) +
				-.0000167*math.Cos(2*l4-ψ-ω4) +
				.0000167*math.Cos(ψ-ω4) +
				-.0000155*math.Cos(2*(l4-Π-G)) +
				.0000142*math.Cos(2*(l4-ψ)) +
				.0000105*math.Cos(l1-l4) +
				.0000092*math.Cos(l2-l4) +
				-.0000089*math.Cos(l4-Π-G) +
				-.0000062*math.Cos(l4+π4-2*Π-3*G) +
				.0000048*math.Cos(2*(l4-ω4))),
		}
		// p. 311
		T0 := (jde - 2433282.423) / base.JulianCentury
		P := (1.3966626*p + .0003088*p*T0) * T0
		for i := range L {
			L[i] += P
		}
		ψ += P
		T := (jde - base.J1900) / base.JulianCentury
		I = 3.120262*p + .0006*p*T
		for i := range L {
			sLψ, cLψ := math.Sincos(L[i] - ψ)
			sB, cB := math.Sincos(B[i])
			X[i] = R[i] * cLψ * cB
			Y[i] = R[i] * sLψ * cB
			Z[i] = R[i] * sB
		}
	}
	Z[4] = 1
	// p. 312
	A := make([]float64, 5)
	B := make([]float64, 5)
	C := make([]float64, 5)
	sI, cI := math.Sincos(I)
	Ω := pe.Node(pe.Jupiter, jde)
	sΩ, cΩ := math.Sincos(Ω)
	sΦ, cΦ := math.Sincos(ψ - Ω)
	si, ci := math.Sincos(pe.Inc(pe.Jupiter, jde))
	sλ0, cλ0 := math.Sincos(λ0)
	sβ0, cβ0 := math.Sincos(β0)
	for i := range A {
		// step 1
		a := X[i]
		b := Y[i]*cI - Z[i]*sI
		c := Y[i]*sI + Z[i]*cI
		// step 2
		a, b =
			a*cΦ-b*sΦ,
			a*sΦ+b*cΦ
		// step 3
		b, c =
			b*ci-c*si,
			b*si+c*ci
		// step 4
		a, b =
			a*cΩ-b*sΩ,
			a*sΩ+b*cΩ
		// step 5
		a, b =
			a*sλ0-b*cλ0,
			a*cλ0+b*sλ0
		// step 6
		A[i] = a
		B[i] = c*sβ0 + b*cβ0
		C[i] = c*cβ0 - b*sβ0
	}
	sD, cD := math.Sincos(math.Atan2(A[4], C[4]))
	// p. 313
	for i := 0; i < 4; i++ {
		x := A[i]*cD - C[i]*sD
		y := A[i]*sD + C[i]*cD
		z := B[i]
		// differential light time
		d := x / R[i]
		x += math.Abs(z) / k[i] * math.Sqrt(1-d*d)
		// perspective effect
		W := Δ / (Δ + z/2095)
		pos[i].X = x * W
		pos[i].Y = y * W
	}
	return
}

var k = [...]float64{17295, 21819, 27558, 36548}
