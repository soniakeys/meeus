// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package apparent_test

import (
	"fmt"

	"github.com/soniakeys/meeus/apparent"
	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/meeus/coord"
	"github.com/soniakeys/meeus/julian"
)

func ExampleNutation() {
	α := base.NewRA(2, 46, 11.331).Rad()
	δ := base.NewAngle(false, 49, 20, 54.54).Rad()
	jd := julian.CalendarGregorianToJD(2028, 11, 13.19)
	Δα1, Δδ1 := apparent.Nutation(α, δ, jd)
	fmt.Printf("%.3s  %.3s\n", base.NewFmtAngle(Δα1), base.NewFmtAngle(Δδ1))
	// Output:
	// 15.843″  6.217″
}

func ExampleAbberation() {
	α := base.NewRA(2, 46, 11.331).Rad()
	δ := base.NewAngle(false, 49, 20, 54.54).Rad()
	jd := julian.CalendarGregorianToJD(2028, 11, 13.19)
	Δα2, Δδ2 := apparent.Abberation(α, δ, jd)
	fmt.Printf("%.3s  %.3s\n", base.NewFmtAngle(Δα2), base.NewFmtAngle(Δδ2))
	// Output:
	// 30.045″  6.697″
}

func ExamplePosition() {
	jd := julian.CalendarGregorianToJD(2028, 11, 13.19)
	eq := &coord.Equatorial{
		base.NewRA(2, 44, 11.986).Rad(),
		base.NewAngle(false, 49, 13, 42.48).Rad(),
	}
	apparent.Position(eq, eq, 2000, base.JDEToJulianYear(jd),
		base.NewHourAngle(false, 0, 0, 0.03425).Rad(),
		base.NewAngle(true, 0, 0, 0.0895).Rad())
	fmt.Printf("α = %0.3d\n", base.NewFmtRA(eq.RA))
	fmt.Printf("δ = %0.2d\n", base.NewFmtAngle(eq.Dec))
	// Output:
	// α = 2ʰ46ᵐ14ˢ.390
	// δ = 49°21′07″.45
}
