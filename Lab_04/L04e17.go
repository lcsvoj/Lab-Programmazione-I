/* Scrivere un programma che legga da standard input un intero soglia>0 e stampi a video tutte le
terne pitagorighe tali che a<soglia , b<soglia e c<soglia . */

package main

import (
	"fmt"
	"math"
)

func main() {
	var soglia int
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&soglia)

	TernePitagoriche(soglia)
}

func ÈTernaPitagoriga(a int, b int, c int) bool {
	// restituisce true se c² è uguale a a² + b² e false altrimenti

	if a == 0 || b == 0 || c == 0 {
		return false
	}

	return math.Pow(float64(c), 2) == math.Pow(float64(a), 2)+math.Pow(float64(b), 2)
}

func TernePitagoriche(soglia int) {
	// stampa tutte le terne pitagoriche inferiori a soglia
	for a := 0; a < soglia; a++ {
		for b := 0; b < soglia; b++ {
			for c := 0; c < soglia; c++ {
				if ÈTernaPitagoriga(a, b, c) {
					fmt.Printf("(%d, %d, %d)\n", a, b, c)
				}
			}
		}
	}
}
