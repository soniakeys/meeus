// DeltaT: Chapter 10, Dynamical Time and Universal Time
package deltat

import (
	"github.com/soniakeys/meeus"
)

func c2000(y float64) float64 {
	return (y - 2000) * .01
}

// DeltaTBefore948 returns a polynomial approximation valid for years
// before 948.
func DeltaTBefore948(year float64) float64 {
	return meeus.Horner(c2000(year), []float64{2177, 497, 44.1})
}

// DeltaT948to1600 returns a polynomial approximation valid for years
// 948 to 1600.
func DeltaT948to1600(year float64) float64 {
	return meeus.Horner(c2000(year), []float64{102, 102, 25.3})
}

// DeltaTAfter2000 returns a polynomial approximation valid for years
// after 2000.
func DeltaTAfter2000(year float64) float64 {
	return DeltaT948to1600(year) + .37*(year-2100)
}

func c1900(y float64) float64 {
	return (y - 1900) * .01
}

// DeltaT1800to1997 returns a polynomial approximation valid for years
// 1800 to 1997.
//
// The accuracy is within 2.3 seconds.
func DeltaT1800to1997(year float64) float64 {
	return meeus.Horner(c1900(year), []float64{
		-1.02, 91.02, 265.90, -839.16, -1545.20,
		3603.62, 4385.98, -6993.23, -6090.04,
		6298.12, 4102.86, -2137.64, -1081.51})
}

// DeltaT1800to1899 returns a polynomial approximation valid for years
// 1800 to 1899.
//
// The accuracy is within 0.9 seconds.
func DeltaT1800to1899(year float64) float64 {
	return meeus.Horner(c1900(year), []float64{
		-2.50, 228.95, 5218.61, 56282.84, 324011.78,
		1061660.75, 2087298.89, 2513807.78,
		1818961.41, 727058.63, 123563.95})
}

// DeltaT1900to1997 returns a polynomial approximation valid for years
// 1900 to 1997.
//
// The accuracy is within 0.9 seconds.
func DeltaT1900to1997(year float64) float64 {
	return meeus.Horner(c1900(year), []float64{
		-2.44, 87.24, 815.20, -2637.80, -18756.33,
		124906.16, -303191.19, 372919.88,
		-232424.66, 58353.42})
}
