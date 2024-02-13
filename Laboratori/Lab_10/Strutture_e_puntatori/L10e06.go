package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Frazione struct {
	numeratore   int
	denominatore int
}

func main() {
	frazioni := LeggiFrazioni()
	moltiplicazione := MoltiplicaN(frazioni)
	Riduci(moltiplicazione)
	fmt.Println("Prodotto:", String(*moltiplicazione))
}

func LeggiFrazioni() []Frazione {
	fmt.Println("Inserisci numeratore e denominatore delle frazioni:")
	scanner := bufio.NewScanner(os.Stdin)
	frazioni := make([]Frazione, 0)
	for scanner.Scan() {
		riga := scanner.Text()
		if riga == "" {
			break
		}
		numeri := strings.Split(riga, " ")
		numeratore, _ := strconv.Atoi(numeri[0])
		denominatore, _ := strconv.Atoi(numeri[1])

		frazioni = append(frazioni, *NuovaFrazione(numeratore, denominatore))
	}
	return frazioni
}

func NuovaFrazione(numeratore, denominatore int) *Frazione {
	f := Frazione{numeratore, denominatore}
	return &f
}

func String(f Frazione) string {
	return fmt.Sprintf("%d/%d", f.numeratore, f.denominatore)
}

func Moltiplica(f1, f2 Frazione) *Frazione {
	f := Frazione{
		numeratore:   f1.numeratore * f2.numeratore,
		denominatore: f1.denominatore * f2.denominatore,
	}
	return &f
}

func MoltiplicaN(fN []Frazione) *Frazione {
	risultato := Frazione{1, 1}
	for _, f := range fN {
		risultato = *Moltiplica(risultato, f)
	}
	return &risultato
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
