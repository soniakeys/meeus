// Copyright 2013 Sonia Keys
// License: MIT

package line_test

import (
	"fmt"
	"math"
	"time"

	"github.com/soniakeys/meeus/v3/julian"
	"github.com/soniakeys/meeus/v3/line"
	"github.com/soniakeys/sexagesimal"
	"github.com/soniakeys/unit"
)

func ExampleTime() {
	// Example 19.a, p. 121.
	r1 := unit.AngleFromDeg(113.56833)
	d1 := unit.AngleFromDeg(31.89756)
	r2 := unit.AngleFromDeg(116.25042)
	d2 := unit.AngleFromDeg(28.03681)
	r3 := make([]unit.Angle, 5)
	for i, ri := range []float64{
		118.98067, 119.59396, 120.20413, 120.81108, 121.41475} {
		r3[i] = unit.AngleFromDeg(ri)
	}
	d3 := make([]unit.Angle, 5)
	for i, di := range []float64{
		21.68417, 21.58983, 21.49394, 21.39653, 21.29761} {
		d3[i] = unit.AngleFromDeg(di)
	}
	// use JD as time to handle month boundary
	jd, err := line.Time(r1, d1, r2, d2, r3, d3,
		julian.CalendarGregorianToJD(1994, 9, 29),
		julian.CalendarGregorianToJD(1994, 10, 3))
	if err != nil {
		fmt.Println(err)
		return
	}
	y, m, d := julian.JDToCalendar(jd)
	dInt, dFrac := math.Modf(d)
	fmt.Printf("%d %s %.4f\n", y, time.Month(m), d)
	fmt.Printf("%d %s %d, at %h TD(UT)\n", y, time.Month(m), int(dInt),
		sexa.FmtTime(unit.TimeFromDay(dFrac)))
	// Output:
	// 1994 October 1.2233
	// 1994 October 1, at 5ʰ TD(UT)
}

func ExampleAngle() {
	// Example p. 123.
	rδ := unit.NewRA(5, 32, 0.40).Angle()
	dδ := unit.NewAngle('-', 0, 17, 56.9)
	rε := unit.NewRA(5, 36, 12.81).Angle()
	dε := unit.NewAngle('-', 1, 12, 7.0)
	rζ := unit.NewRA(5, 40, 45.52).Angle()
	dζ := unit.NewAngle('-', 1, 56, 33.3)

	n := line.Angle(rδ, dδ, rε, dε, rζ, dζ)
	fmt.Printf("%.4f degrees\n", n.Deg())
	fmt.Printf("%m\n", sexa.FmtAngle(n))
	// Output:
	// 172.4830 degrees
	// 172°29′
}

func ExampleError() {
	// Example p. 124.
	rδ := unit.NewRA(5, 32, 0.40).Angle()
	dδ := unit.NewAngle('-', 0, 17, 56.9)
	rε := unit.NewRA(5, 36, 12.81).Angle()
	dε := unit.NewAngle('-', 1, 12, 7.0)
	rζ := unit.NewRA(5, 40, 45.52).Angle()
	dζ := unit.NewAngle('-', 1, 56, 33.3)

	ω := line.Error(rζ, dζ, rδ, dδ, rε, dε)
	e := sexa.FmtAngle(ω)
	fmt.Printf("%.6j\n", e)
	fmt.Printf("%.0f″\n", ω.Sec())
	fmt.Println(e)
	// Output:
	// 0°.089876
	// 324″
	// 5′24″
}

func ExampleAngleError() {
	// Example p. 125.
	rδ := unit.NewRA(5, 32, 0.40).Angle()
	dδ := unit.NewAngle('-', 0, 17, 56.9)
	rε := unit.NewRA(5, 36, 12.81).Angle()
	dε := unit.NewAngle('-', 1, 12, 7.0)
	rζ := unit.NewRA(5, 40, 45.52).Angle()
	dζ := unit.NewAngle('-', 1, 56, 33.3)

	n, ω := line.AngleError(rδ, dδ, rε, dε, rζ, dζ)
	fmt.Printf("%m\n", sexa.FmtAngle(n))
	fmt.Println(sexa.FmtAngle(ω))
	// Output:
	// 7°31′
	// -5′24″
}
