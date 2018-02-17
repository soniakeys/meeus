// Copyright 2013 Sonia Keys
// License: MIT

// Sundial: Chapter 58, Calculation of a Planar Sundial.
package sundial

import (
	"math"

	"github.com/soniakeys/unit"
)

// Point return type represents a point to be used in constructing the sundial.
type Point struct {
	X, Y float64
}

// Line holds data to draw an hour line on the sundial.
type Line struct {
	Hour   int     // 0 to 24.
	Points []Point // One or more points corresponding to the hour.
}

var m = []float64{-23.44, -20.15, -11.47, 0, 11.47, 20.15, 23.44}

// General computes data for the general case of a planar sundial.
//
// Argument φ is geographic latitude at which the sundial will be located.
// D is gnomonic declination, the azimuth of the perpendicular to the plane
// of the sundial, measured from the southern meridian towards the west.
// Argument a is the length of a straight stylus perpendicular to the plane
// of the sundial, z is zenithal distance of the direction defined by the
// stylus.  Units of stylus length a are arbitrary.
//
// Results consist of a set of lines, a center point, u, the length of a
// polar stylus, and ψ, the angle which the polar stylus makes with the plane
// of the sundial.  The center point, the points defining the hour lines, and
// u are in units of a, the stylus length.
func General(φ, D unit.Angle, a float64, z unit.Angle) (lines []Line, center Point, u float64, ψ unit.Angle) {
	sφ, cφ := φ.Sincos()
	tφ := sφ / cφ
	sD, cD := D.Sincos()
	sz, cz := z.Sincos()
	P := sφ*cz - cφ*sz*cD
	for i := 0; i < 24; i++ {
		l := Line{Hour: i}
		H := float64(i-12) * 15 * math.Pi / 180
		aH := math.Abs(H)
		sH, cH := math.Sincos(H)
		for _, d := range m {
			tδ := math.Tan(d * math.Pi / 180)
			H0 := math.Acos(-tφ * tδ)
			if aH > H0 {
				continue // sun below horizon
			}
			Q := sD*sz*sH + (cφ*cz+sφ*sz*cD)*cH + P*tδ
			if Q < 0 {
				continue // sun below plane of sundial
			}
			Nx := cD*sH - sD*(sφ*cH-cφ*tδ)
			Ny := cz*sD*sH - (cφ*sz-sφ*cz*cD)*cH - (sφ*sz+cφ*cz*cD)*tδ
			l.Points = append(l.Points, Point{a * Nx / Q, a * Ny / Q})
		}
		if len(l.Points) > 0 {
			lines = append(lines, l)
		}
	}
	center.X = a / P * cφ * sD
	center.Y = -a / P * (sφ*sz + cφ*cz*cD)
	aP := math.Abs(P)
	u = a / aP
	ψ = unit.Angle(math.Asin(aP))
	return
}

// Equatorial computes data for a sundial level with the equator.
//
// Argument φ is geographic latitude at which the sundial will be located;
// a is the length of a straight stylus perpendicular to the plane of the
// sundial.
//
// The sundial will have two sides, north and south.  Results n and s define
// lines on the north and south sides of the sundial.  Result coordinates
// are in units of a, the stylus length.
func Equatorial(φ unit.Angle, a float64) (n, s []Line) {
	tφ := φ.Tan()
	for i := 0; i < 24; i++ {
		nl := Line{Hour: i}
		sl := Line{Hour: i}
		H := float64(i-12) * 15 * math.Pi / 180
		aH := math.Abs(H)
		sH, cH := math.Sincos(H)
		for _, d := range m {
			tδ := math.Tan(d * math.Pi / 180)
			H0 := math.Acos(-tφ * tδ)
			if aH > H0 {
				continue
			}
			x := -a * sH / tδ
			yy := a * cH / tδ
			if tδ < 0 {
				sl.Points = append(sl.Points, Point{x, yy})
			} else {
				nl.Points = append(nl.Points, Point{x, -yy})
			}
		}
		if len(nl.Points) > 0 {
			n = append(n, nl)
		}
		if len(sl.Points) > 0 {
			s = append(s, sl)
		}
	}
	return
}

// Horizontal computes data for a horizontal sundial.
//
// Argument φ is geographic latitude at which the sundial will be located,
// a is the length of a straight stylus perpendicular to the plane of the
// sundial.
//
// Results consist of a set of lines, a center point, and u, the length of a
// polar stylus.  They are in units of a, the stylus length.
func Horizontal(φ unit.Angle, a float64) (lines []Line, center Point, u float64) {
	sφ, cφ := φ.Sincos()
	tφ := sφ / cφ
	for i := 0; i < 24; i++ {
		l := Line{Hour: i}
		H := float64(i-12) * 15 * math.Pi / 180
		aH := math.Abs(H)
		sH, cH := math.Sincos(H)
		for _, d := range m {
			tδ := math.Tan(d * math.Pi / 180)
			H0 := math.Acos(-tφ * tδ)
			if aH > H0 {
				continue // sun below horizon
			}
			Q := cφ*cH + sφ*tδ
			x := a * sH / Q
			y := a * (sφ*cH - cφ*tδ) / Q
			l.Points = append(l.Points, Point{x, y})
		}
		if len(l.Points) > 0 {
			lines = append(lines, l)
		}
	}
	center.Y = -a / tφ
	u = a / math.Abs(sφ)
	return
}

// Vertical computes data for a vertical sundial.
//
// Argument φ is geographic latitude at which the sundial will be located.
// D is gnomonic declination, the azimuth of the perpendicular to the plane
// of the sundial, measured from the southern meridian towards the west.
// Argument a is the length of a straight stylus perpendicular to the plane
// of the sundial.
//
// Results consist of a set of lines, a center point, and u, the length of a
// polar stylus.  They are in units of a, the stylus length.
func Vertical(φ, D unit.Angle, a float64) (lines []Line, center Point, u float64) {
	sφ, cφ := φ.Sincos()
	tφ := sφ / cφ
	sD, cD := D.Sincos()
	for i := 0; i < 24; i++ {
		l := Line{Hour: i}
		H := float64(i-12) * 15 * math.Pi / 180
		aH := math.Abs(H)
		sH, cH := math.Sincos(H)
		for _, d := range m {
			tδ := math.Tan(d * math.Pi / 180)
			H0 := math.Acos(-tφ * tδ)
			if aH > H0 {
				continue // sun below horizon
			}
			Q := sD*sH + sφ*cD*cH - cφ*cD*tδ
			if Q < 0 {
				continue // sun below plane of sundial
			}
			x := a * (cD*sH - sφ*sD*cH + cφ*sD*tδ) / Q
			y := -a * (cφ*cH + sφ*tδ) / Q
			l.Points = append(l.Points, Point{x, y})
		}
		if len(l.Points) > 0 {
			lines = append(lines, l)
		}
	}
	center.X = -a * sD / cD
	center.Y = a * tφ / cD
	u = a / math.Abs(cφ*cD)
	return
}
