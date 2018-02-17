// Copyright 2013 Sonia Keys
// License: MIT

// Easter: Chapter 8, Date of Easter
package easter

// Gregorian returns month and day of Easter in the Gregorian calendar.
func Gregorian(y int) (m, d int) {
	a := y % 19
	b, c := y/100, y%100
	d, e := b/4, b%4
	f := (b + 8) / 25
	g := (b - f + 1) / 3
	h := (19*a + b - d - g + 15) % 30
	i, k := c/4, c%4
	l := (32 + 2*e + 2*i - h - k) % 7
	m = (a + 11*h + 22*l) / 451
	n := h + l - 7*m + 114
	n, p := n/31, n%31
	return n, p + 1
}

// Julian returns month and day of Easter in the Julian calendar.
func Julian(y int) (m, d int) {
	a := y % 4
	b := y % 7
	c := y % 19
	d = (19*c + 15) % 30
	e := (2*a + 4*b - d + 34) % 7
	f := d + e + 114
	f, g := f/31, f%31
	return f, g + 1
}
