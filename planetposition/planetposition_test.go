// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package planetposition_test

import (
	"fmt"
	"math"

	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/meeus/julian"
	pp "github.com/soniakeys/meeus/planetposition"
)

func ExampleV87Planet_Position2000() {
	// Mars 1899 spherical data from vsop87.chk.
	jd := 2415020.0
	p, err := pp.LoadPlanet(pp.Mars, "")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, b, r := p.Position2000(jd)
	fmt.Printf("L = %.10f rad\n", l)
	fmt.Printf("B = %.10f rad\n", b)
	fmt.Printf("R = %.10f au\n", r)
	// Output:
	// L = 5.0185792656 rad
	// B = -0.0274073500 rad
	// R = 1.4218777718 au
}

func ExampleV87Planet_Position() {
	// Example 32.a, p. 219
	jd := julian.CalendarGregorianToJD(1992, 12, 20)
	p, err := pp.LoadPlanet(pp.Venus, "")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, b, r := p.Position(jd)
	fmt.Printf("L = %s\n",
		base.DecSymAdd(fmt.Sprintf("%+.5f", l*180/math.Pi), '°'))
	fmt.Printf("B = %s\n",
		base.DecSymAdd(fmt.Sprintf("%+.5f", b*180/math.Pi), '°'))
	fmt.Printf("R = %.6f AU\n", r)
	// Meeus results:
	// L = +26°.11428
	// B = -2°.62070
	// R = 0.724603 AU
	// Answers below seem close enough.

	// Output:
	// L = +26°.11412
	// B = -2°.62060
	// R = 0.724602 AU
}
