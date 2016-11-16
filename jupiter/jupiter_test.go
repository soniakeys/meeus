// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package jupiter_test

import (
	"fmt"
	"math"

	"github.com/soniakeys/meeus/jupiter"
)

func ExamplePhysical2() {
	// Example 43.b, p. 299
	DS, DE, ω1, ω2 := jupiter.Physical2(2448972.50068)
	fmt.Printf("DS = %+.3f\n", DS*180/math.Pi)
	fmt.Printf("DE = %+.2f\n", DE*180/math.Pi)
	fmt.Printf("ω1 = %.2f\n", ω1*180/math.Pi)
	fmt.Printf("ω2 = %.2f\n", ω2*180/math.Pi)
	// Output:
	// DS = -2.194
	// DE = -2.50
	// ω1 = 268.12
	// ω2 = 72.79
}
