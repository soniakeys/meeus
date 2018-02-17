// Copyright 2013 Sonia Keys
// License: MIT

// Solardisk: Chapter 29, Ephemeris for Physical Observations of the Sun.
package solardisk

import (
	"math"

	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/meeus/nutation"
	pp "github.com/soniakeys/meeus/planetposition"
	"github.com/soniakeys/meeus/solar"
	"github.com/soniakeys/unit"
)

// Ephemeris returns the apparent orientation of the sun at the given jd.
//
// Results:
//	P:  Position angle of the solar north pole.
//	B0: Heliographic latitude of the center of the solar disk.
//	L0: Heliographic longitude of the center of the solar disk.
func Ephemeris(jd float64, e *pp.V87Planet) (P, B0, L0 unit.Angle) {
	θ := unit.Angle((jd - 2398220) * 2 * math.Pi / 25.38)
	I := unit.AngleFromDeg(7.25)
	K := unit.AngleFromDeg(73.6667) +
		unit.AngleFromDeg(1.3958333).Mul((jd-2396758)/base.JulianCentury)

	L, _, R := solar.TrueVSOP87(e, jd)
	Δψ, Δε := nutation.Nutation(jd)
	ε0 := nutation.MeanObliquity(jd)
	ε := ε0 + Δε
	λ := L - unit.AngleFromSec(20.4898).Div(R)
	λp := λ + Δψ

	sλK, cλK := (λ - K).Sincos()
	sI, cI := I.Sincos()

	tx := -(λp.Cos() * ε.Tan())
	ty := -(cλK * I.Tan())
	P = unit.Angle(math.Atan(tx) + math.Atan(ty))
	B0 = unit.Angle(math.Asin(sλK * sI))
	η := unit.Angle(math.Atan2(-sλK*cI, -cλK))
	L0 = (η - θ).Mod1()
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
