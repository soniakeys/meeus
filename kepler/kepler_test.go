// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package kepler_test

import (
	"fmt"

	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/meeus/kepler"
)

func ExampleKepler1() {
	// Example 30.a, p. 196
	E, err := kepler.Kepler1(.1, base.AngleFromDeg(5), 8)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%.6f\n", E.Deg())
	// Output:
	// 5.554589
}

func ExampleKepler2() {
	// Example 30.b, p. 199
	E, err := kepler.Kepler2(.1, base.AngleFromDeg(5), 11)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%.9f\n", E.Deg())
	// Output:
	// 5.554589254
}

func ExampleKepler2a() {
	// Example data from p. 205
	E, err := kepler.Kepler2a(.99, base.Angle(.2), 14)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%.12f\n", E.Rad())
	fmt.Printf("%.8f\n", E.Deg())
	// Output:
	// 1.066997365282
	// 61.13444578
}

func ExampleKepler2b() {
	// Example data from p. 205
	E, err := kepler.Kepler2b(.99, .2, 14)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%.12f\n", E)
	// Output:
	// 1.066997365282
}

func ExampleKepler3() {
	// Example data from p. 205
	fmt.Printf("%.12f\n", kepler.Kepler3(.99, .2))
	// Output:
	// 1.066997365282
}

func ExampleKepler4() {
	// Input data from example 30.a, p. 196,
	// result from p. 207
	E := kepler.Kepler4(.1, base.AngleFromDeg(5))
	fmt.Printf("%.6f\n", E.Deg())
	// Output:
	// 5.554599
}
