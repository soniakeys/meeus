package interp_test

import (
	"math"
	"testing"

	"github.com/soniakeys/meeus"
	"github.com/soniakeys/meeus/interp"
)

// Example 3.a, p. 25.
func TestDiff2(t *testing.T) {
	x0 := 5.
	xLast := 9.
	yTable := []float64{
		.898013,
		.891109,
		.884226,
		.877366,
		.870531,
	}
	x := 8 + (4+21/60.)/24.
	y := interp.Diff2(yTable, x0, xLast, x)
	if math.Abs(y-.876125) > 1e-6 {
		t.Fatal("Diff2")
	}
}

// Example 3.b p. 26.
func TestExtremum2(t *testing.T) {
	x0 := 12.
	dx := 4.
	yTable := []float64{
		1.3814294,
		1.3812213,
		1.3812453,
	}
	x, y, ok := interp.Extremum2(yTable, x0, dx)
	if !ok {
		t.Fatal("Extremum not found")
	}
	// (Meeus looses a decimal place here by rounding nm before multiplying
	// by dx.)
	if math.Abs(x-17.5864) > 1e-4 {
		t.Fatal("Wrong xm:", x)
	}
	if math.Abs(y-1.3812030) > 1e-7 {
		t.Fatal("Wrong ym", y)
	}
}

func TestZero2(t *testing.T) {
	x0 := 26.
	dx := 1.
	yTable := []float64{
		meeus.DMSToRad(true, 0, 28, 13.4),
		meeus.DMSToRad(false, 0, 6, 46.3),
		meeus.DMSToRad(false, 0, 38, 23.2),
	}
	x, ok := interp.Zero2(yTable, x0, dx)
	if !ok {
		t.Fatal("Zero2 not ok")
	}
	if math.Abs(x-26.79873) > 1e-5 {
		t.Fatal("Zero2", x)
	}
}
