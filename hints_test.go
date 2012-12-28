package meeus_test

import (
	"testing"

	"github.com/soniakeys/meeus"
)

// Meeus gives no test case.
// The test case here is from Wikipedia's entry on Horner's method.
func TestHorner(t *testing.T) {
	y := meeus.Horner(3, []float64{-1, 2, -6, 2})
	if y != 5 {
		t.Fatal("Horner")
	}
}
