// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Solardisk: Chapter 29, Ephemeris for Physical Observations of the Sun.
package solardisk

import (
	"math"

	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/meeus/nutation"
	pp "github.com/soniakeys/meeus/planetposition"
	"github.com/soniakeys/meeus/solar"
)

// Ephemeris returns the apparent orientation of the sun at the given jd.
//
// Results:
//	P:  Position angle of the solar north pole.
//	B0: Heliographic latitude of the center of the solar disk.
//	L0: Heliographic longitude of the center of the solar disk.
//
// All results in radians.
func Ephemeris(jd float64, e *pp.V87Planet) (P, B0, L0 float64) {
	θ := (jd - 2398220) * 2 * math.Pi / 25.38
	I := 7.25 * math.Pi / 180
	K := 73.6667*math.Pi/180 +
		1.3958333*math.Pi/180*(jd-2396758)/base.JulianCentury

	L, _, R := solar.TrueVSOP87(e, jd)
	Δψ, Δε := nutation.Nutation(jd)
	ε0 := nutation.MeanObliquity(jd)
	ε := ε0 + Δε
	λ := L - 20.4898/3600*math.Pi/180/R
	λp := λ + Δψ

	sλK, cλK := math.Sincos(λ - K)
	sI, cI := math.Sincos(I)

	tx := -math.Cos(λp) * math.Tan(ε)
	ty := -cλK * math.Tan(I)
	P = math.Atan(tx) + math.Atan(ty)
	B0 = math.Asin(sλK * sI)
	η := math.Atan2(-sλK*cI, -cλK)
	L0 = base.PMod(η-θ, 2*math.Pi)
	return
}

// Cycle returns the jd of the start of the given synodic rotation.
//
// Argument c is the "Carrington" cycle number.
//
// Result is a dynamical time (not UT).
func Cycle(c int) (jde float64) {
	cf := float64(c)
	jde = 2398140.227 + 27.2752316*cf
	m := 281.96*math.Pi/180 + 26.882476*math.Pi/180*cf
	s2m, c2m := math.Sincos(2 * m)
	return jde + .1454*math.Sin(m) - .0085*s2m - .0141*c2m
}
