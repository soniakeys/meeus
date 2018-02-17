// Copyright 2013 Sonia Keys
// License: MIT

package base

import (
	"math"

	"github.com/soniakeys/unit"
)

// Illuminated returns the illuminated fraction of a body's disk.
//
// The illuminated body can be the Moon or a planet.
//
// Argument i is the phase angle.
func Illuminated(i unit.Angle) float64 {
	// (41.1) p. 283, also (48.1) p. 345.
	return (1 + i.Cos()) * .5
}

// Limb returns the position angle of the midpoint of an illuminated limb.
//
// The illuminated body can be the Moon or a planet.
//
// Arguments α, δ are equatorial coordinates of the body; α0, δ0 are
// apparent coordinates of the Sun.
func Limb(α unit.RA, δ unit.Angle, α0 unit.RA, δ0 unit.Angle) unit.Angle {
	// Mentioned in ch 41, p. 283.  Formula (48.5) p. 346
	sδ, cδ := δ.Sincos()
	sδ0, cδ0 := δ0.Sincos()
	sα0α, cα0α := (α0 - α).Sincos()
	χ := unit.Angle(math.Atan2(cδ0*sα0α, sδ0*cδ-cδ0*sδ*cα0α))
	if χ < 0 {
		χ += 2 * math.Pi
	}
	return χ
}
