// Copyright 2013 Sonia Keys
// License: MIT

// +build !nopp

package jupiter_test

import (
	"fmt"

	"github.com/soniakeys/meeus/jupiter"
	pp "github.com/soniakeys/meeus/planetposition"
)

func ExamplePhysical() {
	// Example 43.a, p. 295
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
	DS, DE, ω1, ω2, P := jupiter.Physical(2448972.50068, e, j)
	fmt.Printf("DS = %+.2f\n", DS.Deg())
	fmt.Printf("DE = %+.2f\n", DE.Deg())
	fmt.Printf("ω1 = %.2f\n", ω1.Deg())
	fmt.Printf("ω2 = %.2f\n", ω2.Deg())
	fmt.Printf("P = %.2f\n", P.Deg())
	// Output:
	// DS = -2.20
	// DE = -2.48
	// ω1 = 268.06
	// ω2 = 72.74
	// P = 24.80
}
