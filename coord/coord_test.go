// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package coord_test

import (
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/meeus/coord"
	"github.com/soniakeys/meeus/globe"
	"github.com/soniakeys/meeus/julian"
	"github.com/soniakeys/meeus/sidereal"
	"github.com/soniakeys/sexagesimal"
)

func ExampleEcliptic_EqToEcl() {
	// Example 13.a, p. 95.
	eq := &coord.Equatorial{
		base.NewRA(7, 45, 18.946),
		base.NewAngle(' ', 28, 1, 34.26),
	}
	obl := coord.NewObliquity(23.4392911 * math.Pi / 180)
	ecl := new(coord.Ecliptic).EqToEcl(eq, obl)
	fmt.Printf("λ = %.5j\n", sexa.Angle(ecl.Lon).Fmt())
	fmt.Printf("β = %+.6j\n", sexa.Angle(ecl.Lat).Fmt())
	// Output:
	// λ = 113°.21563
	// β = +6°.684170
}

func TestEquatorial_EclToEq(t *testing.T) {
	// repeat example above
	eq0 := &coord.Equatorial{
		base.NewRA(7, 45, 18.946),
		base.NewAngle(' ', 28, 1, 34.26),
	}
	obl := coord.NewObliquity(23.4392911 * math.Pi / 180)
	ecl := new(coord.Ecliptic).EqToEcl(eq0, obl)

	// now reverse transform
	eq := new(coord.Equatorial).EclToEq(ecl, obl)
	if math.Abs((eq.RA-eq0.RA).Rad()/eq.RA.Rad()) > 1e-15 {
		t.Fatal("RA:", eq0.RA, eq.RA)
	}
	if math.Abs((eq.Dec-eq0.Dec).Rad()/eq.Dec.Rad()) > 1e-15 {
		t.Fatal("Dec:", eq0.Dec, eq.Dec)
	}
}

func ExampleHorizontal_EqToHz() {
	// Example 13.b, p. 95.
	eq := &coord.Equatorial{
		RA:  base.NewRA(23, 9, 16.641),
		Dec: base.NewAngle('-', 6, 43, 11.61),
	}
	g := &globe.Coord{
		Lat: base.NewAngle(' ', 38, 55, 17),
		Lon: base.NewAngle(' ', 77, 3, 56),
	}
	jd := julian.TimeToJD(time.Date(1987, 4, 10, 19, 21, 0, 0, time.UTC))
	st := sidereal.Apparent(jd)
	hz := new(coord.Horizontal).EqToHz(eq, g, st)
	fmt.Printf("A = %+.3j\n", sexa.Angle(hz.Az).Fmt())
	fmt.Printf("h = %+.3j\n", sexa.Angle(hz.Alt).Fmt())
	// Output:
	// A = +68°.034
	// h = +15°.125
}

func TestEqToGal(t *testing.T) {
	g := new(coord.Galactic).EqToGal(&coord.Equatorial{
		RA:  base.NewRA(17, 48, 59.74),
		Dec: base.NewAngle('-', 14, 43, 8.2),
	})
	if s := fmt.Sprintf("%.4f", g.Lon*180/math.Pi); s != "12.9593" {
		t.Fatal("lon:", s)
	}
	if s := fmt.Sprintf("%+.4f", g.Lat*180/math.Pi); s != "+6.0463" {
		t.Fatal("lat:", s)
	}
}
