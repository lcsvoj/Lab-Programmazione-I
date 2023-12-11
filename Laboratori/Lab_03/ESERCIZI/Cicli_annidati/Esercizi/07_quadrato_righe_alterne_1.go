/*
Scrivere un programma che legga da standard input un numero intero n e che stampi a video un quadrato di n righe costituite ciascuna da n simboli intervallati da spazi, alternando fra loro righe costituite solo da simboli * (asterisco) intervallati da spazi e righe costituite solo da simboli + (più) intervallati da spazi.
*/

package main

import "fmt"

func main() {

	var n int
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			for j := 0; j < n; j++ {
				fmt.Print("* ")
			}
		} else {
			for j := 0; j < n; j++ {
				fmt.Print("+ ")
			}
		}
		fmt.Print("\n")
	}
}
