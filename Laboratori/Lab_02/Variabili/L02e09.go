/* Scrivere un programma che legga da standard input le età di due persone (espresse in anni) e calcoli:
la somma delle età inserite;
la media delle età inserite;
la media delle età inserite arrotondata per difetto all'intero inferiore;
la media delle età inserite arrotondata per eccesso all'intero superiore;
la somma e la media delle età che le due persone avranno fra 10 anni. */

package main

import (
	"fmt"
	"math"
)

func main() {

	var età1, età2 float64
	fmt.Print("Età persona 1 = ")
	fmt.Scan(&età1)
	fmt.Print("Età persona 2 = ")
	fmt.Scan(&età2)

	somma := età1 + età2
	media := (età1 + età2) / 2

	fmt.Println("Somma età =", somma)
	fmt.Println("Media età =", media)
	fmt.Println("Media età arrotondata per difetto all'intero inferiore =", math.Floor(media))
	fmt.Println("Media età arrotondata per eccesso all'intero superiore =", math.Ceil(media))

	età1 += 10
	età2 += 10
	somma = età1 + età2
	media = (età1 + età2) / 2

	fmt.Println("Somma età =", somma)
	fmt.Println("Media età =", media)

}
