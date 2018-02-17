// Copyright 2013 Sonia Keys
// License: MIT

package stellar_test

import (
	"fmt"

	"github.com/soniakeys/meeus/stellar"
)

func ExampleSum() {
	// Example 56.a, p. 393
	fmt.Printf("%.2f\n", stellar.Sum(1.96, 2.89))
	// Output:
	// 1.58
}

func ExampleSumN_triple() {
	// Example 56.b, p. 394
	fmt.Printf("%.2f\n", stellar.SumN(4.73, 5.22, 5.6))
	// Output:
	// 3.93
}

func ExampleSumN_cluster() {
	// Example 56.c, p. 394
	var c []float64
	for i := 0; i < 4; i++ {
		c = append(c, 5)
	}
	for i := 0; i < 14; i++ {
		c = append(c, 6)
	}
	for i := 0; i < 23; i++ {
		c = append(c, 7)
	}
	for i := 0; i < 38; i++ {
		c = append(c, 8)
	}
	fmt.Printf("%.2f\n", stellar.SumN(c...))
	// Output:
	// 2.02
}

func ExampleRatio() {
	// Example 56.d, p. 395
	fmt.Printf("%.2f\n", stellar.Ratio(.14, 2.12))
	// Output:
	// 6.19
}

func ExampleDifference() {
	// Example 56.e, p. 395
	fmt.Printf("%.2f\n", stellar.Difference(500))
	// Output:
	// 6.75
}
