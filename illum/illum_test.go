// Copyright 2013 Sonia Keys
// License: MIT

package illum_test

import (
	"fmt"

	"github.com/soniakeys/meeus/illum"
	"github.com/soniakeys/unit"
)

func ExamplePhaseAngle() {
	// Example 41.a, p. 284
	i := illum.PhaseAngle(.724604, .910947, .983824)
	fmt.Printf("%.5f\n", i.Cos())
	// Output:
	// 0.29312
}

func ExampleFraction() {
	// Example 41.a, p. 284
	k := illum.Fraction(.724604, .910947, .983824)
	fmt.Printf("%.3f\n", k)
	// Output:
	// 0.647
}

func ExamplePhaseAngle2() {
	// Example 41.a, p. 284
	i := illum.PhaseAngle2(
		unit.AngleFromDeg(26.10588),
		unit.AngleFromDeg(-2.62102),
		.724604,
		unit.AngleFromDeg(88.35704),
		.983824, .910947)
	fmt.Printf("%.5f\n", i.Cos())
	// Output:
	// 0.29312
}

func ExamplePhaseAngle3() {
	// Example 41.a, p. 284
	i := illum.PhaseAngle3(
		unit.AngleFromDeg(26.10588),
		unit.AngleFromDeg(-2.62102),
		.621794, -.664905, -.033138, .910947)
	fmt.Printf("%.5f\n", i.Cos())
	// Output:
	// 0.29312
}

func ExampleFractionVenus() {
	// Example 41.b, p. 284
	k := illum.FractionVenus(2448976.5)
	fmt.Printf("%.3f\n", k)
	// Output:
	// 0.640
}

func ExampleVenus() {
	// Example 41.c, p. 285
	v := illum.Venus(.724604, .910947, unit.AngleFromDeg(72.96))
	fmt.Printf("%.1f\n", v)
	// Output:
	// -3.8
}

func ExampleSaturn() {
	// Example 41.d, p. 285
	v := illum.Saturn(9.867882, 10.464606,
		unit.AngleFromDeg(16.442), unit.AngleFromDeg(4.198))
	fmt.Printf("%+.1f\n", v)
	// Output:
	// +0.9
}
