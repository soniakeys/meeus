// Copyright 2012 Sonia Keys
// License: MIT

package base_test

import (
	"fmt"
	"testing"

	"github.com/soniakeys/meeus/v3/base"
)

func ExampleFloorDiv() {
	// compare to / operator examples in Go spec at
	// https://golang.org/ref/spec#Arithmetic_operators
	fmt.Println(base.FloorDiv(+5, +3))
	fmt.Println(base.FloorDiv(-5, +3))
	fmt.Println(base.FloorDiv(+5, -3))
	fmt.Println(base.FloorDiv(-5, -3))
	fmt.Println()
	// exact divisors, no remainders
	fmt.Println(base.FloorDiv(+6, +3))
	fmt.Println(base.FloorDiv(-6, +3))
	fmt.Println(base.FloorDiv(+6, -3))
	fmt.Println(base.FloorDiv(-6, -3))
	// Output:
	// 1
	// -2
	// -2
	// 1
	//
	// 2
	// -2
	// -2
	// 2
}

func ExampleFloorDiv64() {
	// compare to / operator examples in Go spec at
	// https://golang.org/ref/spec#Arithmetic_operators
	fmt.Println(base.FloorDiv64(+5, +3))
	fmt.Println(base.FloorDiv64(-5, +3))
	fmt.Println(base.FloorDiv64(+5, -3))
	fmt.Println(base.FloorDiv64(-5, -3))
	fmt.Println()
	// exact divisors, no remainders
	fmt.Println(base.FloorDiv64(+6, +3))
	fmt.Println(base.FloorDiv64(-6, +3))
	fmt.Println(base.FloorDiv64(+6, -3))
	fmt.Println(base.FloorDiv64(-6, -3))
	// Output:
	// 1
	// -2
	// -2
	// 1
	//
	// 2
	// -2
	// -2
	// 2
}

// Meeus gives no test case.
// The test case here is from Wikipedia's entry on Horner's method.
func TestHorner(t *testing.T) {
	y := base.Horner(3, -1, 2, -6, 2)
	if y != 5 {
		t.Fatal("Horner")
	}
}
