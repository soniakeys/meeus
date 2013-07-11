// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package solarxyz_test

import (
	"fmt"
	//	"math"

	pp "github.com/soniakeys/meeus/planetposition"
	"github.com/soniakeys/meeus/solarxyz"
)

func ExamplePosition() {
	// Example 26.a, p. 172.
	e, err := pp.LoadPlanet(pp.Earth, "")
	if err != nil {
		fmt.Println(err)
		return
	}
	x, y, z := solarxyz.Position(e, 2448908.5)
	fmt.Printf("X = %.7f\n", x)
	fmt.Printf("Y = %.7f\n", y)
	fmt.Printf("Z = %.7f\n", z)
	// Meeus result:
	// X = -0.9379952
	// Y = -0.3116544
	// Z = -0.1351215

	// Output:
	// X = -0.9379963
	// Y = -0.3116537
	// Z = -0.1351207
}

func ExamplePositionJ2000() {
	// Example 26.b, p. 175 but for output see complete VSOP87
	// results given on p. 176.
	e, err := pp.LoadPlanet(pp.Earth, "")
	if err != nil {
		fmt.Println(err)
		return
	}
	x, y, z := solarxyz.PositionJ2000(e, 2448908.5)
	fmt.Printf("X0 = %.8f\n", x)
	fmt.Printf("Y0 = %.8f\n", y)
	fmt.Printf("Z0 = %.8f\n", z)
	// Output:
	// X0 = -0.93739707
	// Y0 = -0.31316724
	// Z0 = -0.13577841
}

func ExamplePositionB1950() {
	// Example 26.b, p. 175
	e, err := pp.LoadPlanet(pp.Earth, "")
	if err != nil {
		fmt.Println(err)
		return
	}
	x, y, z := solarxyz.PositionB1950(e, 2448908.5)
	fmt.Printf("X0 = %.8f\n", x)
	fmt.Printf("Y0 = %.8f\n", y)
	fmt.Printf("Z0 = %.8f\n", z)
	// Output:
	// X0 = -0.94148805
	// Y0 = -0.30266488
	// Z0 = -0.13121349
}

func ExamplePositionEquinox() {
	// Example 26.b, p. 175
	e, err := pp.LoadPlanet(pp.Earth, "")
	if err != nil {
		fmt.Println(err)
		return
	}
	x, y, z := solarxyz.PositionEquinox(e, 2448908.5, 2044)
	fmt.Printf("X0 = %.8f\n", x)
	fmt.Printf("Y0 = %.8f\n", y)
	fmt.Printf("Z0 = %.8f\n", z)
	// Output:
	// X0 = -0.93368100
	// Y0 = -0.32237347
	// Z0 = -0.13977803
}
