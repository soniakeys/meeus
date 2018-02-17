// Copyright 2013 Sonia Keys
// License: MIT

package base_test

import (
	"fmt"
	"math"

	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/unit"
)

func ExampleIlluminated_venus() {
	// Example 41.a, p. 284.
	k := base.Illuminated(unit.Angle(math.Acos(.29312)))
	fmt.Printf("%.3f\n", k)
	// Output:
	// 0.647
}

func ExampleIlluminated_moon() {
	// Example 48.a, p. 347.
	k := base.Illuminated(unit.AngleFromDeg(69.0756))
	fmt.Printf("k = %.4f\n", k)
	// Output:
	// k = 0.6786
}

// The coincidentally similar fractions of these two test cases reminds me
// of a story from Powell Observatory.  The Astronomical Society of Kansas City
// hosts frequent educational programs there, and while most are scheduled
// well in advance, the observatory and it's massive 30" Newtonian reflector
// are famous enough that there are occasional unannounced visitors.  One
// bright sunny day while some society members were performing maintenance,
// a school bus full of children pulled up.  The adults asked if there was
// any way the children could look through the telescope.  The workers didn't
// hesitate.  "Well, Venus is high in the sky right now, we could let them
// see the phase..."  The children were filed off of the bus and up the eight-
// foot warehouse ladder to take turns at the eyepiece, then right back onto
// the bus.  As the last two adults were following them up the steps of the
// bus, one was heard to grumble to the other, "I don't see why you need
// a giant telescope to look at the Moon, you can see it plain as day right
// there."  And of course he pointed to the crescent Moon in the sky, it
// looking just like Venus through the telescope.  The bus door closed and
// the bus rumbled off.

func ExampleLimb() {
	// Example 48.a, p. 347.
	χ := base.Limb(
		unit.RAFromDeg(134.6885),
		unit.AngleFromDeg(13.7684),
		unit.RAFromDeg(20.6579),
		unit.AngleFromDeg(8.6964))
	fmt.Printf("χ = %.1f\n", χ.Deg())
	// Output:
	// χ = 285.0
}
