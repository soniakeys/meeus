package jupitermoons_test

import (
	"fmt"

	"github.com/soniakeys/meeus/jupitermoons"
	pp "github.com/soniakeys/meeus/planetposition"
)

func ExamplePositions() {
	p1, p2, p3, p4 := jupitermoons.Positions(2448972.50068)
	fmt.Printf("X1 = %+.2f  Y1 = %+.2f\n", p1.X, p1.Y)
	fmt.Printf("X2 = %+.2f  Y2 = %+.2f\n", p2.X, p2.Y)
	fmt.Printf("X3 = %+.2f  Y3 = %+.2f\n", p3.X, p3.Y)
	fmt.Printf("X4 = %+.2f  Y4 = %+.2f\n", p4.X, p4.Y)
	// Output:
	// X1 = -3.44  Y1 = +0.21
	// X2 = +7.44  Y2 = +0.25
	// X3 = +1.24  Y3 = +0.65
	// X4 = +7.08  Y4 = +1.10
}

func ExampleE5() {
	e, err := pp.LoadPlanet(pp.Earth, "")
	if err != nil {
		fmt.Println(err)
		return
	}
	j, err := pp.LoadPlanet(pp.Jupiter, "")
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
