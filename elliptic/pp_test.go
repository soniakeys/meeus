// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// +build !nopp

package elliptic_test

import (
	"fmt"
	"math"

	"github.com/soniakeys/meeus/elliptic"
	"github.com/soniakeys/meeus/julian"
	pp "github.com/soniakeys/meeus/planetposition"
	"github.com/soniakeys/sexagesimal"
)

func ExamplePosition() {
	// Example 33.a, p. 225.  VSOP87 result p. 227.
	earth, err := pp.LoadPlanet(pp.Earth)
	if err != nil {
		fmt.Println(err)
		return
	}
	venus, err := pp.LoadPlanet(pp.Venus)
	if err != nil {
		fmt.Println(err)
		return
	}
	α, δ := elliptic.Position(venus, earth, 2448976.5)
	fmt.Printf("α = %.3d\n", sexa.NewFmtRA(α))
	fmt.Printf("δ = %.2d\n", sexa.NewFmtAngle(δ))
	// Output:
	// α = 21ʰ4ᵐ41ˢ.454
	// δ = -18°53′16″.84
}

func ExampleElements_Position() {
	// Example 33.b, p. 232.
	earth, err := pp.LoadPlanet(pp.Earth)
	if err != nil {
		fmt.Println(err)
		return
	}
	k := &elliptic.Elements{
		TimeP: julian.CalendarGregorianToJD(1990, 10, 28.54502),
		Axis:  2.2091404,
		Ecc:   .8502196,
		Inc:   11.94524 * math.Pi / 180,
		Node:  334.75006 * math.Pi / 180,
		ArgP:  186.23352 * math.Pi / 180,
	}
	j := julian.CalendarGregorianToJD(1990, 10, 6)
	α, δ, ψ := k.Position(j, earth)
	fmt.Printf("α = %.1d\n", sexa.NewFmtRA(α))
	fmt.Printf("δ = %.0d\n", sexa.NewFmtAngle(δ))
	fmt.Printf("ψ = %.2f\n", ψ*180/math.Pi)
	// Output:
	// α = 10ʰ34ᵐ14ˢ.2
	// δ = 19°9′31″
	// ψ = 40.51
}
