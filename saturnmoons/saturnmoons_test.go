package saturnmoons_test

import (
	"fmt"

	pp "github.com/soniakeys/meeus/planetposition"
	"github.com/soniakeys/meeus/saturnmoons"
)

func ExamplePositions() {
	// Example 46.a, p. 334.
	earth, err := pp.LoadPlanet(pp.Earth, "")
	if err != nil {
		fmt.Println(err)
		return
	}
	saturn, err := pp.LoadPlanet(pp.Saturn, "")
	if err != nil {
		fmt.Println(err)
		return
	}
	var pos [8]saturnmoons.XY
	saturnmoons.Positions(2451439.50074, earth, saturn, &pos)
	for i := range pos {
		fmt.Printf("%d:  %+7.3f  %+7.3f\n", i+1, pos[i].X, pos[i].Y)
	}
	// Output:
	// 1:   +3.102   -0.204
	// 2:   +3.823   +0.318
	// 3:   +4.027   -1.061
	// 4:   -5.365   -1.148
	// 5:   -0.972   -3.136
	// 6:  +14.568   +4.738
	// 7:  -18.001   -5.328
	// 8:  -48.760   +4.137
}
