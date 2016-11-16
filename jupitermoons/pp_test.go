// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// +build !nopp

package jupitermoons_test

import (
	"fmt"

	"github.com/soniakeys/meeus/jupitermoons"
	pp "github.com/soniakeys/meeus/planetposition"
)

func ExampleE5() {
	e, err := pp.LoadPlanet(pp.Earth)
	if err != nil {
		fmt.Println(err)
		return
	}
	j, err := pp.LoadPlanet(pp.Jupiter)
	if err != nil {
		fmt.Println(err)
		return
	}
	var pos [4]jupitermoons.XY
	jupitermoons.E5(2448972.50068, e, j, &pos)
	fmt.Printf("X  %+.4f  %+.4f  %+.4f  %+.4f\n",
		pos[0].X, pos[1].X, pos[2].X, pos[3].X)
	fmt.Printf("Y  %+.4f  %+.4f  %+.4f  %+.4f\n",
		pos[0].Y, pos[1].Y, pos[2].Y, pos[3].Y)
	// Output:
	// X  -3.4503  +7.4418  +1.2010  +7.0720
	// Y  +0.2093  +0.2500  +0.6480  +1.0956
}
