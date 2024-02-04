/* Legga da standard input un intero n oltre al numero reale. L'intero n specifica che il troncamento
e l'arrotondamento devono avvenire alla n -esima cifra decimale */

package main

import (
	"fmt"
	"math"
)

func main() {
	var n, cifre float64
	fmt.Print("Inserisci il valore: ")
	fmt.Scan(&n)
	fmt.Print("Inserisci il numero di cifre dopo la virgola: ")
	fmt.Scan(&cifre)

	fmt.Println("Valore troncato = ", Troncamento(n, cifre))
	fmt.Println("Valore arrotondato = ", Arrotondamento(n, cifre))
}

func Arrotondamento(numero, cifre float64) float64 {
	// 1. Moltiplicare il numero reale per 10^n
	numero *= math.Pow(10, cifre)
	// 2. Utilizzare la funzione math.Round del package math per arrontondare
	numero = math.Round(numero)
	// 3. Dividere il valore ottenuto 10^n
	numero /= math.Pow(10, cifre)

	return numero
}

func Troncamento(numero, cifre float64) float64 {
	// 1. Moltiplicare il numero reale per 10^n
	numero *= math.Pow(10, cifre)
	// 2. Utilizzare la funzione math.Trunc del package math per troncare
	numero = math.Trunc(numero)
	// 3. Dividere il valore ottenuto per 10^n
	numero /= math.Pow(10, cifre)

	return numero
}
