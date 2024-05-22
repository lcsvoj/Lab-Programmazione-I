/* Scrivere un programma che legga da standard input un numero reale `raggio` e stampi i valori di
circonferenza e area di un cerchio di raggio `raggio` .
Oltre alla funzione main() , il programma deve definire e utilizzare le seguenti funzioni:
- CalcolaArea(raggio float64) float64 che riceve in input il valore reale del raggio di un cerchio
nel parametro raggio e restituisce il valore dell'area del cerchio associato;
- CalcolaCirconferenza(raggio float64) float64 che riceve in input il valore reale del raggio di
un cerchio nel parametro raggio e restituisce il valore della circonferenza del cerchio associato.
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var raggio, circonferenza float64
	fmt.Print("Inserisci il raggio del cerchio: ")
	fmt.Scan(&raggio)

	area := CalcolaArea(raggio)
	circonferenza = CalcolaCirconferenza(raggio)

	fmt.Print("Area del cerchio: ", area, "\n")
	fmt.Print("Circonferenza del cerchio: ", circonferenza, "\n")
}

func CalcolaArea(raggio float64) float64 {
	area := math.Pi * raggio * raggio
	return area
}

func CalcolaCirconferenza(raggio float64) float64 {
	circonferenza := 2 * raggio * math.Pi
	return circonferenza
}
