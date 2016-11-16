// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package sundial_test

import (
	"fmt"
	"math"

	"github.com/soniakeys/meeus/sundial"
)

func ExampleGeneral_a() {
	// Example 58.a, p. 404.
	const p = math.Pi / 180
	ls, c, _, ψ := sundial.General(40*p, 70*p, 1, 50*p)
	fmt.Printf("Hours:  %d", ls[0].Hour)
	for _, l := range ls[1:] {
		fmt.Printf(", %d", l.Hour)
	}
	fmt.Println()
	for _, l := range ls {
		if l.Hour == 11 {
			fmt.Printf("%d:  x = %.4f  y = %.4f\n",
				l.Hour, l.Points[2].X, l.Points[2].Y)
		}
		if l.Hour == 14 {
			fmt.Printf("%d:  x = %.4f  y = %.4f\n",
				l.Hour, l.Points[6].X, l.Points[6].Y)
		}
	}
	fmt.Printf("x0 = %+.4f\n", c.X)
	fmt.Printf("y0 = %+.4f\n", c.Y)
	fmt.Printf("ψ = %.4f\n", ψ/p)
	// Output:
	// Hours:  9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19
	// 11:  x = -2.0007  y = -1.1069
	// 14:  x = -0.0390  y = -0.3615
	// x0 = +3.3880
	// y0 = -3.1102
	// ψ = 12.2672
}

func ExampleGeneral_b() {
	// Example 58.b, p. 404.
	const p = math.Pi / 180
	ls, c, _, ψ := sundial.General(-35*p, 160*p, 1, 90*p)
	for _, l := range ls {
		if l.Hour == 12 {
			fmt.Printf("%d:  x = %+.4f  y = %+.4f\n",
				l.Hour, l.Points[5].X, l.Points[5].Y)
		}
		if l.Hour == 15 {
			fmt.Printf("%d:  x = %+.4f  y = %+.4f\n",
				l.Hour, l.Points[3].X, l.Points[3].Y)
		}
	}
	fmt.Printf("x0 = %+.4f\n", c.X)
	fmt.Printf("y0 = %+.4f\n", c.Y)
	fmt.Printf("ψ = %.4f\n", ψ/p)
	// Output:
	// 12:  x = +0.3640  y = -0.7410
	// 15:  x = -0.8439  y = -0.9298
	// x0 = +0.3640
	// y0 = +0.7451
	// ψ = 50.3315
}

func ExampleGeneral_c() {
	// Example 58.c, p. 405.
	const p = math.Pi / 180
	ls, _, _, _ := sundial.General(40*p, 160*p, 1, 75*p)
	fmt.Printf("Hours:  %d", ls[0].Hour)
	for _, l := range ls[1:] {
		fmt.Printf(", %d", l.Hour)
	}
	fmt.Println()
	// Output:
	// Hours:  5, 6, 13, 14, 15, 16, 17, 18, 19
}
