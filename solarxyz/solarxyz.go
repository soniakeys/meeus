// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Solarxyz: Chapter 26, Rectangular Coordinates of the Sun.
package solarxyz

import (
	"math"

	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/meeus/nutation"
	pp "github.com/soniakeys/meeus/planetposition"
	"github.com/soniakeys/meeus/solar"
	"github.com/soniakeys/unit"
)

// Position returns rectangular coordinates referenced to the mean equinox
// of date.
func Position(e *pp.V87Planet, jde float64) (x, y, z float64) {
	// (26.1) p. 171
	s, β, R := solar.TrueVSOP87(e, jde)
	sε, cε := nutation.MeanObliquity(jde).Sincos()
	ss, cs := s.Sincos()
	sβ := β.Sin()
	x = R * cs
	y = R * (ss*cε - sβ*sε)
	z = R * (ss*sε + sβ*cε)
	return
}

// LongitudeJ2000 returns geometric longitude referenced to equinox J2000.
func LongitudeJ2000(e *pp.V87Planet, jde float64) (l unit.Angle) {
	l, _, _ = e.Position2000(jde)
	return (l + math.Pi - unit.AngleFromSec(.09033)).Mod1()
}

// PositionJ2000 returns rectangular coordinates referenced to equinox J2000.
func PositionJ2000(e *pp.V87Planet, jde float64) (x, y, z float64) {
	x, y, z = xyz(e, jde)
	// (26.3) p. 174
	return x + .00000044036*y - .000000190919*z,
		-.000000479966*x + .917482137087*y - .397776982902*z,
		.397776982902*y + .917482137087*z
}

func xyz(e *pp.V87Planet, jde float64) (x, y, z float64) {
	l, b, r := e.Position2000(jde)
	s := l + math.Pi
	β := -b
	ss, cs := s.Sincos()
	sβ, cβ := β.Sincos()
	// (26.2) p. 172
	x = r * cβ * cs
	y = r * cβ * ss
	z = r * sβ
	return
}

// PositionB1950 returns rectangular coordinates referenced to B1950.
//
// Results are referenced to the mean equator and equinox of the epoch B1950
// in the FK5 system, not FK4.
func PositionB1950(e *pp.V87Planet, jde float64) (x, y, z float64) {
	x, y, z = xyz(e, jde)
	return .999925702634*x + .012189716217*y + .000011134016*z,
		-.011179418036*x + .917413998946*y - .397777041885*z,
		-.004859003787*x + .397747363646*y + .917482111428*z
}

var (
	ζt = []float64{2306.2181, 0.30188, 0.017998}
	zt = []float64{2306.2181, 1.09468, 0.018203}
	θt = []float64{2004.3109, -0.42665, -0.041833}
)

// PositionEquinox returns rectangular coordinates referenced to an arbitrary epoch.
//
// Position will be computed for given Julian day "jde" but referenced to mean
// equinox "epoch" (year).
func PositionEquinox(e *pp.V87Planet, jde, epoch float64) (xp, yp, zp float64) {
	x0, y0, z0 := PositionJ2000(e, jde)
	t := (epoch - 2000) * .01
	ζ := base.Horner(t, ζt...) * t * math.Pi / 180 / 3600
	z := base.Horner(t, zt...) * t * math.Pi / 180 / 3600
	θ := base.Horner(t, θt...) * t * math.Pi / 180 / 3600
	sζ, cζ := math.Sincos(ζ)
	sz, cz := math.Sincos(z)
	sθ, cθ := math.Sincos(θ)
	xx := cζ*cz*cθ - sζ*sz
	xy := sζ*cz + cζ*sz*cθ
	xz := cζ * sθ
	yx := -cζ*sz - sζ*cz*cθ
	yy := cζ*cz - sζ*sz*cθ
	yz := -sζ * sθ
	zx := -cz * sθ
	zy := -sz * sθ
	zz := cθ
	return xx*x0 + yx*y0 + zx*z0,
		xy*x0 + yy*y0 + zy*z0,
		xz*x0 + yz*y0 + zz*z0
}
