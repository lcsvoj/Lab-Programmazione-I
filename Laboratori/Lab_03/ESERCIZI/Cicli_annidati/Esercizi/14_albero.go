/* Scrivere un programma che legga da standard input un intero n e,
stampi a video un albero utilizzando il carattere * (asterisco) per
rappresentarne la chioma ed il carattere + (più) per rappresentarne
il tronco:
- La chioma dell'albero deve essere alta n righe e, nel punto di
larghezza massima, larga 2*n-1 colonne.
- Il tronco dell'albero deve essere rappresentato con un quadrato di
altezza e larghezza pari a 3 caratteri.
- Il tronco dell'albero deve essere centrato rispetto alla chioma dell'albero.
- Se n<=0 , il programma stampa solo il tronco dell'albero.*/

package main

import (
	"fmt"
	"strings"
)

func main() {

	var n int
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)

	stampaChioma(n)
	stampaTronco(n)
}

func stampaTronco(n int) {

	// Per n < 3: formatto minimo (senza spazi prima)
	if n < 3 {
		for i := n; i < n+3; i++ {
			for j := -1; j <= 1; j++ {
				fmt.Print("+")
			}
			fmt.Print("\n")
		}
	} else {
		for i := n; i < n+3; i++ {
			for j := -n + 1; j < n; j++ {
				// Aligna il tronco e il chioma
				if j > 1 || j < -1 {
					fmt.Print(" ")
				} else {
					// Stampa il tronco
					fmt.Print("+")
				}
			}
			fmt.Print("\n")
		}
	}
}

func stampaChioma(n int) {
	// Per n == 1: formatto minimo
	if n == 1 {
		fmt.Print(" * \n")

	} else {
		for i := 0; i < n; i++ {
			// Aligna l'apice del chioma
			fmt.Print(strings.Repeat(" ", n-i-1))
			// Stampa il chioma
			fmt.Print(strings.Repeat("*", 2*i+1))
			fmt.Print("\n")
		}
	}
}
