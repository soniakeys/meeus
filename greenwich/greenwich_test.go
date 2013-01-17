package greenwich_test

import (
	"fmt"
	"time"

	"github.com/soniakeys/meeus"
	"github.com/soniakeys/meeus/greenwich"
	"github.com/soniakeys/meeus/julian"
)

func ExampleMeanSidereal_a() {
	// Example 12.a, p. 88.
	jd := 2446895.5
	s := greenwich.MeanSidereal(jd)
	sa := greenwich.ApparentSidereal(jd)
	fmt.Printf("%.4d\n", meeus.NewFmtTime(s))
	fmt.Printf("%.4d\n", meeus.NewFmtTime(sa))
	// Output:
	// 13ʰ10ᵐ46ˢ.3668
	// 13ʰ10ᵐ46ˢ.1351
}

func ExampleMeanSidereal_b() {
	// Example 12.b, p. 89.
	jd := julian.TimeToJD(time.Date(1987, 4, 10, 19, 21, 0, 0, time.UTC))
	fmt.Printf("%.4d\n", meeus.NewFmtTime(greenwich.MeanSidereal(jd)))
	// Output:
	// 8ʰ34ᵐ57ˢ.0896
}
