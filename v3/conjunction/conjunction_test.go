// Copyright 2013 Sonia Keys
// License: MIT

package conjunction_test

import (
	"fmt"
	"math"
	"time"

	"github.com/soniakeys/meeus/v3/base"
	"github.com/soniakeys/meeus/v3/conjunction"
	"github.com/soniakeys/meeus/v3/deltat"
	"github.com/soniakeys/meeus/v3/julian"
	"github.com/soniakeys/sexagesimal"
	"github.com/soniakeys/unit"
)

func ExamplePlanetary() {
	// Example 18.a, p. 117.

	// Day of month is sufficient for a time scale.
	day1 := 5.
	day5 := 9.

	// Text asks for Mercury-Venus conjunction, so r1, d1 is Venus ephemeris,
	// r2, d2 is Mercury ephemeris.

	// Venus
	r1 := []unit.Angle{
		unit.NewRA(10, 27, 27.175).Angle(),
		unit.NewRA(10, 26, 32.410).Angle(),
		unit.NewRA(10, 25, 29.042).Angle(),
		unit.NewRA(10, 24, 17.191).Angle(),
		unit.NewRA(10, 22, 57.024).Angle(),
	}
	d1 := []unit.Angle{
		unit.NewAngle(' ', 4, 04, 41.83),
		unit.NewAngle(' ', 3, 55, 54.66),
		unit.NewAngle(' ', 3, 48, 03.51),
		unit.NewAngle(' ', 3, 41, 10.25),
		unit.NewAngle(' ', 3, 35, 16.61),
	}
	// Mercury
	r2 := []unit.Angle{
		unit.NewRA(10, 24, 30.125).Angle(),
		unit.NewRA(10, 25, 00.342).Angle(),
		unit.NewRA(10, 25, 12.515).Angle(),
		unit.NewRA(10, 25, 06.235).Angle(),
		unit.NewRA(10, 24, 41.185).Angle(),
	}
	d2 := []unit.Angle{
		unit.NewAngle(' ', 6, 26, 32.05),
		unit.NewAngle(' ', 6, 10, 57.72),
		unit.NewAngle(' ', 5, 57, 33.08),
		unit.NewAngle(' ', 5, 46, 27.07),
		unit.NewAngle(' ', 5, 37, 48.45),
	}
	// compute conjunction
	day, dd, err := conjunction.Planetary(day1, day5, r1, d1, r2, d2)
	if err != nil {
		fmt.Println(err)
		return
	}
	// time of conjunction
	fmt.Printf("1991 August %.5f\n", day)

	// more useful clock format
	dInt, dFrac := math.Modf(day)
	fmt.Printf("1991 August %d at %s TD\n", int(dInt),
		sexa.FmtTime(unit.TimeFromDay(dFrac)))

	// deltat func needs jd
	jd := julian.CalendarGregorianToJD(1991, 8, day)
	// compute UT = TD - ΔT, and separate back into calendar components.
	// (we could use our known calendar components, but this illustrates
	// the more general technique that would allow for rollovers.)
	y, m, d := julian.JDToCalendar(jd - deltat.Interp10A(jd).Day())
	// format as before
	dInt, dFrac = math.Modf(d)
	fmt.Printf("%d %s %d at %s UT\n", y, time.Month(m), int(dInt),
		sexa.FmtTime(unit.TimeFromDay(dFrac)))

	// Δδ
	fmt.Printf("Δδ = %s\n", sexa.FmtAngle(dd))

	// Output:
	// 1991 August 7.23797
	// 1991 August 7 at 5ʰ42ᵐ41ˢ TD
	// 1991 August 7 at 5ʰ41ᵐ43ˢ UT
	// Δδ = 2°8′22″
}

func ExampleStellar() {
	// Exercise, p. 119.
	day1 := 7.
	day5 := 27.
	r2 := []unit.Angle{
		unit.NewRA(15, 3, 51.937).Angle(),
		unit.NewRA(15, 9, 57.327).Angle(),
		unit.NewRA(15, 15, 37.898).Angle(),
		unit.NewRA(15, 20, 50.632).Angle(),
		unit.NewRA(15, 25, 32.695).Angle(),
	}
	d2 := []unit.Angle{
		unit.NewAngle('-', 8, 57, 34.51),
		unit.NewAngle('-', 9, 9, 03.88),
		unit.NewAngle('-', 9, 17, 37.94),
		unit.NewAngle('-', 9, 23, 16.25),
		unit.NewAngle('-', 9, 26, 01.01),
	}
	jd := julian.CalendarGregorianToJD(1996, 2, 17)
	dt := jd - base.J2000
	dy := dt / base.JulianYear
	dc := dy / 100
	fmt.Printf("%.2f years\n", dy)
	fmt.Printf("%.4f century\n", dc)

	pmr := -.649 // sec/cen
	pmd := -1.91 // sec/cen
	r1 := unit.NewRA(15, 17, 0.421) + unit.RAFromSec(pmr*dc)
	d1 := unit.NewAngle('-', 9, 22, 58.54) + unit.AngleFromSec(pmd*dc)
	fmt.Printf("α′ = %.3d, δ′ = %.2d\n", sexa.FmtRA(r1), sexa.FmtAngle(d1))

	day, dd, err := conjunction.Stellar(day1, day5, r1.Angle(), d1, r2, d2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(sexa.FmtAngle(dd))
	dInt, dFrac := math.Modf(day)
	fmt.Printf("1996 February %d at %s TD\n", int(dInt),
		sexa.FmtTime(unit.TimeFromDay(dFrac)))

	// Output:
	// -3.87 years
	// -0.0387 century
	// α′ = 15ʰ17ᵐ0ˢ.446, δ′ = -9°22′58″.47
	// 3′38″
	// 1996 February 18 at 6ʰ36ᵐ55ˢ TD
}
