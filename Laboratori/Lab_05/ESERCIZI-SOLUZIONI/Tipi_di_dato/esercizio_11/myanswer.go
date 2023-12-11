package main

import (
	"fmt"
	"math"
)

func main() {

	var n float64

	fmt.Print("Inserisci il valore da troncare: ")
	fmt.Scan(&n)

	n = (math.Trunc(n*100) / 100)

	fmt.Printf("Valore troncato = %.2f\n", n)

}
