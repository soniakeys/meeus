// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package circle_test

import (
	"fmt"
	"math"

	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/meeus/circle"
)

func ExampleSmallest_a() {
	// Example 20.a, p. 128.
	r1 := base.NewRA(12, 41, 8.64).Rad()
	r2 := base.NewRA(12, 52, 5.21).Rad()
	r3 := base.NewRA(12, 39, 28.11).Rad()
	d1 := base.NewAngle(true, 5, 37, 54.2).Rad()
	d2 := base.NewAngle(true, 4, 22, 26.2).Rad()
	d3 := base.NewAngle(true, 1, 50, 3.7).Rad()
	d, t := circle.Smallest(r1, d1, r2, d2, r3, d3)
	fmt.Printf("Δ = %s = %.62s\n",
		base.DecSymAdd(fmt.Sprintf("%.5f", d*180/math.Pi), '°'),
		base.NewFmtAngle(d))
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
	r1 := base.NewRA(9, 5, 41.44).Rad()
	r2 := base.NewRA(9, 9, 29).Rad()
	r3 := base.NewRA(8, 59, 47.14).Rad()
	d1 := base.NewAngle(false, 18, 30, 30).Rad()
	d2 := base.NewAngle(false, 17, 43, 56.7).Rad()
	d3 := base.NewAngle(false, 17, 49, 36.8).Rad()
	d, t := circle.Smallest(r1, d1, r2, d2, r3, d3)
	fmt.Printf("Δ = %.62s\n", base.NewFmtAngle(d))
	if t {
		fmt.Println("type I")
	} else {
		fmt.Println("type II")
	}
	// Output:
	// Δ = 2°19′
	// type I
}
