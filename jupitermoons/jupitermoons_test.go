// Copyright 2013 Sonia Keys
// License: MIT

package jupitermoons_test

import (
	"fmt"

	"github.com/soniakeys/meeus/deltat"
	"github.com/soniakeys/meeus/julian"
	"github.com/soniakeys/meeus/jupitermoons"
	"github.com/soniakeys/sexagesimal"
	"github.com/soniakeys/unit"
)

func ExamplePositions() {
	// Example 44.a, p. 303.
	p1, p2, p3, p4 := jupitermoons.Positions(2448972.50068)
	fmt.Printf("X1 = %+.2f  Y1 = %+.2f\n", p1.X, p1.Y)
	fmt.Printf("X2 = %+.2f  Y2 = %+.2f\n", p2.X, p2.Y)
	fmt.Printf("X3 = %+.2f  Y3 = %+.2f\n", p3.X, p3.Y)
	fmt.Printf("X4 = %+.2f  Y4 = %+.2f\n", p4.X, p4.Y)
	// Output:
	// X1 = -3.44  Y1 = +0.21
	// X2 = +7.44  Y2 = +0.25
	// X3 = +1.24  Y3 = +0.65
	// X4 = +7.08  Y4 = +1.10
}

// The exercise of finding the zero crossing is not coded here, but computed
// are offsets at the times given by Meeus, showing the X coordinates near
// zero (indicating conjunction) and Y coordinates near the values given by
// Meeus.
func ExamplePositions_conjunction() {
	// Exercise, p. 314.
	jd := julian.CalendarGregorianToJD(1988, 11, 23)
	jd += deltat.Interp10A(jd).Day()
	t3 := unit.NewTime(' ', 7, 28, 0)
	_, _, p3, _ := jupitermoons.Positions(jd + t3.Day())
	fmt.Printf("III  %m  X = %+.2f  Y = %+.2f\n", sexa.FmtTime(t3), p3.X, p3.Y)
	t4 := unit.NewTime(' ', 5, 15, 0)
	_, _, _, p4 := jupitermoons.Positions(jd + t4.Day())
	fmt.Printf("IV   %m  X = %+.2f  Y = %+.2f\n", sexa.FmtTime(t4), p4.X, p4.Y)
	// Output:
	// III  7ʰ28ᵐ  X = -0.00  Y = -0.84
	// IV   5ʰ15ᵐ  X = +0.06  Y = +1.48
}
