package illum_test

import (
	"fmt"
	"math"

	"github.com/soniakeys/meeus/illum"
)

func ExamplePhaseAngle() {
	// Example 41.a, p. 284
	i := illum.PhaseAngle(.724604, .910947, .983824)
	fmt.Printf("%.5f\n", math.Cos(i))
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
	i := illum.PhaseAngle2(26.10588*math.Pi/180, -2.62102*math.Pi/180, .724604,
		88.35704*math.Pi/180, .983824, .910947)
	fmt.Printf("%.5f\n", math.Cos(i))
	// Output:
	// 0.29312
}

func ExamplePhaseAngle3() {
	// Example 41.a, p. 284
	i := illum.PhaseAngle3(26.10588*math.Pi/180, -2.62102*math.Pi/180,
		.621794, -.664905, -.033138, .910947)
	fmt.Printf("%.5f\n", math.Cos(i))
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
	v := illum.Venus(.724604, .910947, 72.96*math.Pi/180)
	fmt.Printf("%.1f\n", v)
	// Output:
	// -3.8
}

func ExampleSaturn() {
	// Example 41.d, p. 285
	v := illum.Saturn(9.867882, 10.464606, 16.442*math.Pi/180, 4.198*math.Pi/180)
	fmt.Printf("%+.1f\n", v)
	// Output:
	// +0.9
}
