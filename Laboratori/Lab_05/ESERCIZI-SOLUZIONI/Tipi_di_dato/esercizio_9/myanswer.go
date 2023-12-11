package main

import (
	"fmt"
	"math"
)

func main() {

	var x float64
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&x)

	radiceQuadrata := math.Sqrt(x)

	if radiceQuadrata*radiceQuadrata == x {
		fmt.Println(x, "uguale a", radiceQuadrata, "*", radiceQuadrata, "\n")
	} else {
		fmt.Println(x, "diverso da", radiceQuadrata, "*", radiceQuadrata, "\n")
	}
}
