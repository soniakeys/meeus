// Copyright 2013 Sonia Keys
// License: MIT

// +build !nopp

package planetposition_test

import (
	"fmt"
	"testing"

	"github.com/soniakeys/meeus/julian"
	pp "github.com/soniakeys/meeus/planetposition"
	"github.com/soniakeys/sexagesimal"
	"github.com/soniakeys/unit"
)

func ExampleV87Planet_Position2000() {
	// Mars 1899 spherical data from vsop87.chk.
	jd := 2415020.0
	p, err := pp.LoadPlanet(pp.Mars)
	if err != nil {
		fmt.Println(err)
		return
	}
	l, b, r := p.Position2000(jd)
	fmt.Printf("L = %.10f rad\n", l)
	fmt.Printf("B = %.10f rad\n", b)
	fmt.Printf("R = %.10f AU\n", r)
	// Output:
	// L = 5.0185792656 rad
	// B = -0.0274073500 rad
	// R = 1.4218777718 AU
}

func ExampleV87Planet_Position() {
	// Example 32.a, p. 219
	jd := julian.CalendarGregorianToJD(1992, 12, 20)
	p, err := pp.LoadPlanet(pp.Venus)
	if err != nil {
		fmt.Println(err)
		return
	}
	l, b, r := p.Position(jd)
	fmt.Printf("L = %+.5j\n", sexa.FmtAngle(l))
	fmt.Printf("B = %+.5j\n", sexa.FmtAngle(b))
	fmt.Printf("R = %.6f AU\n", r)
	// Output:
	// L = +26°.11412
	// B = -2°.62060
	// R = 0.724602 AU
}

func ExampleToFK5() {
	// In example 33.a, p. 226
	jd := 2448976.5
	λ := unit.AngleFromDeg(313.07689) // (the value from mid-page)
	β := unit.AngleFromDeg(-2.08489)
	λ5, β5 := pp.ToFK5(λ, β, jd)
	// recovering Δs,
	Δλ := sexa.FmtAngle(λ5 - λ)
	Δβ := sexa.FmtAngle(β5 - β)
	fmt.Printf("λ = %3.5j\n", sexa.FmtAngle(λ))
	fmt.Printf("β = %3.5j\n", sexa.FmtAngle(β))
	fmt.Printf("Δλ = %+.5d = %+.5j\n", Δλ, Δλ)
	fmt.Printf("Δβ = %+.5d = %+.5j\n", Δβ, Δβ)
	fmt.Printf("FK5 λ = %3.5j\n", sexa.FmtAngle(λ5))
	fmt.Printf("FK5 β = %3.5j\n", sexa.FmtAngle(β5))
	// Output:
	// λ =  313°.07689
	// β =   -2°.08489
	// Δλ = -0″.09027 = -°.00003
	// Δβ = +0″.05535 = +°.00002
	// FK5 λ =  313°.07686
	// FK5 β =   -2°.08487
}

func TestFK5(t *testing.T) {
	// The effect of issue #10 is in like the 10th decimal place.
	// ExampleToFK5 above doesn't show it.  We don't have the correct answer, but we can
	// at least test reproducibility with enough precision to show the effect of the bug fix.
	jd := 2448976.5
	λ := unit.AngleFromDeg(313.07689)
	β := unit.AngleFromDeg(-2.08489)
	λ5, β5 := pp.ToFK5(λ, β, jd)
	Δλ := fmt.Sprintf("%.12s", sexa.FmtAngle(λ5-λ))
	Δβ := fmt.Sprintf("%.12s", sexa.FmtAngle(β5-β))
	Δλwant := "-0.090265798293″"
	Δβwant := "0.055352515765″"
	if Δλ != Δλwant {
		t.Error(Δλ)
	}
	if Δβ != Δβwant {
		t.Error(Δβ)
	}
}
