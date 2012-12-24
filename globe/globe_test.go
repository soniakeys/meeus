package globe_test

import (
	"math"
	"testing"

	"github.com/soniakeys/meeus"
	"github.com/soniakeys/meeus/globe"
)

// p. 83
func TestLatDiff(t *testing.T) {
	φ := meeus.DMSToRad(false, 45, 5, 46.36)
	diff := globe.GeocentricLatitudeDifference(φ)
	answer := meeus.DMSToRad(false, 0, 11, 32.73)
	if math.Abs(answer-diff) > math.Pi/(180*3600*100) {
		t.Fatalf("lat diff %d", diff)
	}
}

// Example 11.a, p 82.
func TestParallax(t *testing.T) {
	// phi = geographic latitude of Palomar
	φ := meeus.DMSToRad(false, 33, 21, 22)
	s, c := globe.Earth76.Parallax(φ, 1706)
	if math.Abs(s-.546861) > 1e-6 || math.Abs(c-.836339) > 1e-6 {
		t.Fatal("parallax")
	}
}

// Example 11.b p 84.
func TestOther(t *testing.T) {
	φ := 42 * math.Pi / 180
	rp := globe.Earth76.RadiusAtLatitude(φ)
	if math.Abs(rp-4747.001) > 1e-3 {
		t.Fatal("radius at lat")
	}
	if math.Abs(globe.OneDegreeOfLongitude(rp)-82.8508) > 1e-4 {
		t.Fatal("degree of long")
	}
	if math.Abs(rp*globe.RotationRate1996_5-.34616) > 1e-5 {
		t.Fatal("linear velocity")
	}
	rm := globe.Earth76.RadiusOfCurvature(φ)
	if math.Abs(rm-6364.033) > 1e-3 {
		t.Fatal("radius of curvature")
	}
	if math.Abs(globe.OneDegreeOfLatitude(rm)-111.0733) > 1e-4 {
		t.Fatal("degree of lat")
	}
}

// Example 11.c p 85.
func TestDistance(t *testing.T) {
	c1 := globe.Coord{
		meeus.DMSToRad(false, 48, 50, 11), // geographic latitude
		meeus.DMSToRad(true, 2, 20, 14),   // geographic longitude
	}
	c2 := globe.Coord{
		meeus.DMSToRad(false, 38, 55, 17),
		meeus.DMSToRad(false, 77, 3, 56),
	}
	if math.Abs(globe.Earth76.Distance(c1, c2)-6181.63) > .01 {
		t.Fatal("distance")
	}
	cos := globe.ApproxAngularDistance(c1, c2)
	if math.Abs(cos-.567146) > 1e-6 {
		t.Fatal("ApproxAngularDistance")
	}
	d := math.Acos(cos)
	if math.Abs(d*180/math.Pi-55.44855) > 1e-5 {
		t.Fatal("Acos")
	}
	if math.Abs(globe.ApproxLinearDistance(d)-6166) > 1 {
		t.Fatal("ApproxLinearDistance", globe.ApproxLinearDistance(d))
	}
}
