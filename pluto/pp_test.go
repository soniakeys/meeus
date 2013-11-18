// +build !nopp

package pluto_test

import (
	"fmt"

	"github.com/soniakeys/meeus/base"
	pp "github.com/soniakeys/meeus/planetposition"
	"github.com/soniakeys/meeus/pluto"
)

func ExampleAstrometric() {
	// Example 37.a, p. 266
	e, err := pp.LoadPlanet(pp.Earth)
	if err != nil {
		fmt.Println(err)
		return
	}
	α, δ := pluto.Astrometric(2448908.5, e)
	fmt.Printf("α: %.1d\n", base.NewFmtRA(α))
	fmt.Printf("δ: %.0d\n", base.NewFmtAngle(δ))
	// Output:
	// α: 15ʰ31ᵐ43ˢ.8
	// δ: -4°27′29″
}
