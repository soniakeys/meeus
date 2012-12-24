package meeus_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/soniakeys/meeus"
)

// For various numbers and symbols, test both Add and Combine.
// See that the functions do something, and that Strip returns
// the original number.
func TestStrip(t *testing.T) {
	var d string
	var sym rune
	t1 := func(fName string, f func(string, rune) string) {
		ad := f(d, sym)
		if ad == d {
			t.Fatalf("%s(%s, %c) had no effect", fName, d, sym)
		}
		if sd := meeus.DecSymStrip(ad, sym); sd != d {
			t.Fatalf("Strip(%s, %c) returned %s expected %s",
				ad, sym, sd, d)
		}
	}
	for _, d = range []string{"1.25", "1.", ".25"} {
		for _, sym = range []rune{'°', '"', 'h', 'ʰ'} {
			t1("DecSymAdd", meeus.DecSymAdd)
			t1("DecSymCombine", meeus.DecSymCombine)
		}
	}
}

func ExampleCombine() {
	formatted := "1.25"
	fmt.Println("Standard decimal symbol:", formatted)
	fmt.Println("Degree units, non combining decimal point: ",
		meeus.DecSymAdd(formatted, '°'))
	// Note that some software may not be capable of combining or even
	// rendering the combining dot.
	fmt.Println("Degree units, combining form of decimal point:",
		meeus.DecSymCombine(formatted, '°'))
	// Output:
	// Standard decimal symbol: 1.25
	// Degree units, non combining decimal point:  1°.25
	// Degree units, combining form of decimal point: 1°̣25
}

func TestFormatter(t *testing.T) {
	three := meeus.NewAngle(3)
	t.Log(three)
	t.Log(three.Overflow())
	t.Logf("v: %v", three)
	t.Log(three.Overflow())
	t.Logf("s: %s", three)
	t.Log(three.Overflow())
	t.Logf("d: %d", three)
	t.Log(three.Overflow())
	t.Logf("c: %c", three)
	t.Log(three.Overflow())
	t.Logf("x: %x", three)
	t.Log(three.Overflow())
	one := meeus.NewAngle(1)
	t.Log(180 / math.Pi)
	t.Logf(".3v: %.3v", one)
	t.Log(one.Overflow())
	t.Logf("0.3s: %0.3s", one)
	t.Log(one.Overflow())
	t.Logf("+0.3d: %+0.3d", one)
	t.Log(one.Overflow())
	t.Logf(" 0.3c: % 0.3c", one)
	t.Log(one.Overflow())
	t.Logf("0.5x: %0.5x", one)
	t.Log(one.Overflow())
	t.Logf(" 015.3c: [% 015.3c]", one)
	t.Log(one.Overflow())
	t.Logf(" 016.3c: [% 016.3c]", one)
	t.Log(one.Overflow())
	t.Logf(" 017.3c: [% 017.3c]", one)
	t.Log(one.Overflow())
	t.Logf("-017.3c: [%-017.3c]", one)
	t.Log(one.Overflow())
	t.Logf(".3c: [%.3c]", one)
	t.Log(one.Overflow())
	t.Logf(".9c: [%.9c]", one)
	t.Log(one.Overflow())
	t.Logf("19.10c: [%19.10c]", one)
	t.Log(one.Overflow())
	t.Logf(".11c: [%.11c]", one)
	t.Log(one.Overflow())
	t.Logf(".12c: [%.12c]", one)
	t.Log(one.Overflow())
	t.Logf(".13c: [%.13c]", one)
	t.Log(one.Overflow())
	t.Logf(".14c: [%.14c]", one)
	t.Log(one.Overflow())
	z := meeus.NewAngle(1e20)
	t.Log(z)
	t.Log(z.Overflow())
	z = meeus.NewAngle(math.NaN())
	t.Log(z)
	t.Log(z.Overflow())
	z = meeus.NewAngle(math.Inf(1))
	t.Log(z)
	t.Log(z.Overflow())
}
