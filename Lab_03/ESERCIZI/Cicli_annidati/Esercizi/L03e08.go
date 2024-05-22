/* Scrivere un programma che legga da standard input un numero intero n e che, come mostrato nell'Esempio di
esecuzione, stampi a video un quadrato di n righe costituite ciascuna da n simboli intervallati da spazi,
alternando fra loro righe costituite solo da simboli * (asterisco) intervallati da spazi, righe costituite
solo da simboli + (più) intervallati da spazi e righe costituite solo da simboli o (lettera o) intervallati da
spazi. */

package main

import (
	"fmt"
)

func main() {

	var n int
	for n <= 0 {
		fmt.Print("Inserisci un numero: ")
		fmt.Scan(&n)
	}

	symbols := []string{"*", "+", "o"}
	for i := 0; i < n; i++ {
		for j := 0; j < 2*n; j++ {
			if j%2 == 0 {
				fmt.Print(symbols[i%3])
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

}
