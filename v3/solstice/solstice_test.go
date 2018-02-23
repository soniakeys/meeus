// Copyright 2013 Sonia Keys
// License: MIT

package solstice_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/soniakeys/meeus/v3/julian"
	"github.com/soniakeys/meeus/v3/solstice"
	"github.com/soniakeys/unit"
)

func ExampleJune() {
	// Example 27.a, p. 180
	fmt.Printf("%.5f\n", solstice.June(1962))
	// Output:
	// 2437837.39245
}

type eq struct {
	y    int
	d    float64
	h, m int
	s    float64
}

var (
	mar = []eq{
		{1996, 20, 8, 4, 7},
		{1997, 20, 13, 55, 42},
		{1998, 20, 19, 55, 35},
		{1999, 21, 1, 46, 53},
		{2000, 20, 7, 36, 19},

		{2001, 20, 13, 31, 47},
		{2002, 20, 19, 17, 13},
		{2003, 21, 1, 0, 50},
		{2004, 20, 6, 49, 42},
		{2005, 20, 12, 34, 29},
	}
	jun = []eq{
		{1996, 21, 2, 24, 46},
		{1997, 21, 8, 20, 59},
		{1998, 21, 14, 3, 38},
		{1999, 21, 19, 50, 11},
		{2000, 21, 1, 48, 46},

		{2001, 21, 7, 38, 48},
		{2002, 21, 13, 25, 29},
		{2003, 21, 19, 11, 32},
		{2004, 21, 0, 57, 57},
		{2005, 21, 6, 47, 12},
	}
	sep = []eq{
		{1996, 22, 18, 1, 8},
		{1997, 22, 23, 56, 49},
		{1998, 23, 5, 38, 15},
		{1999, 23, 11, 32, 34},
		{2000, 22, 17, 28, 40},

		{2001, 22, 23, 5, 32},
		{2002, 23, 4, 56, 28},
		{2003, 23, 10, 47, 53},
		{2004, 22, 16, 30, 54},
		{2005, 22, 22, 24, 14},
	}
	dec = []eq{
		{1996, 21, 14, 6, 56},
		{1997, 21, 20, 8, 5},
		{1998, 22, 1, 57, 31},
		{1999, 22, 7, 44, 52},
		{2000, 21, 13, 38, 30},

		{2001, 21, 19, 22, 34},
		{2002, 22, 1, 15, 26},
		{2003, 22, 7, 4, 53},
		{2004, 21, 12, 42, 40},
		{2005, 21, 18, 36, 1},
	}
)

func Test2000(t *testing.T) {
	for i := range mar {
		e := &mar[i]
		approx := solstice.March(e.y)
		vsop87 := julian.CalendarGregorianToJD(e.y, 3, e.d) +
			unit.NewTime(' ', e.h, e.m, e.s).Day()
		if math.Abs(vsop87-approx) > 1./24/60 {
			t.Logf("mar %d: got %.5f expected %.5f", e.y, approx, vsop87)
			t.Errorf("%.0f second error", math.Abs(vsop87-approx)*24*60*60)
		}
	}
	for i := range jun {
		e := &jun[i]
		approx := solstice.June(e.y)
		vsop87 := julian.CalendarGregorianToJD(e.y, 6, e.d) +
			unit.NewTime(' ', e.h, e.m, e.s).Day()
		if math.Abs(vsop87-approx) > 1./24/60 {
			t.Logf("jun %d: got %.5f expected %.5f", e.y, approx, vsop87)
			t.Errorf("%.0f second error", math.Abs(vsop87-approx)*24*60*60)
		}
	}
	for i := range sep {
		e := &sep[i]
		approx := solstice.September(e.y)
		vsop87 := julian.CalendarGregorianToJD(e.y, 9, e.d) +
			unit.NewTime(' ', e.h, e.m, e.s).Day()
		if math.Abs(vsop87-approx) > 1./24/60 {
			t.Logf("sep %d: got %.5f expected %.5f", e.y, approx, vsop87)
			t.Errorf("%.0f day error", math.Abs(vsop87-approx))
		}
	}
	for i := range dec {
		e := &dec[i]
		approx := solstice.December(e.y)
		vsop87 := julian.CalendarGregorianToJD(e.y, 12, e.d) +
			unit.NewTime(' ', e.h, e.m, e.s).Day()
		if math.Abs(vsop87-approx) > 1./24/60 {
			t.Logf("dec %d: got %.5f expected %.5f", e.y, approx, vsop87)
			t.Errorf("%.0f second error", math.Abs(vsop87-approx)*24*60*60)
		}
	}
}
