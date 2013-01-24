package deltat_test

import (
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/soniakeys/meeus"
	"github.com/soniakeys/meeus/deltat"
	"github.com/soniakeys/meeus/interp"
	"github.com/soniakeys/meeus/julian"
)

func ExampleDeltaT1900to1997_table() {
	// Example 10.a, p. 78.
	calYear := 1977 + float64(julian.DayOfYear(1977, 2, 18, false))/365
	fmt.Printf("calendar year %.1f\n", calYear)
	x1, x3, yTable := interp.Slice(calYear,
		deltat.TableYear1, deltat.TableYearN, deltat.Table10A, 3)
	dt, err := interp.Len3Interpolate(calYear, x1, x3, yTable, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+.1f seconds\n", dt)
	// Output:
	// calendar year 1977.1
	// +46.8 seconds
}

func ExampleDeltaT1900to1997_polynomial() {
	// Example 10.a, p. 78.
	jd := julian.TimeToJD(time.Date(1977, 2, 18, 3, 37, 40, 0, time.UTC))
	year := meeus.JDEToJulianYear(jd)
	fmt.Printf("julian year %.1f\n", year)
	fmt.Printf("%+.1f seconds\n", deltat.DeltaT1900to1997(jd))
	// Output:
	// julian year 1977.1
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
	for _, tp := range []struct {
		year int
		ΔT   float64
	}{
		{1800, 13.1},
		{1900, -2.8},
		{1996, 61.6},
	} {
		jd := julian.CalendarGregorianToJD(tp.year, 0, 0)
		ΔT := deltat.DeltaT1800to1997(jd)
		if math.Abs(ΔT-tp.ΔT) > 2.3 {
			t.Fatalf("%#v, got %.1f", tp, ΔT)
		}
	}
}

func TestDeltaT1800to1899(t *testing.T) {
	for _, tp := range []struct {
		year int
		ΔT   float64
	}{
		{1800, 13.1},
		{1850, 6.8},
		{1898, -4.7},
	} {
		jd := julian.CalendarGregorianToJD(tp.year, 0, 0)
		if math.Abs(deltat.DeltaT1800to1899(jd)-tp.ΔT) > 1 {
			t.Fatalf("%#v", tp)
		}
	}
}

func TestDeltaT1900to1997(t *testing.T) {
	for y := 1900; y < 1998; y += 2 {
		jd := julian.CalendarGregorianToJD(y, 0, 0)
		t.Logf("%d %.2f  %.1f", y, jd, deltat.DeltaT1900to1997(jd))
	}
	for _, tp := range []struct {
		year int
		ΔT   float64
	}{
		{1900, -2.8},
		{1950, 29.1},
		{1996, 61.6},
	} {
		jd := julian.CalendarGregorianToJD(tp.year, 0, 0)
		ΔT := deltat.DeltaT1900to1997(jd)
		if math.Abs(ΔT-tp.ΔT) > 1 {
			t.Fatalf("%#v, got %.1f", tp, ΔT)
		}
	}
}
