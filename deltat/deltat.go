// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// DeltaT: Chapter 10, Dynamical Time and Universal Time.
//
// Functions in this package compute ΔT for various ranges of dates.
//
// ΔT = TD - UT.
//
// TD = "Dynamical Time", which is related to:
//	ET "Ephermis Time", an older term.
//	TDB "Barycentric Dynamical Time", very close to TD.
//	TDT "Terrestrial Dynamical Time", a more correct term.
//	TT "Terrestrial Time", a newer and more correct term.
//
// UT = "Universal Time", which is related (if ambiguously) to GMT "Greenwich
// Mean Time".
//
// The return value for all functions is ΔT in seconds.
package deltat

import (
	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/meeus/interp"
	"github.com/soniakeys/meeus/julian"
	"github.com/soniakeys/unit"
)

// Table10A encodes ΔT for the range of years tableYear1 to tableYearN.
var (
	tableYear1 = 1620.
	tableYearN = 2010.
	table10A   = []float64{
		121.0, 112.0, 103.0, 95.0, 88.0, 82.0, 77.0, 72.0, 68.0, 63.0,
		60.0, 56.0, 53.0, 51.0, 48.0, 46.0, 44.0, 42.0, 40.0, 38.0,
		35.0, 33.0, 31.0, 29.0, 26.0, 24.0, 22.0, 20.0, 18.0, 16.0,
		14.0, 12.0, 11.0, 10.0, 9.0, 8.0, 7.0, 7.0, 7.0, 7.0,

		7.0, 7.0, 8.0, 8.0, 9.0, 9.0, 9.0, 9.0, 9.0, 10.0,
		10.0, 10.0, 10.0, 10.0, 10.0, 10.0, 10.0, 11.0, 11.0, 11.0,
		11.0, 11.0, 12.0, 12.0, 12.0, 12.0, 13.0, 13.0, 13.0, 14.0,
		14.0, 14.0, 14.0, 15.0, 15.0, 15.0, 15.0, 15.0, 16.0, 16.0,

		16.0, 16.0, 16.0, 16.0, 16.0, 16.0, 15.0, 15.0, 14.0, 13.0,
		13.1, 12.5, 12.2, 12.0, 12.0, 12.0, 12.0, 12.0, 12.0, 11.9,
		11.6, 11.0, 10.2, 9.2, 8.2, 7.1, 6.2, 5.6, 5.4, 5.3,
		5.4, 5.6, 5.9, 6.2, 6.5, 6.8, 7.1, 7.3, 7.5, 7.6,

		7.7, 7.3, 6.2, 5.2, 2.7, 1.4, -1.2, -2.8, -3.8, -4.8,
		-5.5, -5.3, -5.6, -5.7, -5.9, -6.0, -6.3, -6.5, -6.2, -4.7,
		-2.8, -0.1, 2.6, 5.3, 7.7, 10.4, 13.3, 16.0, 18.2, 20.2,
		21.1, 22.4, 23.5, 23.8, 24.3, 24.0, 23.9, 23.9, 23.7, 24.0,

		24.3, 25.3, 26.2, 27.3, 28.2, 29.1, 30.0, 30.7, 31.4, 32.2,
		33.1, 34.0, 35.0, 36.5, 38.3, 40.2, 42.2, 44.5, 46.5, 48.5,
		50.5, 52.2, 53.8, 54.9, 55.8, 56.9, 58.3, 60.0, 61.6, 63.0,
		63.8, 64.3, 64.6, 64.8, 65.5, 66.1}
)

// Interp10A returns ΔT at a date, accurate from years 1620 to 2010.
func Interp10A(jde float64) (ΔT unit.Time) {
	// kind of crazy, working in calendar years, but it seems that's what
	// we're supposed to do.
	y, m, d := julian.JDToCalendar(jde)
	l := julian.LeapYearGregorian(y)
	yl := 365.
	if l {
		yl++
	}
	yf := float64(y) + float64(julian.DayOfYear(y, m, int(d+.5), l))/yl
	d3, err := interp.Len3ForInterpolateX(yf, tableYear1, tableYearN, table10A)
	if err != nil {
		panic(err) // error would indicate a bug in interp.Slice.
	}
	return unit.Time(d3.InterpolateX(yf))
}

// c2000 returns centuries from calendar year 2000.0.
//
// Arg should be a calendar year.
func c2000(y float64) float64 {
	return (y - 2000) * .01
}

// PolyBefore948 returns a polynomial approximation of ΔT valid for calendar
// years before 948.
func PolyBefore948(year float64) (ΔT unit.Time) {
	// (10.1) p. 78
	return unit.Time(base.Horner(c2000(year), 2177, 497, 44.1))
}

// Poly948to1600 returns a polynomial approximation of ΔT valid for calendar
// years 948 to 1600.
func Poly948to1600(year float64) (ΔT unit.Time) {
	// (10.2) p. 78
	return unit.Time(base.Horner(c2000(year), 102, 102, 25.3))
}

// PolyAfter2000 returns a polynomial approximation of ΔT valid for calendar
// years after 2000.
func PolyAfter2000(year float64) (ΔT unit.Time) {
	ΔT = Poly948to1600(year)
	if year < 2100 {
		ΔT += unit.Time(.37 * (year - 2100))
	}
	return
}

// jc1900 returns julian centuries from the epoch J1900.0
//
// Arg should be a julian day, technically JDE.
func jc1900(jde float64) float64 {
	return (jde - base.J1900) / base.JulianCentury
}

// Poly1800to1997 returns a polynomial approximation of ΔT valid for years
// 1800 to 1997.
//
// The accuracy is within 2.3 seconds.
func Poly1800to1997(jde float64) (ΔT unit.Time) {
	return unit.Time(base.Horner(jc1900(jde),
		-1.02, 91.02, 265.90, -839.16, -1545.20,
		3603.62, 4385.98, -6993.23, -6090.04,
		6298.12, 4102.86, -2137.64, -1081.51))
}

// Poly1800to1899 returns a polynomial approximation of ΔT valid for years
// 1800 to 1899.
//
// The accuracy is within 0.9 seconds.
func Poly1800to1899(jde float64) (ΔT unit.Time) {
	return unit.Time(base.Horner(jc1900(jde),
		-2.50, 228.95, 5218.61, 56282.84, 324011.78,
		1061660.75, 2087298.89, 2513807.78,
		1818961.41, 727058.63, 123563.95))
}

// Poly1900to1997 returns a polynomial approximation of ΔT valid for years
// 1900 to 1997.
//
// The accuracy is within 0.9 seconds.
func Poly1900to1997(jde float64) (ΔT unit.Time) {
	return unit.Time(base.Horner(jc1900(jde),
		-2.44, 87.24, 815.20, -2637.80, -18756.33,
		124906.15, -303191.19, 372919.88,
		-232424.66, 58353.42))
}
