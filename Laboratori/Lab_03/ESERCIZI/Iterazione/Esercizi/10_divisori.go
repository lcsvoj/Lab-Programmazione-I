/*
Scrivere un programma che legga da standard input un numero intero n e stampi a video i divisori
propri del numero letto, ovvero tutti i suoi divisori escluso il numero stesso. Ad esempio, i divisori del numero 12 sono: 1, 2, 3, 4, 6, 12; quindi i divisori propri di 12 sono: 1, 2, 3, 4, 6.
*/

package main

import (
	"fmt"
	"strconv"
)
func main() {
	
	fmt.Print("Inserisci un numero: ")
	var n int
	fmt.Scan(&n)

	var divisori string
	divisori = ""
	for i := 1; i < n; i++ {
		if n%i == 0 {
			divisori += strconv.Itoa(i) + " "
		}
	}

	fmt.Print("Divisori di ", n, ": ", divisori, "\n")

}

