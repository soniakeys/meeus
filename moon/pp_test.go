// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// +build !nopp

package moon_test

import (
	"fmt"
	"math"
	"time"

	"github.com/soniakeys/meeus/julian"
	"github.com/soniakeys/meeus/moon"
	pp "github.com/soniakeys/meeus/planetposition"
)

func ExamplePhysical() {
	j := julian.CalendarGregorianToJD(1992, 4, 12)
	earth, err := pp.LoadPlanet(pp.Earth)
	if err != nil {
		fmt.Println(err)
		return
	}
	l, b, P, l0, b0 := moon.Physical(j, earth)
	fmt.Printf("l = %.2f\n", l*180/math.Pi)
	fmt.Printf("b = %+.2f\n", b*180/math.Pi)
	fmt.Printf("P = %.2f\n", P*180/math.Pi)
	fmt.Printf("l0 = %.2f\n", l0*180/math.Pi)
	fmt.Printf("b0 = %+.2f\n", b0*180/math.Pi)
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
	h := moon.SunAltitude(-20*math.Pi/180, 9.7*math.Pi/180, l0, b0)
	fmt.Printf("%+.3f\n", h*180/math.Pi)
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
	j := moon.Sunrise(-20*math.Pi/180, 9.7*math.Pi/180, j0, earth)
	y, m, d := julian.JDToCalendar(j)
	fmt.Printf("%d %s %.4f TD\n", y, time.Month(m), d)
	// Output:
	// 1992 April 11.8069 TD
}
