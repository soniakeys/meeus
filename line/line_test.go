// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package line_test

import (
	"fmt"
	"math"
	"time"

	"github.com/soniakeys/meeus/julian"
	"github.com/soniakeys/meeus/line"
	"github.com/soniakeys/sexagesimal"
)

func ExampleTime() {
	// Example 19.a, p. 121.

	// convert degree data to radians
	r1 := 113.56833 * math.Pi / 180
	d1 := 31.89756 * math.Pi / 180
	r2 := 116.25042 * math.Pi / 180
	d2 := 28.03681 * math.Pi / 180
	r3 := make([]float64, 5)
	for i, ri := range []float64{
		118.98067, 119.59396, 120.20413, 120.81108, 121.41475} {
		r3[i] = ri * math.Pi / 180
	}
	d3 := make([]float64, 5)
	for i, di := range []float64{
		21.68417, 21.58983, 21.49394, 21.39653, 21.29761} {
		d3[i] = di * math.Pi / 180
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
		sexa.TimeFromDay(dFrac).Fmt())
	// Output:
	// 1994 October 1.2233
	// 1994 October 1, at 5ʰ TD(UT)
}

func ExampleAngle() {
	// Example p. 123.
	rδ := sexa.NewRA(5, 32, 0.40).Rad()
	dδ := sexa.NewAngle('-', 0, 17, 56.9).Rad()
	rε := sexa.NewRA(5, 36, 12.81).Rad()
	dε := sexa.NewAngle('-', 1, 12, 7.0).Rad()
	rζ := sexa.NewRA(5, 40, 45.52).Rad()
	dζ := sexa.NewAngle('-', 1, 56, 33.3).Rad()

	n := line.Angle(rδ, dδ, rε, dε, rζ, dζ)
	fmt.Printf("%.4f degrees\n", sexa.Angle(n).Deg())
	fmt.Printf("%m\n", sexa.Angle(n).Fmt())
	// Output:
	// 172.4830 degrees
	// 172°29′
}

func ExampleError() {
	// Example p. 124.
	rδ := sexa.NewRA(5, 32, 0.40).Rad()
	dδ := sexa.NewAngle('-', 0, 17, 56.9).Rad()
	rε := sexa.NewRA(5, 36, 12.81).Rad()
	dε := sexa.NewAngle('-', 1, 12, 7.0).Rad()
	rζ := sexa.NewRA(5, 40, 45.52).Rad()
	dζ := sexa.NewAngle('-', 1, 56, 33.3).Rad()

	ω := line.Error(rζ, dζ, rδ, dδ, rε, dε)
	e := sexa.Angle(ω)
	fmt.Printf("%.6j\n", e.Fmt())
	fmt.Printf("%.0f″\n", e.Sec())
	fmt.Println(e.Fmt())
	// Output:
	// 0°.089876
	// 324″
	// 5′24″
}

func ExampleAngleError() {
	// Example p. 125.
	rδ := sexa.NewRA(5, 32, 0.40).Rad()
	dδ := sexa.NewAngle('-', 0, 17, 56.9).Rad()
	rε := sexa.NewRA(5, 36, 12.81).Rad()
	dε := sexa.NewAngle('-', 1, 12, 7.0).Rad()
	rζ := sexa.NewRA(5, 40, 45.52).Rad()
	dζ := sexa.NewAngle('-', 1, 56, 33.3).Rad()

	n, ω := line.AngleError(rδ, dδ, rε, dε, rζ, dζ)
	fmt.Printf("%m\n", sexa.Angle(n).Fmt())
	fmt.Println(sexa.Angle(ω).Fmt())
	// Output:
	// 7°31′
	// -5′24″
}
