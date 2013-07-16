package eqtime_test

import (
	"fmt"

	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/meeus/eqtime"
	"github.com/soniakeys/meeus/julian"
	pp "github.com/soniakeys/meeus/planetposition"
)

func ExampleE() {
	// Example 28.a, p. 184
	earth, err := pp.LoadPlanet(pp.Earth, "")
	if err != nil {
		fmt.Println(err)
		return
	}
	j := julian.CalendarGregorianToJD(1992, 10, 13)
	eq := eqtime.E(j, earth)
	fmt.Printf("%+.1d", base.NewFmtHourAngle(eq))
	// Output:
	// +13ᵐ42ˢ.6
}

func ExampleESmart() {
	// Example 28.b, p. 185
	eq := eqtime.ESmart(julian.CalendarGregorianToJD(1992, 10, 13))
	fmt.Printf("+%.7f rad\n", eq)
	fmt.Printf("%+.1d", base.NewFmtHourAngle(eq))
	// Output:
	// +0.0598256 rad
	// +13ᵐ42ˢ.7
}
