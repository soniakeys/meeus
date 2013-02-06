package sidereal_test

import (
	"fmt"
	"time"

	"github.com/soniakeys/meeus/common"
	"github.com/soniakeys/meeus/julian"
	"github.com/soniakeys/meeus/sidereal"
)

func ExampleMean_a() {
	// Example 12.a, p. 88.
	jd := 2446895.5
	s := sidereal.Mean(jd)
	sa := sidereal.Apparent(jd)
	fmt.Printf("%.4d\n", common.NewFmtTime(s))
	fmt.Printf("%.4d\n", common.NewFmtTime(sa))
	// Output:
	// 13ʰ10ᵐ46ˢ.3668
	// 13ʰ10ᵐ46ˢ.1351
}

func ExampleMean_b() {
	// Example 12.b, p. 89.
	jd := julian.TimeToJD(time.Date(1987, 4, 10, 19, 21, 0, 0, time.UTC))
	fmt.Printf("%.4d\n", common.NewFmtTime(sidereal.Mean(jd)))
	// Output:
	// 8ʰ34ᵐ57ˢ.0896
}
