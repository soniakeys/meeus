// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Coord: Chapter 13, Transformation of Coordinates.
package coord

import (
	"math"

	"github.com/soniakeys/meeus/common"
	"github.com/soniakeys/meeus/globe"
)

// Equatorial coordinates are referenced to the Earth's rotational axis.
type Equatorial struct {
	RA  float64 // Right ascension (α) in radians
	Dec float64 // Declination (δ) in radians
}

// Ecliptic coordinates are referenced to the plane of the ecliptic.
type Ecliptic struct {
	Lat float64 // Latitude (β) in radians
	Lon float64 // Longitude (λ) in radians
}

// Horizontal coordinates are referenced to the local horizon of an observer
// on the surface of the Earth.
type Horizontal struct {
	Alt float64 // Altitude (h) in radians
	Az  float64 // Azimuth (A) in radians
}

// Galactic coordinates are referenced to the plane of the Milky Way.
type Galactic struct {
	Lat float64 // Latitude (b) in radians
	Lon float64 // Longitude (l) in radians
}

// EqToEcl converts equatorial coordinates to ecliptic coordinates.
func (ecl *Ecliptic) EqToEcl(eq *Equatorial, sε, cε float64) *Ecliptic {
	sα, cα := math.Sincos(eq.RA)
	sδ, cδ := math.Sincos(eq.Dec)
	ecl.Lon = math.Atan2(sα*cε+(sδ/cδ)*sε, cα)
	ecl.Lat = math.Asin(sδ*cε - cδ*sε*sα)
	return ecl
}

// EclToEq converts ecliptic coordinates to equatorial coordinates.
func (eq *Equatorial) EclToEq(ecl *Ecliptic, sε, cε float64) *Equatorial {
	sβ, cβ := math.Sincos(ecl.Lat)
	sλ, cλ := math.Sincos(ecl.Lon)
	eq.RA = math.Atan2(sλ*cε-(sβ/cβ)*sε, cλ)
	eq.Dec = math.Asin(sβ*cε + cβ*sε*sλ)
	return eq
}

// EqToHz computes Horizontal coordinates from equatorial coordinates.
//
// Argument g is the location of the observer on the Earth.  Argument st
// is the sidereal time at Greenwich.
//
// Sidereal time must be consistent with the equatorial coordinates.
// If coordinates are apparent, sidereal time must be apparent as well.
func (hz *Horizontal) EqToHz(eq *Equatorial, g *globe.Coord, st float64) *Horizontal {
	H := common.Time(st).Rad() - g.Lon - eq.RA
	sH, cH := math.Sincos(H)
	sφ, cφ := math.Sincos(g.Lat)
	sδ, cδ := math.Sincos(eq.Dec)
	hz.Az = math.Atan2(sH, cH*sφ-(sδ/cδ)*cφ)
	hz.Alt = math.Asin(sφ*sδ + cφ*cδ*cH)
	return hz
}

var galacticNorth = &Equatorial{
	RA:  common.NewRA(12, 49, 0).Rad(),
	Dec: 27.4 * math.Pi / 180,
}

var galacticLon0 = 123 * math.Pi / 180

// EqToGal converts equatorial coordinates to galactic coordinates.
//
// Equatorial coordinates must be referred to the standard equinox of B1950.0.
// For conversion to B1950, see package precess and utility functions in
// package "common".
func (g *Galactic) EqToGal(eq *Equatorial) *Galactic {
	sdα, cdα := math.Sincos(galacticNorth.RA - eq.RA)
	sgδ, cgδ := math.Sincos(galacticNorth.Dec)
	sδ, cδ := math.Sincos(eq.Dec)
	x := math.Atan2(sdα, cdα*sgδ-(sδ/cδ)*cgδ)
	g.Lon = math.Mod(math.Pi+galacticLon0-x, 2*math.Pi)
	g.Lat = math.Asin(sδ*sgδ + cδ*cgδ*cdα)
	return g
}

// GalToEq converts galactic coordinates to equatorial coordinates.
//
// Resulting equatorial coordinates will be referred to the standard equinox of
// B1950.0.  For subsequent conversion to other epochs, see package precess and
// utility functions in package meeus.
func (eq *Equatorial) GalToEq(g *Galactic) *Equatorial {
	sdLon, cdLon := math.Sincos(g.Lon - galacticLon0)
	sgδ, cgδ := math.Sincos(galacticNorth.Dec)
	sb, cb := math.Sincos(g.Lat)
	y := math.Atan2(sdLon, cdLon*sgδ-(sb/cb)*cgδ)
	eq.RA = math.Mod(y+galacticNorth.RA, 2*math.Pi)
	if eq.RA < 0 {
		eq.RA += 2 * math.Pi
	}
	eq.Dec = math.Asin(sb*sgδ + cb*cgδ*cdLon)
	return eq
}
