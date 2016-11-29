// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// +build !nopp

package pluto_test

import (
	"fmt"

	pp "github.com/soniakeys/meeus/planetposition"
	"github.com/soniakeys/meeus/pluto"
	"github.com/soniakeys/sexagesimal"
)

func ExampleAstrometric() {
	// Example 37.a, p. 266
	e, err := pp.LoadPlanet(pp.Earth)
	if err != nil {
		fmt.Println(err)
		return
	}
	α, δ := pluto.Astrometric(2448908.5, e)
	fmt.Printf("α: %.1d\n", sexa.FmtRA(α))
	fmt.Printf("δ: %.0d\n", sexa.FmtAngle(δ))
	// Output:
	// α: 15ʰ31ᵐ43ˢ.8
	// δ: -4°27′29″
}
