// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package base

import "math"

// Illuminated returns the illuminated fraction of a body's disk.
//
// The illuminated body can be the Moon or a planet.
//
// Argument i is the phase angle.
func Illuminated(i Angle) float64 {
	// (41.1) p. 283, also (48.1) p. 345.
	return (1 + math.Cos(i.Rad())) * .5
}

// Limb returns the position angle of the midpoint of an illuminated limb.
//
// The illuminated body can be the Moon or a planet.
//
// Arguments α, δ are equatorial coordinates of the body; α0, δ0 are
// apparent coordinates of the Sun.
func Limb(α RA, δ Angle, α0 RA, δ0 Angle) Angle {
	// Mentioned in ch 41, p. 283.  Formula (48.5) p. 346
	sδ, cδ := math.Sincos(δ.Rad())
	sδ0, cδ0 := math.Sincos(δ0.Rad())
	sα0α, cα0α := math.Sincos((α0 - α).Rad())
	χ := Angle(math.Atan2(cδ0*sα0α, sδ0*cδ-cδ0*sδ*cα0α))
	if χ < 0 {
		χ += 2 * math.Pi
	}
	return χ
}
