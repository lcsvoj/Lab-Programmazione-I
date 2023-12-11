package main

import (
	"fmt"
	"math"
	"math/rand"
)

const EPSILON = 1.0e-9

func main() {
	var (
		s      int
		m1, q1 float64
		m2, q2 float64
		m3, q3 float64
	)

	fmt.Print("Inserisci s: ")
	fmt.Scan(&s)

	fmt.Print("Inserisci m1 e q1: ")
	fmt.Scan(&m1, &q1)
	fmt.Print("Inserisci m2 e q2: ")
	fmt.Scan(&m2, &q2)
	fmt.Print("Inserisci m3 e q3: ")
	fmt.Scan(&m3, &q3)

	rand.Seed(int64(s))
	for i := 0; i < 10; i++ {
		x := rand.Float64() * 10
		y := rand.Float64() * 10

		fmt.Printf("\nPunto (%f, %f) - ", x, y)

		// Il punto sta all'esterno se sta sotto la base || se sta sopra la base ma sta sotto gli altri lati
		if ÈXMinoreDiY(y, m1*x+q1) || (ÈXMaggioreDiY(y, m1*x+q1) && (ÈXMaggioreDiY(y, m2*x+q2) || ÈXMaggioreDiY(y, m3*x+q3))) {
			fmt.Printf("Il punto sta all'esterno del triangolo.")
		} else {
			fmt.Printf("Il punto sta all'interno del triangolo.")
		}
	}
	fmt.Println("\n")
}

/*
	La funzione `func ÈXMaggioreDiY(x, y float64) bool` riceve in input due

valore reali nei parametri `x` e `y` e restituisce `true` se `x` è maggiore
di `y` a meno di una costante EPSILON
*/
func ÈXMaggioreDiY(x, y float64) bool {
	return (x - y) > EPSILON
}

/*
	La funzione `func ÈXUgualeAY(x, y float64) bool` riceve in input due

valore reali nei parametri `x` e `y` e restituisce `true` se `x` è uguale
di `y` a meno di una costante EPSILON
*/
func ÈXUgualeAY(x, y float64) bool {
	return math.Abs(x-y) <= EPSILON
}

/*
	La funzione `func ÈXMaggioreDiY(x, y float64) bool` riceve in input due

valore reali nei parametri `x` e `y` e restituisce `true` se `x` è minore
di `y` a meno di una costante EPSILON
*/
func ÈXMinoreDiY(x, y float64) bool {
	return (x - y) < -EPSILON
}
