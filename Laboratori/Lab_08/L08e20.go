package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	numeri := LeggiNumeri()
	fmt.Println("La somma è", Calcola(numeri))
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

func Calcola(sl []int) int {
	somma := 0
	for i, n := range sl {
		if i%2 != 0 {
			somma += n * sl[i-1]
		}
	}
	return somma
}
