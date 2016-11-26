// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package parallax_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/meeus/julian"
	"github.com/soniakeys/meeus/moonposition"
	"github.com/soniakeys/meeus/parallax"
	"github.com/soniakeys/meeus/sidereal"
	"github.com/soniakeys/sexagesimal"
)

func ExampleHorizontal() {
	// Example 40.a, p. 280
	π := parallax.Horizontal(.37276)
	fmt.Printf("%.3s", sexa.Angle(π).Fmt())
	// Output:
	// 23.592″
}

func TestHorizontal(t *testing.T) {
	// example from moonposition.Parallax, ch 47, p. 342
	_, _, Δ := moonposition.Position(julian.CalendarGregorianToJD(1992, 4, 12))
	π := parallax.Horizontal(Δ / base.AU).Deg()
	want := .99199
	// we don't quite get all the digits here.
	// for close objects we need that Arcsin that's in moonposition.Parallax.
	if math.Abs(π-want) > .0001 {
		t.Fatal("got", π, "want", want)
	}
}

func ExampleTopocentric() {
	// Example 40.a, p. 280
	α, δ := parallax.Topocentric(
		base.RAFromDeg(339.530208),
		base.AngleFromDeg(-15.771083),
		.37276, .546861, .836339,
		base.Angle(base.NewHourAngle(' ', 7, 47, 27)),
		julian.CalendarGregorianToJD(2003, 8, 28+
			base.NewTime(' ', 3, 17, 0).Day()))
	fmt.Printf("α' = %.2d\n", sexa.RA(α).Fmt())
	fmt.Printf("δ' = %.1d\n", sexa.Angle(δ).Fmt())
	// Output:
	// α' = 22ʰ38ᵐ8ˢ.54
	// δ' = -15°46′30″.0
}

func ExampleTopocentric2() {
	// Example 40.a, p. 280
	Δα, Δδ := parallax.Topocentric2(
		base.RAFromDeg(339.530208),
		base.AngleFromDeg(-15.771083),
		.37276, .546861, .836339,
		base.Angle(base.NewHourAngle(' ', 7, 47, 27)),
		julian.CalendarGregorianToJD(2003, 8, 28+
			base.NewTime(' ', 3, 17, 0).Day()))
	fmt.Printf("Δα = %.2s (sec of RA)\n", sexa.RA(Δα).Fmt())
	fmt.Printf("Δδ = %.1s (sec of Dec)\n", sexa.Angle(Δδ).Fmt())
	// Output:
	// Δα = 1.29ˢ (sec of RA)
	// Δδ = -14.1″ (sec of Dec)
}

func ExampleTopocentric3() {
	// same test case as example 40.a, p. 280
	α := base.RAFromDeg(339.530208)
	δ := base.AngleFromDeg(-15.771083)
	Δ := .37276
	ρsφʹ := .546861
	ρcφʹ := .836339
	L := base.Angle(base.NewHourAngle(' ', 7, 47, 27))
	jde := julian.CalendarGregorianToJD(2003, 8, 28+
		base.NewTime(' ', 3, 17, 0).Day())
	Hʹ, δʹ := parallax.Topocentric3(α, δ, Δ, ρsφʹ, ρcφʹ, L, jde)
	fmt.Printf("Hʹ = %.2d\n", sexa.HourAngle(Hʹ).Fmt())
	θ0 := sidereal.Apparent(jde)
	αʹ := base.RAFromRad(θ0.Rad() - L.Rad() - Hʹ.Rad())
	// same result as example 40.a, p. 280
	fmt.Printf("αʹ = %.2d\n", sexa.RA(αʹ).Fmt())
	fmt.Printf("δʹ = %.1d\n", sexa.Angle(δʹ).Fmt())
	// Output:
	// Hʹ = -4ʰ44ᵐ50ˢ.28
	// αʹ = 22ʰ38ᵐ8ˢ.54
	// δʹ = -15°46′30″.0
}

func ExampleTopocentricEcliptical() {
	// exercise, p. 282
	λʹ, βʹ, sʹ := parallax.TopocentricEcliptical(
		base.NewAngle(' ', 181, 46, 22.5),
		base.NewAngle(' ', 2, 17, 26.2),
		base.NewAngle(' ', 0, 16, 15.5),
		base.NewAngle(' ', 50, 5, 7.8),
		0,
		base.NewAngle(' ', 23, 28, 0.8),
		base.NewAngle(' ', 209, 46, 7.9).Time(),
		base.NewAngle(' ', 0, 59, 27.7))
	fmt.Printf("λʹ = %.1s\n", sexa.Angle(λʹ).Fmt())
	fmt.Printf("βʹ = %+.1s\n", sexa.Angle(βʹ).Fmt())
	fmt.Printf("sʹ = %.1s\n", sexa.Angle(sʹ).Fmt())
	// Output:
	// λʹ = 181°48′5.0″
	// βʹ = +1°29′7.1″
	// sʹ = 16′25.5″
}
