package deltat_test

import (
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/soniakeys/meeus/deltat"
	"github.com/soniakeys/meeus/julian"
)

func ExampleDeltaT1900to1997() {
	// Example 10.a, p. 78.
	// Text says ΔT = +48, Table 10.A on p. 79 shows +47.5
	// DeltaT1900to1997 expected accuracy is 0.9 second.
	jd := julian.TimeToJD(time.Date(1977, 2, 18, 3, 37, 40, 0, time.UTC))
	year := 2000 + (jd-julian.J2000)/365.25
	fmt.Printf("year %.1f\n", year)
	fmt.Printf("%+.1f seconds\n", deltat.DeltaT1900to1997(year))
	// Output:
	// year 1977.1
	// +47.1 seconds
}

func ExampleDeltaTBefore948() {
	// Example 10.b, p. 80.
	fmt.Printf("%+.0f seconds\n", deltat.DeltaTBefore948(333.1))
	// Output:
	// +6146 seconds
}

// Table 10.A p. 79 provides a way to test these polynomials
func TestDeltaT1800to1997(t *testing.T) {
	for _, tp := range []struct{ year, ΔT float64 }{
		{1800, 13.1},
		{1900, -2.8},
		{1996, 61.6},
	} {
		if math.Abs(deltat.DeltaT1800to1997(tp.year)-tp.ΔT) > 2.3 {
			t.Fatalf("%#v", tp)
		}
	}
}

func TestDeltaT1800to1899(t *testing.T) {
	for _, tp := range []struct{ year, ΔT float64 }{
		{1800, 13.1},
		{1850, 6.8},
		{1898, -4.7},
	} {
		if math.Abs(deltat.DeltaT1800to1899(tp.year)-tp.ΔT) > .9 {
			t.Fatalf("%#v", tp)
		}
	}
}

func TestDeltaT1900to1997(t *testing.T) {
	for _, tp := range []struct{ year, ΔT float64 }{
		{1900, -2.8},
		{1950, 29.1},
		{1996, 61.6},
	} {
		if math.Abs(deltat.DeltaT1900to1997(tp.year)-tp.ΔT) > .9 {
			t.Fatalf("%#v", tp)
		}
	}
}
