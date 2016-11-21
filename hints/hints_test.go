// Copyright 2016 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package hints_test

import (
	"fmt"
	"math"

	"github.com/soniakeys/meeus/base"
)

func Example() {
	// Example 1.a, p. 8
	h := base.FromSexa(' ', 9, 14, 55.8)
	fmt.Printf("%.9f\n", h)
	α := base.RAFromHours(h)
	fmt.Printf("%.5f\n", α.Deg())
	fmt.Printf("%.6f\n", math.Tan(α.Rad()))
	// Output:
	// 9.248833333
	// 138.73250
	// -0.877517
}
