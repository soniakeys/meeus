// Copyright 2013 Sonia Keys
// License: MIT

package iterate_test

import (
	"fmt"
	"math"

	"github.com/soniakeys/meeus/v3/iterate"
)

func ExampleDecimalPlaces() {
	// Example 5.a, p. 48.
	betterSqrt := func(N float64) iterate.BetterFunc {
		return func(n float64) float64 {
			return (n + N/n) / 2
		}
	}
	start := 12.
	places := 8
	maxIter := 20
	n, err := iterate.DecimalPlaces(betterSqrt(159), start, places, maxIter)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%.*f", places, n)
	// Output:
	// 12.60952021
}

func ExampleFullPrecision() {
	// Example 5.b, p. 48.
	betterRoot := func() iterate.BetterFunc {
		return func(x float64) float64 {
			return (8 - math.Pow(x, 5)) / 17
		}
	}
	x, err := iterate.FullPrecision(betterRoot(), 0, 20)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%.9f\n", x)
	fmt.Println(x, "(full precision)")
	// Output:
	// 0.469249878
	// 0.4692498784547387 (full precision)
}

func ExampleFullPrecision_diverging() {
	// Example 5.c, p. 49.
	betterRoot := func() iterate.BetterFunc {
		return func(x float64) float64 {
			return (8 - math.Pow(x, 5)) / 3
		}
	}
	x, err := iterate.FullPrecision(betterRoot(), 0, 20)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%.9f\n", x)
	fmt.Println(x, "(full precision)")
	// Output:
	// Maximum iterations reached
}

func ExampleFullPrecision_converging() {
	// Example 5.d, p.49.
	betterRoot := func() iterate.BetterFunc {
		return func(x float64) float64 {
			return math.Pow(8-3*x, .2)
		}
	}
	x, err := iterate.FullPrecision(betterRoot(), 0, 30)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%.9f\n", x)
	fmt.Println(x, "(full precision)")
	// Output:
	// 1.321785627
	// 1.321785627117658 (full precision)
}

func ExampleBinaryRoot() {
	// Example from p. 53.
	f := func(x float64) float64 {
		return math.Pow(x, 5) + 17*x - 8
	}
	x := iterate.BinaryRoot(f, 0, 1)
	fmt.Println(x)
	// Output:
	// 0.46924987845473876
}
