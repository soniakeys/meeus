package jupiter_test

import (
	"fmt"
	"math"

	"github.com/soniakeys/meeus/jupiter"
	pp "github.com/soniakeys/meeus/planetposition"
)

func ExamplePhysical() {
	// Example 43.a, p. 295
	e, err := pp.LoadPlanet(pp.Earth, "")
	if err != nil {
		fmt.Println(err)
		return
	}
	j, err := pp.LoadPlanet(pp.Jupiter, "")
	if err != nil {
		fmt.Println(err)
		return
	}
	DS, DE, ω1, ω2, P := jupiter.Physical(2448972.50068, e, j)
	fmt.Printf("DS = %+.2f\n", DS*180/math.Pi)
	fmt.Printf("DE = %+.2f\n", DE*180/math.Pi)
	fmt.Printf("ω1 = %.2f\n", ω1*180/math.Pi)
	fmt.Printf("ω2 = %.2f\n", ω2*180/math.Pi)
	fmt.Printf("P = %.2f\n", P*180/math.Pi)
	// Output:
	// DS = -2.20
	// DE = -2.48
	// ω1 = 268.06
	// ω2 = 72.74
	// P = 24.80
}

func ExamplePhysical2() {
	// Example 43.b, p. 299
	DS, DE, ω1, ω2 := jupiter.Physical2(2448972.50068)
	fmt.Printf("DS = %+.3f\n", DS*180/math.Pi)
	fmt.Printf("DE = %+.2f\n", DE*180/math.Pi)
	fmt.Printf("ω1 = %.2f\n", ω1*180/math.Pi)
	fmt.Printf("ω2 = %.2f\n", ω2*180/math.Pi)
	// Output:
	// DS = -2.194
	// DE = -2.50
	// ω1 = 268.12
	// ω2 = 72.79
}
