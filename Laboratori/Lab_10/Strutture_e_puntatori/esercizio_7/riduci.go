package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	numeratore, _ := strconv.Atoi(os.Args[1])
	denominatore, _ := strconv.Atoi(os.Args[2])

	frazione := NuovaFrazione(numeratore, denominatore)
	Riduci(frazione)
	fmt.Println(String(*frazione))
}

type Frazione struct {
	numeratore, denominatore int
}

func NuovaFrazione(numeratore, denominatore int) *Frazione {
	return &Frazione{numeratore, denominatore}
}

func String(f Frazione) string {
	return fmt.Sprintf("%d/%d", f.numeratore, f.denominatore)
}

func Riduci(frazione *Frazione) {
	mcd := 1
	for i := 1; i <= frazione.numeratore && i <= frazione.denominatore; i++ {
		if frazione.numeratore%i == 0 && frazione.denominatore%i == 0 {
			mcd = i
		}
	}
	frazione.numeratore /= mcd
	frazione.denominatore /= mcd
}
