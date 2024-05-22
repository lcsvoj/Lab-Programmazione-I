/* Scrivere un programma che legga da standard input le misure dell’altezza e della base di un rettangolo o del
raggio di un cerchio e ne calcoli il perimetro/circonferenza e l’area. La forma geometrica è determinata dal CLI
per "c" = cerchio e "r" = retangolo */

package main

import (
	"fmt"
	"math"
	"os"
)

func main() {

	var formaGeometrica string = os.Args[1]
	switch formaGeometrica {
	case "c":
		CalcolaCerchio()
	case "r":
		CalcolaRetangolo()
	}

}

func CalcolaRetangolo() {

	var altezza, base float64

	fmt.Print("Inserisci la altezza: ")
	fmt.Scan(&altezza)

	fmt.Print("Inserisci la base: ")
	fmt.Scan(&base)

	perimetro := base * altezza
	area := 2 * (base + altezza)

	fmt.Printf("Area = %.2f\nPerimetro = %.2f", area, perimetro)
}

func CalcolaCerchio() {

	var raggio float64

	fmt.Print("Inserisci il raggio: ")
	fmt.Scan(&raggio)

	circonferenza := 2 * math.Pi * raggio
	area := math.Pi * math.Pow(raggio, 2)

	fmt.Printf("Area = %.2f\nCirconferenza = %.2f", area, circonferenza)
}
