// Copyright 2013 Sonia Keys
// License: MIT

package pluto_test

import (
	"fmt"

	"github.com/soniakeys/meeus/v3/pluto"
)

func ExampleHeliocentric() {
	// Example 37.a, p. 266
	l, b, r := pluto.Heliocentric(2448908.5)
	fmt.Printf("l: %.5f\n", l.Deg())
	fmt.Printf("b: %.5f\n", b.Deg())
	fmt.Printf("r: %.6f\n", r)
	// Output:
	// l: 232.74071
	// b: 14.58782
	// r: 29.711111
}
