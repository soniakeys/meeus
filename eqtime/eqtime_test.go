// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package eqtime_test

import (
	"fmt"

	"github.com/soniakeys/meeus/eqtime"
	"github.com/soniakeys/meeus/julian"
	"github.com/soniakeys/sexagesimal"
)

func ExampleESmart() {
	// Example 28.b, p. 185
	eq := eqtime.ESmart(julian.CalendarGregorianToJD(1992, 10, 13))
	fmt.Printf("+%.7f rad\n", eq)
	fmt.Printf("%+.1d", sexa.HourAngle(eq).Fmt())
	// Output:
	// +0.0598256 rad
	// +13ᵐ42ˢ.7
}
