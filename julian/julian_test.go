package julian_test

import (
	"math"
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
		dt := julian.CalendarGregorianToJD(tp.y, tp.m, tp.d) - tp.jd
		if math.Abs(dt) > .1 {
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
		dt := julian.CalendarJulianToJD(tp.y, tp.m, tp.d) - tp.jd
		if math.Abs(dt) > .1 {
			t.Logf("%#v", tp)
			t.Fatal("dt:", time.Duration(dt*24*float64(time.Hour)))
		}
	}
}

func TestJuliLeap(t *testing.T) {
	for _, tp := range []struct {
		year int
		leap bool
	}{
		{900, true},
		{1236, true},
		{750, false},
		{1429, false},
	} {
		if julian.LeapYearJulian(tp.year) != tp.leap {
			t.Logf("%#v", tp)
			t.Fatal("JuliLeapYear")
		}
	}
}

func TestGregLeap(t *testing.T) {
	for _, tp := range []struct {
		year int
		leap bool
	}{
		{1700, false},
		{1800, false},
		{1900, false},
		{2100, false},
		{1600, true},
		{2000, true},
		{2400, true},
	} {
		if julian.LeapYearGregorian(tp.year) != tp.leap {
			t.Logf("%#v", tp)
			t.Fatal("JuliLeapYear")
		}
	}
}

func TestYMD(t *testing.T) {
	for _, tp := range []struct {
		jd   float64
		y, m int
		d    float64
	}{
		{2436116.31, 1957, 10, 4.81},
		{1842713, 333, 1, 27.5},
		{1507900.13, -584, 5, 28.63},
	} {
		y, m, d := julian.JDToCalendar(tp.jd)
		if y != tp.y || m != tp.m || math.Abs(d-tp.d) > .01 {
			t.Logf("%#v", tp)
			t.Fatal("JDToYMD", y, m, d)
		}
	}
}

func TestDOW(t *testing.T) {
	if julian.DayOfWeek(2434923.5) != 3 {
		t.Fatal("DOW")
	}
}

var doyTD = []struct {
	y, m, d int
	leap    bool
	doy     int
}{
	{1978, 11, 14, false, 318},
	{1988, 4, 22, true, 113},
}

func TestDOY(t *testing.T) {
	for _, tp := range doyTD {
		doy := julian.DayOfYear(tp.y, tp.m, tp.d, tp.leap)
		if doy != tp.doy {
			t.Logf("%#v", tp)
			t.Fatal("DayOfYear", doy)
		}
	}
}

func TestDOYToCal(t *testing.T) {
	for _, tp := range doyTD {
		m, d := julian.DayOfYearToCalendar(tp.doy, tp.leap)
		if m != tp.m || d != tp.d {
			t.Logf("%#v", tp)
			t.Fatal("DayOfYearToCalendar", m, d)
		}
	}
}
