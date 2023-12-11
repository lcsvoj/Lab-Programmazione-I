/*
Scrivere un programma che legga da standard input un numero intero n e che stampi a video un quadrato di n righe costituite ciascuna da n simboli intervallati da spazi, alternando fra loro colonne costituite da simboli * (asterisco) a colonne costituite da simboli + (più).
*/

package main

import "fmt"

func main() {

	var n int
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j%2 == 0 {
				fmt.Print("* ")
			} else {
				fmt.Print("+ ")
			}
		}
	fmt.Print("\n")
	}
}

