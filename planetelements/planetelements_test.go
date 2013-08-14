package planetelements_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/soniakeys/meeus/julian"
	pe "github.com/soniakeys/meeus/planetelements"
)

func ExampleMean() {
	// Example 31.a, p. 211
	j := julian.CalendarGregorianToJD(2065, 6, 24)
	var e pe.Elements
	pe.Mean(pe.Mercury, j, &e)
	fmt.Printf("L: %.6f\n", e.Lon*180/math.Pi)
	fmt.Printf("a: %.9f\n", e.Axis)
	fmt.Printf("e: %.8f\n", e.Ecc)
	fmt.Printf("i: %.6f\n", e.Inc*180/math.Pi)
	fmt.Printf("Ω: %.6f\n", e.Node*180/math.Pi)
	fmt.Printf("ϖ: %.6f\n", e.Peri*180/math.Pi)
	// Output:
	// L: 203.494701
	// a: 0.387098310
	// e: 0.20564510
	// i: 7.006171
	// Ω: 49.107650
	// ϖ: 78.475382
}

func TestInc(t *testing.T) {
	j := julian.CalendarGregorianToJD(2065, 6, 24)
	var e pe.Elements
	pe.Mean(pe.Mercury, j, &e)
	if i := pe.Inc(pe.Mercury, j); i != e.Inc {
		t.Fatal(i, "!=", e.Inc)
	}
}

func TestNode(t *testing.T) {
	j := julian.CalendarGregorianToJD(2065, 6, 24)
	var e pe.Elements
	pe.Mean(pe.Mercury, j, &e)
	if Ω := pe.Node(pe.Mercury, j); Ω != e.Node {
		t.Fatal(Ω, "!=", e.Node)
	}
}
