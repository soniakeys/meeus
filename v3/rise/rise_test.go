// Copyright 2013 Sonia Keys
// License: MIT

package rise_test

import (
	"fmt"

	"github.com/soniakeys/meeus/v3/globe"
	"github.com/soniakeys/meeus/v3/rise"
	"github.com/soniakeys/sexagesimal"
	"github.com/soniakeys/unit"
)

func ExampleApproxTimes() {
	// Example 15.a, p. 103.
	// Venus on 1988 March 20
	p := globe.Coord{
		Lon: unit.NewAngle(' ', 71, 5, 0),
		Lat: unit.NewAngle(' ', 42, 20, 0),
	}
	Th0 := unit.NewTime(' ', 11, 50, 58.1)
	α := unit.NewRA(2, 46, 55.51)
	δ := unit.NewAngle(' ', 18, 26, 27.3)
	h0 := rise.Stdh0Stellar
	tRise, tTransit, tSet, err := rise.ApproxTimes(p, h0, Th0, α, δ)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Units for "m" values given near top of p. 104 are day fraction.
	fmt.Printf("rising:  %+.5f %02s\n", tRise/86400, sexa.FmtTime(tRise))
	fmt.Printf("transit: %+.5f %02s\n", tTransit/86400, sexa.FmtTime(tTransit))
	fmt.Printf("seting:  %+.5f %02s\n", tSet/86400, sexa.FmtTime(tSet))
	// Output:
	// rising:  +0.51816  12ʰ26ᵐ09ˢ
	// transit: +0.81965  19ʰ40ᵐ17ˢ
	// seting:  +0.12113  02ʰ54ᵐ26ˢ
}

func ExampleTimes() {
	// Example 15.a, p. 103.
	// Venus on 1988 March 20
	p := globe.Coord{
		Lon: unit.NewAngle(' ', 71, 5, 0),
		Lat: unit.NewAngle(' ', 42, 20, 0),
	}
	Th0 := unit.NewTime(' ', 11, 50, 58.1)
	α3 := []unit.RA{
		unit.NewRA(2, 42, 43.25),
		unit.NewRA(2, 46, 55.51),
		unit.NewRA(2, 51, 07.69),
	}
	δ3 := []unit.Angle{
		unit.NewAngle(' ', 18, 02, 51.4),
		unit.NewAngle(' ', 18, 26, 27.3),
		unit.NewAngle(' ', 18, 49, 38.7),
	}
	h0 := unit.AngleFromDeg(-.5667)
	ΔT := unit.Time(56)
	tRise, tTransit, tSet, err := rise.Times(p, ΔT, h0, Th0, α3, δ3)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("rising:  %+.5f %02s\n", tRise/86400, sexa.FmtTime(tRise))
	fmt.Printf("transit: %+.5f %02s\n", tTransit/86400, sexa.FmtTime(tTransit))
	fmt.Printf("seting:  %+.5f %02s\n", tSet/86400, sexa.FmtTime(tSet))
	// Output:
	// rising:  +0.51766  12ʰ25ᵐ26ˢ
	// transit: +0.81980  19ʰ40ᵐ30ˢ
	// seting:  +0.12130  02ʰ54ᵐ40ˢ
}
