/*
Scrivere un programma che legga da standard input le età di due persone (espresse in anni) e calcoli:
- la somma delle età inserite;
- la media delle età inserite;
- la media delle età inserite arrotondata per difetto all'intero inferiore;
- la media delle età inserite arrotondata per eccesso all'intero superiore;
- la somma e la media delle età che le due persone avranno fra 10 anni.
*/

package main

import "fmt"
import "math"

func main() {

	var eta_1, eta_2 float64
	fmt.Print("Età persona 1 = ")
	fmt.Scan(&eta_1)
	fmt.Print("Età persona 2 = ")
	fmt.Scan(&eta_2)
	
	var somma, media, media_difetto, media_eccesso, somma_fra_10_anni, media_fra_10_anni float64
	somma = eta_1 + eta_2
	somma_fra_10_anni = somma + 20
	media = (eta_1 + eta_2)/2
	media_difetto = math.Floor(media)
	media_eccesso = math.Ceil(media)
	media_fra_10_anni = somma_fra_10_anni/2
	
	fmt.Println("Somma età =", somma)
	fmt.Println("Media età =", media)
	fmt.Println("Media età arrotondata per difetto all'intero inferiore =", media_difetto)
	fmt.Println("Media età arrotondata per eccesso all'intero superiore =", media_eccesso)
	fmt.Println("Somma età a 10 anni =", somma_fra_10_anni)
	fmt.Println("Media età a 10 anni =", media_fra_10_anni)
}
