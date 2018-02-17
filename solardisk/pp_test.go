// Copyright 2013 Sonia Keys
// License: MIT

// +build !nopp

package solardisk_test

import (
	"fmt"

	pp "github.com/soniakeys/meeus/planetposition"
	"github.com/soniakeys/meeus/solardisk"
)

func ExampleEphemeris() {
	j := 2448908.50068
	e, err := pp.LoadPlanet(pp.Earth)
	if err != nil {
		fmt.Println(err)
		return
	}
	P, B0, L0 := solardisk.Ephemeris(j, e)
	fmt.Printf("P:  %.2f\n", P.Deg())
	fmt.Printf("B0: %+.2f\n", B0.Deg())
	fmt.Printf("L0: %.2f\n", L0.Deg())
	// Output:
	// P:  26.27
	// B0: +5.99
	// L0: 238.63
}
