package solardisk_test

import (
	"fmt"
	"math"
	"time"

	"github.com/soniakeys/meeus/julian"
	pp "github.com/soniakeys/meeus/planetposition"
	"github.com/soniakeys/meeus/solardisk"
)

func ExampleEphemeris() {
	j := 2448908.50068
	e, err := pp.LoadPlanet(pp.Earth, "")
	if err != nil {
		fmt.Println(err)
		return
	}
	P, B0, L0 := solardisk.Ephemeris(j, e)
	fmt.Printf("P:  %.2f\n", P*180/math.Pi)
	fmt.Printf("B0: %+.2f\n", B0*180/math.Pi)
	fmt.Printf("L0: %.2f\n", L0*180/math.Pi)
	// Output:
	// P:  26.27
	// B0: +5.99
	// L0: 238.63
}

func ExampleCycle() {
	j := solardisk.Cycle(1699)
	fmt.Printf("%.4f\n", j)
	y, m, d := julian.JDToCalendar(j)
	fmt.Printf("%d %s %.2f\n", y, time.Month(m), d)
	// Output:
	// 2444480.7230
	// 1980 August 29.22
}
