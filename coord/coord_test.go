// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package coord_test

import (
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/soniakeys/meeus/coord"
	"github.com/soniakeys/meeus/globe"
	"github.com/soniakeys/meeus/julian"
	"github.com/soniakeys/meeus/sidereal"
	"github.com/soniakeys/sexagesimal"
)

func ExampleEcliptic_EqToEcl() {
	// Example 13.a, p. 95.
	eq := &coord.Equatorial{
		sexa.NewRA(7, 45, 18.946).Rad(),
		sexa.NewAngle(' ', 28, 1, 34.26).Rad(),
	}
	obl := coord.NewObliquity(23.4392911 * math.Pi / 180)
	ecl := new(coord.Ecliptic).EqToEcl(eq, obl)
	λStr := fmt.Sprintf("%.5j", sexa.NewFmtAngle(ecl.Lon))
	βStr := fmt.Sprintf("%+.6j", sexa.NewFmtAngle(ecl.Lat))
	fmt.Println("λ =", λStr)
	fmt.Println("β =", βStr)
	// Output:
	// λ = 113°.21563
	// β = +6°.684170
}

func TestEquatorial_EclToEq(t *testing.T) {
	// repeat example above
	eq0 := &coord.Equatorial{
		sexa.NewRA(7, 45, 18.946).Rad(),
		sexa.NewAngle(' ', 28, 1, 34.26).Rad(),
	}
	obl := coord.NewObliquity(23.4392911 * math.Pi / 180)
	ecl := new(coord.Ecliptic).EqToEcl(eq0, obl)

	// now reverse transform
	eq := new(coord.Equatorial).EclToEq(ecl, obl)
	if math.Abs((eq.RA-eq0.RA)/eq.RA) > 1e-15 {
		t.Fatal("RA:", eq0.RA, eq.RA)
	}
	if math.Abs((eq.Dec-eq0.Dec)/eq.Dec) > 1e-15 {
		t.Fatal("Dec:", eq0.Dec, eq.Dec)
	}
}

func ExampleHorizontal_EqToHz() {
	// Example 13.b, p. 95.
	eq := &coord.Equatorial{
		RA:  sexa.NewRA(23, 9, 16.641).Rad(),
		Dec: sexa.NewAngle('-', 6, 43, 11.61).Rad(),
	}
	g := &globe.Coord{
		Lat: sexa.NewAngle(' ', 38, 55, 17).Rad(),
		Lon: sexa.NewAngle(' ', 77, 3, 56).Rad(),
	}
	jd := julian.TimeToJD(time.Date(1987, 4, 10, 19, 21, 0, 0, time.UTC))
	st := sidereal.Apparent(jd)
	hz := new(coord.Horizontal).EqToHz(eq, g, st)
	AStr := fmt.Sprintf("%+.3j", sexa.NewFmtAngle(hz.Az))
	hStr := fmt.Sprintf("%+.3j", sexa.NewFmtAngle(hz.Alt))
	fmt.Println("A =", AStr)
	fmt.Println("h =", hStr)
	// Output:
	// A = +68°.034
	// h = +15°.125
}

func TestEqToGal(t *testing.T) {
	g := new(coord.Galactic).EqToGal(&coord.Equatorial{
		RA:  sexa.NewRA(17, 48, 59.74).Rad(),
		Dec: sexa.NewAngle('-', 14, 43, 8.2).Rad(),
	})
	if s := fmt.Sprintf("%.4f", g.Lon*180/math.Pi); s != "12.9593" {
		t.Fatal("lon:", s)
	}
	if s := fmt.Sprintf("%+.4f", g.Lat*180/math.Pi); s != "+6.0463" {
		t.Fatal("lat:", s)
	}
}
