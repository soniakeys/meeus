package pluto_test

import (
	"fmt"
	"math"

	"github.com/soniakeys/meeus/base"
	pp "github.com/soniakeys/meeus/planetposition"
	"github.com/soniakeys/meeus/pluto"
)

func ExampleHeliocentric() {
	// Example 37.a, p. 266
	l, b, r := pluto.Heliocentric(2448908.5)
	fmt.Printf("l: %.5f\n", l*180/math.Pi)
	fmt.Printf("b: %.5f\n", b*180/math.Pi)
	fmt.Printf("r: %.6f\n", r)
	// Output:
	// l: 232.74071
	// b: 14.58782
	// r: 29.711111
}

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
