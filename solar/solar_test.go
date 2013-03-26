// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package solar_test

import (
	"fmt"
	"math"

	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/meeus/julian"
	"github.com/soniakeys/meeus/solar"
)

func ExampleTrue() {
	jd := julian.CalendarGregorianToJD(1992, 10, 13)
	fmt.Printf("JDE: %.1f\n", jd)
	T := base.J2000Century(jd)
	fmt.Printf("T:   %.9f\n", T)
	s, _ := solar.True(T)
	fmt.Printf("☉:   %.5f\n", (s * 180 / math.Pi))
	// Output:
	// JDE: 2448908.5
	// T:   -0.072183436
	// ☉:   199.90987
}
