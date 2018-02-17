// Copyright 2013 Sonia Keys
// License: MIT

package solar_test

import (
	"fmt"

	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/meeus/julian"
	"github.com/soniakeys/meeus/solar"
	"github.com/soniakeys/sexagesimal"
)

func ExampleTrue() {
	// Example 25.a, p. 165.
	jd := julian.CalendarGregorianToJD(1992, 10, 13)
	fmt.Printf("JDE: %.1f\n", jd)
	T := base.J2000Century(jd)
	fmt.Printf("T:   %.9f\n", T)
	s, _ := solar.True(T)
	fmt.Printf("☉:   %.5f\n", s.Deg())
	// Output:
	// JDE: 2448908.5
	// T:   -0.072183436
	// ☉:   199.90987
}

func ExampleMeanAnomaly() {
	// Example 25.a, p. 165.
	T := base.J2000Century(julian.CalendarGregorianToJD(1992, 10, 13))
	fmt.Printf("%.5f\n", solar.MeanAnomaly(T).Deg())
	// Output:
	// -2241.00603
}

func ExampleEccentricity() {
	// Example 25.a, p. 165.
	T := base.J2000Century(julian.CalendarGregorianToJD(1992, 10, 13))
	fmt.Printf("%.9f\n", solar.Eccentricity(T))
	// Output:
	// 0.016711668
}

func ExampleRadius() {
	// Example 25.a, p. 165.
	T := base.J2000Century(julian.CalendarGregorianToJD(1992, 10, 13))
	fmt.Printf("%.5f AU\n", solar.Radius(T))
	// Output:
	// 0.99766 AU
}

func ExampleApparentLongitude() {
	// Example 25.a, p. 165.
	T := base.J2000Century(julian.CalendarGregorianToJD(1992, 10, 13))
	fmt.Println("λ:", sexa.FmtAngle(solar.ApparentLongitude(T)))
	// Output:
	// λ: 199°54′32″
}

func ExampleApparentEquatorial() {
	// Example 25.a, p. 165.
	jde := julian.CalendarGregorianToJD(1992, 10, 13)
	α, δ := solar.ApparentEquatorial(jde)
	fmt.Printf("α: %.1d\n", sexa.FmtRA(α))
	fmt.Printf("δ: %d\n", sexa.FmtAngle(δ))
	// Output:
	// α: 13ʰ13ᵐ31ˢ.4
	// δ: -7°47′6″
}
