// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package elliptic_test

import (
	"fmt"

	"github.com/soniakeys/meeus/elliptic"
)

func ExampleVelocity() {
	// Example 33.c, p. 238
	fmt.Printf("%.2f\n", elliptic.Velocity(17.9400782, 1))
	// Output:
	// 41.53
}

func ExampleVPerihelion() {
	// Example 33.c, p. 238
	fmt.Printf("%.2f\n", elliptic.VPerihelion(17.9400782, .96727426))
	// Output:
	// 54.52
}

func ExampleVAphelion() {
	// Example 33.c, p. 238
	fmt.Printf("%.2f\n", elliptic.VAphelion(17.9400782, 0.96727426))
	// Output:
	// 0.91
}

func ExampleLength1() {
	// Example 33.d, p. 239
	fmt.Printf("%.2f\n", elliptic.Length1(17.9400782, 0.96727426))
	// Output:
	// 77.06
}

func ExampleLength2() {
	// Example 33.d, p. 239
	fmt.Printf("%.2f\n", elliptic.Length2(17.9400782, 0.96727426))
	// Output:
	// 77.09
}

/* func ExampleLength3() {
	// Example 33.d, p. 239
	fmt.Printf("%.2f\n", elliptic.Length3(17.9400782, 0.96727426))
	// Output:
	// 77.07
} */

func ExampleLength4() {
	// Example 33.d, p. 239
	fmt.Printf("%.2f\n", elliptic.Length4(17.9400782, 0.96727426))
	// Output:
	// 77.07
}
