// Copyright 2013 Sonia Keys
// License: MIT

package parallactic_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/soniakeys/meeus/parallactic"
	"github.com/soniakeys/sexagesimal"
	"github.com/soniakeys/unit"
)

func ExampleEclipticAtHorizon() {
	ε := unit.AngleFromDeg(23.44)
	φ := unit.AngleFromDeg(51)
	θ := unit.TimeFromHour(5)
	λ1, λ2, I := parallactic.EclipticAtHorizon(ε, φ, θ)
	fmt.Println(sexa.FmtAngle(λ1))
	fmt.Println(sexa.FmtAngle(λ2))
	fmt.Println(sexa.FmtAngle(I))
	// Output:
	// 169°21′30″
	// 349°21′30″
	// 61°53′14″
}

func TestDiurnalPathAtHorizon(t *testing.T) {
	φ := unit.AngleFromDeg(40)
	ε := unit.AngleFromDeg(23.44)
	J := parallactic.DiurnalPathAtHorizon(0, φ)
	Jexp := math.Pi/2 - φ
	if math.Abs((J-Jexp).Rad()/Jexp.Rad()) > 1e-15 {
		t.Fatal("0 dec:", sexa.FmtAngle(J))
	}
	J = parallactic.DiurnalPathAtHorizon(ε, φ)
	Jexp = unit.NewAngle(' ', 45, 31, 0)
	if math.Abs((J-Jexp).Rad()/Jexp.Rad()) > 1e-3 {
		t.Fatal("solstice:", sexa.FmtAngle(J))
	}
}
