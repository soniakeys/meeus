package interp_test

import (
	"math"
	"testing"

	"github.com/soniakeys/meeus"
	"github.com/soniakeys/meeus/interp"
)

// Example 3.a, p. 25.
func TestLen3Interpolate(t *testing.T) {
	x1 := 7.
	x3 := 9.
	yTable := []float64{
		.884226,
		.877366,
		.870531,
	}
	x := 8 + (4+21/60.)/24.
	y, err := interp.Len3Interpolate(x, x1, x3, yTable, false)
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(y-.876125) > 1e-6 {
		t.Fatal("y:", y)
	}
}

// Example 3.b p. 26.
func TestLen3Extremum(t *testing.T) {
	x1 := 12.
	x3 := 20.
	yTable := []float64{
		1.3814294,
		1.3812213,
		1.3812453,
	}
	x, y, err := interp.Len3Extremum(x1, x3, yTable)
	if err != nil {
		t.Fatal(err)
	}
	// (Meeus looses a decimal place here by rounding nm before multiplying
	// by dx.)
	if math.Abs(x-17.5864) > 1e-4 {
		t.Fatal("wrong x:", x)
	}
	if math.Abs(y-1.3812030) > 1e-7 {
		t.Fatal("wrong y", y)
	}
}

func TestLen3Zero(t *testing.T) {
	x1 := 26.
	x3 := 28.
	yTable := []float64{
		meeus.DMSToDeg(true, 0, 28, 13.4) * 180 / math.Pi,
		meeus.DMSToDeg(false, 0, 6, 46.3) * 180 / math.Pi,
		meeus.DMSToDeg(false, 0, 38, 23.2) * 180 / math.Pi,
	}
	x, err := interp.Len3Zero(x1, x3, yTable, false)
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(x-26.79873) > 1e-5 {
		t.Fatal("x:", x)
	}
}

func TestLen3ZeroStrong(t *testing.T) {
	x1 := -1.
	x3 := 1.
	yTable := []float64{-2, 3, 2}
	x, err := interp.Len3Zero(x1, x3, yTable, true)
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(x-(-.720759220056)) > 1e-12 {
		t.Fatal("x:", x)
	}
}
