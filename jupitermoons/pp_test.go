// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// +build !nopp

package jupitermoons_test

import (
	"fmt"

	"github.com/soniakeys/meeus/deltat"
	"github.com/soniakeys/meeus/julian"
	"github.com/soniakeys/meeus/jupitermoons"
	pp "github.com/soniakeys/meeus/planetposition"
	"github.com/soniakeys/sexagesimal"
	"github.com/soniakeys/unit"
)

func ExampleE5() {
	// Example 44.b, p. 314.
	e, err := pp.LoadPlanet(pp.Earth)
	if err != nil {
		fmt.Println(err)
		return
	}
	j, err := pp.LoadPlanet(pp.Jupiter)
	if err != nil {
		fmt.Println(err)
		return
	}
	var pos [4]jupitermoons.XY
	jupitermoons.E5(2448972.50068, e, j, &pos)
	fmt.Printf("X  %+.4f  %+.4f  %+.4f  %+.4f\n",
		pos[0].X, pos[1].X, pos[2].X, pos[3].X)
	fmt.Printf("Y  %+.4f  %+.4f  %+.4f  %+.4f\n",
		pos[0].Y, pos[1].Y, pos[2].Y, pos[3].Y)
	// Output:
	// X  -3.4503  +7.4418  +1.2010  +7.0720
	// Y  +0.2137  +0.2752  +0.5900  +1.0290
}

// The exercise of finding the zero crossing is not coded here, but computed
// are offsets at the times given by Meeus, showing the X coordinates near
// zero (indicating conjunction) and Y coordinates near the values given by
// Meeus.
func ExampleE5_conjunction() {
	// Exercise, p. 314.
	e, err := pp.LoadPlanet(pp.Earth)
	if err != nil {
		fmt.Println(err)
		return
	}
	j, err := pp.LoadPlanet(pp.Jupiter)
	if err != nil {
		fmt.Println(err)
		return
	}
	var pos [4]jupitermoons.XY
	jd := julian.CalendarGregorianToJD(1988, 11, 23)
	jd += deltat.Interp10A(jd).Day()
	t3 := unit.NewTime(' ', 7, 28, 0)
	jupitermoons.E5(jd+t3.Day(), e, j, &pos)
	fmt.Printf("III  %m  X = %+.4f  Y = %+.4f\n",
		sexa.FmtTime(t3), pos[2].X, pos[2].Y)
	t4 := unit.NewTime(' ', 5, 15, 0)
	jupitermoons.E5(jd+t4.Day(), e, j, &pos)
	fmt.Printf("IV   %m  X = %+.4f  Y = %+.4f\n",
		sexa.FmtTime(t4), pos[3].X, pos[3].Y)
	// Output:
	// III  7ʰ28ᵐ  X = +0.0032  Y = -0.8042
	// IV   5ʰ15ᵐ  X = +0.0002  Y = +1.3990
}
