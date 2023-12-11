package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	// Get the integers from os.Args and store them on a slice
	numeri := LeggiNumeri()

	// Now that the slice is populated, we can execute the requested operations:
	fmt.Printf("Minimo: %d\nMassimo: %d\nValore Medio: %.2f\n", Minimo(numeri), Massimo(numeri), Media(numeri))
}

func LeggiNumeri() []int {
	// Create a slice to hold the number values
	var numeri []int

	// Iterate through the CLI arguments and associate them with a correspongind slice position
	for _, c := range os.Args[1:] {
		// Before associating, the error checking makes sure that no values other than numbers are considered
		if n, err := strconv.Atoi(c); err == nil {
			numeri = append(numeri, n)
		}
	}
	return numeri
}

func Minimo(sl []int) int {
	var valoreMinimo int = sl[0]
	for i := 1; i < len(sl); i++ {
		if sl[i] < valoreMinimo {
			valoreMinimo = sl[i]
		}
	}
	return valoreMinimo
}

func Massimo(sl []int) int {
	var valoreMassimo int = sl[0]
	for i := 1; i < len(sl); i++ {
		if sl[i] > valoreMassimo {
			valoreMassimo = sl[i]
		}
	}
	return valoreMassimo
}

func Media(sl []int) float64 {
	var somma int
	for i := 0; i < len(sl); i++ {
		somma += sl[i]
	}
	return float64(somma) / float64(len(sl))
}
