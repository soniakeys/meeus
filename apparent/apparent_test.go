// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package apparent_test

import (
	"fmt"
	//	"math"

	"github.com/soniakeys/meeus/apparent"
	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/meeus/julian"

//	"github.com/soniakeys/meeus/nutation"
//	"github.com/soniakeys/meeus/solar"
)

func ExampleNutationCorrection() {
	α := base.NewRA(2, 46, 11.331).Rad()
	δ := base.NewAngle(false, 49, 20, 54.54).Rad()
	jd := julian.CalendarGregorianToJD(2028, 11, 13.19)
	Δα1, Δδ1 := apparent.Nutation(α, δ, jd)
	fmt.Printf("%.3s  %.3s\n", base.NewFmtAngle(Δα1), base.NewFmtAngle(Δδ1))
	Δα2, Δδ2 := apparent.EquatorialAbberation(α, δ, jd)
	fmt.Printf("%.3s  %.3s\n", base.NewFmtAngle(Δα2), base.NewFmtAngle(Δδ2))
	// Output:
	// 15.843″  6.217″
	// 30.045″  6.697″
}
