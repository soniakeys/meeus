// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package rise_test

import (
	"fmt"

	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/meeus/deltat"
	"github.com/soniakeys/meeus/globe"
	"github.com/soniakeys/meeus/julian"
	"github.com/soniakeys/meeus/rise"
	"github.com/soniakeys/meeus/sidereal"
)

func ExampleApproxTimes() {
	// Example 15.a, p. 103.
	jd := julian.CalendarGregorianToJD(1988, 3, 20)
	p := globe.Coord{
		Lon: base.NewAngle(false, 71, 5, 0).Rad(),
		Lat: base.NewAngle(false, 42, 20, 0).Rad(),
	}
	// Meeus gives us the value of 11h 50m 58.1s but we have a package
	// function for this:
	Th0 := sidereal.Apparent0UT(jd)
	α := base.NewRA(2, 46, 55.51).Rad()
	δ := base.NewAngle(false, 18, 26, 27.3).Rad()
	h0 := rise.Stdh0Stellar
	rise, transit, set, err := rise.ApproxTimes(p, h0, Th0, α, δ)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Units for approximate values given near top of p. 104 are circles.
	fmt.Printf("rising:  %+.5f\n", rise/86400)
	fmt.Printf("transit: %+.5f\n", transit/86400)
	fmt.Printf("seting:  %+.5f\n", set/86400)
	// Output:
	// rising:  +0.51816
	// transit: +0.81965
	// seting:  +0.12113
}

func ExampleTimes() {
	// Example 15.a, p. 103.
	jd := julian.CalendarGregorianToJD(1988, 3, 20)
	p := globe.Coord{
		Lon: base.NewAngle(false, 71, 5, 0).Rad(),
		Lat: base.NewAngle(false, 42, 20, 0).Rad(),
	}
	// Meeus gives us the value of 11h 50m 58.1s but we have a package
	// function for this:
	Th0 := sidereal.Apparent0UT(jd)
	α3 := []float64{
		base.NewRA(2, 42, 43.25).Rad(),
		base.NewRA(2, 46, 55.51).Rad(),
		base.NewRA(2, 51, 07.69).Rad(),
	}
	δ3 := []float64{
		base.NewAngle(false, 18, 02, 51.4).Rad(),
		base.NewAngle(false, 18, 26, 27.3).Rad(),
		base.NewAngle(false, 18, 49, 38.7).Rad(),
	}
	h0 := rise.Stdh0Stellar
	// Similarly as with Th0, Meeus gives us the value of 56 for ΔT but
	// let's use our package function.
	ΔT := deltat.Interp10A(jd)
	rise, transit, set, err := rise.Times(p, ΔT, h0, Th0, α3, δ3)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("rising: ", base.NewFmtTime(rise))
	fmt.Println("transit:", base.NewFmtTime(transit))
	fmt.Println("seting: ", base.NewFmtTime(set))
	// Output:
	// rising:  12ʰ26ᵐ9ˢ
	// transit: 19ʰ40ᵐ30ˢ
	// seting:  2ʰ54ᵐ26ˢ
}
