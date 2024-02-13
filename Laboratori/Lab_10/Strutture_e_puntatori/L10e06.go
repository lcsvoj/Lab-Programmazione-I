package main

import (
	"fmt"
	"os"
	"strconv"
)

type Frazione struct {
	numeratore   int
	denominatore int
}

func main() {
	n, d := LeggiNumeri()
	f := NuovaFrazione(n, d)
	Riduci(f)
	fmt.Println(String(*f))
}

func LeggiNumeri() (n, d int) {
	input := os.Args[1:]
	n, _ = strconv.Atoi(input[0])
	d, _ = strconv.Atoi(input[1])
	return n, d
}

func NuovaFrazione(numeratore, denominatore int) *Frazione {
	f := Frazione{numeratore, denominatore}
	return &f
}

func String(f Frazione) string {
	return fmt.Sprintf("%d/%d", f.numeratore, f.denominatore)
}

func Riduci(f *Frazione) {

	if f.denominatore == 1 || f.numeratore == 1 {
		return
	}

	var maggiore int
	if f.numeratore > f.denominatore {
		maggiore = f.numeratore
	} else {
		maggiore = f.denominatore
	}
	for i := maggiore; i > 1; i-- {
		if f.numeratore%i == 0 && f.denominatore%i == 0 {
			f.numeratore /= i
			f.denominatore /= i
			Riduci(f)
		}
	}
}
