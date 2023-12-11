package main

import (
	"fmt"
	"math/rand"
)

const EPSILON = 1.0e-9

func main() {

	var s int
	var m, q float64
	fmt.Print("Inserisci s: ")
	fmt.Scan(&s)
	fmt.Print("Inserisci m e q: ")
	fmt.Scan(&m, &q)

	/* inizializzazione del generatore di numeri casuali */
	rand.Seed(int64(s))

	for n := 0; n < 10; n++ {
		x := rand.Float64()
		y := rand.Float64()

		var posizione string
		if ÈXMaggioreDiY(y, m*x+q) {
			posizione = "sopra"
		} else if ÈXMinoreDiY(y, m*x+q) {
			posizione = "sotto"
		}

		fmt.Printf("Punto (%f, %f) - Il punto sta %s la retta\n", x, y, posizione)
	}
}

/* La funzione `func ÈXMaggioreDiY(x, y float64) bool` riceve in input due
valore reali nei parametri `x` e `y` e restituisce `true` se `x` è maggiore
di `y` a meno di una costante EPSILON */
/* ÈXMaggioreDiY(5.0,4.999) -> true */
/* ÈXMaggioreDiY(5.0,4.9999999999) -> false */
func ÈXMaggioreDiY(x, y float64) bool {
	return (x - y) > EPSILON
}

/* La funzione `func ÈXMaggioreDiY(x, y float64) bool` riceve in input due
valore reali nei parametri `x` e `y` e restituisce `true` se `x` è minore
di `y` a meno di una costante EPSILON */
/* ÈXMinoreDiY(4.999,5.0) -> true */
/* ÈXMinoreDiY(4.9999999999,5.0) -> false */
func ÈXMinoreDiY(x, y float64) bool {
	return (x - y) < -EPSILON
}
