// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package rise_test

import (
	"fmt"

	"github.com/soniakeys/meeus/globe"
	"github.com/soniakeys/meeus/rise"
	"github.com/soniakeys/sexagesimal"
)

func ExampleApproxTimes() {
	// Example 15.a, p. 103.
	// Venus on 1988 March 20
	p := globe.Coord{
		Lon: sexa.NewAngle(false, 71, 5, 0).Rad(),
		Lat: sexa.NewAngle(false, 42, 20, 0).Rad(),
	}
	Th0 := sexa.NewTime(false, 11, 50, 58.1).Sec()
	α := sexa.NewRA(2, 46, 55.51).Rad()
	δ := sexa.NewAngle(false, 18, 26, 27.3).Rad()
	h0 := rise.Stdh0Stellar
	tRise, tTransit, tSet, err := rise.ApproxTimes(p, h0, Th0, α, δ)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Units for "m" values given near top of p. 104 are day fraction.
	fmt.Printf("rising:  %+.5f  %02s\n", tRise/86400, sexa.NewFmtTime(tRise))
	fmt.Printf("transit: %+.5f  %02s\n", tTransit/86400, sexa.NewFmtTime(tTransit))
	fmt.Printf("seting:  %+.5f  %02s\n", tSet/86400, sexa.NewFmtTime(tSet))
	// Output:
	// rising:  +0.51816  12ʰ26ᵐ09ˢ
	// transit: +0.81965  19ʰ40ᵐ17ˢ
	// seting:  +0.12113  02ʰ54ᵐ26ˢ
}

func ExampleTimes() {
	// Example 15.a, p. 103.
	// Venus on 1988 March 20
	p := globe.Coord{
		Lon: sexa.NewAngle(false, 71, 5, 0).Rad(),
		Lat: sexa.NewAngle(false, 42, 20, 0).Rad(),
	}
	Th0 := sexa.NewTime(false, 11, 50, 58.1).Sec()
	α3 := []float64{
		sexa.NewRA(2, 42, 43.25).Rad(),
		sexa.NewRA(2, 46, 55.51).Rad(),
		sexa.NewRA(2, 51, 07.69).Rad(),
	}
	δ3 := []float64{
		sexa.NewAngle(false, 18, 02, 51.4).Rad(),
		sexa.NewAngle(false, 18, 26, 27.3).Rad(),
		sexa.NewAngle(false, 18, 49, 38.7).Rad(),
	}
	h0 := rise.Stdh0Stellar
	ΔT := 56.
	tRise, tTransit, tSet, err := rise.Times(p, ΔT, h0, Th0, α3, δ3)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("rising:  %+.5f  %02s\n", tRise/86400, sexa.NewFmtTime(tRise))
	fmt.Printf("transit: %+.5f  %02s\n", tTransit/86400, sexa.NewFmtTime(tTransit))
	fmt.Printf("seting:  %+.5f  %02s\n", tSet/86400, sexa.NewFmtTime(tSet))
	// Output:
	// rising:  +0.51766  12ʰ25ᵐ26ˢ
	// transit: +0.81980  19ʰ40ᵐ30ˢ
	// seting:  +0.12130  02ʰ54ᵐ40ˢ
}
