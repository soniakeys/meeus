package meeus_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/soniakeys/meeus"
)

func ExampleDecSymAdd() {
	formatted := "1.25"
	fmt.Println("Standard decimal symbol:", formatted)
	fmt.Println("Degree units, non combining decimal point: ",
		meeus.DecSymAdd(formatted, '°'))
	// Output:
	// Standard decimal symbol: 1.25
	// Degree units, non combining decimal point:  1°.25
}

func ExampleDecSymCombine() {
	formatted := "1.25"
	fmt.Println("Standard decimal symbol:", formatted)
	// Note that some software may not be capable of combining or even
	// rendering the combining dot.
	fmt.Println("Degree units, combining form of decimal point:",
		meeus.DecSymCombine(formatted, '°'))
	// Output:
	// Standard decimal symbol: 1.25
	// Degree units, combining form of decimal point: 1°̣25
}

// For various numbers and symbols, test both Add and Combine.
// See that the functions do something, and that Strip returns
// the original number.
func TestStrip(t *testing.T) {
	var d string
	var sym rune
	t1 := func(fName string, f func(string, rune) string) {
		ad := f(d, sym)
		if ad == d {
			t.Fatalf("%s(%s, %c) had no effect", fName, d, sym)
		}
		if sd := meeus.DecSymStrip(ad, sym); sd != d {
			t.Fatalf("Strip(%s, %c) returned %s expected %s",
				ad, sym, sd, d)
		}
	}
	for _, d = range []string{"1.25", "1.", "1", ".25"} {
		for _, sym = range []rune{'°', '"', 'h', 'ʰ'} {
			t1("DecSymAdd", meeus.DecSymAdd)
			t1("DecSymCombine", meeus.DecSymCombine)
		}
	}
}

func TestDMSToDeg(t *testing.T) {
	// example p. 7
	a := meeus.DMSToDeg(false, 23, 26, 49)
	if math.Abs(a-23.44694444) > 1e-8 {
		t.Fatal("DMSToRad")
	}
}

func ExampleRA_SetHMS() {
	// Example 1.a, p. 8.
	a := new(meeus.RA).SetHMS(9, 14, 55.8)
	fmt.Printf("%.6f\n", math.Tan(a.Rad))
	// Output:
	// -0.877517
}

func TestAngle_SetDMS(t *testing.T) {
	// examples p. 9
	a := new(meeus.Angle).SetDMS(true, 13, 47, 22) // negative angle
	if a.String() != "-13°47′22″" {
		t.Fatal(a.String())
	}
	a.SetDMS(true, 0, 32, 41) // negative, < 1 degree
	if f := a.String(); f != "-32′41″" {
		t.Fatal(f)
	}
}

// example p. 6
func TestAngle_Format(t *testing.T) {
	a := new(meeus.Angle).SetDMS(false, 23, 26, 44)
	if f := a.String(); f != "23°26′44″" {
		t.Fatal(f)
	}
}

// example p. 6
func TestHourAngle_Format(t *testing.T) {
	a := new(meeus.HourAngle).SetHMS(false, 15, 22, 7)
	if f := fmt.Sprintf("%0s", a); f != "15ʰ22ᵐ07ˢ" {
		t.Fatalf(f)
	}
}

func TestOverflow(t *testing.T) {
	a := new(meeus.Angle).SetDMS(false, 23, 26, 44)
	if f := fmt.Sprintf("%03s", a); f != "023°26′44″" {
		t.Fatal(f)
	}
	a.SetDMS(false, 4423, 26, 44)
	if f := fmt.Sprintf("%03s", a); f != "***" {
		t.Fatal(f)
	}
}

func ExampleSplit60() {
	neg, x60, seg, err := meeus.Split60(-123.456, 2, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	s := ""
	if neg {
		s = "-"
	}
	fmt.Printf("%s%02d° %s′\n", s, x60, seg)
	// Output:
	// -02° 03.46′
}

func TestSplit60(t *testing.T) {
	for _, tc := range []struct {
		x    float64
		prec int
		neg  bool
		quo  int64
		rem  string
		err  error
	}{
		// warm up
		{75, 0, false, 1, "15", nil},
		{75, 1, false, 1, "15.0", nil},
		// smallest valid with prec = 15 is about 4.5 seconds.
		{4.500000123456789, 15, false, 0, "4.500000123456789", nil},
		{9, 16, false, 0, "9", meeus.WidthErrorInvalidPrecision},
		{10, 15, false, 0, "10", meeus.WidthErrorLossOfPrecision},
		// one degree can have 12 digits of precision without loss.
		{3600, 12, false, 60, "0.000000000000", nil},
		// 360 degrees (21600 minutes) can have 9.
		{360 * 3600, 9, false, 21600, "0.000000000", nil},
	} {
		neg, quo, rem, err := meeus.Split60(tc.x, tc.prec, false)
		if err != tc.err {
			t.Logf("%#v", tc)
			t.Fatal("err", err)
		}
		if err != nil {
			continue
		}
		if neg != tc.neg {
			t.Logf("%#v", tc)
			t.Fatal("neg", neg)
		}
		if quo != tc.quo {
			t.Logf("%#v", tc)
			t.Fatal("quo", quo)
		}
		if rem != tc.rem {
			t.Logf("%#v", tc)
			t.Fatal("rem", rem)
		}
	}
}
