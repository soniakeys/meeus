// Copyright 2013 Sonia Keys
// License: MIT

package angle_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/soniakeys/meeus/angle"
	"github.com/soniakeys/meeus/julian"
	"github.com/soniakeys/sexagesimal"
	"github.com/soniakeys/unit"
)

func ExampleSep() {
	// Example 17.a, p. 110.
	r1 := unit.NewRA(14, 15, 39.7).Angle()
	d1 := unit.NewAngle(' ', 19, 10, 57)
	r2 := unit.NewRA(13, 25, 11.6).Angle()
	d2 := unit.NewAngle('-', 11, 9, 41)
	d := angle.Sep(r1, d1, r2, d2)
	fmt.Println(sexa.FmtAngle(d))
	// Output:
	// 32°47′35″
}

// First exercise, p. 110.
func TestSep(t *testing.T) {
	r1 := unit.NewRA(4, 35, 55.2).Angle()
	d1 := unit.NewAngle(' ', 16, 30, 33)
	r2 := unit.NewRA(16, 29, 24).Angle()
	d2 := unit.NewAngle('-', 26, 25, 55)
	d := angle.Sep(r1, d1, r2, d2)
	answer := unit.NewAngle(' ', 169, 58, 0)
	if math.Abs((d - answer).Rad()) > 1e-4 {
		t.Fatal(d, answer)
	}
}

var (
	r1 = []unit.Angle{
		unit.NewRA(10, 29, 44.27).Angle(),
		unit.NewRA(10, 36, 19.63).Angle(),
		unit.NewRA(10, 43, 01.75).Angle(),
	}
	d1 = []unit.Angle{
		unit.NewAngle(' ', 11, 02, 05.9),
		unit.NewAngle(' ', 10, 29, 51.7),
		unit.NewAngle(' ', 9, 55, 16.7),
	}
	r2 = []unit.Angle{
		unit.NewRA(10, 33, 29.64).Angle(),
		unit.NewRA(10, 33, 57.97).Angle(),
		unit.NewRA(10, 34, 26.22).Angle(),
	}
	d2 = []unit.Angle{
		unit.NewAngle(' ', 10, 40, 13.2),
		unit.NewAngle(' ', 10, 37, 33.4),
		unit.NewAngle(' ', 10, 34, 53.9),
	}
	jd1 = julian.CalendarGregorianToJD(1978, 9, 13)
	jd3 = julian.CalendarGregorianToJD(1978, 9, 15)
)

// Second exercise, p. 110.
func TestMinSep(t *testing.T) {
	sep, err := angle.MinSep(jd1, jd3, r1, d1, r2, d2)
	if err != nil {
		t.Fatal(err)
	}
	answer := unit.AngleFromDeg(.5017) // on p. 111
	if math.Abs((sep-answer).Rad()/sep.Rad()) > 1e-3 {
		t.Fatal(sep, answer)
	}
}

// "rectangular coordinate" solution, p. 113.
func TestMinSepRect(t *testing.T) {
	sep, err := angle.MinSepRect(jd1, jd3, r1, d1, r2, d2)
	if err != nil {
		t.Fatal(err)
	}
	answer := unit.AngleFromSec(224) // on p. 111
	if math.Abs((sep-answer).Rad()/sep.Rad()) > 1e-2 {
		t.Fatal(sep, answer)
	}

}

func TestSepHav(t *testing.T) {
	// Example 17.a, p. 110.
	r1 := unit.NewRA(14, 15, 39.7).Angle()
	d1 := unit.NewAngle(' ', 19, 10, 57)
	r2 := unit.NewRA(13, 25, 11.6).Angle()
	d2 := unit.NewAngle('-', 11, 9, 41)
	d := angle.SepHav(r1, d1, r2, d2)
	s := fmt.Sprint(sexa.FmtAngle(d))
	if s != "32°47′35″" {
		t.Fatal(s)
	}
}

func ExampleSepPauwels() {
	// Example 17.b, p. 116.
	r1 := unit.NewRA(14, 15, 39.7).Angle()
	d1 := unit.NewAngle(' ', 19, 10, 57)
	r2 := unit.NewRA(13, 25, 11.6).Angle()
	d2 := unit.NewAngle('-', 11, 9, 41)
	d := angle.SepPauwels(r1, d1, r2, d2)
	fmt.Println(sexa.FmtAngle(d))
	// Output:
	// 32°47′35″
}
