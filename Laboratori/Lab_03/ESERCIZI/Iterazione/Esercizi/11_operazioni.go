/*
Create un programma che legge da standard input due numeri interi, che chiameremo x e y . Letti i due numeri, il programma stampa:
- il maggiore tra x e y
- il minore tra x e y
- il risultato della somma tra x e y
- il risultato della differenza tra il maggiore e il minore
- il risultato della divisione x/y
- il risultato del prodotto x*y
- il valore medio tra x e y
- il risultato di x elevato alla y (utilizzando sia un ciclo for sia la funzione math.Pow )
*/

package main

import (
	"fmt"
	"math"
)

func main() {

	var x, y float64
	fmt.Print("Inserisci due numeri interi: ")
	fmt.Scan(&x, &y)
	var maggiore, minore float64

	if x > y {
		maggiore = x
		minore = y
	} else {
		maggiore = y
		minore = x
	}

	fmt.Println("Maggiore:", maggiore)
	fmt.Println("Minore:", minore)
	fmt.Println("Somma:", x+ y)
	fmt.Println("Differenza:", maggiore - minore)
	fmt.Println("Prodotto:", x * y)
	fmt.Println("Divisione:", x / y)
	fmt.Println("Valore medio:", ((x+y)/2) )
	fmt.Println("Potenza:",  math.Pow(x, y))

}

