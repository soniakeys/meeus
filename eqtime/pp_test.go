// +build !nopp

package eqtime_test

import (
	"fmt"

	"github.com/soniakeys/meeus/eqtime"
	"github.com/soniakeys/meeus/julian"
	pp "github.com/soniakeys/meeus/planetposition"
	"github.com/soniakeys/sexagesimal"
)

func ExampleE() {
	// Example 28.a, p. 184
	earth, err := pp.LoadPlanet(pp.Earth)
	if err != nil {
		fmt.Println(err)
		return
	}
	j := julian.CalendarGregorianToJD(1992, 10, 13)
	eq := eqtime.E(j, earth)
	fmt.Printf("%+.1d", sexa.NewFmtHourAngle(eq))
	// Output:
	// +13ᵐ42ˢ.6
}
