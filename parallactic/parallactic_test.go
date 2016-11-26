// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package parallactic_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/meeus/parallactic"
	"github.com/soniakeys/sexagesimal"
)

func ExampleEclipticAtHorizon() {
	ε := base.AngleFromDeg(23.44)
	φ := base.AngleFromDeg(51)
	θ := base.TimeFromHour(5)
	λ1, λ2, I := parallactic.EclipticAtHorizon(ε, φ, θ)
	fmt.Println(sexa.Angle(λ1).Fmt())
	fmt.Println(sexa.Angle(λ2).Fmt())
	fmt.Println(sexa.Angle(I).Fmt())
	// Output:
	// 169°21′30″
	// 349°21′30″
	// 61°53′14″
}

func TestDiurnalPathAtHorizon(t *testing.T) {
	φ := base.AngleFromDeg(40)
	ε := base.AngleFromDeg(23.44)
	J := parallactic.DiurnalPathAtHorizon(0, φ).Rad()
	Jexp := math.Pi/2 - φ.Rad()
	if math.Abs((J-Jexp)/Jexp) > 1e-15 {
		t.Fatal("0 dec:", sexa.Angle(J).Fmt())
	}
	J = parallactic.DiurnalPathAtHorizon(ε, φ).Rad()
	Jexp = base.NewAngle(' ', 45, 31, 0).Rad()
	if math.Abs((J-Jexp)/Jexp) > 1e-3 {
		t.Fatal("solstace:", sexa.Angle(J).Fmt())
	}
}
