// Copyright 2013 Sonia Keys
// License: MIT

// +build !nopp

package mars_test

import (
	"fmt"

	"github.com/soniakeys/meeus/mars"
	pp "github.com/soniakeys/meeus/planetposition"
)

func ExamplePhysical() {
	// Example 42.a, p. 291
	e, err := pp.LoadPlanet(pp.Earth)
	if err != nil {
		fmt.Println(err)
		return
	}
	m, err := pp.LoadPlanet(pp.Mars)
	if err != nil {
		fmt.Println(err)
		return
	}
	DE, DS, ω, P, Q, d, q, k := mars.Physical(2448935.500683, e, m)
	fmt.Printf("DE = %+.2f\n", DE.Deg())
	fmt.Printf("DS = %+.2f\n", DS.Deg())
	fmt.Printf("ω = %.2f\n", ω.Deg())
	fmt.Printf("P = %.2f\n", P.Deg())
	fmt.Printf("Q = %.2f\n", Q.Deg())
	fmt.Printf("d = %.2f\n", d.Sec()) // display as arc sec
	fmt.Printf("k = %.4f\n", k)
	fmt.Printf("q = %.2f\n", q.Sec()) // display as arc sec
	// Output:
	// DE = +12.44
	// DS = -2.76
	// ω = 111.55
	// P = 347.64
	// Q = 279.91
	// d = 10.75
	// k = 0.9012
	// q = 1.06
}
