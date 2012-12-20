package decsym_test

import (
	"fmt"
	"testing"

	"github.com/soniakeys/meeus/decsym"
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
		if sd := decsym.Strip(ad, sym); sd != d {
			t.Fatalf("Strip(%s, %c) returned %s expected %s",
				ad, sym, sd, d)
		}
	}
	for _, d = range []string{"1.25", "1.", ".25"} {
		for _, sym = range []rune{'°', '"', 'h', 'ʰ'} {
			t1("Add", decsym.Add)
			t1("Combine", decsym.Combine)
		}
	}
}

func ExampleCombine() {
	formatted := "1.25"
	fmt.Println("Standard decimal symbol:", formatted)
	fmt.Println("Degree units, non combining decimal point: ",
		decsym.Add(formatted, '°'))
	// Note that some software may not be capable of combining or even
	// redering the combining dot.
	fmt.Println("Degree units, combining form of decimal point:",
		decsym.Combine(formatted, '°'))
	// Output:
	// Standard decimal symbol: 1.25
	// Degree units, non combining decimal point:  1°.25
	// Degree units, combining form of decimal point: 1°̣25
}
