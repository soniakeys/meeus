// Copyright 2013 Sonia Keys
// License: MIT

package circle_test

import (
	"fmt"

	"github.com/soniakeys/meeus/circle"
	"github.com/soniakeys/sexagesimal"
	"github.com/soniakeys/unit"
)

func ExampleSmallest_a() {
	// Example 20.a, p. 128.
	r1 := unit.NewRA(12, 41, 8.64).Angle()
	r2 := unit.NewRA(12, 52, 5.21).Angle()
	r3 := unit.NewRA(12, 39, 28.11).Angle()
	d1 := unit.NewAngle('-', 5, 37, 54.2)
	d2 := unit.NewAngle('-', 4, 22, 26.2)
	d3 := unit.NewAngle('-', 1, 50, 3.7)
	d, t := circle.Smallest(r1, d1, r2, d2, r3, d3)
	fd := sexa.FmtAngle(d)
	fmt.Printf("Δ = %.5j = %m\n", fd, fd)
	if t {
		fmt.Println("type I")
	} else {
		fmt.Println("type II")
	}
	// Output:
	// Δ = 4°.26363 = 4°16′
	// type II
}

func ExampleSmallest_b() {
	// Exercise, p. 128.
	r1 := unit.NewRA(9, 5, 41.44).Angle()
	r2 := unit.NewRA(9, 9, 29).Angle()
	r3 := unit.NewRA(8, 59, 47.14).Angle()
	d1 := unit.NewAngle(' ', 18, 30, 30)
	d2 := unit.NewAngle(' ', 17, 43, 56.7)
	d3 := unit.NewAngle(' ', 17, 49, 36.8)
	d, t := circle.Smallest(r1, d1, r2, d2, r3, d3)
	fmt.Printf("Δ = %m\n", sexa.FmtAngle(d))
	if t {
		fmt.Println("type I")
	} else {
		fmt.Println("type II")
	}
	// Output:
	// Δ = 2°19′
	// type I
}
