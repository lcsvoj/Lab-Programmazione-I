/* Scrivere un programma che legga da standard input un numero intero n e che stampi a video un
quadrato di n righe costituite ciascuna da n simboli intervallati da spazi, alternando fra loro 2
colonne costituite da simboli * (asterisco) a 2 colonne costituite da simboli + (più). In particolare,
se n è dispari solo una delle due colonne più a destra sarà stampata.
*/

package main

import (
	"fmt"
	"strings"
)

func main() {

	var n int
	fmt.Print("Inserisci un intero: ")
	fmt.Scan(&n)

	rigaSlc := make([]string, 2*n-1)
	for i := range rigaSlc {
		// La prima colonna è sempre "*"
		if i == 0 {
			rigaSlc[i] = "*"
			continue
		}

		// Le colonne di indice dispari sono sempre " "
		if i%2 != 0 {
			rigaSlc[i] = " "
		} else {
			// Nelle colonne multiple di 4 sempre si cambia simbolo
			if i >= 4 && i%4 == 0 {
				// Il simbolo è definito dalla precedente colonna di indice multiplo di 4
				if rigaSlc[i-4] == "*" {
					rigaSlc[i] = "+"
				} else {
					rigaSlc[i] = "*"
				}
			} else {
				// Nelle colonne che non sono multiplo di 4, invece si ripete il simbolo precedente
				rigaSlc[i] = rigaSlc[i-2]
			}
		}
	}

	rigaStr := strings.Join(rigaSlc, "")
	for i := 0; i < n; i++ {
		fmt.Println(rigaStr)
	}

}
