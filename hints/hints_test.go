package hints_test

import (
	"math"
	"testing"

	"github.com/soniakeys/meeus/hints"
)

func TestSexagesimal(t *testing.T) {
	// example p. 7
	a := hints.DMSToDeg(false, 23, 26, 49)
	if math.Abs(a-23.44694444) > 1e-8 {
		t.Fatal("DMSToRad")
	}
	// example p. 8
	a = hints.HMSToRad(false, 9, 14, 55.8)
	if math.Abs(math.Tan(a) - -.877517) > 1e-6 {
		t.Fatal("HMSToRad")
	}
}

// Meeus gives no test case.
// The test case here is from Wikipedia's entry on Horner's method.
func TestHorner(t *testing.T) {
	y := hints.Horner([]float64{-1, 2, -6, 2}, 3)
	if y != 5 {
		t.Fatal("Horner")
	}
}
