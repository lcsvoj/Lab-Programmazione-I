/*
Scrivere un programma che legga da standard input due numeri interi che chiameremo n e l e calcoli
l’area di un poligono regolare con n lati di lunghezza l 
*/

package main

import "fmt"
import "math"

func main() {
	
	var lati, lunghezza float64
	fmt.Print("Inserisci il numero di lati del poligono: ")
	fmt.Scan(&lati)
	fmt.Print("Inserisci la lunghezza dei lati del poligono: ")
	fmt.Scan(&lunghezza)

	var area float64 = (lati * math.Pow(lunghezza, 2)) / (4 * math.Tan(math.Pi/lati))
	fmt.Println("Area calcolata =", area)

}
