/*
Scrivere un programma che legga da standard input le misure dell’altezza e della base di un rettangolo e ne calcoli il perimetro e l’area
*/

package main

import "fmt"

func main() {

	var base, altezza int
	fmt.Print("Inserisci la base: ")
	fmt.Scan(&base)
	fmt.Print("Inserisci l'altezza: ")
	fmt.Scan(&altezza)

	var perimetro, area int
	area = base * altezza
	perimetro = 2*(base + altezza)

	fmt.Print("Perimetro = ", perimetro, "\n", "Area = ", area, "\n")
}
