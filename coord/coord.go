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
//
// A number of the functions take sine and cosine of the obliquity of the
// ecliptic.  This becomes an advantage when you doing multiple transformations
// with the same obliquity.  The efficiency of computing sine and cosine once
// and reuse these values far outweighs the overhead of passing one number as
// opposed to two.
package coord

import (
	"math"

	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/meeus/globe"
)

// Obliquity represents the obliquity of the ecliptic.
type Obliquity struct {
	S, C float64
}

// NewObliquity constructs a new Obliquity.
//
// Struct members are initialized from the given value ε of the obliquity of
// the ecliptic.
func NewObliquity(ε base.Angle) *Obliquity {
	r := &Obliquity{}
	r.S, r.C = math.Sincos(ε.Rad())
	return r
}

// Ecliptic coordinates are referenced to the plane of the ecliptic.
type Ecliptic struct {
	Lon base.Angle // Longitude (λ)
	Lat base.Angle // Latitude (β)
}

// EqToEcl converts equatorial coordinates to ecliptic coordinates.
func (ecl *Ecliptic) EqToEcl(eq *Equatorial, ε *Obliquity) *Ecliptic {
	sα, cα := math.Sincos(eq.RA.Rad())
	sδ, cδ := math.Sincos(eq.Dec.Rad())
	ecl.Lon = base.Angle(math.Atan2(sα*ε.C+(sδ/cδ)*ε.S, cα)) // (13.1) p. 93
	ecl.Lat = base.Angle(math.Asin(sδ*ε.C - cδ*ε.S*sα))      // (13.2) p. 93
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
func EqToEcl(α base.RA, δ base.Angle, sε, cε float64) (λ, β base.Angle) {
	sα, cα := math.Sincos(α.Rad())
	sδ, cδ := math.Sincos(δ.Rad())
	λ = base.Angle(math.Atan2(sα*cε+(sδ/cδ)*sε, cα)) // (13.1) p. 93
	β = base.Angle(math.Asin(sδ*cε - cδ*sε*sα))      // (13.2) p. 93
	return
}

// Equatorial coordinates are referenced to the Earth's rotational axis.
type Equatorial struct {
	RA  base.RA    // Right ascension (α)
	Dec base.Angle // Declination (δ)
}

// EclToEq converts ecliptic coordinates to equatorial coordinates.
func (eq *Equatorial) EclToEq(ecl *Ecliptic, ε *Obliquity) *Equatorial {
	sβ, cβ := math.Sincos(ecl.Lat.Rad())
	sλ, cλ := math.Sincos(ecl.Lon.Rad())
	eq.RA = base.RAFromRad(math.Atan2(sλ*ε.C-(sβ/cβ)*ε.S, cλ)) // (13.3) p. 93
	eq.Dec = base.Angle(math.Asin(sβ*ε.C + cβ*ε.S*sλ))         // (13.4) p. 93
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
func EclToEq(λ, β base.Angle, sε, cε float64) (α base.RA, δ base.Angle) {
	sλ, cλ := math.Sincos(λ.Rad())
	sβ, cβ := math.Sincos(β.Rad())
	α = base.RAFromRad(math.Atan2(sλ*cε-(sβ/cβ)*sε, cλ)) // (13.3) p. 93
	δ = base.Angle(math.Asin(sβ*cε + cβ*sε*sλ))          // (13.4) p. 93
	return
}

// HzToEq transforms horizontal coordinates to equatorial coordinates.
//
// Sidereal time st must be consistent with the equatorial coordinates
// in the sense that if coordinates are apparent, sidereal time must be
// apparent as well.
func (eq *Equatorial) HzToEq(hz *Horizontal, g globe.Coord, st base.Time) *Equatorial {
	sA, cA := math.Sincos(hz.Az.Rad())
	sh, ch := math.Sincos(hz.Alt.Rad())
	sφ, cφ := math.Sincos(g.Lat.Rad())
	H := math.Atan2(sA, cA*sφ+sh/ch*cφ)
	eq.RA = base.RAFromRad(st.Rad() - g.Lon.Rad() - H)
	eq.Dec = base.Angle(math.Asin(sφ*sh - cφ*ch*cA))
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
func HzToEq(A, h, φ, ψ base.Angle, st base.Time) (α base.RA, δ base.Angle) {
	sA, cA := math.Sincos(A.Rad())
	sh, ch := math.Sincos(h.Rad())
	sφ, cφ := math.Sincos(φ.Rad())
	H := math.Atan2(sA, cA*sφ+sh/ch*cφ)
	α = base.RAFromRad(st.Rad() - ψ.Rad() - H)
	δ = base.Angle(math.Asin(sφ*sh - cφ*ch*cA))
	return
}

// GalToEq converts galactic coordinates to equatorial coordinates.
//
// Resulting equatorial coordinates will be referred to the standard equinox of
// B1950.0.  For subsequent conversion to other epochs, see package precess and
// utility functions in package meeus.
func (eq *Equatorial) GalToEq(g *Galactic) *Equatorial {
	sdLon, cdLon := math.Sincos((g.Lon - galacticLon0).Rad())
	sgδ, cgδ := math.Sincos(galacticNorth.Dec.Rad())
	sb, cb := math.Sincos(g.Lat.Rad())
	y := math.Atan2(sdLon, cdLon*sgδ-(sb/cb)*cgδ)
	eq.RA = base.RAFromRad(y + galacticNorth.RA.Rad())
	eq.Dec = base.Angle(math.Asin(sb*sgδ + cb*cgδ*cdLon))
	return eq
}

// GalToEq converts galactic coordinates to equatorial coordinates.
//
// Resulting equatorial coordinates will be referred to the standard equinox of
// B1950.0.  For subsequent conversion to other epochs, see package precess and
// utility functions in package meeus.
func GalToEq(l, b base.Angle) (α base.RA, δ base.Angle) {
	sdLon, cdLon := math.Sincos((l - galacticLon0).Rad())
	sgδ, cgδ := math.Sincos(galacticNorth.Dec.Rad())
	sb, cb := math.Sincos(b.Rad())
	y := math.Atan2(sdLon, cdLon*sgδ-(sb/cb)*cgδ)
	α = base.RAFromRad(y + galacticNorth.RA.Rad())
	δ = base.Angle(math.Asin(sb*sgδ + cb*cgδ*cdLon))
	return
}

// Horizontal coordinates are referenced to the local horizon of an observer
// on the surface of the Earth.
type Horizontal struct {
	Az  base.Angle // Azimuth (A)
	Alt base.Angle // Altitude (h)
}

// EqToHz computes Horizontal coordinates from equatorial coordinates.
//
// Argument g is the location of the observer on the Earth.  Argument st
// is the sidereal time at Greenwich.
//
// Sidereal time must be consistent with the equatorial coordinates.
// If coordinates are apparent, sidereal time must be apparent as well.
func (hz *Horizontal) EqToHz(eq *Equatorial, g *globe.Coord, st base.Time) *Horizontal {
	H := st.Rad() - g.Lon.Rad() - eq.RA.Rad()
	sH, cH := math.Sincos(H)
	sφ, cφ := math.Sincos(g.Lat.Rad())
	sδ, cδ := math.Sincos(eq.Dec.Rad())
	hz.Az = base.Angle(math.Atan2(sH, cH*sφ-(sδ/cδ)*cφ)) // (13.5) p. 93
	hz.Alt = base.Angle(math.Asin(sφ*sδ + cφ*cδ*cH))     // (13.6) p. 93
	return hz
}

// EqToHz computes Horizontal coordinates from equatorial coordinates.
//
//	α: right ascension coordinate to transform, in radians
//	δ: declination coordinate to transform, in radians
//	φ: latitude of observer on Earth
//	ψ: longitude of observer on Earth
//	st: sidereal time at Greenwich at time of observation.
//
// Sidereal time must be consistent with the equatorial coordinates.
// If coordinates are apparent, sidereal time must be apparent as well.
//
// Results:
//
//	A: azimuth of observed point in radians, measured westward from the South.
//	h: elevation, or height of observed point in radians above horizon.
func EqToHz(α, δ, φ, ψ, st float64) (A, h float64) {
	H := base.Time(st).Rad() - ψ - α
	sH, cH := math.Sincos(H)
	sφ, cφ := math.Sincos(φ)
	sδ, cδ := math.Sincos(ψ)
	A = math.Atan2(sH, cH*sφ-(sδ/cδ)*cφ) // (13.5) p. 93
	h = math.Asin(sφ*sδ + cφ*cδ*cH)      // (13.6) p. 93
	return
}

// Galactic coordinates are referenced to the plane of the Milky Way.
type Galactic struct {
	Lat base.Angle // Latitude (b) in radians
	Lon base.Angle // Longitude (l) in radians
}

var galacticNorth = &Equatorial{
	RA:  base.RAFromHour(base.FromSexa(0, 12, 49, 0)),
	Dec: base.AngleFromDeg(27.4),
}

var galacticLon0 = base.AngleFromDeg(123)

// EqToGal converts equatorial coordinates to galactic coordinates.
//
// Equatorial coordinates must be referred to the standard equinox of B1950.0.
// For conversion to B1950, see package precess and utility functions in
// package "base".
func (g *Galactic) EqToGal(eq *Equatorial) *Galactic {
	sdα, cdα := math.Sincos((galacticNorth.RA - eq.RA).Rad())
	sgδ, cgδ := math.Sincos(galacticNorth.Dec.Rad())
	sδ, cδ := math.Sincos(eq.Dec.Rad())
	// (13.7) p. 94
	x := math.Atan2(sdα, cdα*sgδ-(sδ/cδ)*cgδ)
	g.Lon = base.Angle(math.Mod(math.Pi+galacticLon0.Rad()-x, 2*math.Pi))
	// (13.8) p. 94
	g.Lat = base.Angle(math.Asin(sδ*sgδ + cδ*cgδ*cdα))
	return g
}

// EqToGal converts equatorial coordinates to galactic coordinates.
//
// Equatorial coordinates must be referred to the standard equinox of B1950.0.
// For conversion to B1950, see package precess and utility functions in
// package "common".
func EqToGal(α base.RA, δ base.Angle) (l, b base.Angle) {
	sdα, cdα := math.Sincos((galacticNorth.RA - α).Rad())
	sgδ, cgδ := math.Sincos(galacticNorth.Dec.Rad())
	sδ, cδ := math.Sincos(δ.Rad())
	// (13.7) p. 94
	x := math.Atan2(sdα, cdα*sgδ-(sδ/cδ)*cgδ)
	l = base.Angle(math.Mod(math.Pi+galacticLon0.Rad()-x, 2*math.Pi))
	// (13.8) p. 94
	b = base.Angle(math.Asin(sδ*sgδ + cδ*cgδ*cdα))
	return
}
