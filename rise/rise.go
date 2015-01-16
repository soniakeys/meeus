// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Rise: Chapter 15, Rising, Transit, and Setting.
package rise

import (
	"errors"
	"math"

	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/meeus/globe"
	"github.com/soniakeys/meeus/interp"
	"github.com/soniakeys/sexagesimal"
)

var meanRefraction = sexa.NewAngle(false, 0, 34, 0).Rad()

// "Standard altitudes" for various bodies.
//
// The standard altitude is the geometric altitude of the center of body
// at the time of apparent rising or setting.
var (
	Stdh0Stellar   = sexa.NewAngle(true, 0, 34, 0).Rad()
	Stdh0Solar     = sexa.NewAngle(true, 0, 50, 0).Rad()
	Stdh0LunarMean = sexa.NewAngle(false, 0, 0, .125).Rad()
)

// Stdh0Lunar is the standard altitude of the Moon considering π, the
// Moon's horizontal parallax.
func Stdh0Lunar(π float64) float64 {
	return sexa.NewAngle(false, 0, 0, .7275).Rad()*π - meanRefraction
}

// ErrorCircumpolar returned by Times when the object does not rise and
// set on the day of interest.
var ErrorCircumpolar = errors.New("Circumpolar")

// ApproxTimes computes approximate UT rise, transit and set times for
// a celestial object on a day of interest.
//
// The function argurments do not actually include the day, but do include
// values computed from the day.
//
//	p is geographic coordinates of observer.
//	h0 is "standard altitude" of the body.
//	Th0 is apparent sidereal time at 0h UT at Greenwich.
//	α, δ are right ascension and declination of the body.
//
// h0 unit is radians.
//
// Th0 must be the time on the day of interest, in seconds.
// See sidereal.Apparent0UT.
//
// α, δ must be values at 0h dynamical time for the day of interest.
// Units are radians.
//
// Result units are seconds and are in the range [0,86400)
func ApproxTimes(p globe.Coord, h0, Th0 float64, α, δ float64) (mRise, mTransit, mSet float64, err error) {
	// Meeus works in a crazy mix of units.
	// This function and Times work with seconds of time as much as possible.

	// approximate local hour angle
	sLat, cLat := math.Sincos(p.Lat)
	sδ1, cδ1 := math.Sincos(δ)
	cH0 := (math.Sin(h0) - sLat*sδ1) / (cLat * cδ1) // (15.1) p. 102
	if cH0 < -1 || cH0 > 1 {
		err = ErrorCircumpolar
		return
	}
	H0 := math.Acos(cH0) * 43200 / math.Pi

	// approximate transit, rise, set times.
	// (15.2) p. 102.
	mt := (α+p.Lon)*43200/math.Pi - Th0
	mTransit = base.PMod(mt, 86400)
	mRise = base.PMod(mt-H0, 86400)
	mSet = base.PMod(mt+H0, 86400)
	return
}

// Times computes UT rise, transit and set times for a celestial object on
// a day of interest.
//
// The function argurments do not actually include the day, but do include
// a number of values computed from the day.
//
//	p is geographic coordinates of observer.
//	ΔT is delta T.
//	h0 is "standard altitude" of the body.
//	Th0 is apparent sidereal time at 0h UT at Greenwich.
//	α3, δ3 are slices of three right ascensions and declinations.
//
// ΔT unit is seconds.  See package deltat.
//
// h0 unit is radians.
//
// Th0 must be the time on the day of interest, in seconds.
// See sidereal.Apparent0UT.
//
// α3, δ3 must be values at 0h dynamical time for the day before, the day of,
// and the day after the day of interest.  Units are radians.
//
// Result units are seconds and are in the range [0,86400)
func Times(p globe.Coord, ΔT, h0, Th0 float64, α3, δ3 []float64) (mRise, mTransit, mSet float64, err error) {
	mRise, mTransit, mSet, err = ApproxTimes(p, h0, Th0, α3[1], δ3[1])
	if err != nil {
		return
	}
	var d3α, d3δ *interp.Len3
	d3α, err = interp.NewLen3(-86400, 86400, α3)
	if err != nil {
		return
	}
	d3δ, err = interp.NewLen3(-86400, 86400, δ3)
	if err != nil {
		return
	}
	// adjust mTransit
	{
		th0 := base.PMod(Th0+mTransit*360.985647/360, 86400)
		α := d3α.InterpolateX(mTransit + ΔT)
		H := th0 - (p.Lon+α)*43200/math.Pi
		mTransit -= H
	}
	// adjust mRise, mSet
	sLat, cLat := math.Sincos(p.Lat)
	adjustRS := func(m float64) (float64, error) {
		th0 := base.PMod(Th0+m*360.985647/360, 86400)
		ut := m + ΔT
		α := d3α.InterpolateX(ut)
		δ := d3δ.InterpolateX(ut)
		H := th0 - (p.Lon+α)*43200/math.Pi
		sδ, cδ := math.Sincos(δ)
		h := sLat*sδ + cLat*cδ*math.Cos(H)
		return m + (h-h0)/cδ*cLat*math.Sin(H), nil
	}
	mRise, err = adjustRS(mRise)
	if err != nil {
		return
	}
	mSet, err = adjustRS(mSet)
	return
}
