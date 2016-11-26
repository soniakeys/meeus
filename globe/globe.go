// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Globe: Chapter 11, The Earth's Globe.
//
// Globe contains functions concerning the surface of the Earth idealized as
// an ellipsoid of revolution.
package globe

import (
	"math"

	"github.com/soniakeys/meeus/base"
)

// Ellipsoid represents an ellipsoid of revolution.
//
// Typical unit for Er is Km.
type Ellipsoid struct {
	Er float64 // equatorial radius
	Fl float64 // flattening
}

// IAU 1976 values.  Radius in Km.
var Earth76 = Ellipsoid{Er: 6378.14, Fl: 1 / 298.257}

// A returns equatorial radius in units of e.Er.
//
// A is a common identifier for equatorial radius.
func (e Ellipsoid) A() float64 {
	return e.Er
}

// B returns polar radius in units of e.ER.
//
// B is a common identifier for polar radius.
func (e Ellipsoid) B() float64 {
	return e.Er * (1 - e.Fl)
}

// Eccentricity of a meridian.
func (e Ellipsoid) Eccentricity() float64 {
	return math.Sqrt((2 - e.Fl) * e.Fl)
}

// ParallaxConstants computes parallax constants ρ sin φ′ and ρ cos φ′.
//
// Arguments are geographic latitude φ and height h above the ellipsoid.
// For e.Er in Km, h must be in meters.
func (e Ellipsoid) ParallaxConstants(φ base.Angle, h float64) (s, c float64) {
	boa := 1 - e.Fl
	su, cu := math.Sincos(math.Atan(boa * math.Tan(φ.Rad())))
	s, c = math.Sincos(φ.Rad())
	hoa := h * 1e-3 / e.Er
	return su*boa + hoa*s, cu + hoa*c
}

// Rho is distance from Earth center to a point on the ellipsoid at latitude φ.
//
// Result unit is fraction of the equatorial radius.
func Rho(φ base.Angle) float64 {
	// Magic numbers...
	return .9983271 + .0016764*math.Cos(2*φ.Rad()) -
		.0000035*math.Cos(4*φ.Rad())
}

// RadiusAtLatitude returns the radius of the circle that is the parallel of
// latitude φ.
//
// Result unit is same as e.Er (typically Km.)
func (e Ellipsoid) RadiusAtLatitude(φ base.Angle) float64 {
	s, c := math.Sincos(φ.Rad())
	return e.A() * c / math.Sqrt(1-(2-e.Fl)*e.Fl*s*s)
}

// OneDegreeOfLongitude returns the length of one degree of longitude.
//
// Argument rp is the radius of a circle that is a parallel of latitude
// (as returned by Ellipsoid.RadiusAtLatitude.)
//
// Result is distance along one degree of the circle, in same units as rp.
func OneDegreeOfLongitude(rp float64) float64 {
	return rp * math.Pi / 180
}

// RotationRate1996_5 is the rotational angular velocity of the Earth
// with respect to the stars at the epoch 1996.5.
//
// Unit is radian/second.
const RotationRate1996_5 = 7.292114992e-5

// RadiusOfCurvature of meridian at latitude φ.
//
// Result in units of e.ER, typically Km.
func (e Ellipsoid) RadiusOfCurvature(φ base.Angle) float64 {
	s := math.Sin(φ.Rad())
	e2 := (2 - e.Fl) * e.Fl
	return e.A() * (1 - e2) / math.Pow(1-e2*s*s, 1.5)
}

// OneDegreeOfLatitude returns the length of one degree of latitude.
//
// Argument rm is the radius of curvature along a meridian.
// (as returned by Ellipsoid.RadiusOfCurvature.)
// Result is distance in units of rm along one degree of the meridian.
func OneDegreeOfLatitude(rm float64) float64 {
	return rm * math.Pi / 180
}

// GeocentricLatitudeDifference returns geographic latitude - geocentric
// latitude (φ - φ′) given geographic latitude (φ).
func GeocentricLatitudeDifference(φ base.Angle) base.Angle {
	// This appears to be an approximation with hard coded magic numbers.
	// No explanation is given in the text. The ellipsoid is not specified.
	// Perhaps the approximation works well enough for all ellipsoids?
	return base.AngleFromSec(692.73*math.Sin(2*φ.Rad()) -
		1.16*math.Sin(4*φ.Rad()))
}

// Coord represents geographic coordinates on the Earth.
//
// Longitude is measured positively westward from the Greenwich meridian.
type Coord struct {
	Lat base.Angle // latitude (φ)
	Lon base.Angle // longitude (ψ, or L)
}

// ApproxAngularDistance returns the cosine of the angle between two points.
//
// The accuracy deteriorates at small angles.
func ApproxAngularDistance(p1, p2 Coord) float64 {
	s1, c1 := math.Sincos(p1.Lat.Rad())
	s2, c2 := math.Sincos(p2.Lat.Rad())
	return s1*s2 + c1*c2*math.Cos((p1.Lon-p2.Lon).Rad())
}

// ApproxLinearDistance computes a distance across the surface of the Earth.
//
// Approximating the Earth as a sphere, the function takes a geocentric angular
// distance and returns the corresponding linear distance in Km.
func ApproxLinearDistance(d base.Angle) float64 {
	return 6371 * d.Rad()
}

// Distance is distance between two points measured along the surface
// of an ellipsoid.
//
// Accuracy is much better than that of ApproxAngularDistance or
// ApproxLinearDistance.
//
// Result unit is units of e.Er, typically Km.
func (e Ellipsoid) Distance(c1, c2 Coord) float64 {
	// From AA, ch 11, p 84.
	s2f, c2f := sincos2((c1.Lat + c2.Lat) / 2)
	s2g, c2g := sincos2((c1.Lat - c2.Lat) / 2)
	s2λ, c2λ := sincos2((c1.Lon - c2.Lon) / 2)
	s := s2g*c2λ + c2f*s2λ
	c := c2g*c2λ + s2f*s2λ
	ω := math.Atan(math.Sqrt(s / c))
	r := math.Sqrt(s*c) / ω
	d := 2 * ω * e.Er
	h1 := (3*r - 1) / (2 * c)
	h2 := (3*r + 1) / (2 * s)
	return d * (1 + e.Fl*(h1*s2f*c2g-h2*c2f*s2g))
}

// small function should expand inline
func sincos2(x base.Angle) (s2, c2 float64) {
	s, c := math.Sincos(x.Rad())
	return s * s, c * c
}
