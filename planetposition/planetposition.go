// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Planetposition: Chapter 32, Positions of the Planets.
//
// Incomplete:
//
// 1.  The package does not implement algorithms that use appendix III,
// but instead implements a full VSOP87 solution.  I do not have a copy
// of the supplimentary disk with appendix III in machine readable form
// and as the appendix is rather large, retyping it by hand is problematic.
// The full VSOP87 data set on the other hand is freely downloadable from
// the internet, so I implement here code that can use that data directly.
//
// 2.  The formula for accuracy of results is not implemented.  It is
// not needed for full VSOP87 solutions.
//
// 3.  Polynomial expressions are not implemented.  Again, implementation
// would involve typing rather large tables of numbers with associated
// risk of typographical errors.
package planetposition

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/meeus/coord"
	"github.com/soniakeys/meeus/precess"
	"github.com/soniakeys/unit"
)

// Mercury-Neptune planet constants suitable for first argument to LoadPlanet.
const (
	Mercury = iota
	Venus
	Earth
	Mars
	Jupiter
	Saturn
	Uranus
	Neptune
	nPlanets // sad practicality
)

// parallel arrays, indexed by planet constants.
var (
	// extensions of VSOP87B files
	ext = [nPlanets]string{
		"mer", "ven", "ear", "mar", "jup", "sat", "ura", "nep"}

	// planet names as found in VSOP87B files
	b7 = [nPlanets]string{
		"MERCURY",
		"VENUS  ",
		"EARTH  ",
		"MARS   ",
		"JUPITER",
		"SATURN ",
		"URANUS ",
		"NEPTUNE",
	}
)

type abc struct {
	a, b, c float64
}

type coeff [6][]abc

// V87Planet holds VSOP87 coefficients for computing planetary
// positions in spherical coorditates.
type V87Planet struct {
	l, b, r coeff
}

// code tested with version 2.  other versions unknown.
const fileVersion = '2'

// LoadPlanet constructs a V87Planet object from a VSOP87 file.
//
// Argument ibody should be one of the planet constants.
//
// The directory containing the VSOP87 must be indicated by environment
// variable VSOP87.
func LoadPlanet(ibody int) (*V87Planet, error) {
	path := os.Getenv("VSOP87")
	if path == "" {
		return nil, errors.New("No path assigned to environment variable VSOP87")
	}
	return LoadPlanetPath(ibody, path)
}

// LoadPlanetPath constructs a V87Planet object from a VSOP87 file.
//
// Argument ibody should be one of the planet constants; path should be
// a directory containing the VSOP87 files.
func LoadPlanetPath(ibody int, path string) (*V87Planet, error) {
	if ibody < 0 || ibody >= nPlanets {
		return nil, errors.New("Invalid planet.")
	}
	data, err := ioutil.ReadFile(path + "/VSOP87B." + ext[ibody])
	if err != nil {
		return nil, err
	}
	v := &V87Planet{}
	lines := strings.Split(string(data), "\n")
	n := 0
	n, err = v.l.parse('1', ibody, lines, n, false)
	if err != nil {
		return nil, err
	}
	n, err = v.b.parse('2', ibody, lines, n, false)
	if err != nil {
		return nil, err
	}
	n, err = v.r.parse('3', ibody, lines, n, true)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (c *coeff) parse(ic byte, ibody int, lines []string, n int, au bool) (int, error) {
	var cbuf [2047]abc
	for n < len(lines) {
		line := lines[n]
		if len(line) < 132 {
			break
		}
		if line[41] != ic {
			break
		}
		if iv := line[17]; iv != fileVersion {
			return n, fmt.Errorf("Line %d: expected version %c, "+
				"found %c.", n+1, fileVersion, iv)
		}
		if bo := line[22:29]; bo != b7[ibody] {
			return n, fmt.Errorf("Line %d: expected body %s, "+
				"found %s.", n+1, b7[ibody], bo)
		}
		it := line[59] - '0'
		in, err := strconv.Atoi(strings.TrimSpace(line[60:67]))
		if err != nil {
			return n, fmt.Errorf("Line %d: %v.", n+1, err)
		}
		if in == 0 {
			continue
		}
		if in > len(lines)-n {
			return n, errors.New("Unexpected end of file.")
		}
		n++
		cx := 0
		for _, line := range lines[n : n+in] {
			a := &cbuf[cx]
			a.a, err =
				strconv.ParseFloat(strings.TrimSpace(line[79:97]), 64)
			if err != nil {
				goto parseError
			}
			a.b, err = strconv.ParseFloat(line[98:111], 64)
			if err != nil {
				goto parseError
			}
			a.c, err =
				strconv.ParseFloat(strings.TrimSpace(line[111:131]), 64)
			if err != nil {
				goto parseError
			}
			cx++
			continue
		parseError:
			return n, fmt.Errorf("Line %d: %v.", n+cx+1, err)
		}
		c[it] = append([]abc{}, cbuf[:cx]...)
		n += in
	}
	return n, nil
}

// Position2000 returns ecliptic position of planets by full VSOP87 theory.
//
// Argument jde is the date for which positions are desired.
//
// Results are for the dynamical equinox and ecliptic J2000.
//
//	L is heliocentric longitude.
//	B is heliocentric latitude.
//	R is heliocentric range in AU.
func (vt *V87Planet) Position2000(jde float64) (L, B unit.Angle, R float64) {
	T := base.J2000Century(jde)
	τ := T * .1
	cf := make([]float64, 6)
	sum := func(series coeff) float64 {
		for x, terms := range series {
			cf[x] = 0
			// sum terms in reverse order to preserve accuracy
			for y := len(terms) - 1; y >= 0; y-- {
				term := &terms[y]
				cf[x] += term.a * math.Cos(term.b+term.c*τ)
			}
		}
		return base.Horner(τ, cf[:len(series)]...)
	}
	L = unit.Angle(unit.PMod(sum(vt.l), 2*math.Pi))
	B = unit.Angle(sum(vt.b))
	R = sum(vt.r)
	return
}

// Position returns ecliptic position of planets at equinox and ecliptic of date.
//
// Argument jde is the date for which positions are desired.
//
// Results are positions consistent with those from Meeus's Apendix III,
// that is, at equinox and ecliptic of date.
//
//  L is heliocentric longitude.
//  B is heliocentric latitude.
//  R is heliocentric range in AU.
func (vt *V87Planet) Position(jde float64) (L, B unit.Angle, R float64) {
	L, B, R = vt.Position2000(jde)
	eclFrom := &coord.Ecliptic{
		Lat: B,
		Lon: L,
	}
	eclTo := &coord.Ecliptic{}
	epochFrom := 2000.0
	epochTo := base.JDEToJulianYear(jde)
	precess.EclipticPosition(eclFrom, eclTo, epochFrom, epochTo, 0, 0)
	return eclTo.Lon, eclTo.Lat, R
}

// ToFK5 converts ecliptic longitude and latitude from dynamical frame to FK5.
func ToFK5(L, B unit.Angle, jde float64) (L5, B5 unit.Angle) {
	// formula 32.3, p. 219.
	T := base.J2000Century(jde)
	Lp := L - unit.AngleFromDeg(1.397*T-.00031*T*T)
	sLp, cLp := math.Sincos(Lp.Rad())
	// (32.3) p. 219
	L5 = L + unit.AngleFromSec(-.09033+.03916*(cLp+sLp)*math.Tan(B.Rad()))
	B5 = B + unit.AngleFromSec(.03916*(cLp-sLp))
	return
}
