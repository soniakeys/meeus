// Fit: Chapter 4, Curve Fitting
package fit

import (
	"math"
)

// Linear fits a line to sample data.
//
// Argument p is a list of data points.  Results a and b are coefficients
// of the best fit line y = ax + b.
func Linear(p []struct{ X, Y float64 }) (a, b float64) {
	var sx, sy, sx2, sxy float64
	for i := range p {
		x := p[i].X
		y := p[i].Y
		sx += x
		sy += y
		sx2 += x * x
		sxy += x * y
	}
	n := float64(len(p))
	d := n*sx2 - sx*sx
	return (n*sxy - sx*sy) / d, (sy*sx2 - sx*sxy) / d
}

// CorrelationCoefficient returns a correlation coefficient for sample data.
func CorrelationCoefficient(p []struct{ X, Y float64 }) float64 {
	var sx, sy, sx2, sy2, sxy float64
	for i := range p {
		x := p[i].X
		y := p[i].Y
		sx += x
		sy += y
		sx2 += x * x
		sy2 += y * y
		sxy += x * y
	}
	n := float64(len(p))
	return (n*sxy - sx*sy) / (math.Sqrt(n*sx2-sx*sx) * math.Sqrt(n*sy2-sy*sy))
}

// Quadratic fits y = ax² + bx + c to sample data.
//
// Argument p is a list of data points.  Results a, b, and c are coefficients
// of the best fit quadratic y = ax² + bx + c.
func Quadratic(p []struct{ X, Y float64 }) (a, b, c float64) {
	var P, Q, R, S, T, U, V float64
	for i := range p {
		x := p[i].X
		y := p[i].Y
		x2 := x * x
		P += x
		Q += x2
		R += x * x2
		S += x2 * x2
		T += y
		U += x * y
		V += x2 * y
	}
	N := float64(len(p))
	D := N*Q*S + 2*P*Q*R - Q*Q*Q - P*P*S - N*R*R
	a = (N*Q*V + P*R*T + P*Q*U - Q*Q*T - P*P*V - N*R*U) / D
	b = (N*S*U + P*Q*V + Q*R*T - Q*Q*U - P*S*T - N*R*V) / D
	c = (Q*S*T + Q*R*U + P*R*V - Q*Q*V - P*S*U - R*R*T) / D
	return
}

// Multiple3 implements multiple linear regression for a linear combination
// of three functions.
//
// Given sample data and three functions in x, Multiple3 returns coefficients
// a, b, and c fitting y = aƒ₀(x) + bƒ₁(x) + cƒ₂(x) to sample data.
func Multiple3(p []struct{ X, Y float64 }, f0, f1, f2 func(float64) float64) (a, b, c float64) {
	var M, P, Q, R, S, T, U, V, W float64
	for i := range p {
		x := p[i].X
		y := p[i].Y
		y0 := f0(x)
		y1 := f1(x)
		y2 := f2(x)
		M += y0 * y0
		P += y0 * y1
		Q += y0 * y2
		R += y1 * y1
		S += y1 * y2
		T += y2 * y2
		U += y * y0
		V += y * y1
		W += y * y2
	}
	D := M*R*T + 2*P*Q*S - M*S*S - R*Q*Q - T*P*P
	a = (U*(R*T-S*S) + V*(Q*S-P*T) + W*(P*S-Q*R)) / D
	b = (U*(S*Q-P*T) + V*(M*T-Q*Q) + W*(P*Q-M*S)) / D
	c = (U*(P*S-R*Q) + V*(P*Q-M*S) + W*(M*R-P*P)) / D
	return
}
