// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package refraction_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/soniakeys/meeus/refraction"
)

func ExampleBennett() {
	// Example 16.a, p. 107.
	h0 := .5 * math.Pi / 180
	R := refraction.Bennett(h0)
	const cMin = 60 * 180 / math.Pi
	fmt.Printf("R Lower: %.3f\n", R*cMin)
	hLower := h0 - R
	fmt.Printf("h Lower: %.3f\n", hLower*cMin)
	hUpper := hLower + 32*math.Pi/(180*60)
	fmt.Printf("h Upper: %.3f\n", hUpper*cMin)
	Rh := refraction.Saemundsson(hUpper)
	fmt.Printf("R Upper: %.3f\n", Rh*cMin)
	// Output:
	// R Lower: 28.754
	// h Lower: 1.246
	// h Upper: 33.246
	// R Upper: 24.618
}

// Test two values for zenith given on p. 106.
func TestBennett(t *testing.T) {
	R := refraction.Bennett(math.Pi / 2)
	const cSec = 3600 * 180 / math.Pi
	if math.Abs(.08+R*cSec) > .01 {
		t.Fatal(R * cSec)
	}
	R = refraction.Bennett2(math.Pi / 2)
	if math.Abs(.89+R*cSec) > .01 {
		t.Fatal(R * cSec)
	}

}
