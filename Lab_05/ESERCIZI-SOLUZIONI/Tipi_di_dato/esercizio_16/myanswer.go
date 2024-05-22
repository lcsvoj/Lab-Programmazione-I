package main

import (
	"fmt"
	"math"
	"math/rand"
)

const EPSILON = 1.0e-6

func main() {

	var (
		s      int64
		xC, yC float64
		raggio float64
	)

	fmt.Print("\nInserisci s: ")
	fmt.Scan(&s)

	fmt.Print("Inserisci i valori dell'ascissa e dell'ordinata del punto C (il centro del cerchio): ")
	fmt.Scan(&xC, &yC)

	fmt.Print("Inserisci il valore del raggio:")
	fmt.Scan(&raggio)

	xMin, yMin := xC+raggio, yC+raggio // un punto sulla circonferenza
	xMax, yMax := xC+raggio, yC+raggio // un punto sulla circonferenza
	distanzaMin, distanzaMax := raggio, raggio

	rand.Seed(s)
	for i := 0; i < 1000000; i++ {
		x := xC - raggio + rand.Float64()*(2*raggio)
		y := yC - raggio + rand.Float64()*(2*raggio)

		distanza := Distanza(xC, yC, x, y)

		if ÈXUgualeAY(raggio, distanza) {
			fmt.Printf("\nIl punto (%f, %f) giace sulla circonferenza del cerchio.\n", x, y)
		}
		if ÈXMaggioreDiY(distanza, raggio) {
			if distanza > distanzaMax {
				xMax, yMax = x, y
				distanzaMax = distanza
			}
		}
		if ÈXMinoreDiY(distanza, raggio) {
			if distanza < distanzaMin {
				xMin, yMin = x, y
				distanzaMin = distanza
			}
		}
	}
	fmt.Printf("\nIl punto (%f, %f) è quello all'esterno del che ha distanza maxima del centro.\n", xMax, yMax)
	fmt.Printf("Distanza: %f\n", distanzaMax)

	fmt.Printf("\nIl punto (%f, %f) è quello all'interno del che ha distanza minima del centro.\n", xMin, yMin)
	fmt.Printf("Distanza: %f\n", distanzaMin)
}

func Distanza(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2))
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
