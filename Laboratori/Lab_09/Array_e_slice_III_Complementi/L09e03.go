package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	n, _ := strconv.Atoi(os.Args[1])
	StampaTavolaPitagorica(CreaTavolaPitagorica(n))
}

func CreaTavolaPitagorica(n int) [][]int {
	// Inizializza la slice bidimensionale
	tavola := make([][]int, n)
	for i := range tavola {
		tavola[i] = make([]int, n)
	}

	// Riempila con i valori
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			tavola[i][j] = (i + 1) * (j + 1)
		}
	}
	return tavola
}

func StampaTavolaPitagorica(s [][]int) {
	for i := range s {
		for j := range s[i] {
			fmt.Printf("%5d", s[i][j])
		}
		fmt.Println()
	}
}
