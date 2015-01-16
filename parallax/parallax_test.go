package parallax_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/meeus/julian"
	"github.com/soniakeys/meeus/parallax"
	"github.com/soniakeys/meeus/sidereal"
	"github.com/soniakeys/sexagesimal"
)

func ExampleHorizontal() {
	// Example 40.a, p. 280
	π := parallax.Horizontal(.37276)
	fmt.Printf("%.3f\n", π*180/math.Pi*60*60)
	// Output:
	// 23.592
}

func ExampleTopocentric() {
	// Example 40.a, p. 280
	α, δ := parallax.Topocentric(339.530208*math.Pi/180,
		-15.771083*math.Pi/180,
		.37276, .546861, .836339,
		sexa.NewHourAngle(false, 7, 47, 27).Rad(),
		julian.CalendarGregorianToJD(2003, 8, 28+(3+17./60)/24))
	fmt.Printf("α' = %.2d\n", sexa.NewFmtRA(α))
	fmt.Printf("δ' = %.1d\n", sexa.NewFmtAngle(δ))
	// Output:
	// α' = 22ʰ38ᵐ8ˢ.54
	// δ' = -15°46′30″.0
}

func ExampleTopocentric2() {
	// Example 40.a, p. 280
	Δα, Δδ := parallax.Topocentric2(339.530208*math.Pi/180,
		-15.771083*math.Pi/180,
		.37276, .546861, .836339,
		sexa.NewHourAngle(false, 7, 47, 27).Rad(),
		julian.CalendarGregorianToJD(2003, 8, 28+(3+17./60)/24))
	fmt.Printf("Δα = %.2f sec of RA\n", Δα*180/math.Pi*60*60/15)
	fmt.Printf("Δδ = %.1f sec\n", Δδ*180/math.Pi*60*60)
	// Output:
	// Δα = 1.29 sec of RA
	// Δδ = -14.1 sec
}

func TestTopocentric3(t *testing.T) {
	// same test case as example 40.a, p. 280
	α := 339.530208 * math.Pi / 180
	δ := -15.771083 * math.Pi / 180
	Δ := .37276
	ρsφʹ := .546861
	ρcφʹ := .836339
	L := sexa.NewHourAngle(false, 7, 47, 27).Rad()
	jde := julian.CalendarGregorianToJD(2003, 8, 28+(3+17./60)/24)
	// reference result
	αʹ, δʹ1 := parallax.Topocentric(α, δ, Δ, ρsφʹ, ρcφʹ, L, jde)
	// result to test
	Hʹ, δʹ3 := parallax.Topocentric3(α, δ, Δ, ρsφʹ, ρcφʹ, L, jde)
	// test
	θ0 := sexa.Time(sidereal.Apparent(jde)).Rad()
	if math.Abs(base.PMod(Hʹ-(θ0-L-αʹ)+1, 2*math.Pi)-1) > 1e-15 {
		t.Fatal(Hʹ, θ0-L-αʹ)
	}
	if math.Abs(δʹ3-δʹ1) > 1e-15 {
		t.Fatal(δʹ3, δʹ1)
	}
}

func TestTopocentricEcliptical(t *testing.T) {
	// exercise, p. 282
	λʹ, βʹ, sʹ := parallax.TopocentricEcliptical(sexa.NewAngle(false,
		181, 46, 22.5).Rad(),
		sexa.NewAngle(false, 2, 17, 26.2).Rad(),
		sexa.NewAngle(false, 0, 16, 15.5).Rad(),
		sexa.NewAngle(false, 50, 5, 7.8).Rad(), 0,
		sexa.NewAngle(false, 23, 28, 0.8).Rad(),
		sexa.NewAngle(false, 209, 46, 7.9).Rad(),
		sexa.NewAngle(false, 0, 59, 27.7).Rad())
	λʹa := sexa.NewAngle(false, 181, 48, 5).Rad()
	βʹa := sexa.NewAngle(false, 1, 29, 7.1).Rad()
	sʹa := sexa.NewAngle(false, 0, 16, 25.5).Rad()
	if math.Abs(λʹ-λʹa) > .1/60/60*math.Pi/180 {
		t.Fatal(λʹ, λʹa)
	}
	if math.Abs(βʹ-βʹa) > .1/60/60*math.Pi/180 {
		t.Fatal(βʹ, βʹa)
	}
	if math.Abs(sʹ-sʹa) > .1/60/60*math.Pi/180 {
		t.Fatal(sʹ, sʹa)
	}
}
