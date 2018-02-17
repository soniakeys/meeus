// Copyright 2012 Sonia Keys
// License: MIT

package interp_test

import (
	"fmt"
	"math"

	"github.com/soniakeys/meeus/interp"
	"github.com/soniakeys/sexagesimal"
	"github.com/soniakeys/unit"
)

func ExampleLen3_InterpolateN() {
	// Example 3.a, p. 25.
	d3, err := interp.NewLen3(7, 9, []float64{
		.884226,
		.877366,
		.870531,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	h := unit.FromSexa(0, 4, 21, 0)
	fmt.Println(h, "hours")
	n := h / 24
	y := d3.InterpolateN(n)
	fmt.Printf("%.6f\n", y)
	// Output:
	// 4.35 hours
	// 0.876125
}

func ExampleLen3_InterpolateX() {
	// Example 3.a, p. 25.
	d3, err := interp.NewLen3(7, 9, []float64{
		.884226,
		.877366,
		.870531,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	x := 8 + unit.NewTime(' ', 4, 21, 0).Day() // 8th day at 4:21
	y := d3.InterpolateX(x)
	fmt.Printf("%.6f\n", y)
	// Output:
	// 0.876125
}

func ExampleLen3_Extremum() {
	// Example 3.b, p. 26.
	d3, err := interp.NewLen3(12, 20, []float64{
		1.3814294,
		1.3812213,
		1.3812453,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	x, y, err := d3.Extremum()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("distance:  %.7f AU\n", y)
	fmt.Printf("date:     %.4f\n", x)
	i, frac := math.Modf(x)
	fmt.Printf("1992 May %d, at %h TD",
		int(i), sexa.FmtTime(unit.TimeFromDay(frac)))
	// Output:
	// distance:  1.3812030 AU
	// date:     17.5864
	// 1992 May 17, at 14ʰ TD
}

func ExampleLen3_Zero() {
	// Example 3.c, p. 26.
	x1 := 26.
	x3 := 28.
	// the y unit doesn't matter.  working in degrees is fine
	yTable := []float64{
		unit.FromSexa('-', 0, 28, 13.4),
		unit.FromSexa(' ', 0, 6, 46.3),
		unit.FromSexa(' ', 0, 38, 23.2),
	}
	d3, err := interp.NewLen3(x1, x3, yTable)
	if err != nil {
		fmt.Println(err)
		return
	}
	x, err := d3.Zero(false)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("February %.5f\n", x)
	i, frac := math.Modf(x)
	fmt.Printf("February %d, at %m TD",
		int(i), sexa.FmtTime(unit.TimeFromDay(frac)))
	// Output:
	// February 26.79873
	// February 26, at 19ʰ10ᵐ TD
}

func ExampleLen3_Zero_strong() {
	// Example 3.d, p. 27.
	x1 := -1.
	x3 := 1.
	yTable := []float64{-2, 3, 2}
	d3, err := interp.NewLen3(x1, x3, yTable)
	if err != nil {
		fmt.Println(err)
		return
	}
	x, err := d3.Zero(true)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%.12f\n", x)
	// Output:
	// -0.720759220056
}

func ExampleLen5_InterpolateX() {
	// Example 3.e, p. 28.
	x1 := 27.
	x5 := 29.
	// work in degrees
	yTable := []float64{
		unit.FromSexa(' ', 0, 54, 36.125),
		unit.FromSexa(' ', 0, 54, 24.606),
		unit.FromSexa(' ', 0, 54, 15.486),
		unit.FromSexa(' ', 0, 54, 08.694),
		unit.FromSexa(' ', 0, 54, 04.133),
	}
	d5, err := interp.NewLen5(x1, x5, yTable)
	if err != nil {
		fmt.Println(err)
		return
	}
	n := (3 + 20./60) / 24
	x := 28 + n
	y := d5.InterpolateX(x)
	fmt.Printf("%.3d", sexa.FmtAngle(unit.AngleFromDeg(y)))
	// Output:
	// 54′13″.369
}

func ExampleLen5_Zero() {
	// Exercise, p. 30.
	x1 := 25.
	x5 := 29.
	yTable := []float64{
		unit.FromSexa('-', 1, 11, 21.23),
		unit.FromSexa('-', 0, 28, 12.31),
		unit.FromSexa(' ', 0, 16, 07.02),
		unit.FromSexa(' ', 1, 01, 00.13),
		unit.FromSexa(' ', 1, 45, 46.33),
	}
	d5, err := interp.NewLen5(x1, x5, yTable)
	if err != nil {
		fmt.Println(err)
		return
	}
	z, err := d5.Zero(false)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("1988 January %.6f\n", z)
	zInt, zFrac := math.Modf(z)
	fmt.Printf("1988 January %d at %m TD\n", int(zInt),
		sexa.FmtTime(unit.TimeFromDay(zFrac)))

	// compare result to that from just three central values
	d3, err := interp.NewLen3(26, 28, yTable[1:4])
	if err != nil {
		fmt.Println(err)
		return
	}
	z3, err := d3.Zero(false)
	if err != nil {
		fmt.Println(err)
		return
	}
	dz := z - z3
	fmt.Printf("%.6f day\n", dz)
	fmt.Printf("%.1f minute\n", dz*24*60)
	// Output:
	// 1988 January 26.638587
	// 1988 January 26 at 15ʰ20ᵐ TD
	// 0.000753 day
	// 1.1 minute
}

func ExampleLen4Half() {
	// Example 3.f, p. 32.
	half, err := interp.Len4Half([]float64{
		unit.FromSexa(0, 10, 18, 48.732),
		unit.FromSexa(0, 10, 23, 22.835),
		unit.FromSexa(0, 10, 27, 57.247),
		unit.FromSexa(0, 10, 32, 31.983),
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%.3d", sexa.FmtRA(unit.RAFromHour(half)))
	// Output:
	// 10ʰ25ᵐ40ˢ.001
}

// exercise, p. 34.
func ExampleLagrange() {
	table := []struct{ X, Y float64 }{
		{29.43, .4913598528},
		{30.97, .5145891926},
		{27.69, .4646875083},
		{28.11, .4711658342},
		{31.58, .5236885653},
		{33.05, .5453707057},
	}
	// 10 significant digits in input, no more than 10 expected in output
	fmt.Printf("30: %.10f\n", interp.Lagrange(30, table))
	fmt.Printf("0:  %.10f\n", interp.Lagrange(0, table))
	fmt.Printf("90: %.10f\n", interp.Lagrange(90, table))
	// Output:
	// 30: 0.5000000000
	// 0:  0.0000512249
	// 90: 0.9999648100
}

func ExampleLagrangePoly() {
	// Example 3.g, p, 34.
	table := []struct{ X, Y float64 }{
		{1, -6},
		{3, 6},
		{4, 9},
		{6, 15},
	}
	p := interp.LagrangePoly(table)
	// output format contrived to fit expected result
	for _, c := range p {
		fmt.Printf("%.0f\n", c*5)
	}
	// Output:
	// -87
	// 69
	// -13
	// 1
}
