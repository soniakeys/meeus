package coord_test

import (
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/soniakeys/meeus"
	"github.com/soniakeys/meeus/coord"
	"github.com/soniakeys/meeus/globe"
	"github.com/soniakeys/meeus/greenwich"
	"github.com/soniakeys/meeus/julian"
)

func ExampleEcliptic_EqToEcl() {
	// Example 13.a, p. 95.
	eq := &coord.Equatorial{
		meeus.NewRA(7, 45, 18.946).Rad(),
		meeus.NewAngle(false, 28, 1, 34.26).Rad(),
	}
	sε, cε := math.Sincos(23.4392911 * math.Pi / 180)
	ecl := new(coord.Ecliptic).EqToEcl(eq, sε, cε)
	λStr := meeus.DecSymAdd(fmt.Sprintf("%.5f", ecl.Lon*180/math.Pi), '°')
	βStr := meeus.DecSymAdd(fmt.Sprintf("%+.6f", ecl.Lat*180/math.Pi), '°')
	fmt.Println("λ =", λStr)
	fmt.Println("β =", βStr)
	// Output:
	// λ = 113°.21563
	// β = +6°.684170
}

func TestEquatorial_EclToEq(t *testing.T) {
	// repeat example above
	eq0 := &coord.Equatorial{
		meeus.NewRA(7, 45, 18.946).Rad(),
		meeus.NewAngle(false, 28, 1, 34.26).Rad(),
	}
	sε, cε := math.Sincos(23.4392911 * math.Pi / 180)
	ecl := new(coord.Ecliptic).EqToEcl(eq0, sε, cε)

	// now reverse transform
	eq := new(coord.Equatorial).EclToEq(ecl, sε, cε)
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
		RA:  meeus.NewRA(23, 9, 16.641).Rad(),
		Dec: meeus.NewAngle(true, 6, 43, 11.61).Rad(),
	}
	g := &globe.Coord{
		Lat: meeus.NewAngle(false, 38, 55, 17).Rad(),
		Lon: meeus.NewAngle(false, 77, 3, 56).Rad(),
	}
	jd := julian.TimeToJD(time.Date(1987, 4, 10, 19, 21, 0, 0, time.UTC))
	st := greenwich.ApparentSidereal(jd)
	hz := new(coord.Horizontal).EqToHz(eq, g, st)
	AStr := meeus.DecSymAdd(fmt.Sprintf("%+.3f", hz.Az*180/math.Pi), '°')
	hStr := meeus.DecSymAdd(fmt.Sprintf("%+.3f", hz.Alt*180/math.Pi), '°')
	fmt.Println("A =", AStr)
	fmt.Println("h =", hStr)
	// Output:
	// A = +68°.034
	// h = +15°.125
}

func TestEqToGal(t *testing.T) {
	g := new(coord.Galactic).EqToGal(&coord.Equatorial{
		RA:  meeus.NewRA(17, 48, 59.74).Rad(),
		Dec: meeus.NewAngle(true, 14, 43, 8.2).Rad(),
	})
	if s := fmt.Sprintf("%.4f", g.Lon*180/math.Pi); s != "12.9593" {
		t.Fatal("lon:", s)
	}
	if s := fmt.Sprintf("%+.4f", g.Lat*180/math.Pi); s != "+6.0463" {
		t.Fatal("lat:", s)
	}
}