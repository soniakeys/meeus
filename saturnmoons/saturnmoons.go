// Copyright 2013 Sonia Keys
// License: MIT

// Saturnmoons: Chapter 46, Positions of the Satellites of Saturn
package saturnmoons

import (
	"math"

	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/meeus/coord"
	pp "github.com/soniakeys/meeus/planetposition"
	"github.com/soniakeys/meeus/precess"
	"github.com/soniakeys/meeus/solar"
	"github.com/soniakeys/unit"
)

// XY holds coordinates returned from Positions().
type XY struct{ X, Y float64 }

const d = math.Pi / 180

// Positions returns positions of the eight major moons of Saturn.
//
// Results returned in argument pos, which must not be nil.
//
// Result units are Saturn radii.
func Positions(jde float64, earth, saturn *pp.V87Planet, pos *[8]XY) {
	s, β, R := solar.TrueVSOP87(earth, jde)
	ss, cs := s.Sincos()
	sβ := β.Sin()
	Δ := 9.
	var x, y, z float64
	var JDE float64
	f := func() {
		τ := base.LightTime(Δ)
		JDE = jde - τ
		l, b, r := saturn.Position(JDE)
		l, b = pp.ToFK5(l, b, JDE)
		sl, cl := l.Sincos()
		sb, cb := b.Sincos()
		x = r*cb*cl + R*cs
		y = r*cb*sl + R*ss
		z = r*sb + R*sβ
		Δ = math.Sqrt(x*x + y*y + z*z)
	}
	f()
	f()
	λ0 := unit.Angle(math.Atan2(y, x))
	β0 := unit.Angle(math.Atan(z / math.Hypot(x, y)))
	ecl := &coord.Ecliptic{λ0, β0}
	precess.EclipticPosition(ecl, ecl,
		base.JDEToJulianYear(jde), base.JDEToJulianYear(base.B1950), 0, 0)
	λ0, β0 = ecl.Lon, ecl.Lat
	q := newQs(JDE)
	s4 := [9]r4{{}, // 0 unused
		q.mimas(),
		q.enceladus(),
		q.tethys(),
		q.dione(),
		q.rhea(),
		q.titan(),
		q.hyperion(),
		q.iapetus(),
	}
	var X, Y, Z [9]float64
	for j := 1; j <= 8; j++ {
		u := s4[j].λ - s4[j].Ω
		w := s4[j].Ω - 168.8112*d
		su, cu := math.Sincos(u)
		sw, cw := math.Sincos(w)
		sγ, cγ := math.Sincos(s4[j].γ)
		r := s4[j].r
		X[j] = r * (cu*cw - su*cγ*sw)
		Y[j] = r * (su*cw*cγ + cu*sw)
		Z[j] = r * su * sγ
	}
	Z[0] = 1
	sλ0, cλ0 := λ0.Sincos()
	sβ0, cβ0 := β0.Sincos()
	var A, B, C [9]float64
	for j := range X {
		a := X[j]
		b := q.c1*Y[j] - q.s1*Z[j]
		c := q.s1*Y[j] + q.c1*Z[j]
		a, b =
			q.c2*a-q.s2*b,
			q.s2*a+q.c2*b
		A[j], b =
			a*sλ0-b*cλ0,
			a*cλ0+b*sλ0
		B[j], C[j] =
			b*cβ0+c*sβ0,
			c*cβ0-b*sβ0
	}
	D := math.Atan2(A[0], C[0])
	sD, cD := math.Sincos(D)
	for j := 1; j <= 8; j++ {
		X[j] = A[j]*cD - C[j]*sD
		Y[j] = A[j]*sD + C[j]*cD
		Z[j] = B[j]
		d := X[j] / s4[j].r
		X[j] += math.Abs(Z[j]) / k[j] * math.Sqrt(1-d*d)
		W := Δ / (Δ + Z[j]/2475)
		pos[j-1].X = X[j] * W
		pos[j-1].Y = Y[j] * W
	}
	return
}

var k = [...]float64{0, 20947, 23715, 26382, 29876, 35313, 53800, 59222, 91820}

type qs struct {
	t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11  float64
	W0, W1, W2, W3, W4, W5, W6, W7, W8            float64
	s1, c1, s2, c2, e1                            float64
	sW0, s3W0, s5W0, sW1, sW2, sW3, cW3, sW4, cW4 float64
	sW7, cW7                                      float64
}

func newQs(JDE float64) *qs {
	var q qs
	q.t1 = JDE - 2411093
	q.t2 = q.t1 / 365.25
	q.t3 = (JDE-2433282.423)/365.25 + 1950
	q.t4 = JDE - 2411368
	q.t5 = q.t4 / 365.25
	q.t6 = JDE - 2415020
	q.t7 = q.t6 / 36525
	q.t8 = q.t6 / 365.25
	q.t9 = (JDE - 2442000.5) / 365.25
	q.t10 = JDE - 2409786
	q.t11 = q.t10 / 36525
	q.W0 = 5.095 * d * (q.t3 - 1866.39)
	q.W1 = 74.4*d + 32.39*d*q.t2
	q.W2 = 134.3*d + 92.62*d*q.t2
	q.W3 = 42*d - .5118*d*q.t5
	q.W4 = 276.59*d + .5118*d*q.t5
	q.W5 = 267.2635*d + 1222.1136*d*q.t7
	q.W6 = 175.4762*d + 1221.5515*d*q.t7
	q.W7 = 2.4891*d + .002435*d*q.t7
	q.W8 = 113.35*d - .2597*d*q.t7
	q.s1, q.c1 = math.Sincos(28.0817 * d)
	q.s2, q.c2 = math.Sincos(168.8112 * d)
	q.e1 = .05589 - .000346*q.t7
	q.sW0 = math.Sin(q.W0)
	q.s3W0 = math.Sin(3 * q.W0)
	q.s5W0 = math.Sin(5 * q.W0)
	q.sW1 = math.Sin(q.W1)
	q.sW2 = math.Sin(q.W2)
	q.sW3, q.cW3 = math.Sincos(q.W3)
	q.sW4, q.cW4 = math.Sincos(q.W4)
	q.sW7, q.cW7 = math.Sincos(q.W7)
	return &q
}

type r4 struct{ λ, r, γ, Ω float64 }

func (q *qs) mimas() (r r4) {
	L := 127.64*d + 381.994497*d*q.t1 -
		43.57*d*q.sW0 - .72*d*q.s3W0 - .02144*d*q.s5W0
	p := 106.1*d + 365.549*d*q.t2
	M := L - p
	C := 2.18287*d*math.Sin(M) +
		.025988*d*math.Sin(2*M) + .00043*d*math.Sin(3*M)
	r.λ = L + C
	r.r = 3.06879 / (1 + .01905*math.Cos(M+C))
	r.γ = 1.563 * d
	r.Ω = 54.5*d - 365.072*d*q.t2
	return
}

func (q *qs) enceladus() (r r4) {
	L := 200.317*d + 262.7319002*d*q.t1 + .25667*d*q.sW1 + .20883*d*q.sW2
	p := 309.107*d + 123.44121*d*q.t2
	M := L - p
	C := .55577*d*math.Sin(M) + .00168*d*math.Sin(2*M)
	r.λ = L + C
	r.r = 3.94118 / (1 + .00485*math.Cos(M+C))
	r.γ = .0262 * d
	r.Ω = 348*d - 151.95*d*q.t2
	return
}
func (q *qs) tethys() (r r4) {
	r.λ = 285.306*d + 190.69791226*d*q.t1 +
		2.063*d*q.sW0 + .03409*d*q.s3W0 + .001015*d*q.s5W0
	r.r = 4.880998
	r.γ = 1.0976 * d
	r.Ω = 111.33*d - 72.2441*d*q.t2
	return
}
func (q *qs) dione() (r r4) {
	L := 254.712*d + 131.53493193*d*q.t1 - .0215*d*q.sW1 - .01733*d*q.sW2
	p := 174.8*d + 30.82*d*q.t2
	M := L - p
	C := .24717*d*math.Sin(M) + .00033*d*math.Sin(2*M)
	r.λ = L + C
	r.r = 6.24871 / (1 + .002157*math.Cos(M+C))
	r.γ = .0139 * d
	r.Ω = 232*d - 30.27*d*q.t2
	return
}

func (q *qs) rhea() (r r4) {
	pʹ := 342.7*d + 10.057*d*q.t2
	spʹ, cpʹ := math.Sincos(pʹ)
	a1 := .000265*spʹ + .001*q.sW4
	a2 := .000265*cpʹ + .001*q.cW4
	e := math.Hypot(a1, a2)
	p := math.Atan2(a1, a2)
	N := 345*d - 10.057*d*q.t2
	sN, cN := math.Sincos(N)
	λʹ := 359.244*d + 79.6900472*d*q.t1 + .086754*d*sN
	i := 28.0362*d + .346898*d*cN + .0193*d*q.cW3
	Ω := 168.8034*d + .736936*d*sN + .041*d*q.sW3
	a := 8.725924
	return q.subr(λʹ, p, e, a, Ω, i)
}

func (q *qs) subr(λʹ, p, e, a, Ω, i float64) (r r4) {
	M := λʹ - p
	e2 := e * e
	e3 := e2 * e
	e4 := e2 * e2
	e5 := e3 * e2
	C := (2*e-.25*e3+.0520833333*e5)*math.Sin(M) +
		(1.25*e2-.458333333*e4)*math.Sin(2*M) +
		(1.083333333*e3-.671875*e5)*math.Sin(3*M) +
		1.072917*e4*math.Sin(4*M) + 1.142708*e5*math.Sin(5*M)
	r.r = a * (1 - e2) / (1 + e*math.Cos(M+C)) // return value
	g := Ω - 168.8112*d
	si, ci := math.Sincos(i)
	sg, cg := math.Sincos(g)
	a1 := si * sg
	a2 := q.c1*si*cg - q.s1*ci
	r.γ = math.Asin(math.Hypot(a1, a2)) // return value
	u := math.Atan2(a1, a2)
	r.Ω = 168.8112*d + u // return value (w)
	h := q.c1*si - q.s1*ci*cg
	ψ := math.Atan2(q.s1*sg, h)
	r.λ = λʹ + C + u - g - ψ // return value
	return
}

func (q *qs) titan() (r r4) {
	L := 261.1582*d + 22.57697855*d*q.t4 + .074025*d*q.sW3
	iʹ := 27.45141*d + .295999*d*q.cW3
	Ωʹ := 168.66925*d + .628808*d*q.sW3
	siʹ, ciʹ := math.Sincos(iʹ)
	sΩʹW8, cΩʹW8 := math.Sincos(Ωʹ - q.W8)
	a1 := q.sW7 * sΩʹW8
	a2 := q.cW7*siʹ - q.sW7*ciʹ*cΩʹW8
	g0 := 102.8623 * d
	ψ := math.Atan2(a1, a2)
	s := math.Hypot(a1, a2)
	g := q.W4 - Ωʹ - ψ
	var ϖ float64
	s2g0, c2g0 := math.Sincos(2 * g0)
	f := func() {
		ϖ = q.W4 + .37515*d*(math.Sin(2*g)-s2g0)
		g = ϖ - Ωʹ - ψ
	}
	f()
	f()
	f()
	eʹ := .029092 + .00019048*(math.Cos(2*g)-c2g0)
	qq := 2 * (q.W5 - ϖ)
	b1 := siʹ * sΩʹW8
	b2 := q.cW7*siʹ*cΩʹW8 - q.sW7*ciʹ
	θ := math.Atan2(b1, b2) + q.W8
	sq, cq := math.Sincos(qq)
	e := eʹ + .002778797*eʹ*cq
	p := ϖ + .159215*d*sq
	u := 2*q.W5 - 2*θ + ψ
	su, cu := math.Sincos(u)
	h := .9375*eʹ*eʹ*sq + .1875*s*s*math.Sin(2*(q.W5-θ))
	λʹ := L - .254744*d*
		(q.e1*math.Sin(q.W6)+.75*q.e1*q.e1*math.Sin(2*q.W6)+h)
	i := iʹ + .031843*d*s*cu
	Ω := Ωʹ + .031843*d*s*su/siʹ
	a := 20.216193
	return q.subr(λʹ, p, e, a, Ω, i)
}

func (q *qs) hyperion() (r r4) {
	η := 92.39*d + .5621071*d*q.t6
	ζ := 148.19*d - 19.18*d*q.t8
	θ := 184.8*d - 35.41*d*q.t9
	θʹ := θ - 7.5*d
	as := 176*d + 12.22*d*q.t8
	bs := 8*d + 24.44*d*q.t8
	cs := bs + 5*d
	ϖ := 69.898*d - 18.67088*d*q.t8
	φ := 2 * (ϖ - q.W5)
	χ := 94.9*d - 2.292*d*q.t8
	sη, cη := math.Sincos(η)
	sζ, cζ := math.Sincos(ζ)
	s2ζ, c2ζ := math.Sincos(2 * ζ)
	s3ζ, c3ζ := math.Sincos(3 * ζ)
	sζpη, cζpη := math.Sincos(ζ + η)
	sζmη, cζmη := math.Sincos(ζ - η)
	sφ, cφ := math.Sincos(φ)
	sχ, cχ := math.Sincos(χ)
	scs, ccs := math.Sincos(cs)
	a := 24.50601 - .08686*cη - .00166*cζpη + .00175*cζmη
	e := .103458 - .004099*cη - .000167*cζpη + .000235*cζmη +
		.02303*cζ - .00212*c2ζ + 0.000151*c3ζ + .00013*cφ
	p := ϖ + .15648*d*sχ - .4457*d*sη - .2657*d*sζpη - .3573*d*sζmη -
		12.872*d*sζ + 1.668*d*s2ζ - .2419*d*s3ζ - .07*d*sφ
	λʹ := 177.047*d + 16.91993829*d*q.t6 + .15648*d*sχ + 9.142*d*sη +
		.007*d*math.Sin(2*η) - .014*d*math.Sin(3*η) + .2275*d*sζpη +
		.2112*d*sζmη - .26*d*sζ - .0098*d*s2ζ -
		.013*d*math.Sin(as) + .017*d*math.Sin(bs) - .0303*d*sφ
	i := 27.3347*d + .6434886*d*cχ + .315*d*q.cW3 + .018*d*math.Cos(θ) -
		.018*d*ccs
	Ω := 168.6812*d + 1.40136*d*cχ + .68599*d*q.sW3 - .0392*d*scs +
		.0366*d*math.Sin(θʹ)
	return q.subr(λʹ, p, e, a, Ω, i)
}

func (q *qs) iapetus() (r r4) {
	L := 261.1582*d + 22.57697855*d*q.t4
	ϖʹ := 91.796*d + .562*d*q.t7
	ψ := 4.367*d - .195*d*q.t7
	θ := 146.819*d - 3.198*d*q.t7
	φ := 60.47*d + 1.521*d*q.t7
	Φ := 205.055*d - 2.091*d*q.t7
	eʹ := .028298 + .001156*q.t11
	ϖ0 := 352.91*d + 11.71*d*q.t11
	μ := 76.3852*d + 4.53795125*d*q.t10
	iʹ := base.Horner(q.t11, 18.4602*d, -.9518*d, -.072*d, .0054*d)
	Ωʹ := base.Horner(q.t11, 143.198*d, -3.919*d, .116*d, .008*d)
	l := μ - ϖ0
	g := ϖ0 - Ωʹ - ψ
	g1 := ϖ0 - Ωʹ - φ
	ls := q.W5 - ϖʹ
	gs := ϖʹ - θ
	lT := L - q.W4
	gT := q.W4 - Φ
	u1 := 2 * (l + g - ls - gs)
	u2 := l + g1 - lT - gT
	u3 := l + 2*(g-ls-gs)
	u4 := lT + gT - g1
	u5 := 2 * (ls + gs)
	sl, cl := math.Sincos(l)
	su1, cu1 := math.Sincos(u1)
	su2, cu2 := math.Sincos(u2)
	su3, cu3 := math.Sincos(u3)
	su4, cu4 := math.Sincos(u4)
	slu2, clu2 := math.Sincos(l + u2)
	sg1gT, cg1gT := math.Sincos(g1 - gT)
	su52g, cu52g := math.Sincos(u5 - 2*g)
	su5ψ, cu5ψ := math.Sincos(u5 + ψ)
	su2φ, cu2φ := math.Sincos(u2 + φ)
	s5, c5 := math.Sincos(l + g1 + lT + gT + φ)
	a := 58.935028 + .004638*cu1 + .058222*cu2
	e := eʹ - .0014097*cg1gT + .0003733*cu52g +
		.000118*cu3 + .0002408*cl + .0002849*clu2 + .000619*cu4
	w := .08077*d*sg1gT + .02139*d*su52g - .00676*d*su3 +
		.0138*d*sl + .01632*d*slu2 + .03547*d*su4
	p := ϖ0 + w/eʹ
	λʹ := μ - .04299*d*su2 - .00789*d*su1 - .06312*d*math.Sin(ls) -
		.00295*d*math.Sin(2*ls) - .02231*d*math.Sin(u5) + .0065*d*su5ψ
	i := iʹ + .04204*d*cu5ψ + .00235*d*c5 + .0036*d*cu2φ
	wʹ := .04204*d*su5ψ + .00235*d*s5 + .00358*d*su2φ
	Ω := Ωʹ + wʹ/math.Sin(iʹ)
	return q.subr(λʹ, p, e, a, Ω, i)
}
