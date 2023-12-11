package main

import (
	"fmt"
	"math"
)

func main() {

	var n float64
	var f int

	fmt.Print("Inserisci il valore da troncare: ")
	fmt.Scan(&n)
	fmt.Print("Inserisci il numero di cifre dopo la virgola: ")
	fmt.Scan(&f)

	nRound := math.Round(n*math.Pow(10, float64(f))) / math.Pow(10, float64(f))
	nTrunc := math.Trunc(n*math.Pow(10, float64(f))) / math.Pow(10, float64(f))

	fmt.Printf("Valore troncato = %.*f\nValore arrotondato = %.*f\n", f, nTrunc, f, nRound)

}
