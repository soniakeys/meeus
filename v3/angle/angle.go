// Copyright 2013 Sonia Keys
// License: MIT

// Angle: Chapter 17: Angular Separation.
//
// Functions in this package are useful for Ecliptic, Equatorial, or any
// similar coordinate frame.  To avoid suggestion of a particular frame,
// function parameters are specified simply as "r1, d1" to correspond to a
// right ascenscion, declination pair or to a longitude, latitude pair.
//
// In function Sep, Meeus recommends 10 arc min as a threshold.  This
// value is in package base as base.SmallAngle because it has general utility.
package angle

import (
	"errors"
	"math"

	"github.com/soniakeys/meeus/v3/base"
	"github.com/soniakeys/meeus/v3/interp"
	"github.com/soniakeys/unit"
)

// Sep returns the angular separation between two celestial bodies.
//
// The algorithm is numerically naïve, and while patched up a bit for
// small separations, remains unstable for separations near π.
func Sep(r1, d1, r2, d2 unit.Angle) unit.Angle {
	sd1, cd1 := d1.Sincos()
	sd2, cd2 := d2.Sincos()
	cd := sd1*sd2 + cd1*cd2*(r1-r2).Cos() // (17.1) p. 109
	if cd < base.CosSmallAngle {
		return unit.Angle(math.Acos(cd))
	}
	// (17.2) p. 109
	dm := (d1 + d2) / 2
	return unit.Angle(math.Hypot((r2-r1).Rad()*dm.Cos(), (d2 - d1).Rad()))
}

// MinSep returns the minimum separation between two moving objects.
//
// The motion is represented as an ephemeris of three rows, equally spaced
// in time.  Jd1, jd3 are julian day times of the first and last rows.
// R1, d1, r2, d2 are coordinates at the three times.  They must each be
// slices of length 3.
//
// Result is obtained by computing separation at each of the three times
// and interpolating a minimum.  This may be invalid for sufficiently close
// approaches.
func MinSep(jd1, jd3 float64, r1, d1, r2, d2 []unit.Angle) (unit.Angle, error) {
	if len(r1) != 3 || len(d1) != 3 || len(r2) != 3 || len(d2) != 3 {
		return 0, interp.ErrorNot3
	}
	y := make([]float64, 3)
	for x, r := range r1 {
		y[x] = Sep(r, d1[x], r2[x], d2[x]).Rad()
	}
	d3, err := interp.NewLen3(jd1, jd3, y)
	if err != nil {
		return 0, err
	}
	_, dMin, err := d3.Extremum()
	return unit.Angle(dMin), err
}

// MinSepRect returns the minimum separation between two moving objects.
//
// Like MinSep, but using a method of rectangular coordinates that gives
// accurate results even for close approaches.
func MinSepRect(jd1, jd3 float64, r1, d1, r2, d2 []unit.Angle) (unit.Angle, error) {
	if len(r1) != 3 || len(d1) != 3 || len(r2) != 3 || len(d2) != 3 {
		return 0, interp.ErrorNot3
	}
	uv := func(r1, d1, r2, d2 unit.Angle) (u, v float64) {
		sd1, cd1 := d1.Sincos()
		Δr := r2 - r1
		tΔr := Δr.Tan()
		thΔr := (Δr / 2).Tan()
		K := 1 / (1 + sd1*sd1*tΔr*thΔr)
		sΔd := (d2 - d1).Sin()
		u = -K * (1 - (sd1/cd1)*sΔd) * cd1 * tΔr
		v = K * (sΔd + sd1*cd1*tΔr*thΔr)
		return
	}
	us := make([]float64, 3, 6)
	vs := us[3:6]
	for x, r := range r1 {
		us[x], vs[x] = uv(r, d1[x], r2[x], d2[x])
	}
	u3, err := interp.NewLen3(-1, 1, us)
	if err != nil {
		panic(err) // bug not caller's fault.
	}
	v3, err := interp.NewLen3(-1, 1, vs)
	if err != nil {
		panic(err) // bug not caller's fault.
	}
	up0 := (us[2] - us[0]) / 2
	vp0 := (vs[2] - vs[0]) / 2
	up1 := us[0] + us[2] - 2*us[1]
	vp1 := vs[0] + vs[2] - 2*vs[1]
	up := up0
	vp := vp0
	dn := -(us[1]*up + vs[1]*vp) / (up*up + vp*vp)
	n := dn
	var u, v float64
	for limit := 0; limit < 10; limit++ {
		u = u3.InterpolateN(n)
		v = v3.InterpolateN(n)
		if math.Abs(dn) < 1e-5 {
			return unit.Angle(math.Hypot(u, v)), nil // success
		}
		up := up0 + n*up1
		vp := vp0 + n*vp1
		dn = -(u*up + v*vp) / (up*up + vp*vp)
		n += dn
	}
	return 0, errors.New("MinSepRect: failure to converge")
}

// SepHav returns the angular separation between two celestial bodies.
//
// The algorithm uses the haversine function and is superior to the naïve
// algorithm of the Sep function.
func SepHav(r1, d1, r2, d2 unit.Angle) unit.Angle {
	// using (17.5) p. 115
	return unit.Angle(2 * math.Asin(math.Sqrt(base.Hav(d2-d1)+
		d1.Cos()*d2.Cos()*base.Hav(r2-r1))))
}

// SepPauwels returns the angular separation between two celestial bodies.
//
// The algorithm is a numerically stable form of that used in Sep.
func SepPauwels(r1, d1, r2, d2 unit.Angle) unit.Angle {
	sd1, cd1 := d1.Sincos()
	sd2, cd2 := d2.Sincos()
	cdr := (r2 - r1).Cos()
	x := cd1*sd2 - sd1*cd2*cdr
	y := cd2 * (r2 - r1).Sin()
	z := sd1*sd2 + cd1*cd2*cdr
	return unit.Angle(math.Atan2(math.Hypot(x, y), z))
}

// RelativePosition returns the position angle of one body with respect to
// another.
//
// The position angle result is measured counter-clockwise from North.
func RelativePosition(r1, d1, r2, d2 unit.Angle) unit.Angle {
	sΔr, cΔr := (r1 - r2).Sincos()
	sd2, cd2 := d2.Sincos()
	return unit.Angle(math.Atan2(sΔr, cd2*d1.Tan()-sd2*cΔr))
}
