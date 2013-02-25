// Copyright 2012 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Angle: Chapter 17: Angular Separation
//
// Functions in this package are useful for Ecliptic, Equatorial, or any
// similar coordinate frame.  To avoid suggestion of a particular frame,
// function parameters are specified simply as "r1, d1" to correspond to a
// right ascenscion, declination pair or to a longitude, latitude pair.
//
// All angles are in radians.
package angle

import (
	"errors"
	"math"

	"github.com/soniakeys/meeus/interp"
)

// Sep returns the angular separation between two celestial bodies.
//
// The algorithm is numerically naïve, and while patched up a bit for
// small separations, remains unstable for separations near π.
func Sep(r1, d1, r2, d2 float64) float64 {
	sd1, cd1 := math.Sincos(d1)
	sd2, cd2 := math.Sincos(d2)
	d := math.Acos(sd1*sd2 + cd1*cd2*math.Cos(r1-r2))
	// Meeus recommends 10 arc min as crossover.  0.003 rad is close.
	if d > .003 {
		return d
	}
	return math.Hypot((r2-r1)*cd1, d2-d1)
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
func MinSep(jd1, jd3 float64, r1, d1, r2, d2 []float64) (float64, error) {
	if len(r1) != 3 || len(d1) != 3 || len(r2) != 3 || len(d2) != 3 {
		return 0, interp.ErrorNot3
	}
	y := make([]float64, 3)
	for x, r := range r1 {
		y[x] = Sep(r, d1[x], r2[x], d2[x])
	}
	d3, err := interp.NewLen3(jd1, jd3, y)
	if err != nil {
		return 0, err
	}
	_, dMin, err := d3.Extremum()
	return dMin, err
}

// MinSepRect returns the minimum separation between two moving objects.
//
// Like MinSep, but using a method of rectangular coordinates that gives
// accurate results even for close approaches.
func MinSepRect(jd1, jd3 float64, r1, d1, r2, d2 []float64) (float64, error) {
	if len(r1) != 3 || len(d1) != 3 || len(r2) != 3 || len(d2) != 3 {
		return 0, interp.ErrorNot3
	}
	uv := func(r1, d1, r2, d2 float64) (u, v float64) {
		sd1, cd1 := math.Sincos(d1)
		Δr := r2 - r1
		tΔr := math.Tan(Δr)
		thΔr := math.Tan(Δr / 2)
		K := 1 / (1 + sd1*sd1*tΔr*thΔr)
		sΔd := math.Sin(d2 - d1)
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
		u, err = u3.InterpolateN(n, false)
		if err != nil {
			return 0, err
		}
		v, err = v3.InterpolateN(n, false)
		if err != nil {
			return 0, err
		}
		if math.Abs(dn) < 1e-5 {
			return math.Hypot(u, v), nil // success
		}
		up := up0 + n*up1
		vp := vp0 + n*vp1
		dn = -(u*up + v*vp) / (up*up + vp*vp)
		n += dn
	}
	return 0, errors.New("MinSepRect: failure to converge")
}

func hav(a float64) float64 {
	return .5 * (1 - math.Cos(a))
}

// SepHav returns the angular separation between two celestial bodies.
//
// The algorithm uses the haversine function and is superior to the naïve
// algorithm of the Sep function.
func SepHav(r1, d1, r2, d2 float64) float64 {
	return 2 * math.Asin(math.Sqrt(hav(d2-d1)+
		math.Cos(d1)*math.Cos(d2)*hav(r2-r1)))
}

// SepPauwels returns the angular separation between two celestial bodies.
//
// The algorithm is a numerically stable form of that used in Sep.
func SepPauwels(r1, d1, r2, d2 float64) float64 {
	sd1, cd1 := math.Sincos(d1)
	sd2, cd2 := math.Sincos(d2)
	cdr := math.Cos(r2 - r1)
	x := cd1*sd2 - sd1*cd2*cdr
	y := cd2 * math.Sin(r2-r1)
	z := sd1*sd2 + cd1*cd2*cdr
	return math.Atan2(math.Hypot(x, y), z)
}

// RelativePosition returns the position angle of one body with respect to
// another.
//
// The position angle result is measured counter-clockwise from North.
func RelativePosition(r1, d1, r2, d2 float64) float64 {
	sΔr, cΔr := math.Sincos(r2 - r1)
	sd2, cd2 := math.Sincos(d2)
	return math.Atan2(sΔr, cd2*math.Tan(d1)-sd2*cΔr)
}
