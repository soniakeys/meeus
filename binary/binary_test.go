// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package binary_test

import (
	"fmt"
	"math"

	"github.com/soniakeys/meeus/binary"
	"github.com/soniakeys/meeus/kepler"
)

func ExamplePosition() {
	// Example 57.1, p. 398
	M := binary.M(1980, 1934.008, 41.623)
	fmt.Printf("M = %.3f\n", M*180/math.Pi)
	E, err := kepler.Kepler1(.2763, M, 6)
	if err != nil {
		fmt.Println(err)
		return
	}
	θ, ρ := binary.Position(.907, .2763, 59.025*math.Pi/180,
		23.717*math.Pi/180, 219.907*math.Pi/180, E)
	fmt.Printf("θ = %.1f\n", θ*180/math.Pi)
	fmt.Printf("ρ = %.3f\n", ρ)
	// Output:
	// M = 37.788
	// θ = 318.4
	// ρ = 0.411
}

func ExampleApparentEccentricity() {
	// Example 57.b, p. 400
	fmt.Printf("%.3f\n", binary.ApparentEccentricity(.2763,
		59.025*math.Pi/180, 219.907*math.Pi/180))
	// Output:
	// 0.860
}
