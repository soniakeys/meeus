// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package moonmaxdec_test

import (
	"fmt"
	"math"
	"time"

	"github.com/soniakeys/meeus/julian"
	"github.com/soniakeys/meeus/moonmaxdec"
	"github.com/soniakeys/sexagesimal"
)

func ExampleNorth() {
	// Example 52.a, p. 370.
	j, δ := moonmaxdec.North(1988.95)
	fmt.Printf("JDE = %.4f\n", j)
	y, m, d := julian.JDToCalendar(j)
	d, f := math.Modf(d)
	fmt.Printf("%d %s %d at %0m TD\n", y, time.Month(m), int(d),
		sexa.NewFmtTime(f*24*3600))
	fmt.Printf("δ = %.4f\n", δ*180/math.Pi)
	fmt.Printf("%+0d\n", sexa.NewFmtAngle(δ))
	// Output:
	// JDE = 2447518.3346
	// 1988 December 22 at 20ʰ02ᵐ TD
	// δ = 28.1562
	// +28°09′22″
}

func ExampleSouth() {
	// Example 52.b, p. 370.
	j, δ := moonmaxdec.South(2049.3)
	fmt.Printf("JDE = %.4f\n", j)
	y, m, d := julian.JDToCalendar(j)
	d, f := math.Modf(d)
	fmt.Printf("%d %s %d at %0h TD\n", y, time.Month(m), int(d),
		sexa.NewFmtTime(f*24*3600))
	fmt.Printf("δ = %.4f\n", δ*180/math.Pi)
	fmt.Printf("%+0m\n", sexa.NewFmtAngle(δ))
	// Output:
	// JDE = 2469553.0834
	// 2049 April 21 at 14ʰ TD
	// δ = -22.1384
	// -22°08′
}

func ExampleNorth_c() {
	// Example 52.c, p. 370.
	j, δ := moonmaxdec.North(-3.8)
	fmt.Printf("JDE = %.4f\n", j)
	y, m, d := julian.JDToCalendar(j)
	d, f := math.Modf(d)
	fmt.Printf("%d %s %d at %0h TD\n", y, time.Month(m), int(d),
		sexa.NewFmtTime(f*24*3600))
	fmt.Printf("δ = %.4f\n", δ*180/math.Pi)
	fmt.Printf("%+0m\n", sexa.NewFmtAngle(δ))
	// Output:
	// JDE = 1719672.1412
	// -4 March 16 at 15ʰ TD
	// δ = 28.9739
	// +28°58′
}
