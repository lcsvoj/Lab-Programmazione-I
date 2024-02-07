package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	numeri := LeggiNumeri()
	sort.Ints(numeri)
	fmt.Println("Minimo:", numeri[0])
	fmt.Println("Massimo:", numeri[len(numeri)-1])
	fmt.Printf("Valore medio: %.2f\n", Media(numeri))
}

func LeggiNumeri() (numeri []int) {
	input := os.Args[1:]
	for _, el := range input {
		n, err := strconv.Atoi(el)
		if err == nil {
			numeri = append(numeri, n)
		}
	}
	return numeri
}

// func Minimo(sl []int) int {}

// func Massimo(sl []int) int {}

func Media(sl []int) float64 {
	var somma int
	for _, n := range sl {
		somma += n
	}
	return float64(somma) / float64(len(sl))
}
