// Copyright 2013 Sonia Keys
// License: MIT

package eclipse_test

import (
	"fmt"

	"github.com/soniakeys/meeus/eclipse"
)

func ExampleSolar_1993() {
	// Example 54.a, p. 384.
	t, c, jm, γ, u, p, mag := eclipse.Solar(1993.38)
	switch t {
	case eclipse.None:
		fmt.Println("No eclipse")
		return
	case eclipse.Partial:
		fmt.Println("Partial eclipse")
	case eclipse.Annular:
		fmt.Println("Annular eclipse")
	case eclipse.AnnularTotal:
		fmt.Println("Annular-total eclipse")
	case eclipse.Total:
		fmt.Println("Total eclipse")
	default:
		panic(t)
	}
	if t == eclipse.Partial {
		fmt.Printf("Partial eclipse magnitude:       %.3f\n", mag)
	}
	if c {
		fmt.Println("Central")
	} else {
		fmt.Println("Non-central")
	}

	fmt.Printf("Time of maximum eclipse:  %.4f\n", jm)
	fmt.Printf("Minimum distance, γ:           %+.4f\n", γ)
	fmt.Printf("Umbral radius, u:              %+.4f\n", u)
	fmt.Printf("Penumbral radius:              %+.4f\n", p)
	// Output:
	// Partial eclipse
	// Partial eclipse magnitude:       0.740
	// Non-central
	// Time of maximum eclipse:  2449129.0978
	// Minimum distance, γ:           +1.1348
	// Umbral radius, u:              +0.0097
	// Penumbral radius:              +0.5558
}

func ExampleSolar_2009() {
	// Example 54.b, p. 385.
	t, c, jm, γ, u, p, mag := eclipse.Solar(2009.56)
	switch t {
	case eclipse.None:
		fmt.Println("No eclipse")
		return
	case eclipse.Partial:
		fmt.Println("Partial eclipse")
	case eclipse.Annular:
		fmt.Println("Annular eclipse")
	case eclipse.AnnularTotal:
		fmt.Println("Annular-total eclipse")
	case eclipse.Total:
		fmt.Println("Total eclipse")
	default:
		panic(t)
	}
	if t == eclipse.Partial {
		fmt.Printf("Partial eclipse magnitude:       %.3f\n", mag)
	}
	if c {
		fmt.Println("Central")
	} else {
		fmt.Println("Non-central")
	}

	fmt.Printf("Time of maximum eclipse:  %.4f\n", jm)
	fmt.Printf("Minimum distance, γ:           %+.4f\n", γ)
	fmt.Printf("Umbral radius, u:              %+.4f\n", u)
	fmt.Printf("Penumbral radius:              %+.4f\n", p)
	// Output:
	// Total eclipse
	// Central
	// Time of maximum eclipse:  2455034.6088
	// Minimum distance, γ:           +0.0695
	// Umbral radius, u:              -0.0157
	// Penumbral radius:              +0.5304
}

func ExampleLunar_1973() {
	// Example 54.c, p. 385.
	t, jm, γ, ρ, σ, mag, sdTotal, sdPartial, sdPenumbral :=
		eclipse.Lunar(1973.46)
	switch t {
	case eclipse.None:
		fmt.Println("No eclipse")
		return
	case eclipse.Penumbral:
		fmt.Println("Penumbral eclipse")
	case eclipse.Umbral:
		fmt.Println("Umbral eclipse")
	case eclipse.Total:
		fmt.Println("Total eclipse")
	default:
		panic(t)
	}
	fmt.Printf("Magnitude:                     %+.4f\n", mag)
	fmt.Printf("Time of maximum eclipse:  %.4f\n", jm)
	fmt.Printf("Minimum distance, γ:           %+.4f\n", γ)
	if t >= eclipse.Umbral {
		fmt.Printf("Umbral radius, σ:              %+.4f\n", σ)
	}
	fmt.Printf("Penumbral radius, ρ:           %+.4f\n", ρ)
	switch t {
	case eclipse.Total:
		fmt.Printf("Totality semiduration:         %3.0f min\n",
			sdTotal.Min())
		fallthrough
	case eclipse.Umbral:
		fmt.Printf("Partial phase semiduration:    %3.0f min\n",
			sdPartial.Min())
		fallthrough
	default:
		fmt.Printf("Penumbral semiduration:        %3.0f min\n",
			sdPenumbral.Min())
	}
	// Output:
	// Penumbral eclipse
	// Magnitude:                     +0.4625
	// Time of maximum eclipse:  2441849.3687
	// Minimum distance, γ:           -1.3249
	// Penumbral radius, ρ:           +1.3045
	// Penumbral semiduration:        101 min
}

func ExampleLunar_1997() {
	// Example 54.d, p. 386.
	t, jm, γ, ρ, σ, mag, sdTotal, sdPartial, sdPenumbral :=
		eclipse.Lunar(1997.7)
	switch t {
	case eclipse.None:
		fmt.Println("No eclipse")
		return
	case eclipse.Penumbral:
		fmt.Println("Penumbral eclipse")
	case eclipse.Umbral:
		fmt.Println("Umbral eclipse")
	case eclipse.Total:
		fmt.Println("Total eclipse")
	default:
		panic(t)
	}
	fmt.Printf("Magnitude:                     %+.4f\n", mag)
	fmt.Printf("Time of maximum eclipse:  %.4f\n", jm)
	fmt.Printf("Minimum distance, γ:           %+.4f\n", γ)
	if t >= eclipse.Umbral {
		fmt.Printf("Umbral radius, σ:              %+.4f\n", σ)
	}
	fmt.Printf("Penumbral radius, ρ:           %+.4f\n", ρ)
	switch t {
	case eclipse.Total:
		fmt.Printf("Totality semiduration:         %3.0f min\n",
			sdTotal.Min())
		fallthrough
	case eclipse.Umbral:
		fmt.Printf("Partial phase semiduration:    %3.0f min\n",
			sdPartial.Min())
		fallthrough
	default:
		fmt.Printf("Penumbral semiduration:        %3.0f min\n",
			sdPenumbral.Min())
	}
	// Output:
	// Total eclipse
	// Magnitude:                     +1.1868
	// Time of maximum eclipse:  2450708.2835
	// Minimum distance, γ:           -0.3791
	// Umbral radius, σ:              +0.7534
	// Penumbral radius, ρ:           +1.2717
	// Totality semiduration:          30 min
	// Partial phase semiduration:     98 min
	// Penumbral semiduration:        153 min
}
