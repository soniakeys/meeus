package moonnode_test

import (
	"fmt"
	"math"
	"time"

	"github.com/soniakeys/meeus/julian"
	"github.com/soniakeys/meeus/moonnode"
	"github.com/soniakeys/sexagesimal"
)

func ExampleAscending() {
	// Example 51.a, p. 365.
	j := moonnode.Ascending(1987.37)
	fmt.Printf("%.5f\n", j)
	y, m, d := julian.JDToCalendar(j)
	d, f := math.Modf(d)
	fmt.Printf("%d %s %d, at %d TD\n", y, time.Month(m), int(d),
		sexa.NewFmtTime(f*24*3600))
	// Output:
	// 2446938.76803
	// 1987 May 23, at 6ʰ25ᵐ58ˢ TD
}
