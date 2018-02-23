// Copyright 2016 Sonia Keys
// License: MIT

package hints_test

import (
	"fmt"
	"math"

	"github.com/soniakeys/unit"
)

// Section "Trigonometric functions of large angles":
//
// Meeus makes his point, but an example with integer values is a bit unfair
// when trigonometric functions inherently work on floating point numbers.
func Example_pMod_integer() {
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

	// Use of unit.PMod on integer constants produces results identical to
	// above.
	fmt.Println("PMod int:  ", math.Sin(unit.PMod(large, 360)*math.Pi/180))

	// As soon as the large integer is scaled to a non-integer value though,
	// precision is lost and PMod is of no help recovering at this point.
	fmt.Println("PMod float:", math.Sin(unit.PMod(large*math.Pi/180, 2*math.Pi)))
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
func Example_pMod_mars() {
	// Angle W from step 9 of example 42.a, as suggested.
	const W = 5492522.4593
	const reduced = 2.4593

	// Direct function call.  It's a number.  How correct is it?
	fmt.Println("Direct:  ", math.Sin(W*math.Pi/180))

	// Manually reduced to range 0-360.  This is presumably the "correct"
	// answer, but note that the reduced number has a reduced number of
	// significant digits.  The answer cannot have any more significant digits.
	fmt.Println("Reduced: ", math.Sin(reduced*math.Pi/180))

	// Accordingly, PMod cannot rescue any precision, whether done on degrees
	// or radians.
	fmt.Println("PMod deg:", math.Sin(unit.PMod(W, 360)*math.Pi/180))
	fmt.Println("PMod rad:", math.Sin(unit.PMod(W*math.Pi/180, 2*math.Pi)))
	// Output:
	// Direct:   0.04290970350270464
	// Reduced:  0.04290970350923273
	// PMod deg: 0.04290970351307828
	// PMod rad: 0.04290970350643808
}

func Example_rA() {
	// Example 1.a, p. 8
	h := unit.FromSexa(' ', 9, 14, 55.8)
	fmt.Printf("%.9f\n", h)
	α := unit.RAFromHour(h)
	fmt.Printf("%.5f\n", α.Deg())
	fmt.Printf("%.6f\n", α.Tan())
	// Output:
	// 9.248833333
	// 138.73250
	// -0.877517
}
