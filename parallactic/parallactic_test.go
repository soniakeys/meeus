package parallactic_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/soniakeys/meeus"
	"github.com/soniakeys/meeus/parallactic"
)

func ExampleEclipticAtHorizon() {
	ε := 23.44 * math.Pi / 180
	φ := 51 * math.Pi / 180
	θ := 75 * math.Pi / 180
	λ1, λ2, I := parallactic.EclipticAtHorizon(ε, φ, θ)
	fmt.Println(meeus.NewFmtAngle(λ1))
	fmt.Println(meeus.NewFmtAngle(λ2))
	fmt.Println(meeus.NewFmtAngle(I))
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
		t.Fatal("0 dec:", meeus.NewFmtAngle(J))
	}
	J = parallactic.DiurnalPathAtHorizon(ε, φ)
	Jexp = meeus.NewAngle(false, 45, 31, 0).Rad()
	if math.Abs((J-Jexp)/Jexp) > 1e-3 {
		t.Fatal("solstace:", meeus.NewFmtAngle(J))
	}
}
