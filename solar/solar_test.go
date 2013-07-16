// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package solar_test

import (
	"fmt"
	"math"

	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/meeus/julian"
	pp "github.com/soniakeys/meeus/planetposition"
	"github.com/soniakeys/meeus/solar"
)

func ExampleTrue() {
	// Example 25.a, p. 165.
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

func ExampleMeanAnomaly() {
	// Example 25.a, p. 165.
	T := base.J2000Century(julian.CalendarGregorianToJD(1992, 10, 13))
	fmt.Printf("%.5f\n", solar.MeanAnomaly(T)*180/math.Pi)
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
	fmt.Println("λ:", base.NewFmtAngle(solar.ApparentLongitude(T)))
	// Output:
	// λ: 199°54′32″
}

func ExampleApparentEquatorial() {
	// Example 25.a, p. 165.
	jde := julian.CalendarGregorianToJD(1992, 10, 13)
	α, δ := solar.ApparentEquatorial(jde)
	fmt.Printf("α: %.1d\n", base.NewFmtRA(α))
	fmt.Printf("δ: %d\n", base.NewFmtAngle(δ))
	// Output:
	// α: 13ʰ13ᵐ31ˢ.4
	// δ: -7°47′6″
}

func ExampleApparentEquatorialVSOP87() {
	// Example 25.b, p. 169, but as this code uses the full VSOP87 theory,
	// results match those at bottom of p. 165.
	e, err := pp.LoadPlanet(pp.Earth, "")
	if err != nil {
		fmt.Println(err)
		return
	}
	jde := julian.CalendarGregorianToJD(1992, 10, 13)
	α, δ, _ := solar.ApparentEquatorialVSOP87(e, jde)
	fmt.Printf("α: %.3d\n", base.NewFmtRA(α))
	fmt.Printf("δ: %+.2d\n", base.NewFmtAngle(δ))
	// Output:
	// α: 13ʰ13ᵐ30ˢ.749
	// δ: -7°47′1″.74
}
