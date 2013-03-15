// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Line: Chapter 19, Bodies in Straight Line
package line

import (
	"errors"
	"math"

	"github.com/soniakeys/meeus/interp"
)

// Time computes the time at which a moving body is on a straight line (great
// circle) between two fixed points, such as stars.
//
// Coordinates may be right ascensions and declinations or longitudes and
// latitudes.  Fixed points are r1, d1, r2, d2.  Moving body is an ephemeris
// of 5 rows, r3, d3, starting at time t1 and ending at time t5.  Time scale
// is arbitrary.
//
// Result is time of alignment.
func Time(r1, d1, r2, d2 float64, r3, d3 []float64, t1, t5 float64) (float64, error) {
	if len(r3) != 5 || len(d3) != 5 {
		return 0, errors.New("r3, d3 must be length 5")
	}
	gc := make([]float64, 5)
	for i, r3i := range r3 {
		gc[i] = math.Tan(d1)*math.Sin(r2-r3i) +
			math.Tan(d2)*math.Sin(r3i-r1) +
			math.Tan(d3[i])*math.Sin(r1-r2)
	}
	l5, err := interp.NewLen5(t1, t5, gc)
	if err != nil {
		return 0, err
	}
	return l5.Zero(false)
}

// Angle returns the angle between great circles defined by three points.
//
// Coordinates may be right ascensions and declinations or longitudes and
// latitudes.  If r1, d1, r2, d2 defines one line and r2, d2, r3, d3 defines
// another, the result is the angle between the two lines.
//
// Algorithm by Meeus.
func Angle(r1, d1, r2, d2, r3, d3 float64) float64 {
	sd2, cd2 := math.Sincos(d2)
	sr21, cr21 := math.Sincos(r2 - r1)
	sr32, cr32 := math.Sincos(r3 - r2)
	C1 := math.Atan2(sr21, cd2*math.Tan(d1)-sd2*cr21)
	C2 := math.Atan2(sr32, cd2*math.Tan(d3)-sd2*cr32)
	return C1 + C2
}

// Error returns an error angle of three nearly co-linear points.
//
// For the line defined by r1, d1, r2, d2, the result is the anglular distance
// between that line and r0, d0.
//
// Algorithm by Meeus.
func Error(r1, d1, r2, d2, r0, d0 float64) float64 {
	sr1, cr1 := math.Sincos(r1)
	sd1, cd1 := math.Sincos(d1)
	sr2, cr2 := math.Sincos(r2)
	sd2, cd2 := math.Sincos(d2)
	X1 := cd1 * cr1
	X2 := cd2 * cr2
	Y1 := cd1 * sr1
	Y2 := cd2 * sr2
	Z1 := sd1
	Z2 := sd2
	A := Y1*Z2 - Z1*Y2
	B := Z1*X2 - X1*Z2
	C := X1*Y2 - Y1*X2
	m := math.Tan(r0)
	n := math.Tan(d0) / math.Cos(r0)
	return math.Asin((A + B*m + C*n) /
		(math.Sqrt(A*A+B*B+C*C) * math.Sqrt(1+m*m+n*n)))
}

// AngleError returns both an angle as in the function Angle, and an error
// as in the function Error.
//
// The algorithm is by B. Pessens.
func AngleError(r1, d1, r2, d2, r3, d3 float64) (ψ, ω float64) {
	sr1, cr1 := math.Sincos(r1)
	sd1, cd1 := math.Sincos(d1)
	sr2, cr2 := math.Sincos(r2)
	sd2, cd2 := math.Sincos(d2)
	sr3, cr3 := math.Sincos(r3)
	sd3, cd3 := math.Sincos(d3)
	a1 := cd1 * cr1
	a2 := cd2 * cr2
	a3 := cd3 * cr3
	b1 := cd1 * sr1
	b2 := cd2 * sr2
	b3 := cd3 * sr3
	c1 := sd1
	c2 := sd2
	c3 := sd3
	l1 := b1*c2 - b2*c1
	l2 := b2*c3 - b3*c2
	l3 := b1*c3 - b3*c1
	m1 := c1*a2 - c2*a1
	m2 := c2*a3 - c3*a2
	m3 := c1*a3 - c3*a1
	n1 := a1*b2 - a2*b1
	n2 := a2*b3 - a3*b2
	n3 := a1*b3 - a3*b1
	ψ = math.Acos((l1*l2 + m1*m2 + n1*n2) /
		(math.Sqrt(l1*l1+m1*m1+n1*n1) * math.Sqrt(l2*l2+m2*m2+n2*n2)))
	ω = math.Asin((a2*l3 + b2*m3 + c2*n3) /
		(math.Sqrt(a2*a2+b2*b2+c2*c2) * math.Sqrt(l3*l3+m3*m3+n3*n3)))
	return
}
