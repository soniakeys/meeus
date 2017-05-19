// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Coord: Chapter 13, Transformation of Coordinates.
//
// Transforms in this package are provided in two forms, function and method.
// The results of the two forms should be identical.
//
// The function forms pass all arguments and results as single values.  These
// forms are best used when you are transforming a single pair of coordinates
// and wish to avoid memory allocation.
//
// The method forms take and return pointers to structs.  These forms are best
// used when you are transforming multiple coordinates and can reuse one or
// more of the structs.  In this case reuse of structs will minimize
// allocations, and the struct pointers will pass more efficiently on the
// stack.  These methods transform their arguments, placing the result in
// the receiver.  The receiver is then returned for convenience.
package coord

import (
	"math"

	"github.com/soniakeys/meeus/globe"
	"github.com/soniakeys/unit"
)

// Obliquity represents the obliquity of the ecliptic.
type Obliquity struct {
	S, C float64 // sine and cosine of obliquity
}

// NewObliquity constructs a new Obliquity.
//
// Struct members are initialized from the given value ε of the obliquity of
// the ecliptic.
func NewObliquity(ε unit.Angle) *Obliquity {
	r := &Obliquity{}
	r.S, r.C = ε.Sincos()
	return r
}

// Ecliptic coordinates are referenced to the plane of the ecliptic.
type Ecliptic struct {
	Lon unit.Angle // Longitude (λ)
	Lat unit.Angle // Latitude (β)
}

// EqToEcl converts equatorial coordinates to ecliptic coordinates.
func (ecl *Ecliptic) EqToEcl(eq *Equatorial, ε *Obliquity) *Ecliptic {
	ecl.Lon, ecl.Lat = EqToEcl(eq.RA, eq.Dec, ε.S, ε.C)
	return ecl
}

// EqToEcl converts equatorial coordinates to ecliptic coordinates.
//
//	α: right ascension coordinate to transform
//	δ: declination coordinate to transform
//	sε: sine of obliquity of the ecliptic
//	cε: cosine of obliquity of the ecliptic
//
// Results:
//
//	λ: ecliptic longitude
//	β: ecliptic latitude
func EqToEcl(α unit.RA, δ unit.Angle, sε, cε float64) (λ, β unit.Angle) {
	sα, cα := α.Sincos()
	sδ, cδ := δ.Sincos()
	λ = unit.Angle(math.Atan2(sα*cε+(sδ/cδ)*sε, cα)) // (13.1) p. 93
	β = unit.Angle(math.Asin(sδ*cε - cδ*sε*sα))      // (13.2) p. 93
	return
}

// Equatorial coordinates are referenced to the Earth's rotational axis.
type Equatorial struct {
	RA  unit.RA    // Right ascension (α)
	Dec unit.Angle // Declination (δ)
}

// EclToEq converts ecliptic coordinates to equatorial coordinates.
func (eq *Equatorial) EclToEq(ecl *Ecliptic, ε *Obliquity) *Equatorial {
	eq.RA, eq.Dec = EclToEq(ecl.Lon, ecl.Lat, ε.S, ε.C)
	return eq
}

// EclToEq converts ecliptic coordinates to equatorial coordinates.
//
//	λ: ecliptic longitude coordinate to transform
//	β: ecliptic latitude coordinate to transform
//	sε: sine of obliquity of the ecliptic
//	cε: cosine of obliquity of the ecliptic
//
// Results:
//	α: right ascension
//	δ: declination
func EclToEq(λ, β unit.Angle, sε, cε float64) (α unit.RA, δ unit.Angle) {
	sλ, cλ := λ.Sincos()
	sβ, cβ := β.Sincos()
	α = unit.RAFromRad(math.Atan2(sλ*cε-(sβ/cβ)*sε, cλ)) // (13.3) p. 93
	δ = unit.Angle(math.Asin(sβ*cε + cβ*sε*sλ))          // (13.4) p. 93
	return
}

// HzToEq transforms horizontal coordinates to equatorial coordinates.
//
// Sidereal time st must be consistent with the equatorial coordinates
// in the sense that if coordinates are apparent, sidereal time must be
// apparent as well.
func (eq *Equatorial) HzToEq(hz *Horizontal, g globe.Coord, st unit.Time) *Equatorial {
	eq.RA, eq.Dec = HzToEq(hz.Az, hz.Alt, g.Lat, g.Lon, st)
	return eq
}

// HzToEq transforms horizontal coordinates to equatorial coordinates.
//
//	A: azimuth
//	h: elevation
//	φ: latitude of observer on Earth
//	ψ: longitude of observer on Earth
//	st: sidereal time at Greenwich at time of observation.
//
// Sidereal time must be consistent with the equatorial coordinates
// in the sense that tf coordinates are apparent, sidereal time must be
// apparent as well.
//
// Results:
//
//	α: right ascension
//	δ: declination
func HzToEq(A, h, φ, ψ unit.Angle, st unit.Time) (α unit.RA, δ unit.Angle) {
	sA, cA := A.Sincos()
	sh, ch := h.Sincos()
	sφ, cφ := φ.Sincos()
	H := math.Atan2(sA, cA*sφ+sh/ch*cφ)
	α = unit.RAFromRad(st.Rad() - ψ.Rad() - H)
	δ = unit.Angle(math.Asin(sφ*sh - cφ*ch*cA))
	return
}

// GalToEq converts galactic coordinates to equatorial coordinates.
//
// Resulting equatorial coordinates will be referred to the standard equinox of
// B1950.0.  For subsequent conversion to other epochs, see package precess and
// utility functions in package meeus.
func (eq *Equatorial) GalToEq(g *Galactic) *Equatorial {
	eq.RA, eq.Dec = GalToEq(g.Lon, g.Lat)
	return eq
}

var (
	// IAU B1950.0 coordinates of galactic North Pole
	GalacticNorth1950 = &Equatorial{
		RA:  unit.NewRA(12, 49, 0),
		Dec: unit.AngleFromDeg(27.4),
	}
	// Meeus gives 33 as the origin of galactic longitudes relative to the
	// ascending node of of the galactic equator.  33 + 90 = 123, the IAU
	// value for origin relative to the equatorial pole.
	Galactic0Lon1950 = unit.AngleFromDeg(33)
)

// GalToEq converts galactic coordinates to equatorial coordinates.
//
// Resulting equatorial coordinates will be referred to the standard equinox of
// B1950.0.  For subsequent conversion to other epochs, see package precess and
// utility functions in package meeus.
func GalToEq(l, b unit.Angle) (α unit.RA, δ unit.Angle) {
	// (-Galactic0Lon1950 - math.Pi/2) = magic number of -123 deg
	sdLon, cdLon := (l - Galactic0Lon1950 - math.Pi/2).Sincos()
	sgδ, cgδ := GalacticNorth1950.Dec.Sincos()
	sb, cb := b.Sincos()
	y := math.Atan2(sdLon, cdLon*sgδ-(sb/cb)*cgδ)
	// (GalacticNorth1950.RA.Rad() - math.Pi) = magic number of 12.25 deg
	α = unit.RAFromRad(y + GalacticNorth1950.RA.Rad() - math.Pi)
	δ = unit.Angle(math.Asin(sb*sgδ + cb*cgδ*cdLon))
	return
}

// Horizontal coordinates are referenced to the local horizon of an observer
// on the surface of the Earth.
type Horizontal struct {
	Az  unit.Angle // Azimuth (A)
	Alt unit.Angle // Altitude (h)
}

// EqToHz computes Horizontal coordinates from equatorial coordinates.
//
// Argument g is the location of the observer on the Earth.  Argument st
// is the sidereal time at Greenwich.
//
// Sidereal time must be consistent with the equatorial coordinates.
// If coordinates are apparent, sidereal time must be apparent as well.
func (hz *Horizontal) EqToHz(eq *Equatorial, g *globe.Coord, st unit.Time) *Horizontal {
	hz.Az, hz.Alt = EqToHz(eq.RA, eq.Dec, g.Lat, g.Lon, st)
	return hz
}

// EqToHz computes Horizontal coordinates from equatorial coordinates.
//
//	α: right ascension coordinate to transform
//	δ: declination coordinate to transform
//	φ: latitude of observer on Earth
//	ψ: longitude of observer on Earth
//	st: sidereal time at Greenwich at time of observation.
//
// Sidereal time must be consistent with the equatorial coordinates.
// If coordinates are apparent, sidereal time must be apparent as well.
//
// Results:
//
//	A: azimuth of observed point, measured westward from the South.
//	h: elevation, or height of observed point above horizon.
func EqToHz(α unit.RA, δ, φ, ψ unit.Angle, st unit.Time) (A, h unit.Angle) {
	H := st.Rad() - ψ.Rad() - α.Rad()
	sH, cH := math.Sincos(H)
	sφ, cφ := φ.Sincos()
	sδ, cδ := δ.Sincos()
	A = unit.Angle(math.Atan2(sH, cH*sφ-(sδ/cδ)*cφ)) // (13.5) p. 93
	h = unit.Angle(math.Asin(sφ*sδ + cφ*cδ*cH))      // (13.6) p. 93
	return
}

// Galactic coordinates are referenced to the plane of the Milky Way.
type Galactic struct {
	Lat unit.Angle // Latitude (b) in radians
	Lon unit.Angle // Longitude (l) in radians
}

// EqToGal converts equatorial coordinates to galactic coordinates.
//
// Equatorial coordinates must be referred to the standard equinox of B1950.0.
// For conversion to B1950, see package precess and utility functions in
// package "unit".
func (g *Galactic) EqToGal(eq *Equatorial) *Galactic {
	g.Lon, g.Lat = EqToGal(eq.RA, eq.Dec)
	return g
}

// EqToGal converts equatorial coordinates to galactic coordinates.
//
// Equatorial coordinates must be referred to the standard equinox of B1950.0.
// For conversion to B1950, see package precess and utility functions in
// package "common".
func EqToGal(α unit.RA, δ unit.Angle) (l, b unit.Angle) {
	sdα, cdα := (GalacticNorth1950.RA - α).Sincos()
	sgδ, cgδ := GalacticNorth1950.Dec.Sincos()
	sδ, cδ := δ.Sincos()
	// (13.7) p. 94
	x := unit.Angle(math.Atan2(sdα, cdα*sgδ-(sδ/cδ)*cgδ))
	// (Galactic0Lon1950 + 1.5*math.Pi) = magic number of 303 deg
	l = (Galactic0Lon1950 + 1.5*math.Pi - x).Mod1()
	// (13.8) p. 94
	b = unit.Angle(math.Asin(sδ*sgδ + cδ*cgδ*cdα))
	return
}
