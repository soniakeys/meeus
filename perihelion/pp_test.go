// +build !nopp

package perihelion_test

import (
	"math"
	"testing"

	"github.com/soniakeys/meeus/julian"
	pa "github.com/soniakeys/meeus/perihelion"
	pp "github.com/soniakeys/meeus/planetposition"
)

func TestJS2(t *testing.T) {
	// p. 270
	v, err := pp.LoadPlanet(pp.Jupiter)
	if err != nil {
		t.Fatal(err)
	}
	j, _ := pa.Aphelion2(pa.Jupiter, 1981.5, 1, v)
	y, m, d := julian.JDToCalendar(j)
	if y != 1981 || m != 7 || int(d) != 28 {
		t.Fatal(y, m, d)
	}
	v, err = pp.LoadPlanet(pp.Saturn)
	if err != nil {
		t.Fatal(err)
	}
	s, _ := pa.Perihelion2(pa.Saturn, 1944.5, 1, v)
	y, m, d = julian.JDToCalendar(s)
	if y != 1944 || m != 9 || int(d) != 8 {
		t.Fatal(y, m, d)
	}
}

type su struct {
	ap      byte
	y, m, d int
	r       float64
}

var sd = []su{
	{'a', 1929, 11, 11, 10.0467},
	{'p', 1944, 9, 8, 9.0288},
	/* test cases commented out just to minimize testing time
	{'a', 1959, 5, 29, 10.0664},
	{'p', 1974, 1, 8, 9.0153},
	{'a', 1988, 9, 11, 10.0444},
	{'p', 2003, 7, 26, 9.0309},
	{'a', 2018, 4, 17, 10.0656},
	{'p', 2032, 11, 28, 9.0149},
	{'a', 2047, 7, 15, 10.0462},
	*/
}

func TestS2(t *testing.T) {
	// p. 271
	v, err := pp.LoadPlanet(pp.Saturn)
	if err != nil {
		t.Fatal(err)
	}
	for _, d := range sd {
		yf := float64(d.y) + (float64(d.m)-.5)/12
		var s, r float64
		if d.ap == 'a' {
			s, r = pa.Aphelion2(pa.Saturn, yf, 1, v)
		} else {
			s, r = pa.Perihelion2(pa.Saturn, yf, 1, v)
		}
		y, m, df := julian.JDToCalendar(s)
		if y != d.y || m != d.m || int(df) != d.d || math.Abs(r-d.r) > .0001 {
			t.Log(d)
			t.Fatal(y, m, df, r)
		}
	}
}

var sr = []su{
	{'a', 1756, 11, 27, 20.0893},
	/*
		{'p', 1798, 3, 3, 18.289},
		{'a', 1841, 3, 16, 20.0976},
		{'p', 1882, 3, 23, 18.2807},
		{'a', 1925, 4, 1, 20.0973},
		{'p', 1966, 5, 21, 18.2848},
		{'a', 2009, 2, 27, 20.0989},
		{'p', 2050, 8, 17, 18.283},
		{'a', 2092, 11, 23, 20.0994},
	*/
}

func TestU2(t *testing.T) {
	// p. 271
	v, err := pp.LoadPlanet(pp.Uranus)
	if err != nil {
		t.Fatal(err)
	}
	for _, d := range sr {
		yf := float64(d.y) + (float64(d.m)-.5)/12
		var u, r float64
		if d.ap == 'a' {
			u, r = pa.Aphelion2(pa.Uranus, yf, 1, v)
		} else {
			u, r = pa.Perihelion2(pa.Uranus, yf, 1, v)
		}
		y, m, df := julian.JDToCalendar(u)
		if y != d.y || m != d.m || int(df) != d.d || math.Abs(r-d.r) > .0001 {
			t.Log(d)
			t.Fatal(y, m, df, r)
		}
	}
}

/*
var sn = []su{
	{'p', 1876, 8, 28, 29.8148}, // p. 271
	{'a', 1959, 7, 13, 30.3317}, // p. 271
	{'p', 2042, 9, 5, 29.8064},  // p. 272
}

func TestN2(t *testing.T) {
	v, err := pp.LoadPlanet(pp.Neptune)
	if err != nil {
		t.Fatal(err)
	}
	for _, d := range sn {
		yf := float64(d.y) + (float64(d.m)-.5)/12
		var u, r float64
		if d.ap == 'a' {
			u, r = pa.Aphelion2(pa.Neptune, yf, 1, v)
		} else {
			u, r = pa.Perihelion2(pa.Neptune, yf, 1, v)
		}
		y, m, df := julian.JDToCalendar(u)
		if y != d.y || m != d.m || int(df) != d.d || math.Abs(r-d.r) > .0001 {
			t.Log(d)
			t.Fatal(y, m, df, r)
		}
	}
}
*/

type ed struct {
	y, m, d int
	h, r    float64
}

var ep = []ed{
	{1991, 1, 3, 3, .983281},
	/*
		{1992, 1, 3, 15.06, .983324},
		{1993, 1, 4, 3.08, .983283},
		{1994, 1, 2, 5.92, .983301},
		{1995, 1, 4, 11.1, .983302},

		{1996, 1, 4, 7.43, .983223},
		{1997, 1, 1, 23.29, .983267},
		{1998, 1, 4, 21.27, .9833},
		{1999, 1, 3, 13.02, .983281},
		{2000, 1, 3, 5.31, .983321},

		{2001, 1, 4, 8.89, .983286},
		{2002, 1, 2, 14.17, .98329},
		{2003, 1, 4, 5.04, .98332},
		{2004, 1, 4, 17.72, .983265},
		{2005, 1, 2, .61, .983297},

		{2006, 1, 4, 15.52, .983327},
		{2007, 1, 3, 19.74, .983260},
		{2008, 1, 2, 23.87, .983280},
		{2009, 1, 4, 15.51, .983273},
		{2010, 1, 3, .18, .98329},
	*/
}

var ea = []ed{
	{1991, 7, 6, 15.46, 1.016703},
	/*
		{1992, 7, 3, 12.14, 1.01674},
		{1993, 7, 4, 22.37, 1.016666},
		{1994, 7, 5, 19.3, 1.016724},
		{1995, 7, 4, 2.29, 1.016742},

		{1996, 7, 5, 19.02, 1.016717},
		{1997, 7, 4, 19.34, 1.016754},
		{1998, 7, 3, 23.86, 1.016696},
		{1999, 7, 6, 22.86, 1.016718},
		{2000, 7, 3, 23.84, 1.016741},

		{2001, 7, 4, 13.65, 1.016643},
		{2002, 7, 6, 3.8, 1.016688},
		{2003, 7, 4, 5.67, 1.016728},
		{2004, 7, 5, 10.9, 1.016694},
		{2005, 7, 5, 4.98, 1.016742},

		{2006, 7, 3, 23.18, 1.016697},
		{2007, 7, 6, 23.89, 1.016706},
		{2008, 7, 4, 7.71, 1.016754},
		{2009, 7, 4, 1.69, 1.016666},
		{2010, 7, 6, 11.52, 1.016702},
	*/
}

func TestEarth2(t *testing.T) {
	// p. 274
	v, err := pp.LoadPlanet(pp.Earth)
	if err != nil {
		t.Fatal(err)
	}
	for _, d := range ep {
		yf := float64(d.y) + (float64(d.m)-.5)/12
		u, r := pa.Perihelion2(pa.Earth, yf, .0004, v)
		y, m, df := julian.JDToCalendar(u)
		dd, f := math.Modf(df)
		if y != d.y || m != d.m || int(dd) != d.d ||
			math.Abs(f*24-d.h) > .01 ||
			math.Abs(r-d.r) > .000001 {
			t.Log(d)
			t.Fatal(y, m, int(dd), f*24, r)
		}
	}
	for _, d := range ea {
		yf := float64(d.y) + (float64(d.m)-.5)/12
		u, r := pa.Aphelion2(pa.Earth, yf, .0004, v)
		y, m, df := julian.JDToCalendar(u)
		dd, f := math.Modf(df)
		if y != d.y || m != d.m || int(dd) != d.d ||
			math.Abs(f*24-d.h) > .01 ||
			math.Abs(r-d.r) > .000001 {
			t.Log(d)
			t.Fatal(y, m, int(dd), f*24, r)
		}
	}
}
