package angle_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/soniakeys/meeus/angle"
	"github.com/soniakeys/meeus/common"
	"github.com/soniakeys/meeus/julian"
)

func ExampleSep() {
	// Example 17.a, p. 110.
	r1 := common.NewRA(14, 15, 39.7).Rad()
	d1 := common.NewAngle(false, 19, 10, 57).Rad()
	r2 := common.NewRA(13, 25, 11.6).Rad()
	d2 := common.NewAngle(true, 11, 9, 41).Rad()
	d := angle.Sep(r1, d1, r2, d2)
	fmt.Println(common.NewFmtAngle(d))
	// Output:
	// 32°47′35″
}

// First exercise, p. 110.
func TestSep(t *testing.T) {
	r1 := common.NewRA(4, 35, 55.2).Rad()
	d1 := common.NewAngle(false, 16, 30, 33).Rad()
	r2 := common.NewRA(16, 29, 24).Rad()
	d2 := common.NewAngle(true, 26, 25, 55).Rad()
	d := angle.Sep(r1, d1, r2, d2)
	answer := common.NewAngle(false, 169, 58, 0).Rad()
	if math.Abs(d-answer) > 1e-4 {
		t.Fatal(common.NewFmtAngle(d))
	}
}

var (
	r1 = []float64{
		common.NewRA(10, 29, 44.27).Rad(),
		common.NewRA(10, 36, 19.63).Rad(),
		common.NewRA(10, 43, 01.75).Rad(),
	}
	d1 = []float64{
		common.NewAngle(false, 11, 02, 05.9).Rad(),
		common.NewAngle(false, 10, 29, 51.7).Rad(),
		common.NewAngle(false, 9, 55, 16.7).Rad(),
	}
	r2 = []float64{
		common.NewRA(10, 33, 29.64).Rad(),
		common.NewRA(10, 33, 57.97).Rad(),
		common.NewRA(10, 34, 26.22).Rad(),
	}
	d2 = []float64{
		common.NewAngle(false, 10, 40, 13.2).Rad(),
		common.NewAngle(false, 10, 37, 33.4).Rad(),
		common.NewAngle(false, 10, 34, 53.9).Rad(),
	}
	jd1 = julian.CalendarGregorianToJD(1978, 9, 13)
	jd3 = julian.CalendarGregorianToJD(1978, 9, 15)
)

// Second exercise, p. 110.
func TestMinSep(t *testing.T) {
	sep, err := angle.MinSep(jd1, jd3, r1, d1, r2, d2)
	if err != nil {
		t.Fatal(err)
	}
	answer := .5017 * math.Pi / 180 // on p. 111
	if math.Abs((sep-answer)/sep) > 1e-3 {
		t.Fatal(common.NewFmtAngle(sep))
	}
}

// "rectangular coordinate" solution, p. 113.
func TestMinSepRect(t *testing.T) {
	sep, err := angle.MinSepRect(jd1, jd3, r1, d1, r2, d2)
	if err != nil {
		t.Fatal(err)
	}
	answer := 224 * math.Pi / 180 / 3600 // on p. 111
	if math.Abs((sep-answer)/sep) > 1e-2 {
		t.Fatal(common.NewFmtAngle(sep))
	}

}

func TestSepHav(t *testing.T) {
	// Example 17.a, p. 110.
	r1 := common.NewRA(14, 15, 39.7).Rad()
	d1 := common.NewAngle(false, 19, 10, 57).Rad()
	r2 := common.NewRA(13, 25, 11.6).Rad()
	d2 := common.NewAngle(true, 11, 9, 41).Rad()
	d := angle.SepHav(r1, d1, r2, d2)
	s := fmt.Sprint(common.NewFmtAngle(d))
	if s != "32°47′35″" {
		t.Fatal(s)
	}
}

func ExampleSepPauwels() {
	// Example 17.b, p. 116.
	r1 := common.NewRA(14, 15, 39.7).Rad()
	d1 := common.NewAngle(false, 19, 10, 57).Rad()
	r2 := common.NewRA(13, 25, 11.6).Rad()
	d2 := common.NewAngle(true, 11, 9, 41).Rad()
	d := angle.SepPauwels(r1, d1, r2, d2)
	fmt.Println(common.NewFmtAngle(d))
	// Output:
	// 32°47′35″
}
