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
)

func ExampleApproxTimes_computed() {
	// Example 15.a, p. 103.
	jd := julian.CalendarGregorianToJD(1988, 3, 20)
	p := globe.Coord{
		Lon: sexa.NewAngle(false, 71, 5, 0).Rad(),
		Lat: sexa.NewAngle(false, 42, 20, 0).Rad(),
	}

	// Th0 computed rather than taken from the text.
	Th0 := sidereal.Apparent0UT(jd)
	fmt.Printf("Th0: %.2s\n", sexa.NewFmtTime(Th0))

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
	fmt.Printf("α: %.2s\n", sexa.NewFmtRA(α))
	fmt.Printf("δ: %.1s\n", sexa.NewFmtAngle(δ))

	h0 := rise.Stdh0Stellar
	rise, transit, set, err := rise.ApproxTimes(p, h0, Th0, α, δ)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("rising:   %+.5f  %02s\n", rise/86400, sexa.NewFmtTime(rise))
	fmt.Printf("transit:  %+.5f  %02s\n", transit/86400, sexa.NewFmtTime(transit))
	fmt.Printf("seting:   %+.5f  %02s\n", set/86400, sexa.NewFmtTime(set))
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
	jd := julian.CalendarGregorianToJD(1988, 3, 20)
	p := globe.Coord{
		Lon: sexa.NewAngle(false, 71, 5, 0).Rad(),
		Lat: sexa.NewAngle(false, 42, 20, 0).Rad(),
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
	mRise, mTransit, mSet, err := rise.ApproxPlanet(jd, p, e, v)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Units for approximate values given near top of p. 104 are circles.
	fmt.Printf("rising:   %+.5f  %02s\n", mRise/86400, sexa.NewFmtTime(mRise))
	fmt.Printf("transit:  %+.5f  %02s\n", mTransit/86400, sexa.NewFmtTime(mTransit))
	fmt.Printf("seting:   %+.5f  %02s\n", mSet/86400, sexa.NewFmtTime(mSet))
	// Output:
	// rising:   +0.51816  12ʰ26ᵐ09ˢ
	// transit:  +0.81965  19ʰ40ᵐ17ˢ
	// seting:   +0.12113  02ʰ54ᵐ26ˢ
}

func ExampleTimes_computed() {
	// Example 15.a, p. 103.
	jd := julian.CalendarGregorianToJD(1988, 3, 20)
	p := globe.Coord{
		Lon: sexa.NewAngle(false, 71, 5, 0).Rad(),
		Lat: sexa.NewAngle(false, 42, 20, 0).Rad(),
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
	α := make([]float64, 3)
	δ := make([]float64, 3)
	α[0], δ[0] = elliptic.Position(v, e, jd-1)
	α[1], δ[1] = elliptic.Position(v, e, jd)
	α[2], δ[2] = elliptic.Position(v, e, jd+1)
	for i, j := range []float64{jd - 1, jd, jd + 1} {
		_, m, d := julian.JDToCalendar(j)
		fmt.Printf("%s %.0f  α: %0.2s  δ: %0.1s\n",
			time.Month(m), d, sexa.NewFmtRA(α[i]), sexa.NewFmtAngle(δ[i]))
	}

	// ΔT computed rather than taken from the text.
	ΔT := deltat.Interp10A(jd)
	fmt.Printf("ΔT: %.1f\n", ΔT)

	h0 := rise.Stdh0Stellar
	rise, transit, set, err := rise.Times(p, ΔT, h0, Th0, α, δ)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("rising: ", sexa.NewFmtTime(rise))
	fmt.Println("transit:", sexa.NewFmtTime(transit))
	fmt.Println("seting:  ", sexa.NewFmtTime(set))
	// Output:
	// March 19  α: 2ʰ42ᵐ43.25ˢ  δ: 18°02′51.4″
	// March 20  α: 2ʰ46ᵐ55.51ˢ  δ: 18°26′27.3″
	// March 21  α: 2ʰ51ᵐ07.69ˢ  δ: 18°49′38.7″
	// ΔT: 55.9
	// rising:  12ʰ25ᵐ26ˢ
	// transit: 19ʰ40ᵐ30ˢ
	// seting:   2ʰ54ᵐ40ˢ
}

func ExamplePlanet() {
	// Example 15.a, p. 103.
	jd := julian.CalendarGregorianToJD(1988, 3, 20)
	p := globe.Coord{
		Lon: sexa.NewAngle(false, 71, 5, 0).Rad(),
		Lat: sexa.NewAngle(false, 42, 20, 0).Rad(),
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
	mRise, mTransit, mSet, err := rise.Planet(jd, p, e, v)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("rising: ", sexa.NewFmtTime(mRise))
	fmt.Println("transit:", sexa.NewFmtTime(mTransit))
	fmt.Println("seting:  ", sexa.NewFmtTime(mSet))
	// Output:
	// rising:  12ʰ25ᵐ26ˢ
	// transit: 19ʰ40ᵐ30ˢ
	// seting:   2ʰ54ᵐ40ˢ
}
