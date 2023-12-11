/*
Scrivere un programma che legga da standard input una sequenza di numeri reali strettamente
positivi (un numero è strettamente positivo se è maggiore di 0; se un numero è minore o uguale 0
non è strettamente positivo). La lettura termina quando viene letto un numero minore o uguale a 0.
Il programma deve stampare a video il risultato della media aritmetica dei valori inseriti.
*/

package main

import (
	"fmt"
)

func main() {

	var n float64
	var somma, media, numeroNumeri = 0.0, 0.0, 0.0

	fmt.Print("Inserisci una sequenza di numeri (interrompi con numero <= 0): ")
	for {

		fmt.Scan(&n)

		if n <= 0 {
			break
		}

		somma += n
		numeroNumeri += 1

	}

	media = somma / numeroNumeri

	fmt.Println("Media aritmetica =", media)

}
