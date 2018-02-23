// Copyright 2013 Sonia Keys
// License: MIT

package refraction_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/soniakeys/meeus/v3/refraction"
	"github.com/soniakeys/sexagesimal"
	"github.com/soniakeys/unit"
)

func Example() {
	// Example 16.a, p. 107.
	h0 := unit.AngleFromDeg(.5) // apparent lower limb of Sun
	R := refraction.Bennett(h0)
	fmt.Printf("R:  %.3m\n", sexa.FmtAngle(R))
	tL := h0 - R // true lower
	fmt.Printf("L:  %.3m\n", sexa.FmtAngle(tL))
	tU := tL + unit.AngleFromMin(32) // add true diameter of sun
	fmt.Printf("U:  %.3m\n", sexa.FmtAngle(tU))
	R = refraction.Saemundsson(tU)
	fmt.Printf("R:  %.3m\n", sexa.FmtAngle(R))
	// Output:
	// R:  28.754′
	// L:  1.246′
	// U:  33.246′
	// R:  24.618′
}

// Test two values for zenith given on p. 106.
func TestBennett(t *testing.T) {
	R := refraction.Bennett(math.Pi / 2)
	const cSec = 3600 * 180 / math.Pi
	if math.Abs(.08+R.Rad()*cSec) > .01 {
		t.Fatal(R * cSec)
	}
	R = refraction.Bennett2(math.Pi / 2)
	if math.Abs(.89+R.Rad()*cSec) > .01 {
		t.Fatal(R * cSec)
	}

}
