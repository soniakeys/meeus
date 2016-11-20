// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// +build !nopp

package saturnring_test

import (
	"fmt"
	"math"
	"testing"

	pp "github.com/soniakeys/meeus/planetposition"
	"github.com/soniakeys/meeus/saturnring"
	"github.com/soniakeys/sexagesimal"
)

func ExampleRing() {
	// Example 45.a, p. 320
	earth, err := pp.LoadPlanet(pp.Earth)
	if err != nil {
		fmt.Println(err)
		return
	}
	saturn, err := pp.LoadPlanet(pp.Saturn)
	if err != nil {
		fmt.Println(err)
		return
	}
	B, Bʹ, ΔU, P, a, b := saturnring.Ring(2448972.50068, earth, saturn)
	fmt.Printf("B  = %.3f\n", B*180/math.Pi)
	fmt.Printf("Bʹ = %.3f\n", Bʹ*180/math.Pi)
	fmt.Printf("ΔU = %.3f\n", ΔU*180/math.Pi)
	fmt.Printf("P  = %.3f\n", P*180/math.Pi)
	fmt.Printf("a  = %.2d\n", sexa.Angle(a).Fmt())
	fmt.Printf("b  = %.2d\n", sexa.Angle(b).Fmt())
	// Output:
	// B  = 16.442
	// Bʹ = 14.679
	// ΔU = 4.198
	// P  = 6.741
	// a  = 35″.87
	// b  = 10″.15
}

func TestUB(t *testing.T) {
	earth, err := pp.LoadPlanet(pp.Earth)
	if err != nil {
		fmt.Println(err)
		return
	}
	saturn, err := pp.LoadPlanet(pp.Saturn)
	if err != nil {
		fmt.Println(err)
		return
	}
	B, _, ΔU, _, _, _ := saturnring.Ring(2448972.50068, earth, saturn)
	ubΔU, ubB := saturnring.UB(2448972.50068, earth, saturn)
	if ubΔU != ΔU || ubB != B {
		t.Fatal()
	}
}
