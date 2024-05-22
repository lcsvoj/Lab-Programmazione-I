/* Scrivere un programma che legga da standard input un numero reale n e stampi a video il valore
della radice quadrata di n solo se n >= 0 . Se n < 0 il programma deve stampare Il valore inserito non appartiene al dominio della funzione .

Oltre alla funzione main() , il programma deve definire e utilizzare le seguenti funzioni:
- RadiceQuadrata(numero float64) (float64, bool) che riceve in input un valore reale nel
parametro numero . Se numero >= 0 la funzione restituisce il valore della radice quadrata di
numero e un valore logico true come certificato della corretta esecuzione del calcolo.
Altrimenti, la funzione restituisce un valore reale 0 e un valore logico false , per segnalare che
non è stato possibile calcolare la radice quadrata di numero nel campo dei reali.
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var numero float64
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&numero)

	var radice float64
	var ok bool

	radice, ok = RadiceQuadrata(numero)
	if ok {
		fmt.Print("Radice quadrata: ", radice, "\n")
	} else {
		fmt.Println("Il valore inserito non appartiene al dominio della funzione")
	}
}

func RadiceQuadrata(numero float64) (float64, bool) {
	if numero >= 0 {
		radice := math.Sqrt(numero)
		return radice, true
	}
	return 0, false
}
