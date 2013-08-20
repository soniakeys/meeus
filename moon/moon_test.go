package moon_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/meeus/julian"
	"github.com/soniakeys/meeus/moon"
)

func ExamplePosition() {
	// Example 47.a, p. 342.
	λ, β, Δ := moon.Position(julian.CalendarGregorianToJD(1992, 4, 12))
	fmt.Printf("λ = %.6f\n", λ*180/math.Pi)
	fmt.Printf("β = %.6f\n", β*180/math.Pi)
	fmt.Printf("Δ = %.1f\n", Δ)
	// Output:
	// λ = 133.162655
	// β = -3.229126
	// Δ = 368409.7
}

func ExampleParallax() {
	// Example 47.a, p. 342.
	_, _, Δ := moon.Position(julian.CalendarGregorianToJD(1992, 4, 12))
	π := moon.Parallax(Δ)
	fmt.Printf("π = %.6f\n", π*180/math.Pi)
	// Output:
	// π = 0.991990
}

// Test data p. 344.
var n0 = []float64{
	julian.CalendarGregorianToJD(1913, 5, 27),
	julian.CalendarGregorianToJD(1932, 1, 6),
	julian.CalendarGregorianToJD(1950, 8, 17),
	julian.CalendarGregorianToJD(1969, 3, 29),
	julian.CalendarGregorianToJD(1987, 11, 8),
	julian.CalendarGregorianToJD(2006, 6, 19),
	julian.CalendarGregorianToJD(2025, 1, 29),
	julian.CalendarGregorianToJD(2043, 9, 10),
	julian.CalendarGregorianToJD(2062, 4, 22),
	julian.CalendarGregorianToJD(2080, 12, 1),
	julian.CalendarGregorianToJD(2099, 7, 13),
}

var n180 = []float64{
	julian.CalendarGregorianToJD(1922, 9, 16),
	julian.CalendarGregorianToJD(1941, 4, 27),
	julian.CalendarGregorianToJD(1959, 12, 7),
	julian.CalendarGregorianToJD(1978, 7, 19),
	julian.CalendarGregorianToJD(1997, 2, 27),
	julian.CalendarGregorianToJD(2015, 10, 10),
	julian.CalendarGregorianToJD(2034, 5, 21),
	julian.CalendarGregorianToJD(2052, 12, 30),
	julian.CalendarGregorianToJD(2071, 8, 12),
	julian.CalendarGregorianToJD(2090, 3, 23),
	julian.CalendarGregorianToJD(2108, 11, 3),
}

func TestNode0(t *testing.T) {
	for i, j := range n0 {
		if e := math.Abs(base.PMod(moon.Node(j)+1, 2*math.Pi) - 1); e > 1e-3 {
			t.Error(i, e)
		}
	}
}

func TestNode180(t *testing.T) {
	for i, j := range n180 {
		if e := math.Abs(moon.Node(j) - math.Pi); e > 1e-3 {
			t.Error(i, e)
		}
	}
}
