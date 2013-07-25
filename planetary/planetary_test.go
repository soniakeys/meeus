package planetary_test

import (
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/meeus/julian"
	"github.com/soniakeys/meeus/planetary"
)

func ExampleMercuryInfConj() {
	// Example 36.a, p. 252
	j := planetary.MercuryInfConj(1993.75)
	fmt.Printf("%.3f\n", j)
	y, m, df := julian.JDToCalendar(j)
	d, f := math.Modf(df)
	fmt.Printf("%d %s %d, at %dʰ\n", y, time.Month(m), int(d), int(f*24+.5))
	// Output:
	// 2449297.645
	// 1993 November 6, at 3ʰ
}

func ExampleSaturnConj() {
	// Example 36.b, p. 252
	j := planetary.SaturnConj(2125.5)
	fmt.Printf("%.3f\n", j)
	y, m, df := julian.JDToCalendar(j)
	d, f := math.Modf(df)
	fmt.Printf("%d %s %d, at %dʰ\n", y, time.Month(m), int(d), int(f*24+.5))
	// Output:
	// 2497437.904
	// 2125 August 26, at 10ʰ
}

func ExampleMercuryWestElongation() {
	// Example 36.c, p. 253
	j, e := planetary.MercuryWestElongation(1993.9)
	fmt.Printf("%.2f\n", j)
	y, m, df := julian.JDToCalendar(j)
	d, f := math.Modf(df)
	fmt.Printf("%d %s %d, at %dʰ\n", y, time.Month(m), int(d), int(f*24+.5))
	fmt.Printf("%.4f deg\n", e*180/math.Pi)
	fmt.Printf("%.62d\n", base.NewFmtAngle(e))
	// Output:
	// 2449314.14
	// 1993 November 22, at 15ʰ
	// 19.7506 deg
	// 19°45′
}

func ExampleMarsStation2() {
	// Example 36.d, p. 254
	j := planetary.MarsStation2(1997.3)
	fmt.Printf("%.3f\n", j)
	y, m, df := julian.JDToCalendar(j)
	d, f := math.Modf(df)
	fmt.Printf("%d %s %d, at %dʰ\n", y, time.Month(m), int(d), int(f*24+.5))
	// Output:
	// 2450566.255
	// 1997 April 27, at 18ʰ
}

type tc struct {
	f    func(float64) float64
	jNom float64
	hour int
}

var td = []tc{
	{planetary.MercuryInfConj, julian.CalendarGregorianToJD(1631, 11, 7), 7},
	{planetary.VenusInfConj, julian.CalendarGregorianToJD(1882, 12, 6), 17},
	{planetary.MarsOpp, julian.CalendarGregorianToJD(2729, 9, 9), 3},
	{planetary.JupiterOpp, julian.CalendarJulianToJD(-6, 9, 15), 7},
	{planetary.SaturnOpp, julian.CalendarJulianToJD(-6, 9, 14), 9},
	{planetary.UranusOpp, julian.CalendarGregorianToJD(1780, 12, 17), 14},
	{planetary.NeptuneOpp, julian.CalendarGregorianToJD(1846, 8, 20), 4},
}

func Test255(t *testing.T) {
	for _, d := range td {
		_, f := math.Modf(.5 + d.f(base.JDEToJulianYear(d.jNom)))
		if int(f*24+.5) != d.hour {
			t.Errorf("got %d, expected %d", int(f*24+.5), d.hour)
		}
	}
}
