// Copyright 2013 Sonia Keys
// License: MIT

package easter_test

import (
	"fmt"
	"time"

	"github.com/soniakeys/meeus/v3/easter"
)

func ExampleGregorian() {
	// Example values from p. 68.
	for _, y := range []int{1991, 1992, 1993, 1954, 2000, 1818} {
		m, d := easter.Gregorian(y)
		fmt.Println(y, ":", time.Month(m), d)
	}
	// Output:
	// 1991 : March 31
	// 1992 : April 19
	// 1993 : April 11
	// 1954 : April 18
	// 2000 : April 23
	// 1818 : March 22
}

func ExampleJulian() {
	// Example value from p. 69.
	y := 1243
	m, d := easter.Julian(y)
	fmt.Println(y, ":", time.Month(m), d)
	// Output:
	// 1243 : April 12
}
