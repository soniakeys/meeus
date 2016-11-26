// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package apparent_test

import (
	"fmt"

	"github.com/soniakeys/meeus/apparent"
	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/meeus/coord"
	"github.com/soniakeys/meeus/julian"
	"github.com/soniakeys/sexagesimal"
)

func ExampleNutation() {
	// Example 23.a, p. 152
	α := base.NewRA(2, 46, 11.331)
	δ := base.NewAngle(' ', 49, 20, 54.54)
	jd := julian.CalendarGregorianToJD(2028, 11, 13.19)
	Δα1, Δδ1 := apparent.Nutation(α, δ, jd)
	fmt.Printf("%.3s  %.3s\n", sexa.Angle(Δα1).Fmt(), sexa.Angle(Δδ1).Fmt())
	// Output:
	// 15.843″  6.217″
}

func ExampleAberration() {
	// Example 23.a, p. 152
	α := base.NewRA(2, 46, 11.331)
	δ := base.NewAngle(' ', 49, 20, 54.54)
	jd := julian.CalendarGregorianToJD(2028, 11, 13.19)
	Δα2, Δδ2 := apparent.Aberration(α, δ, jd)
	fmt.Printf("%.3s  %.3s\n", sexa.Angle(Δα2).Fmt(), sexa.Angle(Δδ2).Fmt())
	// Output:
	// 30.045″  6.697″
}

func ExamplePosition() {
	// Example 23.a, p. 152
	jd := julian.CalendarGregorianToJD(2028, 11, 13.19)
	eq := &coord.Equatorial{
		base.NewRA(2, 44, 11.986),
		base.NewAngle(' ', 49, 13, 42.48),
	}
	apparent.Position(eq, eq, 2000, base.JDEToJulianYear(jd),
		base.HourAngleFromSec(.03425),
		base.AngleFromSec(-.0895))
	fmt.Printf("α = %0.3d\n", sexa.RA(eq.RA).Fmt())
	fmt.Printf("δ = %0.2d\n", sexa.Angle(eq.Dec).Fmt())
	// Output:
	// α = 2ʰ46ᵐ14ˢ.390
	// δ = 49°21′07″.45
}

func ExampleAberrationRonVondrak() {
	// Example 23.b, p. 156
	α := base.NewRA(2, 44, 12.9747)
	δ := base.NewAngle(' ', 49, 13, 39.896)
	jd := julian.CalendarGregorianToJD(2028, 11, 13.19)
	Δα, Δδ := apparent.AberrationRonVondrak(α, δ, jd)
	fmt.Printf("Δα = %+.9f radian\n", Δα)
	fmt.Printf("Δδ = %+.9f radian\n", Δδ)
	// Output:
	// Δα = +0.000145252 radian
	// Δδ = +0.000032723 radian
}

func ExamplePositionRonVondrak() {
	// Example 23.b, p. 156
	jd := julian.CalendarGregorianToJD(2028, 11, 13.19)
	eq := &coord.Equatorial{
		RA:  base.NewRA(2, 44, 11.986),
		Dec: base.NewAngle(' ', 49, 13, 42.48),
	}
	apparent.PositionRonVondrak(eq, eq, base.JDEToJulianYear(jd),
		base.HourAngleFromSec(.03425),
		base.AngleFromSec(-.0895))
	fmt.Printf("α = %0.3d\n", sexa.RA(eq.RA).Fmt())
	fmt.Printf("δ = %0.2d\n", sexa.Angle(eq.Dec).Fmt())
	// Output:
	// α = 2ʰ46ᵐ14ˢ.392
	// δ = 49°21′07″.45
}
