package greenwich_test

import (
	"fmt"
	"time"

	"github.com/soniakeys/meeus/greenwich"
	"github.com/soniakeys/meeus/julian"
)

func ExampleMeanSidereal82_a() {
	// Example 12.a, p. 88.
	jd := 2446895.5
	s := greenwich.MeanSidereal82(jd)
	t := time.Time{}.UTC().Add(time.Duration(s * float64(time.Second)))
	fmt.Println(t.Format("15h 4m 5.0000s"))
	// Output:
	// 13h 10m 46.3668s
}

func ExampleMeanSidereal82_b() {
	// Example 12.b, p. 89.
	jd := julian.TimeToJD(time.Date(1987, 4, 10, 19, 21, 0, 0, time.UTC))
	s := greenwich.MeanSidereal82(jd)
	t := time.Time{}.UTC().Add(time.Duration(s * float64(time.Second)))
	fmt.Println(t.Format("15h 4m 5.00000s"))
	// Output:
	// 08h 34m 57.08958s
}

// Note above, time 5.00 format truncates rather than rounds.
// could be considered a bug.
