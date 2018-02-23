// Copyright 2013 Sonia Keys
// License: MIT

package binary_test

import (
	"fmt"

	"github.com/soniakeys/meeus/v3/binary"
	"github.com/soniakeys/meeus/v3/kepler"
	"github.com/soniakeys/unit"
)

func ExamplePosition() {
	// Example 57.a, p. 398
	M := binary.M(1980, 1934.008, 41.623)
	fmt.Printf("M = %.3f\n", M.Deg())
	E, err := kepler.Kepler1(.2763, M, 4)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("E = %.3f\n", E.Deg())
	θ, ρ := binary.Position(.2763, unit.AngleFromSec(.907),
		unit.AngleFromDeg(59.025), unit.AngleFromDeg(23.717),
		unit.AngleFromDeg(219.907), E)
	fmt.Printf("θ = %.1f\n", θ.Deg())
	fmt.Printf("ρ = %.3f\n", ρ.Sec())
	// Output:
	// M = 37.788
	// E = 49.896
	// θ = 318.4
	// ρ = 0.411
}

func ExampleApparentEccentricity() {
	// Example 57.b, p. 400
	fmt.Printf("%.3f\n", binary.ApparentEccentricity(.2763,
		unit.AngleFromDeg(59.025), unit.AngleFromDeg(219.907)))
	// Output:
	// 0.860
}
