/*
Scrivere un programma che legga da standard input un numero intero n e che stampi a video un quadrato di n righe costituite ciascuna da n simboli intervallati da spazi, alternando fra loro 2 colonne costituite da simboli * (asterisco) a 2 colonne costituite da simboli + (più). In particolare, se n è dispari solo una delle due colonne più a destra sarà stampata.
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
		for j := 0; j < n; j++ {
			if (j%4 == 0) || (j%4 == 1) {
				fmt.Print("* ")
			} else {
				fmt.Print("+ ")
			}
		}
		fmt.Print("\n")
	}
}
