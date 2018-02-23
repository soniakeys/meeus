// Copyright 2013 Sonia Keys
// License: MIT

package globe_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/soniakeys/meeus/v3/globe"
	"github.com/soniakeys/sexagesimal"
	"github.com/soniakeys/unit"
)

func ExampleEllipsoid_ParallaxConstants() {
	// Example 11.a, p 82.
	// phi = geographic latitude of Palomar
	φ := unit.NewAngle(' ', 33, 21, 22)
	s, c := globe.Earth76.ParallaxConstants(φ, 1706)
	fmt.Printf("ρ sin φ′ = %+.6f\n", s)
	fmt.Printf("ρ cos φ′ = %+.6f\n", c)
	// Output:
	// ρ sin φ′ = +0.546861
	// ρ cos φ′ = +0.836339
}

// p. 83
func TestLatDiff(t *testing.T) {
	φ0 := unit.NewAngle(' ', 45, 5, 46.36)
	diff := globe.GeocentricLatitudeDifference(φ0)
	if f := fmt.Sprintf("%.2d", sexa.FmtAngle(diff)); f != "11′32″.73" {
		t.Fatal(f)
	}
}

func ExampleEllipsoid_RadiusAtLatitude() {
	// Example 11.b p 84.
	φ := unit.AngleFromDeg(42)
	rp := globe.Earth76.RadiusAtLatitude(φ)
	fmt.Printf("Rp = %.3f km\n", rp)
	fmt.Printf("1° of longitude = %.4f km\n", globe.OneDegreeOfLongitude(rp))
	fmt.Printf("linear velocity = ωRp = %.5f km/second\n",
		rp*globe.RotationRate1996_5)
	rm := globe.Earth76.RadiusOfCurvature(φ)
	fmt.Printf("Rm = %.3f km\n", rm)
	fmt.Printf("1° of latitude = %.4f km\n", globe.OneDegreeOfLatitude(rm))
	// Output:
	// Rp = 4747.001 km
	// 1° of longitude = 82.8508 km
	// linear velocity = ωRp = 0.34616 km/second
	// Rm = 6364.033 km
	// 1° of latitude = 111.0733 km
}

func ExampleEllipsoid_Distance() {
	// Example 11.c p 85.
	c1 := globe.Coord{
		unit.NewAngle(' ', 48, 50, 11), // geographic latitude
		unit.NewAngle('-', 2, 20, 14),  // geographic longitude
	}
	c2 := globe.Coord{
		unit.NewAngle(' ', 38, 55, 17),
		unit.NewAngle(' ', 77, 3, 56),
	}
	fmt.Printf("%.2f km\n", globe.Earth76.Distance(c1, c2))
	cos := globe.ApproxAngularDistance(c1, c2)
	fmt.Printf("cos d = %.6f\n", cos)
	d := unit.Angle(math.Acos(cos))
	fmt.Printf("    d = %.5j\n", sexa.FmtAngle(d))
	fmt.Printf("    s = %.0f km\n", globe.ApproxLinearDistance(d))
	// Output:
	// 6181.63 km
	// cos d = 0.567146
	//     d = 55°.44855
	//     s = 6166 km
}
