// Copyright 2013 Sonia Keys
// License: MIT

package coord_test

import (
	"fmt"
	"time"

	"github.com/soniakeys/meeus/v3/base"
	"github.com/soniakeys/meeus/v3/coord"
	"github.com/soniakeys/meeus/v3/globe"
	"github.com/soniakeys/meeus/v3/julian"
	"github.com/soniakeys/meeus/v3/sidereal"
	"github.com/soniakeys/sexagesimal"
	"github.com/soniakeys/unit"
)

func ExampleEclToEq() {
	// Exercise, end of Example 13.a, p. 95.
	α, δ := coord.EclToEq(
		unit.AngleFromDeg(113.21563),
		unit.AngleFromDeg(6.68417),
		base.SOblJ2000,
		base.COblJ2000)
	fmt.Printf("α = %.3d, δ = %+.2d\n", sexa.FmtRA(α), sexa.FmtAngle(δ))
	// Output:
	// α = 7ʰ45ᵐ18ˢ.946, δ = +28°1′34″.26
}

func ExampleEcliptic_EqToEcl() {
	// Example 13.a, p. 95.
	eq := &coord.Equatorial{
		RA:  unit.NewRA(7, 45, 18.946),
		Dec: unit.NewAngle(' ', 28, 1, 34.26),
	}
	obl := coord.NewObliquity(unit.AngleFromDeg(23.4392911))
	ecl := new(coord.Ecliptic).EqToEcl(eq, obl)
	fmt.Printf("λ = %.5j\n", sexa.FmtAngle(ecl.Lon))
	fmt.Printf("β = %+.6j\n", sexa.FmtAngle(ecl.Lat))
	// Output:
	// λ = 113°.21563
	// β = +6°.684170
}

func ExampleEqToEcl() {
	// Example 13.a, p. 95 but using precomputed obliquity sine and cosine.
	λ, β := coord.EqToEcl(
		unit.NewRA(7, 45, 18.946),
		unit.NewAngle(' ', 28, 1, 34.26),
		base.SOblJ2000, base.COblJ2000)
	fmt.Printf("λ = %.5j\n", sexa.FmtAngle(λ))
	fmt.Printf("β = %+.6j\n", sexa.FmtAngle(β))
	// Output:
	// λ = 113°.21563
	// β = +6°.684170
}

func ExampleEqToGal() {
	// Exercise, p. 96.
	l, b := coord.EqToGal(
		unit.NewRA(17, 48, 59.74),
		unit.NewAngle('-', 14, 43, 8.2))
	fmt.Printf("l = %.4j, b = %+.4j\n", sexa.FmtAngle(l), sexa.FmtAngle(b))
	// Output:
	// l = 12°.9593, b = +6°.0463
}

func ExampleEqToHz() {
	// Example 13.b, p. 95.
	jd := julian.TimeToJD(time.Date(1987, 4, 10, 19, 21, 0, 0, time.UTC))
	A, h := coord.EqToHz(
		unit.NewRA(23, 9, 16.641),
		unit.NewAngle('-', 6, 43, 11.61),
		unit.NewAngle(' ', 38, 55, 17),
		unit.NewAngle(' ', 77, 3, 56),
		sidereal.Apparent(jd))
	fmt.Printf("A = %+.3j\n", sexa.FmtAngle(A))
	fmt.Printf("h = %+.3j\n", sexa.FmtAngle(h))
	// Output:
	// A = +68°.034
	// h = +15°.125
}

func ExampleEquatorial_EclToEq() {
	// Exercise, end of Example 13.a, p. 95.
	ecl := &coord.Ecliptic{
		Lon: unit.AngleFromDeg(113.21563),
		Lat: unit.AngleFromDeg(6.68417),
	}
	ε := coord.NewObliquity(unit.AngleFromDeg(23.4392911))
	eq := new(coord.Equatorial).EclToEq(ecl, ε)
	fmt.Printf("α = %.3d, δ = %+.2d\n",
		sexa.FmtRA(eq.RA), sexa.FmtAngle(eq.Dec))
	// Output:
	// α = 7ʰ45ᵐ18ˢ.946, δ = +28°1′34″.26
}

func ExampleEquatorial_GalToEq() {
	// Exercise, p. 96, inverse
	g := &coord.Galactic{
		Lon: unit.AngleFromDeg(12.9593),
		Lat: unit.AngleFromDeg(6.0463),
	}
	eq := new(coord.Equatorial).GalToEq(g)
	fmt.Printf("α = %.1d, δ = %+d\n", sexa.FmtRA(eq.RA), sexa.FmtAngle(eq.Dec))
	// Output:
	// α = 17ʰ48ᵐ59ˢ.7, δ = -14°43′8″
}

func ExampleEquatorial_HzToEq() {
	// Example 13.b, p. 95, inverse.
	hz := &coord.Horizontal{
		Az:  unit.AngleFromDeg(68.0337),
		Alt: unit.AngleFromDeg(15.1249),
	}
	g := globe.Coord{
		Lat: unit.NewAngle(' ', 38, 55, 17),
		Lon: unit.NewAngle(' ', 77, 3, 56),
	}
	jd := julian.TimeToJD(time.Date(1987, 4, 10, 19, 21, 0, 0, time.UTC))
	eq := new(coord.Equatorial).HzToEq(hz, g, sidereal.Apparent(jd))
	fmt.Printf("α = %+.1d, δ = %+d\n", sexa.FmtRA(eq.RA), sexa.FmtAngle(eq.Dec))
	// Output:
	// α = +23ʰ9ᵐ16ˢ.6, δ = -6°43′12″
}
func ExampleHorizontal_EqToHz() {
	// Example 13.b, p. 95.
	eq := &coord.Equatorial{
		RA:  unit.NewRA(23, 9, 16.641),
		Dec: unit.NewAngle('-', 6, 43, 11.61),
	}
	g := &globe.Coord{
		Lat: unit.NewAngle(' ', 38, 55, 17),
		Lon: unit.NewAngle(' ', 77, 3, 56),
	}
	jd := julian.TimeToJD(time.Date(1987, 4, 10, 19, 21, 0, 0, time.UTC))
	st := sidereal.Apparent(jd)
	hz := new(coord.Horizontal).EqToHz(eq, g, st)
	fmt.Printf("A = %+.3j\n", sexa.FmtAngle(hz.Az))
	fmt.Printf("h = %+.3j\n", sexa.FmtAngle(hz.Alt))
	// Output:
	// A = +68°.034
	// h = +15°.125
}

func ExampleGalToEq() {
	// Exercise, p. 96, inverse
	α, δ := coord.GalToEq(
		unit.AngleFromDeg(12.9593), unit.AngleFromDeg(6.0463))
	fmt.Printf("α = %.1d, δ = %+d\n", sexa.FmtRA(α), sexa.FmtAngle(δ))
	// Output:
	// α = 17ʰ48ᵐ59ˢ.7, δ = -14°43′8″
}

func ExampleHzToEq() {
	// Example 13.b, p. 95, inverse.
	jd := julian.TimeToJD(time.Date(1987, 4, 10, 19, 21, 0, 0, time.UTC))
	α, δ := coord.HzToEq(
		unit.AngleFromDeg(68.0337),
		unit.AngleFromDeg(15.1249),
		unit.NewAngle(' ', 38, 55, 17),
		unit.NewAngle(' ', 77, 3, 56),
		sidereal.Apparent(jd))
	fmt.Printf("α = %+.1d, δ = %+d\n", sexa.FmtRA(α), sexa.FmtAngle(δ))
	// Output:
	// α = +23ʰ9ᵐ16ˢ.6, δ = -6°43′12″
}

func ExampleGalactic_EqToGal() {
	// Exercise, p. 96.
	eq := &coord.Equatorial{
		RA:  unit.NewRA(17, 48, 59.74),
		Dec: unit.NewAngle('-', 14, 43, 8.2),
	}
	g := new(coord.Galactic).EqToGal(eq)
	fmt.Printf("l = %.4j, b = %+.4j\n",
		sexa.FmtAngle(g.Lon), sexa.FmtAngle(g.Lat))
	// Output:
	// l = 12°.9593, b = +6°.0463
}
