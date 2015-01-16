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
	α := sexa.NewRA(2, 46, 11.331).Rad()
	δ := sexa.NewAngle(false, 49, 20, 54.54).Rad()
	jd := julian.CalendarGregorianToJD(2028, 11, 13.19)
	Δα1, Δδ1 := apparent.Nutation(α, δ, jd)
	fmt.Printf("%.3s  %.3s\n", sexa.NewFmtAngle(Δα1), sexa.NewFmtAngle(Δδ1))
	// Output:
	// 15.843″  6.217″
}

func ExampleAberration() {
	// Example 23.a, p. 152
	α := sexa.NewRA(2, 46, 11.331).Rad()
	δ := sexa.NewAngle(false, 49, 20, 54.54).Rad()
	jd := julian.CalendarGregorianToJD(2028, 11, 13.19)
	Δα2, Δδ2 := apparent.Aberration(α, δ, jd)
	fmt.Printf("%.3s  %.3s\n", sexa.NewFmtAngle(Δα2), sexa.NewFmtAngle(Δδ2))
	// Output:
	// 30.045″  6.697″
}

func ExamplePosition() {
	// Example 23.a, p. 152
	jd := julian.CalendarGregorianToJD(2028, 11, 13.19)
	eq := &coord.Equatorial{
		sexa.NewRA(2, 44, 11.986).Rad(),
		sexa.NewAngle(false, 49, 13, 42.48).Rad(),
	}
	apparent.Position(eq, eq, 2000, base.JDEToJulianYear(jd),
		sexa.NewHourAngle(false, 0, 0, 0.03425),
		sexa.NewAngle(true, 0, 0, 0.0895))
	fmt.Printf("α = %0.3d\n", sexa.NewFmtRA(eq.RA))
	fmt.Printf("δ = %0.2d\n", sexa.NewFmtAngle(eq.Dec))
	// Output:
	// α = 2ʰ46ᵐ14ˢ.390
	// δ = 49°21′07″.45
}

func ExampleAberrationRonVondrak() {
	// Example 23.b, p. 156
	α := sexa.NewRA(2, 44, 12.9747).Rad()
	δ := sexa.NewAngle(false, 49, 13, 39.896).Rad()
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
		RA:  sexa.NewRA(2, 44, 11.986).Rad(),
		Dec: sexa.NewAngle(false, 49, 13, 42.48).Rad(),
	}
	apparent.PositionRonVondrak(eq, eq, base.JDEToJulianYear(jd),
		sexa.NewHourAngle(false, 0, 0, 0.03425),
		sexa.NewAngle(true, 0, 0, 0.0895))
	fmt.Printf("α = %0.3d\n", sexa.NewFmtRA(eq.RA))
	fmt.Printf("δ = %0.2d\n", sexa.NewFmtAngle(eq.Dec))
	// Output:
	// α = 2ʰ46ᵐ14ˢ.392
	// δ = 49°21′07″.45
}
