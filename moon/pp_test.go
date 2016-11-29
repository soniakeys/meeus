// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// +build !nopp

package moon_test

import (
	"fmt"
	"time"

	"github.com/soniakeys/meeus/julian"
	"github.com/soniakeys/meeus/moon"
	pp "github.com/soniakeys/meeus/planetposition"
	"github.com/soniakeys/unit"
)

func ExamplePhysical() {
	j := julian.CalendarGregorianToJD(1992, 4, 12)
	earth, err := pp.LoadPlanet(pp.Earth)
	if err != nil {
		fmt.Println(err)
		return
	}
	l, b, P, l0, b0 := moon.Physical(j, earth)
	fmt.Printf("l = %.2f\n", l.Deg())
	fmt.Printf("b = %+.2f\n", b.Deg())
	fmt.Printf("P = %.2f\n", P.Deg())
	fmt.Printf("l0 = %.2f\n", l0.Deg())
	fmt.Printf("b0 = %+.2f\n", b0.Deg())
	// Output:
	// l = -1.23
	// b = +4.20
	// P = 15.08
	// l0 = 67.90
	// b0 = +1.46
}

func ExampleSunAltitude() {
	j := julian.CalendarGregorianToJD(1992, 4, 12)
	earth, err := pp.LoadPlanet(pp.Earth)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, _, _, l0, b0 := moon.Physical(j, earth)
	h := moon.SunAltitude(
		unit.AngleFromDeg(-20), unit.AngleFromDeg(9.7), l0, b0)
	fmt.Printf("%+.3f\n", h.Deg())
	// Output:
	// +2.318
}

func ExampleSunrise() {
	earth, err := pp.LoadPlanet(pp.Earth)
	if err != nil {
		fmt.Println(err)
		return
	}
	j0 := julian.CalendarGregorianToJD(1992, 4, 15)
	j := moon.Sunrise(
		unit.AngleFromDeg(-20), unit.AngleFromDeg(9.7), j0, earth)
	y, m, d := julian.JDToCalendar(j)
	fmt.Printf("%d %s %.4f TD\n", y, time.Month(m), d)
	// Output:
	// 1992 April 11.8069 TD
}
