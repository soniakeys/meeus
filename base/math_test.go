// Copyright 2012 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package base_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/soniakeys/meeus/base"
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

// Section "Trigonometric functions of large angles":
//
// Meeus makes his point, but an example with integer values is a bit unfair
// when trigonometric functions inherently work on floating point numbers.
func ExamplePMod_integer() {
	const large = 36000030

	// The direct function call loses precition as expected.
	fmt.Println("Direct:    ", math.Sin(large*math.Pi/180))

	// When the value is manually reduced to the integer 30, the Go constant
	// evaluaton does a good job of delivering a radian value to math.Sin that
	// evaluates to .5 exactly.
	fmt.Println("Integer 30:", math.Sin(30*math.Pi/180))

	// Math.Mod takes float64s and returns float64s.  The integer constants
	// here however can be represented exactly as float64s, and the returned
	// result is exact as well.
	fmt.Println("Math.Mod:  ", math.Mod(large, 360))

	// But when math.Mod is substituted into the Sin function, float64s
	// are multiplied instead of the high precision constants, and the result
	// comes back slightly inexact.
	fmt.Println("Sin Mod:   ", math.Sin(math.Mod(large, 360)*math.Pi/180))

	// Use of PMod on integer constants produces results identical to above.
	fmt.Println("PMod int:  ", math.Sin(base.PMod(large, 360)*math.Pi/180))

	// As soon as the large integer is scaled to a non-integer value though,
	// precision is lost and PMod is of no help recovering at this point.
	fmt.Println("PMod float:", math.Sin(base.PMod(large*math.Pi/180, 2*math.Pi)))
	// Output:
	// Direct:     0.49999999995724154
	// Integer 30: 0.5
	// Math.Mod:   30
	// Sin Mod:    0.49999999999999994
	// PMod int:   0.49999999999999994
	// PMod float: 0.49999999997845307
}

// Section "Trigonometric functions of large angles":
//
// A non integer example better illustrates that reduction to a range
// does not rescue precision.
func ExamplePMod_mars() {
	// Angle W from step 9 of example 42.a, as suggested.
	const W = 5492522.4593
	const reduced = 2.4593

	// Direct function call.  It's a number.  How correct is it?
	fmt.Println("Direct:  ", math.Sin(W*math.Pi/180))

	// Manually reduced to range 0-360.  This is presumably the "correct"
	// answer, but note that the reduced number has a reduced number of
	// significat digits.  The answer cannot have any more significant digits.
	fmt.Println("Reduced: ", math.Sin(reduced*math.Pi/180))

	// Accordingly, PMod cannot rescue any precision, whether done on degrees
	// or radians.
	fmt.Println("PMod deg:", math.Sin(base.PMod(W, 360)*math.Pi/180))
	fmt.Println("PMod rad:", math.Sin(base.PMod(W*math.Pi/180, 2*math.Pi)))
	// Output:
	// Direct:   0.04290970350270464
	// Reduced:  0.04290970350923273
	// PMod deg: 0.04290970351307828
	// PMod rad: 0.04290970350643808
}

// Meeus gives no test case.
// The test case here is from Wikipedia's entry on Horner's method.
func TestHorner(t *testing.T) {
	y := base.Horner(3, -1, 2, -6, 2)
	if y != 5 {
		t.Fatal("Horner")
	}
}
