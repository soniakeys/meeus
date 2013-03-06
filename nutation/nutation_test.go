// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package nutation_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/soniakeys/meeus/common"
	"github.com/soniakeys/meeus/julian"
	"github.com/soniakeys/meeus/nutation"
)

func ExampleNutation() {
	// Example 22.a, p. 148.
	jd := julian.CalendarGregorianToJD(1987, 4, 10)
	Δψ, Δε := nutation.Nutation(jd)
	ε0 := nutation.MeanObliquity(jd)
	ε := ε0 + Δε
	fmt.Printf("%+.3d\n", common.NewFmtAngle(Δψ))
	fmt.Printf("%+.3d\n", common.NewFmtAngle(Δε))
	fmt.Printf("%.3d\n", common.NewFmtAngle(ε0))
	fmt.Printf("%.3d\n", common.NewFmtAngle(ε))
	// Output:
	// -3″.788
	// +9″.443
	// 23°26′27″.407
	// 23°26′36″.850
}

func TestApproxNutation(t *testing.T) {
	jd := julian.CalendarGregorianToJD(1987, 4, 10)
	Δψ, Δε := nutation.ApproxNutation(jd)
	if math.Abs(Δψ*(180/math.Pi)*3600+3.788) > .5 {
		t.Fatal(Δψ * (180 / math.Pi) * 3600)
	}
	if math.Abs(Δε*(180/math.Pi)*3600-9.443) > .1 {
		t.Fatal(Δε * (180 / math.Pi) * 3600)
	}
}

func TestIAUvsLaskar(t *testing.T) {
	for _, y := range []int{1000, 2000, 3000} {
		jd := julian.CalendarGregorianToJD(y, 0, 0)
		i := nutation.MeanObliquity(jd)
		l := nutation.MeanObliquityLaskar(jd)
		if math.Abs(i-l)*(180/math.Pi)*3600 > 1 {
			t.Fatal(y)
		}
	}
	for _, y := range []int{0, 4000} {
		jd := julian.CalendarGregorianToJD(y, 0, 0)
		i := nutation.MeanObliquity(jd)
		l := nutation.MeanObliquityLaskar(jd)
		if math.Abs(i-l)*(180/math.Pi)*3600 > 10 {
			t.Fatal(y)
		}
	}
}
