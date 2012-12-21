package julian_test

import (
	"testing"
	"time"

	"github.com/soniakeys/meeus/julian"
)

func TestGreg(t *testing.T) {
	for _, tp := range []struct {
		y, m  int
		d, jd float64
	}{
		{1957, 10, 4.81, 2436116.31}, // Sputnik, Ex 7.a, p. 61
		{2000, 1, 1.5, 2451545},      // more examples, p. 62
		{1999, 1, 1, 2451179.5},
		{1987, 1, 27, 2446822.5},
		{1987, 6, 19.5, 2446966},
		{1988, 1, 27, 2447187.5},
		{1988, 6, 19.5, 2447332},
		{1900, 1, 1, 2415020.5},
		{1600, 1, 1, 2305447.5},
		{1600, 12, 31, 2305812.5},
	} {
		dt := julian.GregYMDToJD(tp.y, tp.m, tp.d) - tp.jd
		if dt != 0 { // pretty strict!
			t.Logf("%#v", tp)
			t.Fatal("dt:", time.Duration(dt*24*float64(time.Hour)))
		}
	}
}

func TestJuli(t *testing.T) {
	for _, tp := range []struct {
		y, m  int
		d, jd float64
	}{
		{333, 1, 27.5, 1842713},   // Ex 7.b, p. 61
		{837, 4, 10.3, 2026871.8}, // more examples, p. 62
		{-123, 12, 31, 1676496.5},
		{-122, 1, 1, 1676497.5},
		{-1000, 7, 12.5, 1356001},
		{-1000, 2, 29, 1355866.5},
		{-1001, 8, 17.9, 1355671.4},
		{-4712, 1, 1.5, 0},
	} {
		dt := julian.JuliYMDToJD(tp.y, tp.m, tp.d) - tp.jd
		if dt != 0 { // pretty strict!
			t.Logf("%#v", tp)
			t.Fatal("dt:", time.Duration(dt*24*float64(time.Hour)))
		}
	}
}
