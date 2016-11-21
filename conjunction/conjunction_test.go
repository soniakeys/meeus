// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package conjunction_test

import (
	"fmt"
	"math"
	"time"

	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/meeus/conjunction"
	"github.com/soniakeys/meeus/deltat"
	"github.com/soniakeys/meeus/julian"
	"github.com/soniakeys/sexagesimal"
)

func ExamplePlanetary() {
	// Example 18.a, p. 117.

	// Day of month is sufficient for a time scale.
	day1 := 5.
	day5 := 9.

	// Text asks for Mercury-Venus conjunction, so r1, d1 is Venus ephemeris,
	// r2, d2 is Mercury ephemeris.

	// Venus
	r1 := []float64{
		sexa.NewRA(10, 27, 27.175).Rad(),
		sexa.NewRA(10, 26, 32.410).Rad(),
		sexa.NewRA(10, 25, 29.042).Rad(),
		sexa.NewRA(10, 24, 17.191).Rad(),
		sexa.NewRA(10, 22, 57.024).Rad(),
	}
	d1 := []float64{
		sexa.NewAngle(' ', 4, 04, 41.83).Rad(),
		sexa.NewAngle(' ', 3, 55, 54.66).Rad(),
		sexa.NewAngle(' ', 3, 48, 03.51).Rad(),
		sexa.NewAngle(' ', 3, 41, 10.25).Rad(),
		sexa.NewAngle(' ', 3, 35, 16.61).Rad(),
	}
	// Mercury
	r2 := []float64{
		sexa.NewRA(10, 24, 30.125).Rad(),
		sexa.NewRA(10, 25, 00.342).Rad(),
		sexa.NewRA(10, 25, 12.515).Rad(),
		sexa.NewRA(10, 25, 06.235).Rad(),
		sexa.NewRA(10, 24, 41.185).Rad(),
	}
	d2 := []float64{
		sexa.NewAngle(' ', 6, 26, 32.05).Rad(),
		sexa.NewAngle(' ', 6, 10, 57.72).Rad(),
		sexa.NewAngle(' ', 5, 57, 33.08).Rad(),
		sexa.NewAngle(' ', 5, 46, 27.07).Rad(),
		sexa.NewAngle(' ', 5, 37, 48.45).Rad(),
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
		sexa.TimeFromDay(dFrac).Fmt())

	// deltat func needs jd
	jd := julian.CalendarGregorianToJD(1991, 8, day)
	// compute UT = TD - ΔT, and separate back into calendar components.
	// (we could use our known calendar components, but this illustrates
	// the more general technique that would allow for rollovers.)
	y, m, d := julian.JDToCalendar(jd - deltat.Interp10A(jd)/(3600*24))
	// format as before
	dInt, dFrac = math.Modf(d)
	fmt.Printf("%d %s %d at %s UT\n", y, time.Month(m), int(dInt),
		sexa.TimeFromDay(dFrac).Fmt())

	// Δδ
	fmt.Printf("Δδ = %s\n", sexa.Angle(dd).Fmt())

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
	r2 := []float64{
		sexa.NewRA(15, 3, 51.937).Rad(),
		sexa.NewRA(15, 9, 57.327).Rad(),
		sexa.NewRA(15, 15, 37.898).Rad(),
		sexa.NewRA(15, 20, 50.632).Rad(),
		sexa.NewRA(15, 25, 32.695).Rad(),
	}
	d2 := []float64{
		sexa.NewAngle('-', 8, 57, 34.51).Rad(),
		sexa.NewAngle('-', 9, 9, 03.88).Rad(),
		sexa.NewAngle('-', 9, 17, 37.94).Rad(),
		sexa.NewAngle('-', 9, 23, 16.25).Rad(),
		sexa.NewAngle('-', 9, 26, 01.01).Rad(),
	}
	jd := julian.CalendarGregorianToJD(1996, 2, 17)
	dt := jd - base.J2000
	dy := dt / base.JulianYear
	dc := dy / 100
	fmt.Printf("%.2f years\n", dy)
	fmt.Printf("%.4f century\n", dc)

	pmr := -.649 // sec/cen
	pmd := -1.91 // sec/cen
	r1 := sexa.NewRA(15, 17, 0.421+pmr*dc).Rad()
	// Careful with quick and dirty way of applying correction to seconds
	// component before converting to radians.  The dec here is negative
	// so correction must be subtracted.  Alternative, less error-prone,
	// way would be to convert both to radians, then add.
	d1 := sexa.NewAngle('-', 9, 22, 58.54-pmd*dc).Rad()
	fmt.Printf("α′ = %.3d, δ′ = %.2d\n",
		sexa.RA(r1).Fmt(), sexa.Angle(d1).Fmt())

	day, dd, err := conjunction.Stellar(day1, day5, r1, d1, r2, d2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(sexa.Angle(dd).Fmt())
	dInt, dFrac := math.Modf(day)
	fmt.Printf("1996 February %d at %s TD\n", int(dInt),
		sexa.TimeFromDay(dFrac).Fmt())

	// Output:
	// -3.87 years
	// -0.0387 century
	// α′ = 15ʰ17ᵐ0ˢ.446, δ′ = -9°22′58″.47
	// 3′38″
	// 1996 February 18 at 6ʰ36ᵐ55ˢ TD
}
