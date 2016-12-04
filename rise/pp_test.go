// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// +build !nopp

package rise_test

import (
	"fmt"
	"time"

	"github.com/soniakeys/meeus/deltat"
	"github.com/soniakeys/meeus/elliptic"
	"github.com/soniakeys/meeus/globe"
	"github.com/soniakeys/meeus/julian"
	pp "github.com/soniakeys/meeus/planetposition"
	"github.com/soniakeys/meeus/rise"
	"github.com/soniakeys/meeus/sidereal"
	"github.com/soniakeys/sexagesimal"
	"github.com/soniakeys/unit"
)

func ExampleApproxTimes_computed() {
	// Example 15.a, p. 103, but using meeus packages to compute values
	// given in the text.
	jd := julian.CalendarGregorianToJD(1988, 3, 20)
	p := globe.Coord{
		Lon: unit.NewAngle(' ', 71, 5, 0),
		Lat: unit.NewAngle(' ', 42, 20, 0),
	}

	// Th0 computed rather than taken from the text.
	Th0 := sidereal.Apparent0UT(jd)
	fmt.Printf("Th0: %.2s\n", sexa.FmtTime(Th0))

	// Venus α, δ computed rather than taken from the text.
	e, err := pp.LoadPlanet(pp.Earth)
	if err != nil {
		fmt.Println(err)
		return
	}
	v, err := pp.LoadPlanet(pp.Venus)
	if err != nil {
		fmt.Println(err)
		return
	}
	α, δ := elliptic.Position(v, e, jd)
	fmt.Printf("α: %.2s\n", sexa.FmtRA(α))
	fmt.Printf("δ: %.1s\n", sexa.FmtAngle(δ))

	h0 := rise.Stdh0Stellar
	tRise, tTransit, tSet, err := rise.ApproxTimes(p, h0, Th0, α, δ)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("rising:   %+.5f %02s\n", tRise/86400, sexa.FmtTime(tRise))
	fmt.Printf("transit:  %+.5f %02s\n", tTransit/86400, sexa.FmtTime(tTransit))
	fmt.Printf("seting:   %+.5f %02s\n", tSet/86400, sexa.FmtTime(tSet))
	// Output:
	// Th0: 11ʰ50ᵐ58.09ˢ
	// α: 2ʰ46ᵐ55.51ˢ
	// δ: 18°26′27.3″
	// rising:   +0.51816  12ʰ26ᵐ09ˢ
	// transit:  +0.81965  19ʰ40ᵐ17ˢ
	// seting:   +0.12113  02ʰ54ᵐ26ˢ
}

func ExampleApproxPlanet() {
	// Example 15.a, p. 103.
	p := globe.Coord{
		Lon: unit.NewAngle(' ', 71, 5, 0),
		Lat: unit.NewAngle(' ', 42, 20, 0),
	}
	e, err := pp.LoadPlanet(pp.Earth)
	if err != nil {
		fmt.Println(err)
		return
	}
	v, err := pp.LoadPlanet(pp.Venus)
	if err != nil {
		fmt.Println(err)
		return
	}
	tRise, tTransit, tSet, err := rise.ApproxPlanet(1988, 3, 20, p, e, v)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Units for "m" values given near top of p. 104 are day fraction.
	fmt.Printf("rising:   %+.5f %02s\n", tRise/86400, sexa.FmtTime(tRise))
	fmt.Printf("transit:  %+.5f %02s\n", tTransit/86400, sexa.FmtTime(tTransit))
	fmt.Printf("seting:   %+.5f %02s\n", tSet/86400, sexa.FmtTime(tSet))
	// Output:
	// rising:   +0.51816  12ʰ26ᵐ09ˢ
	// transit:  +0.81965  19ʰ40ᵐ17ˢ
	// seting:   +0.12113  02ʰ54ᵐ26ˢ
}

func ExampleTimes_computed() {
	// Example 15.a, p. 103, but using meeus packages to compute values
	// given in the text.
	jd := julian.CalendarGregorianToJD(1988, 3, 20)
	p := globe.Coord{
		Lon: unit.NewAngle(' ', 71, 5, 0),
		Lat: unit.NewAngle(' ', 42, 20, 0),
	}
	// Th0 computed rather than taken from the text.
	Th0 := sidereal.Apparent0UT(jd)

	// Venus α, δ computed rather than taken from the text.
	e, err := pp.LoadPlanet(pp.Earth)
	if err != nil {
		fmt.Println(err)
		return
	}
	v, err := pp.LoadPlanet(pp.Venus)
	if err != nil {
		fmt.Println(err)
		return
	}
	α := make([]unit.RA, 3)
	δ := make([]unit.Angle, 3)
	α[0], δ[0] = elliptic.Position(v, e, jd-1)
	α[1], δ[1] = elliptic.Position(v, e, jd)
	α[2], δ[2] = elliptic.Position(v, e, jd+1)
	for i, j := range []float64{jd - 1, jd, jd + 1} {
		_, m, d := julian.JDToCalendar(j)
		fmt.Printf("%s %.0f  α: %0.2s  δ: %0.1s\n",
			time.Month(m), d, sexa.FmtRA(α[i]), sexa.FmtAngle(δ[i]))
	}

	// ΔT computed rather than taken from the text.
	ΔT := deltat.Interp10A(jd)
	fmt.Printf("ΔT: %.1f\n", ΔT)

	h0 := rise.Stdh0Stellar
	tRise, tTransit, tSet, err := rise.Times(p, ΔT, h0, Th0, α, δ)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("rising:   %+.5f %02s\n", tRise/86400, sexa.FmtTime(tRise))
	fmt.Printf("transit:  %+.5f %02s\n", tTransit/86400, sexa.FmtTime(tTransit))
	fmt.Printf("seting:   %+.5f %02s\n", tSet/86400, sexa.FmtTime(tSet))
	// Output:
	// March 19  α: 2ʰ42ᵐ43.25ˢ  δ: 18°02′51.4″
	// March 20  α: 2ʰ46ᵐ55.51ˢ  δ: 18°26′27.3″
	// March 21  α: 2ʰ51ᵐ07.69ˢ  δ: 18°49′38.7″
	// ΔT: 55.9
	// rising:   +0.51766  12ʰ25ᵐ26ˢ
	// transit:  +0.81980  19ʰ40ᵐ30ˢ
	// seting:   +0.12130  02ʰ54ᵐ40ˢ
}

func ExamplePlanet() {
	// Example 15.a, p. 103.
	p := globe.Coord{
		Lon: unit.NewAngle(' ', 71, 5, 0),
		Lat: unit.NewAngle(' ', 42, 20, 0),
	}
	e, err := pp.LoadPlanet(pp.Earth)
	if err != nil {
		fmt.Println(err)
		return
	}
	v, err := pp.LoadPlanet(pp.Venus)
	if err != nil {
		fmt.Println(err)
		return
	}
	tRise, tTransit, tSet, err := rise.Planet(1988, 3, 20, p, e, v)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("rising:   %+.5f %02s\n", tRise/86400, sexa.FmtTime(tRise))
	fmt.Printf("transit:  %+.5f %02s\n", tTransit/86400, sexa.FmtTime(tTransit))
	fmt.Printf("seting:   %+.5f %02s\n", tSet/86400, sexa.FmtTime(tSet))
	// Output:
	// rising:   +0.51766  12ʰ25ᵐ26ˢ
	// transit:  +0.81980  19ʰ40ᵐ30ˢ
	// seting:   +0.12130  02ʰ54ᵐ40ˢ
}
