// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package nearparabolic_test

import (
	"math"
	"math/rand"
	"testing"

	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/meeus/nearparabolic"
)

type tc struct {
	q, e, t, ν, r float64
}

var tdat = []tc{
	// test data p. 247
	{.921326, 1, 138.4783, 102.74426, 2.364192},
	{.1, .987, 254.9, 164.50029, 4.063777},
	{.123456, .99997, -30.47, 221.91190, .965053},
	{3.363943, 1.05731, 1237.1, 109.40598, 10.668551},
	{.5871018, .9672746, 20, 52.85331, .729116},
	{.5871018, .9672746, 0, 0, .5871018},
}

func TestAnomalyDistance(t *testing.T) {
	var e nearparabolic.Elements
	for _, d := range tdat {
		e.TimeP = base.J2000 + rand.Float64()*base.JulianCentury
		e.PDis = d.q
		e.Ecc = d.e
		ν, r, err := e.AnomalyDistance(e.TimeP + d.t)
		if err != nil {
			t.Error(err)
			continue
		}
		if math.Abs(ν.Deg()-d.ν) > 1e-5 {
			t.Errorf("got ν = %.6f expected %.6f", ν.Deg(), d.ν)
		}
		if math.Abs(r-d.r) > 1e-6 {
			t.Errorf("got r = %.7f expected %.7f", r, d.r)
		}
	}
}

type tc2 struct {
	q, e, t, ν float64
	p          int
	c          bool
}

var tdat2 = []tc2{
	// test data p. 248
	{.1, .9, 10, 126, 0, true},
	{.1, .9, 20, 142, 0, true},
	{.1, .9, 30, 0, 0, false},
	{.1, .987, 10, 123, 0, true},
	{.1, .987, 20, 137, 0, true},
	{.1, .987, 30, 143, 0, true},
	{.1, .987, 60, 152, 0, true},
	{.1, .987, 100, 157, 0, true},
	{.1, .987, 200, 163, 0, true},
	{.1, .987, 400, 167, 0, true},
	{.1, .987, 500, 0, 0, false},
	{.1, .999, 100, 156, 0, true},
	{.1, .999, 200, 161, 0, true},
	{.1, .999, 500, 166, 0, true},
	{.1, .999, 1000, 169, 0, true},
	{.1, .999, 5000, 174, 0, true},
	{1, .99999, 100000, 172.5, 1, true},
	{1, .99999, 10000000, 178.41, 2, true},
	{1, .99999, 14000000, 178.58, 2, true},
	{1, .99999, 17000000, 178.68, 2, true},
	{1, .99999, 18000000, 0, 2, false},
}

func TestAnomalyDistance2(t *testing.T) {
	var e nearparabolic.Elements
	for _, d := range tdat2 {
		e.TimeP = base.J2000 + rand.Float64()*base.JulianCentury
		e.PDis = d.q
		e.Ecc = d.e
		ν, _, err := e.AnomalyDistance(e.TimeP + d.t)
		if (err == nil) != d.c {
			t.Errorf("%#v", d)
			continue
		}
		if math.Abs(ν.Deg()-d.ν) > math.Pow(10, float64(-d.p)) {
			t.Errorf("got ν = %.*f expected %.*f",
				d.p+1, ν.Deg(), d.p+1, d.ν)
		}
	}
}
