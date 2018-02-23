// Copyright 2013 Sonia Keys
// License: MIT

package apparent_test

import (
	"fmt"

	"github.com/soniakeys/meeus/v3/apparent"
	"github.com/soniakeys/meeus/v3/base"
	"github.com/soniakeys/meeus/v3/coord"
	"github.com/soniakeys/meeus/v3/julian"
	"github.com/soniakeys/sexagesimal"
	"github.com/soniakeys/unit"
)

func ExampleNutation() {
	// Example 23.a, p. 152
	α := unit.NewRA(2, 46, 11.331)
	δ := unit.NewAngle(' ', 49, 20, 54.54)
	jd := julian.CalendarGregorianToJD(2028, 11, 13.19)
	Δα1, Δδ1 := apparent.Nutation(α, δ, jd)
	fmt.Printf("%.3s  %.3s\n",
		sexa.FmtAngle(unit.Angle(Δα1)), // (Δα1 is in HourAngle)
		sexa.FmtAngle(Δδ1))
	// Output:
	// 15.843″  6.217″
}

func ExampleAberration() {
	// Example 23.a, p. 152
	α := unit.NewRA(2, 46, 11.331)
	δ := unit.NewAngle(' ', 49, 20, 54.54)
	jd := julian.CalendarGregorianToJD(2028, 11, 13.19)
	Δα2, Δδ2 := apparent.Aberration(α, δ, jd)
	fmt.Printf("%.3s  %.3s\n",
		sexa.FmtAngle(unit.Angle(Δα2)), // (Δα2 is in HourAngle)
		sexa.FmtAngle(Δδ2))
	// Output:
	// 30.045″  6.697″
}

func ExamplePosition() {
	// Example 23.a, p. 152
	jd := julian.CalendarGregorianToJD(2028, 11, 13.19)
	eq := &coord.Equatorial{
		unit.NewRA(2, 44, 11.986),
		unit.NewAngle(' ', 49, 13, 42.48),
	}
	apparent.Position(eq, eq, 2000, base.JDEToJulianYear(jd),
		unit.HourAngleFromSec(.03425),
		unit.AngleFromSec(-.0895))
	fmt.Printf("α = %0.3d\n", sexa.FmtRA(eq.RA))
	fmt.Printf("δ = %0.2d\n", sexa.FmtAngle(eq.Dec))
	// Output:
	// α = 2ʰ46ᵐ14ˢ.390
	// δ = 49°21′07″.45
}

func ExampleAberrationRonVondrak() {
	// Example 23.b, p. 156
	α := unit.NewRA(2, 44, 12.9747)
	δ := unit.NewAngle(' ', 49, 13, 39.896)
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
		RA:  unit.NewRA(2, 44, 11.986),
		Dec: unit.NewAngle(' ', 49, 13, 42.48),
	}
	apparent.PositionRonVondrak(eq, eq, base.JDEToJulianYear(jd),
		unit.HourAngleFromSec(.03425),
		unit.AngleFromSec(-.0895))
	fmt.Printf("α = %0.3d\n", sexa.FmtRA(eq.RA))
	fmt.Printf("δ = %0.2d\n", sexa.FmtAngle(eq.Dec))
	// Output:
	// α = 2ʰ46ᵐ14ˢ.392
	// δ = 49°21′07″.45
}
