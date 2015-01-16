// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package parallactic_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/soniakeys/meeus/parallactic"
	"github.com/soniakeys/sexagesimal"
)

func ExampleEclipticAtHorizon() {
	ε := 23.44 * math.Pi / 180
	φ := 51 * math.Pi / 180
	θ := 75 * math.Pi / 180
	λ1, λ2, I := parallactic.EclipticAtHorizon(ε, φ, θ)
	fmt.Println(sexa.NewFmtAngle(λ1))
	fmt.Println(sexa.NewFmtAngle(λ2))
	fmt.Println(sexa.NewFmtAngle(I))
	// Output:
	// 169°21′30″
	// 349°21′30″
	// 61°53′14″
}

func TestDiurnalPathAtHorizon(t *testing.T) {
	φ := 40 * math.Pi / 180
	ε := 23.44 * math.Pi / 180
	J := parallactic.DiurnalPathAtHorizon(0, φ)
	Jexp := math.Pi/2 - φ
	if math.Abs((J-Jexp)/Jexp) > 1e-15 {
		t.Fatal("0 dec:", sexa.NewFmtAngle(J))
	}
	J = parallactic.DiurnalPathAtHorizon(ε, φ)
	Jexp = sexa.NewAngle(false, 45, 31, 0).Rad()
	if math.Abs((J-Jexp)/Jexp) > 1e-3 {
		t.Fatal("solstace:", sexa.NewFmtAngle(J))
	}
}
