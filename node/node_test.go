// Copyright 2013 Sonia Keys
// License: MIT

package node_test

import (
	"fmt"
	"time"

	"github.com/soniakeys/meeus/julian"
	"github.com/soniakeys/meeus/node"
	"github.com/soniakeys/meeus/perihelion"
	"github.com/soniakeys/meeus/planetelements"
	"github.com/soniakeys/unit"
)

func ExampleEllipticAscending() {
	// Example 39.a, p. 276
	t, r := node.EllipticAscending(17.9400782, .96727426,
		unit.AngleFromDeg(111.84644),
		julian.CalendarGregorianToJD(1986, 2, 9.45891))
	y, m, d := julian.JDToCalendar(t)
	fmt.Printf("%d %s %.2f\n", y, time.Month(m), d)
	fmt.Printf("%.4f AU\n", r)
	// Output:
	// 1985 November 9.16
	// 1.8045 AU
}

func ExampleEllipticDescending() {
	// Example 39.a, p. 276
	t, r := node.EllipticDescending(17.9400782, .96727426,
		unit.AngleFromDeg(111.84644),
		julian.CalendarGregorianToJD(1986, 2, 9.45891))
	y, m, d := julian.JDToCalendar(t)
	fmt.Printf("%d %s %.2f\n", y, time.Month(m), d)
	fmt.Printf("%.4f AU\n", r)
	// Output:
	// 1986 March 10.37
	// 0.8493 AU
}

func ExampleParabolicAscending() {
	// Example 29.b, p. 277
	t, r := node.ParabolicAscending(1.324502,
		unit.AngleFromDeg(154.9103),
		julian.CalendarGregorianToJD(1989, 8, 20.291))
	y, m, d := julian.JDToCalendar(t)
	fmt.Printf("%d %s %d\n", y, time.Month(m), int(d))
	fmt.Printf("%.2f AU\n", r)
	// Output:
	// 1977 September 17
	// 28.07 AU
}

func ExampleParabolicDescending() {
	// Example 29.b, p. 277
	t, r := node.ParabolicDescending(1.324502,
		unit.AngleFromDeg(154.9103),
		julian.CalendarGregorianToJD(1989, 8, 20.291))
	y, m, d := julian.JDToCalendar(t)
	fmt.Printf("%d %s %.3f\n", y, time.Month(m), d)
	fmt.Printf("%.4f AU\n", r)
	// Output:
	// 1989 September 17.636
	// 1.3901 AU
}

func ExampleEllipticAscending_venus() {
	// Example 39.c, p. 278
	var k planetelements.Elements
	planetelements.Mean(planetelements.Venus,
		julian.CalendarGregorianToJD(1979, 1, 1), &k)
	t, _ := node.EllipticAscending(k.Axis, k.Ecc,
		k.Peri-k.Node,
		perihelion.Perihelion(perihelion.Venus, 1979))
	y, m, d := julian.JDToCalendar(t)
	fmt.Printf("%d %s %.3f\n", y, time.Month(m), d)
	// Output:
	// 1978 November 27.409
}
