// Copyright 2013 Sonia Keys
// License: MIT

package perihelion_test

import (
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/soniakeys/meeus/v3/julian"
	pa "github.com/soniakeys/meeus/v3/perihelion"
)

func ExamplePerihelion() {
	// Example 38.a, p. 270
	j := pa.Perihelion(pa.Venus, 1978.79)
	fmt.Printf("%.3f\n", j)
	y, m, df := julian.JDToCalendar(j)
	d, f := math.Modf(df)
	fmt.Printf("%d %s %d at %dʰ\n", y, time.Month(m), int(d), int(f*24+.5))
	// Output:
	// 2443873.704
	// 1978 December 31 at 5ʰ
}

func ExampleAphelion() {
	// Example 38.b, p. 270
	j := pa.Aphelion(pa.Mars, 2032.5)
	fmt.Printf("%.3f\n", j)
	y, m, df := julian.JDToCalendar(j)
	d, f := math.Modf(df)
	fmt.Printf("%d %s %d at %dʰ\n", y, time.Month(m), int(d), int(f*24+.5))
	// Output:
	// 2463530.456
	// 2032 October 24 at 23ʰ
}

func TestJS(t *testing.T) {
	// p. 270
	j := pa.Aphelion(pa.Jupiter, 1981.5)
	y, m, d := julian.JDToCalendar(j)
	if y != 1981 || m != 7 || int(d) != 19 {
		t.Fatal(y, m, d)
	}
	s := pa.Perihelion(pa.Saturn, 1944.5)
	y, m, d = julian.JDToCalendar(s)
	if y != 1944 || m != 7 || int(d) != 30 {
		t.Fatal(y, m, d)
	}
}

func TestEarth(t *testing.T) {
	// p. 273
	j := pa.Perihelion(pa.EMBary, 1990)
	y, m, d := julian.JDToCalendar(j)
	if y != 1990 || m != 1 || int(d) != 3 {
		t.Fatal(y, m, d)
	}
	j = pa.Perihelion(pa.Earth, 1990)
	y, m, df := julian.JDToCalendar(j)
	d, f := math.Modf(df)
	if y != 1990 || m != 1 || int(d) != 4 || int(f*24+.5) != 16 {
		t.Fatal(y, m, df)
	}
}
