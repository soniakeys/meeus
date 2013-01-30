package precess_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/soniakeys/meeus"
	"github.com/soniakeys/meeus/coord"
	"github.com/soniakeys/meeus/julian"
	"github.com/soniakeys/meeus/precess"
)

func ExampleApproxAnnualPrecession() {
	// Example 21.a, p. 132.
	eq := &coord.Equatorial{
		meeus.NewRA(10, 8, 22.3).Rad(),
		meeus.NewAngle(false, 11, 58, 2).Rad(),
	}
	epochFrom := 2000.0
	epochTo := 1978.0
	Δα, Δδ := precess.ApproxAnnualPrecession(eq, epochFrom, epochTo)
	fmt.Printf("%+.3d\n", meeus.NewFmtHourAngle(Δα))
	fmt.Printf("%+.2d\n", meeus.NewFmtAngle(Δδ))

	mα := meeus.NewHourAngle(true, 0, 0, 0.0169).Rad()
	mδ := meeus.NewAngle(false, 0, 0, 0.006).Rad()
	precess.PrecessApprox(eq, eq, epochFrom, epochTo, mα, mδ)
	fmt.Printf("%0.1d\n", meeus.NewFmtRA(eq.RA))
	fmt.Printf("%+0d\n", meeus.NewFmtAngle(eq.Dec))
	// Output:
	// +3ˢ.207
	// -17″.71
	// 10ʰ07ᵐ12ˢ.1
	// +12°04′32″
}

// test example epochs on p. 133 that are not constants in meeus/julian.go
func TestEpoch(t *testing.T) {
	if math.Abs(meeus.BesselianYearToJDE(1950)-2433282.4235) > 1e-4 {
		t.Fatal("B1950")
	}
	if math.Abs((meeus.JulianYearToJDE(2050)-2469807.5)/2469807.5) > 1e-15 {
		t.Fatal("J2050")
	}
}

func ExamplePrecess() {
	// Example 21.b, p. 135.
	eq := &coord.Equatorial{
		meeus.NewRA(2, 44, 11.986).Rad(),
		meeus.NewAngle(false, 49, 13, 42.48).Rad(),
	}
	epochFrom := 2000.0
	jdTo := julian.CalendarGregorianToJD(2028, 11, 13.19)
	epochTo := meeus.JDEToJulianYear(jdTo)
	precess.Precess(eq, eq, epochFrom, epochTo,
		meeus.NewHourAngle(false, 0, 0, 0.03425).Rad(),
		meeus.NewAngle(true, 0, 0, 0.0895).Rad())
	fmt.Printf("%0.3d\n", meeus.NewFmtRA(eq.RA))
	fmt.Printf("%+0.2d\n", meeus.NewFmtAngle(eq.Dec))
	// Output:
	// 2ʰ46ᵐ11ˢ.331
	// +49°20′54″.54
}

// Exercise, p. 136.
func TestPrecess(t *testing.T) {
	eqFrom := &coord.Equatorial{
		meeus.NewRA(2, 31, 48.704).Rad(),
		meeus.NewAngle(false, 89, 15, 50.72).Rad(),
	}
	eqTo := &coord.Equatorial{}
	mα := meeus.NewHourAngle(false, 0, 0, 0.19877).Rad()
	mδ := meeus.NewAngle(true, 0, 0, 0.0152).Rad()
	for _, tc := range []struct {
		α, δ string
		jde  float64
	}{
		{"1 22 33.90", "88 46 26.18", meeus.BesselianYearToJDE(1900)},
		{"3 48 16.43", "89 27 15.38", meeus.JulianYearToJDE(2050)},
		{"5 53 29.17", "89 32 22.18", meeus.JulianYearToJDE(2100)},
	} {
		epochTo := meeus.JDEToJulianYear(tc.jde)
		precess.Precess(eqFrom, eqTo, 2000.0, epochTo, mα, mδ)
		αStr := fmt.Sprintf("%.2x", meeus.NewFmtRA(eqTo.RA))
		δStr := fmt.Sprintf("%.2x", meeus.NewFmtAngle(eqTo.Dec))
		if αStr != tc.α {
			t.Fatal(αStr)
		}
		if δStr != tc.δ {
			t.Fatal(δStr)
		}
	}
}

func ExamplePrecessEcl() {
	// Example 21.c, p. 137.
	eclFrom := &coord.Ecliptic{
		Lat: 1.76549 * math.Pi / 180,
		Lon: 149.48194 * math.Pi / 180,
	}
	eclTo := &coord.Ecliptic{}
	epochFrom := 2000.0
	epochTo := meeus.JDEToJulianYear(julian.CalendarJulianToJD(-214, 6, 30))
	precess.PrecessEcl(eclFrom, eclTo, epochFrom, epochTo, 0, 0)
	fmt.Printf("%.3f\n", eclTo.Lon*180/math.Pi)
	fmt.Printf("%+.3f\n", eclTo.Lat*180/math.Pi)
	// Output:
	// 118.704
	// +1.615
}

func ExampleProperMotion3D() {
	// Example 21.d, p. 141.
	eqFrom := &coord.Equatorial{
		RA:  meeus.NewRA(6, 45, 8.871).Rad(),
		Dec: meeus.NewAngle(true, 16, 42, 57.99).Rad(),
	}
	// 13751 (= 3600*180/pi/15) is conversion factor from seconds (of time)
	// per year to (arc) radians per year.  see top of p. 141.
	mra := -0.03847 / 13751
	mdec := -1.2053 / 206265 // 206265 = 3600*180/pi
	r := 2.64                // given in correct unit
	mr := -7.6 / 977792      // another magic conversion factor
	eqTo := &coord.Equatorial{}
	for _, epoch := range []float64{1000, 0, -1000, -2000, -10000} {
		precess.ProperMotion3D(eqFrom, eqTo, 2000, epoch, r, mr, mra, mdec)
		fmt.Printf("%8.1f  %0.2d  %-0.1d\n", epoch,
			meeus.NewFmtRA(eqTo.RA), meeus.NewFmtAngle(eqTo.Dec))
	}
	// Output:
	//   1000.0  6ʰ45ᵐ47ˢ.19  -16°22′57″.5
	//      0.0  6ʰ46ᵐ25ˢ.32  -16°03′02″.9
	//  -1000.0  6ʰ47ᵐ03ˢ.23  -15°43′15″.4
	//  -2000.0  6ʰ47ᵐ40ˢ.92  -15°23′35″.3
	// -10000.0  6ʰ52ᵐ34ˢ.56  -12°50′37″.8

}
