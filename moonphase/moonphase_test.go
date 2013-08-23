package moonphase_test

import (
	"fmt"

	"github.com/soniakeys/meeus/moonphase"
)

func ExampleMeanNew() {
	// Example 49.a, p. 353.
	fmt.Printf("JDE = %.5f\n", moonphase.MeanNew(1977.13))
	// Output:
	// JDE = 2443192.94102
}

func ExampleNew() {
	// Example 49.a, p. 353.
	fmt.Printf("JDE = %.5f\n", moonphase.New(1977.13))
	// Output:
	// JDE = 2443192.65118
}

func ExampleMeanLast() {
	// Example 49.b, p. 353.
	fmt.Printf("JDE = %.5f\n", moonphase.MeanLast(2044.04))
	// Output:
	// JDE = 2467636.88597
}

func ExampleLast() {
	// Example 49.b, p. 353.
	fmt.Printf("JDE = %.5f\n", moonphase.Last(2044.04))
	// Output:
	// JDE = 2467636.49186
}
