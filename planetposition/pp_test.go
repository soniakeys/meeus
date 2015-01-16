// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// +build !nopp

package planetposition_test

import (
	"fmt"

	"github.com/soniakeys/meeus/julian"
	pp "github.com/soniakeys/meeus/planetposition"
	"github.com/soniakeys/sexagesimal"
)

func ExampleV87Planet_Position2000() {
	// Mars 1899 spherical data from vsop87.chk.
	jd := 2415020.0
	p, err := pp.LoadPlanet(pp.Mars)
	if err != nil {
		fmt.Println(err)
		return
	}
	l, b, r := p.Position2000(jd)
	fmt.Printf("L = %.10f rad\n", l)
	fmt.Printf("B = %.10f rad\n", b)
	fmt.Printf("R = %.10f AU\n", r)
	// Output:
	// L = 5.0185792656 rad
	// B = -0.0274073500 rad
	// R = 1.4218777718 AU
}

func ExampleV87Planet_Position() {
	// Example 32.a, p. 219
	jd := julian.CalendarGregorianToJD(1992, 12, 20)
	p, err := pp.LoadPlanet(pp.Venus)
	if err != nil {
		fmt.Println(err)
		return
	}
	l, b, r := p.Position(jd)
	fmt.Printf("L = %+.5j\n", sexa.NewFmtAngle(l))
	fmt.Printf("B = %+.5j\n", sexa.NewFmtAngle(b))
	fmt.Printf("R = %.6f AU\n", r)
	// Output:
	// L = +26°.11412
	// B = -2°.62060
	// R = 0.724602 AU
}
