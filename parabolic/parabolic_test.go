// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package parabolic_test

import (
	"fmt"

	"github.com/soniakeys/meeus/julian"
	"github.com/soniakeys/meeus/parabolic"
)

func ExampleElements_AnomalyDistance() {
	// Example 34.a, p. 243
	e := &parabolic.Elements{
		TimeP: julian.CalendarGregorianToJD(1998, 4, 14.4358),
		PDis:  1.487469,
	}
	j := julian.CalendarGregorianToJD(1998, 8, 5)
	ν, r := e.AnomalyDistance(j)
	fmt.Printf("%.5f deg\n", ν.Deg())
	fmt.Printf("%.6f AU\n", r)
	// Output:
	// 66.78862 deg
	// 2.133911 AU
}
