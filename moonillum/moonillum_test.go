// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package moonillum_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/meeus/julian"
	"github.com/soniakeys/meeus/moonillum"
	"github.com/soniakeys/meeus/moonposition"
	"github.com/soniakeys/meeus/solar"
)

func ExamplePhaseAngleEq() {
	i := moonillum.PhaseAngleEq(
		base.RAFromDeg(134.6885),
		base.AngleFromDeg(13.7684),
		368410,
		base.RAFromDeg(20.6579),
		base.AngleFromDeg(8.6964),
		149971520)
	fmt.Printf("i = %.4f\n", i.Deg())
	// Output:
	// i = 69.0756
}

func ExamplePhaseAngleEq2() {
	const p = math.Pi / 180
	i := moonillum.PhaseAngleEq2(
		base.RAFromDeg(134.6885),
		base.AngleFromDeg(13.7684),
		base.RAFromDeg(20.6579),
		base.AngleFromDeg(8.6964))
	k := base.Illuminated(i)
	fmt.Printf("k = %.4f\n", k)
	// Output:
	// k = 0.6775
}

func TestPhaseAngleEcl(t *testing.T) {
	j := julian.CalendarGregorianToJD(1992, 4, 12)
	λ, β, Δ := moonposition.Position(j)
	T := base.J2000Century(j)
	λ0 := solar.ApparentLongitude(T)
	R := solar.Radius(T) * base.AU
	i := moonillum.PhaseAngleEcl(λ, β, Δ, λ0, R)
	ref := base.AngleFromDeg(69.0756)
	if math.Abs(((i - ref) / ref).Rad()) > 1e-4 {
		t.Errorf("i = %.4f", i.Deg())
	}
}

func TestPhaseAngleEcl2(t *testing.T) {
	j := julian.CalendarGregorianToJD(1992, 4, 12)
	λ, β, _ := moonposition.Position(j)
	λ0 := solar.ApparentLongitude(base.J2000Century(j))
	i := moonillum.PhaseAngleEcl2(λ, β, λ0)
	k := base.Illuminated(i)
	ref := .6775
	if math.Abs(k-ref) > 1e-4 {
		t.Errorf("k = %.4f", k)
	}
}

func ExamplePhaseAngle3() {
	i := moonillum.PhaseAngle3(julian.CalendarGregorianToJD(1992, 4, 12))
	k := base.Illuminated(i)
	fmt.Printf("i = %.2f\n", i.Deg())
	fmt.Printf("k = %.4f\n", k)
	// Output:
	// i = 68.88
	// k = 0.6801
}
