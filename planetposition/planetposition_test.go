// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package planetposition_test

import (
	"fmt"
	"math"

	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/meeus/julian"
	"github.com/soniakeys/meeus/planetposition"
)

func ExampleVSOP87() {
	// Example 32.a, p. 219
	l, b, r := planetposition.VSOP87(planetposition.Venus,
		julian.CalendarGregorianToJD(1992, 12, 20))
	fmt.Printf("L = %s\n",
		base.DecSymAdd(fmt.Sprintf("%+.5f", l*180/math.Pi), '°'))
	fmt.Printf("B = %s\n",
		base.DecSymAdd(fmt.Sprintf("%+.5f", b*180/math.Pi), '°'))
	fmt.Printf("R = %.6f AU\n", r)
	// Output:
	// L = +26°.11428
	// B = -2°.62070
	// R = 0.724603 AU
}
