// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Planetary: Chapter 36, The Calculation of some Planetary Phenomena.
//
// Incomplete: Some functions unimplemented for lack of test data.
package planetary

import (
	"math"

	"github.com/soniakeys/meeus/base"
)

// Mean computes some intermediate values for a mean planetary configuration
// given a year and a row of coefficients from Table 36.A, p. 250.
func mean(y float64, a *ca) (J, M, T float64) {
	// (36.1) p. 250
	k := math.Floor((365.2425*y+1721060-a.A)/a.B + .5)
	J = a.A + k*a.B
	M = base.PMod(a.M0+k*a.M1, 360) * math.Pi / 180
	T = base.J2000Century(J)
	return
}

// Sum computes a sum of periodic terms.
func sum(T, M float64, c [][]float64) float64 {
	j := base.Horner(T, c[0]...)
	mm := 0.
	for i := 1; i < len(c); i++ {
		mm += M
		smm, cmm := math.Sincos(mm)
		j += smm * base.Horner(T, c[i]...)
		i++
		j += cmm * base.Horner(T, c[i]...)
	}
	return j
}

// Ms returns a mean time corrected by a sum.
func ms(y float64, a *ca, c [][]float64) float64 {
	J, M, T := mean(y, a)
	return J + sum(T, M, c)
}

// MercuryInfConj returns the time of an inferior conjunction of Mercury.
//
// Result is time (as a jde) of the event nearest the given time (as a
// decimal year.)
func MercuryInfConj(y float64) (jde float64) {
	return ms(y, micA, micB)
}

// MercurySupConj returns the time of a superior conjunction of Mercury.
//
// Result is time (as a jde) of the event nearest the given time (as a
// decimal year.)
func MercurySupConj(y float64) (jde float64) {
	return ms(y, mscA, mscB)
}

// VenusInfConj returns the time of an inferior conjunction of Venus.
//
// Result is time (as a jde) of the event nearest the given time (as a
// decimal year.)
func VenusInfConj(y float64) (jde float64) {
	return ms(y, vicA, vicB)
}

// MarsOpp returns the time of an opposition of Mars.
//
// Result is time (as a jde) of the event nearest the given time (as a
// decimal year.)
func MarsOpp(y float64) (jde float64) {
	return ms(y, moA, moB)
}

// SumA computes the sum of periodic terms with "additional angles"
func sumA(T, M float64, c [][]float64, aa []caa) float64 {
	i := len(c) - 2*len(aa)
	j := sum(T, M, c[:i])
	for k := 0; k < len(aa); k++ {
		saa, caa := math.Sincos((aa[k].c + aa[k].f*T) * math.Pi / 180)
		j += saa * base.Horner(T, c[i]...)
		i++
		j += caa * base.Horner(T, c[i]...)
		i++
	}
	return j
}

// Msa returns a mean time corrected by a sum.
func msa(y float64, a *ca, c [][]float64, aa []caa) float64 {
	J, M, T := mean(y, a)
	return J + sumA(T, M, c, aa)
}

// JupiterOpp returns the time of an opposition of Jupiter.
//
// Result is time (as a jde) of the event nearest the given time (as a
// decimal year.)
func JupiterOpp(y float64) (jde float64) {
	return msa(y, joA, joB, jaa)
}

// SaturnOpp returns the time of an opposition of Saturn.
//
// Result is time (as a jde) of the event nearest the given time (as a
// decimal year.)
func SaturnOpp(y float64) (jde float64) {
	return msa(y, soA, soB, saa)
}

// SaturnConj returns the time of a conjunction of Saturn.
//
// Result is time (as a jde) of the event nearest the given time (as a
// decimal year.)
func SaturnConj(y float64) (jde float64) {
	return msa(y, scA, scB, saa)
}

// UranusOpp returns the time of an opposition of Uranus.
//
// Result is time (as a jde) of the event nearest the given time (as a
// decimal year.)
func UranusOpp(y float64) (jde float64) {
	return msa(y, uoA, uoB, uaa)
}

// NeptuneOpp returns the time of an opposition of Neptune.
//
// Result is time (as a jde) of the event nearest the given time (as a
// decimal year.)
func NeptuneOpp(y float64) (jde float64) {
	return msa(y, noA, noB, naa)
}

// El computes time and elongation of a greatest elongation event.
func el(y float64, a *ca, t, e [][]float64) (jde, elongation float64) {
	J, M, T := mean(y, micA)
	return J + sum(T, M, t), sum(T, M, e) * math.Pi / 180
}

// MercuryEastElongation returns the time and elongation of a greatest eastern elongation of Mercury.
//
// Result is time (as a jde) of the event nearest the given time (as a
// decimal year.)
func MercuryEastElongation(y float64) (jde, elongation float64) {
	return el(y, micA, met, mee)
}

// MercuryWestElongation returns the time and elongation of a greatest western elongation of Mercury.
//
// Result is time (as a jde) of the event nearest the given time (as a
// decimal year.)
func MercuryWestElongation(y float64) (jde, elongation float64) {
	return el(y, micA, mwt, mwe)
}

func MarsStation2(y float64) (jde float64) {
	J, M, T := mean(y, moA)
	return J + sum(T, M, ms2)
}

// ca holds coefficients from one line of table 36.A, p. 250
type ca struct {
	A, B, M0, M1 float64
}

// Table 36.A, p. 250
var (
	micA = &ca{2451612.023, 115.8774771, 63.5867, 114.2088742}
	mscA = &ca{2451554.084, 115.8774771, 6.4822, 114.2088742}
	vicA = &ca{2451996.706, 583.921361, 82.7311, 215.513058}

	moA = &ca{2452097.382, 779.936104, 181.9573, 48.705244}

	joA = &ca{2451870.628, 398.884046, 318.4681, 33.140229}

	soA = &ca{2451870.17, 378.091904, 318.0172, 12.647487}
	scA = &ca{2451681.124, 378.091904, 131.6934, 12.647487}
	uoA = &ca{2451764.317, 369.656035, 213.6884, 4.333093}

	noA = &ca{2451753.122, 367.486703, 202.6544, 2.194998}
)

// caa holds coefficents for "additional angles" for outer planets
// as given on p. 251
type caa struct {
	c, f float64
}

var jaa = []caa{
	{82.74, 40.76},
}

var saa = []caa{
	{82.74, 40.76},
	{29.86, 1181.36},
	{14.13, 590.68},
	{220.02, 1262.87},
}

var uaa = []caa{
	{207.83, 8.51},
	{108.84, 419.96},
}

var naa = []caa{
	{207.83, 8.51},
	{276.74, 209.98},
}

// Table 33.B, p. 256

// Mercury inferior conjunction
var micB = [][]float64{
	{.0545, .0002},
	{-6.2008, .0074, .00003},
	{-3.275, -.0197, .00001},
	{.4737, -.0052, -.00001},
	{.8111, .0033, -.00002},
	{.0037, .0018},
	{-.1768, 0, .00001},
	{-.0211, -.0004},
	{.0326, -.0003},
	{.0083, .0001},
	{-.004, .0001},
}

// Mercury superior conjunction
var mscB = [][]float64{
	{-.0548, -.0002},
	{7.3894, -.01, -.00003},
	{3.22, .0197, -.00001},
	{.8383, -.0064, -.00001},
	{.9666, .0039, -.00003},
	{.077, -.0026},
	{.2758, .0002, -.00002},
	{-.0128, -.0008},
	{.0734, -.0004, -.00001},
	{-.0122, -.0002},
	{.0173, -.0002},
}

// Venus inferior conjunction
var vicB = [][]float64{
	{-.0096, .0002, -.00001},
	{2.0009, -.0033, -.00001},
	{.598, -.0104, .00001},
	{.0967, -.0018, -.00003},
	{.0913, .0009, -.00002},
	{.0046, -.0002},
	{.0079, .0001},
}

// Mars opposition
var moB = [][]float64{
	{-.3088, 0, .00002},
	{-17.6965, .0363, .00005},
	{18.3131, .0467, -.00006},
	{-.2162, -.0198, -.00001},
	{-4.5028, -.0019, .00007},
	{.8987, .0058, -.00002},
	{.7666, -.005, -.00003},
	{-.3636, -.0001, .00002},
	{.0402, .0032},
	{.0737, -.0008},
	{-.098, -.0011},
}

// Jupiter opposition
var joB = [][]float64{
	{-.1029, 0, -.00009},
	{-1.9658, -.0056, .00007},
	{6.1537, .021, -.00006},
	{-.2081, -.0013},
	{-.1116, -.001},
	{.0074, .0001},
	{-.0097, -.0001},
	{0, .0144, -.00008},
	{.3642, -.0019, -.00029},
}

// Saturn opposition
var soB = [][]float64{
	{-.0209, .0006, .00023},
	{4.5795, -.0312, -.00017},
	{1.1462, -.0351, .00011},
	{.0985, -.0015},
	{.0733, -.0031, .00001},
	{.0025, -.0001},
	{.005, -.0002},
	{0, -.0337, .00018},
	{-.851, .0044, .00068},
	{0, -.0064, .00004},
	{.2397, -.0012, -.00008},
	{0, -.001},
	{.1245, .0006},
	{0, .0024, -.00003},
	{.0477, -.0005, -.00006},
}

// Saturn conjunction
var scB = [][]float64{
	{.0172, -.0006, .00023},
	{-8.5885, .0411, .00020},
	{-1.147, .0352, -.00011},
	{.3331, -.0034, -.00001},
	{.1145, -.0045, .00002},
	{-.0169, .0002},
	{-.0109, .0004},
	{0, -.0337, .00018},
	{-.851, .0044, .00068},
	{0, -.0064, .00004},
	{.2397, -.0012, -.00008},
	{0, -.001},
	{.1245, .0006},
	{0, .0024, -.00003},
	{.0477, -.0005, -.00006},
}

// Uranus opposition
var uoB = [][]float64{
	{.0844, -.0006},
	{-.1048, .0246},
	{-5.1221, .0104, .00003},
	{-.1428, .0005},
	{-.0148, -.0013},
	{0},
	{.0055},
	{0},
	{.885},
	{0},
	{.2153},
}

// Neptune opposition {
var noB = [][]float64{
	{-.014, 0, .00001},
	{-1.3486, .001, .00001},
	{.8597, .0037},
	{-.0082, -.0002, .00001},
	{.0037, -.0003},
	{0},
	{-.5964},
	{0},
	{.0728},
}

// Table 36.C, p. 259

// Mercury east time correction
var met = [][]float64{
	{-21.6106, .0002},
	{-1.9803, -.006, .00001},
	{1.4151, -.0072, -.00001},
	{.5528, -.0005, -.00001},
	{.2905, .0034, .00001},
	{-.1121, -.0001, .00001},
	{-.0098, -.0015},
	{.0192},
	{.0111, .0004},
	{-.0061},
	{-.0032, -.0001},
}

// Mercury east elongation
var mee = [][]float64{
	{22.4697},
	{-4.2666, .0054, .00002},
	{-1.8537, -.0137},
	{.3598, .0008, -.00001},
	{-.068, .0026},
	{-.0524, -.0003},
	{.0052, -.0006},
	{.0107, .0001},
	{-.0013, .0001},
	{-.0021},
	{.0003},
}

// Mercury west time correction
var mwt = [][]float64{
	{21.6249, -.0002},
	{.1306, .0065},
	{-2.7661, -.0011, .00001},
	{.2438, -.0024, -.00001},
	{.5767, .0023},
	{.1041},
	{-.0184, .0007},
	{-.0051, -.0001},
	{.0048, .0001},
	{.0026},
	{.0037},
}

// Mercury west elongation
var mwe = [][]float64{
	{22.4143, -.0001},
	{4.3651, -.0048, -.00002},
	{2.3787, .0121, -.00001},
	{.2674, .0022},
	{-.3873, .0008, .00001},
	{-.0369, -.0001},
	{.0017, -.0001},
	{.0059},
	{.0061, .0001},
	{.0007},
	{-.0011},
}

// Table 36.D, p. 261

// Mars Station 2
var ms2 = [][]float64{
	{36.7191, .0016, .00003},
	{-12.6163, .0417, -.00001},
	{20.1218, .0379, -.00006},
	{-1.636, -.019},
	{-3.9657, .0045, .00007},
	{1.1546, .0029, -.00003},
	{.2888, -.0073, -.00002},
	{-.3128, .0017, .00002},
	{.2513, .0026, -.00002},
	{-.0021, -.0016},
	{-.1497, -.0006},
}
