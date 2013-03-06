// Copyright 2012 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package hints_test

import (
	"testing"

	"github.com/soniakeys/meeus/hints"
)

// Meeus gives no test case.
// The test case here is from Wikipedia's entry on Horner's method.
func TestHorner(t *testing.T) {
	y := hints.Horner(3, []float64{-1, 2, -6, 2})
	if y != 5 {
		t.Fatal("Horner")
	}
}
