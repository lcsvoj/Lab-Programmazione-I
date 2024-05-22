// Legga da riga di comando un numero intero n e stampi a video la corrispondente tavola pitagorica n x n

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	n, _ := strconv.Atoi(os.Args[1])
	tavola := CreaTavolaPitagorica(n)
	StampaTavolaPitagorica(tavola)

}

func CreaTavolaPitagorica(n int) (tavola [][]int) {
	// Restituisce un valore in cui sono memorizzati i valori di una tavola pitagorica n x n

	tavola = make([][]int, n)
	for i := 0; i < n; i++ {
		tavola[i] = make([]int, n)
	}

	// Popula la tavola con i valori corrispondenti

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			tavola[i-1][j-1] = i * j
		}
	}
	return tavola
}

func StampaTavolaPitagorica(s [][]int) {
	// Stampa la tavola pitagorica corrispondete ai valori memorizzati s

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			fmt.Printf("%4d ", s[i][j])
		}
		fmt.Println()
	}
}
