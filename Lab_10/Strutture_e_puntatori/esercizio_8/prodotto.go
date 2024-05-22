package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	frazioni := LeggiFrazioni()
	prodotto := MoltiplicaN(frazioni)
	Riduci(prodotto)

	fmt.Println("Prodotto:", String(*prodotto))
}

func LeggiFrazioni() (frazioni []Frazione) {
	fmt.Println("Inserisci numeratore e denominatore delle frazioni:")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		riga := scanner.Text()
		if riga == "" {
			break
		}
		var numeratore, denominatore int
		fmt.Sscan(riga, &numeratore, &denominatore)
		frazioni = append(frazioni, *NuovaFrazione(numeratore, denominatore))
	}
	return
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

func Moltiplica(f1, f2 Frazione) *Frazione {
	return NuovaFrazione(f1.numeratore*f2.numeratore, f1.denominatore*f2.denominatore)
}

func MoltiplicaN(frazioni []Frazione) (prodotto *Frazione) {
	if len(frazioni) > 0 {
		prodotto = &frazioni[0]
		for _, f := range frazioni[1:] {
			prodotto = Moltiplica(*prodotto, f)
		}
	}
	return
}
