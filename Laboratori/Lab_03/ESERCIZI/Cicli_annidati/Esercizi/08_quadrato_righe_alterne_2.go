/*
Scrivere un programma che legga da standard input un numero intero n e che stampi a video un quadrato di n righe costituite ciascuna da n simboli intervallati da spazi, alternando fra loro righe costituite solo da simboli * (asterisco) intervallati da spazi, righe costituite solo da simboli + (più) intervallati da spazi e righe costituite solo da simboli o (lettera o) intervallati da spazi.
*/

package main

import (
	"fmt"
)

func main() {

	var n int
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		if i%3 == 2 {
			for j := 0; j < n; j++ {
				fmt.Print("* ")
			}
		} else if i%3 == 1 {
			for j := 0; j < n; j++ {
				fmt.Print("+ ")
			}
		} else if i%3 == 0 {
			for j := 0; j < n; j++ {
				fmt.Print("o ")
			}
		}
		fmt.Print("\n")
	}
}
