package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := LeggiNumeri()
	fmt.Println("La somma è", Calcola(input))
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

	coppie := make([][2]int, 0)
	TrovaCoppie(sl, &coppie)

	somma := 0
	for _, coppia := range coppie {
		prodottoCoppia := coppia[0] * coppia[1]
		if prodottoCoppia%2 == 0 {
			somma += prodottoCoppia
		}
	}
	return somma
}

func TrovaCoppie(sl []int, coppie *[][2]int) {
	// Caso base
	if len(sl) == 1 {
		return
	}

	// Caso ricorsivo
	n1 := sl[0]
	for _, n2 := range sl[1:] {
		*coppie = append(*coppie, [2]int{n1, n2})
	}
	TrovaCoppie(sl[1:], coppie)
	return
}
