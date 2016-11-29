// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Rise: Chapter 15, Rising, Transit, and Setting.
//
// Formulas in the chapter are general enough to handle various astronomical
// objects.  The methods ApproxTimes and Times implement this general math.
//
// The function signatures aren't very friendly though, requiring a number of
// precomputed values.  The example worked in the text gives these values for
// the planet Venus.  With these example values as test data, methods
// ApproxPlanet and Planet are also given here.  Similar methods for stars,
// the Sun, Moon, Pluto, or asteroids might also be developed using other
// packages from this library.
package rise

import (
	"errors"
	"math"

	"github.com/soniakeys/meeus/deltat"
	"github.com/soniakeys/meeus/elliptic"
	"github.com/soniakeys/meeus/globe"
	"github.com/soniakeys/meeus/interp"
	"github.com/soniakeys/meeus/julian"
	pp "github.com/soniakeys/meeus/planetposition"
	"github.com/soniakeys/meeus/sidereal"
	"github.com/soniakeys/unit"
)

var meanRefraction = unit.AngleFromMin(34)

// "Standard altitudes" for various bodies.
//
// The standard altitude is the geometric altitude of the center of body
// at the time of apparent rising or setting.
var (
	Stdh0Stellar   = unit.AngleFromMin(-34)
	Stdh0Solar     = unit.AngleFromMin(-50)
	Stdh0LunarMean = unit.AngleFromDeg(.125)
)

// Stdh0Lunar is the standard altitude of the Moon considering π, the
// Moon's horizontal parallax.
func Stdh0Lunar(π unit.Angle) unit.Angle {
	return π.Mul(.7275) - meanRefraction
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
// Th0 must be the time on the day of interest.
// See sidereal.Apparent0UT.
//
// α, δ must be values at 0h dynamical time for the day of interest.
func ApproxTimes(p globe.Coord, h0 unit.Angle, Th0 unit.Time, α unit.RA, δ unit.Angle) (tRise, tTransit, tSet unit.Time, err error) {
	// approximate local hour angle
	sLat, cLat := p.Lat.Sincos()
	sδ1, cδ1 := δ.Sincos()
	cH0 := (h0.Sin() - sLat*sδ1) / (cLat * cδ1) // (15.1) p. 102
	if cH0 < -1 || cH0 > 1 {
		err = ErrorCircumpolar
		return
	}
	H0 := unit.TimeFromRad(math.Acos(cH0))

	// approximate transit, rise, set times.
	// (15.2) p. 102.
	mt := unit.TimeFromRad(α.Rad()+p.Lon.Rad()) - Th0
	tTransit = mt.Mod1()
	tRise = (mt - H0).Mod1()
	tSet = (mt + H0).Mod1()
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
// h0 unit is radians.
//
// Th0 must be the time on the day of interest, in seconds.
// See sidereal.Apparent0UT.
//
// α3, δ3 must be values at 0h dynamical time for the day before, the day of,
// and the day after the day of interest.  Units are radians.
//
// Result units are seconds of day and are in the range [0,86400).
func Times(p globe.Coord, ΔT unit.Time, h0 unit.Angle, Th0 unit.Time, α3 []unit.RA, δ3 []unit.Angle) (tRise, tTransit, tSet unit.Time, err error) {
	tRise, tTransit, tSet, err = ApproxTimes(p, h0, Th0, α3[1], δ3[1])
	if err != nil {
		return
	}
	αf := make([]float64, 3)
	for i, α := range α3 {
		αf[i] = α.Rad()
	}
	δf := make([]float64, 3)
	for i, δ := range δ3 {
		δf[i] = δ.Rad()
	}
	var d3α, d3δ *interp.Len3
	d3α, err = interp.NewLen3(-86400, 86400, αf)
	if err != nil {
		return
	}
	d3δ, err = interp.NewLen3(-86400, 86400, δf)
	if err != nil {
		return
	}
	// adjust tTransit
	{
		th0 := (Th0 + tTransit.Mul(360.985647/360)).Mod1()
		α := d3α.InterpolateX((tTransit + ΔT).Sec())
		// local hour angle as Time
		H := th0 - unit.TimeFromRad(p.Lon.Rad()+α)
		tTransit -= H
	}
	// adjust tRise, tSet
	sLat, cLat := p.Lat.Sincos()
	adjustRS := func(m unit.Time) (unit.Time, error) {
		th0 := (Th0 + m.Mul(360.985647/360)).Mod1()
		ut := (m + ΔT).Sec()
		α := d3α.InterpolateX(ut)
		δ := d3δ.InterpolateX(ut)
		Hrad := th0.Rad() - p.Lon.Rad() - α
		sδ, cδ := math.Sincos(δ)
		sH, cH := math.Sincos(Hrad)
		h := math.Asin(sLat*sδ + cLat*cδ*cH)
		md := (unit.TimeFromRad(h) - h0.Time()).Div(cδ * cLat * sH)
		return m + md, nil
	}
	tRise, err = adjustRS(tRise)
	if err != nil {
		return
	}
	tSet, err = adjustRS(tSet)
	return
}

// ApproxPlanet computes approximate UT rise, transit and set times for
// a planet on a day of interest.
//
//  yr, mon, day are the Gregorian date.
//  pos is geographic coordinates of observer.
//  e must be a V87Planet object for Earth
//  pl must be a V87Planet object for another planet.
//
// Obtain V87Planet objects with the planetposition package.
//
// Result units are seconds of day and are in the range [0,86400).
func ApproxPlanet(yr, mon, day int, pos globe.Coord, e, pl *pp.V87Planet) (tRise, tTransit, tSet unit.Time, err error) {
	jd := julian.CalendarGregorianToJD(yr, mon, float64(day))
	α, δ := elliptic.Position(pl, e, jd)
	return ApproxTimes(pos, Stdh0Stellar, sidereal.Apparent0UT(jd), α, δ)
}

// Planet computes UT rise, transit and set times for a planet on a day of
// interest.
//
//  yr, mon, day are the Gregorian date.
//  pos is geographic coordinates of observer.
//  e must be a V87Planet object for Earth
//  pl must be a V87Planet object for another planet.
//
// Obtain V87Planet objects with the planetposition package.
//
// Result units are seconds of day and are in the range [0,86400).
func Planet(yr, mon, day int, pos globe.Coord, e, pl *pp.V87Planet) (tRise, tTransit, tSet unit.Time, err error) {
	jd := julian.CalendarGregorianToJD(yr, mon, float64(day))
	α := make([]unit.RA, 3)
	δ := make([]unit.Angle, 3)
	α[0], δ[0] = elliptic.Position(pl, e, jd-1)
	α[1], δ[1] = elliptic.Position(pl, e, jd)
	α[2], δ[2] = elliptic.Position(pl, e, jd+1)
	return Times(pos, deltat.Interp10A(jd), Stdh0Stellar,
		sidereal.Apparent0UT(jd), α, δ)
}
