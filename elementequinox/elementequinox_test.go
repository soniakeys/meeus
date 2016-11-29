// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package elementequinox_test

import (
	"fmt"

	"github.com/soniakeys/meeus/elementequinox"
	"github.com/soniakeys/unit"
)

// See package precess for example 24.a.

func ExampleReduceB1950ToJ2000() {
	// Example 24.b, p. 161.
	ele := &elementequinox.Elements{
		Inc:  unit.AngleFromDeg(11.93911),
		Node: unit.AngleFromDeg(334.04096),
		Peri: unit.AngleFromDeg(186.24444),
	}
	elementequinox.ReduceB1950ToJ2000(ele, ele)
	fmt.Printf("i  %.5f\n", ele.Inc.Deg())
	fmt.Printf("Ω  %.5f\n", ele.Node.Deg())
	fmt.Printf("ω  %.5f\n", ele.Peri.Deg())
	// Output:
	// i  11.94524
	// Ω  334.75006
	// ω  186.23352
}

func ExampleReduceB1950FK4ToJ2000FK5() {
	// Example 24.c, p. 162.
	ele := &elementequinox.Elements{
		Inc:  unit.AngleFromDeg(11.93911),
		Node: unit.AngleFromDeg(334.04096),
		Peri: unit.AngleFromDeg(186.24444),
	}
	elementequinox.ReduceB1950FK4ToJ2000FK5(ele, ele)
	fmt.Printf("i  %.5f\n", ele.Inc.Deg())
	fmt.Printf("Ω  %.5f\n", ele.Node.Deg())
	fmt.Printf("ω  %.5f\n", ele.Peri.Deg())
	// Output:
	// i  11.94521
	// Ω  334.75043
	// ω  186.23327
}
