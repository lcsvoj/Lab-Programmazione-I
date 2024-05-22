package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// get the sequence from CLI input
	sequenza := make([]int, len(os.Args)-1)
	for i, c := range os.Args[1:] {
		sequenza[i], _ = strconv.Atoi(c)
	}
	fmt.Println("La somma è", Calcola(sequenza))
}

func Calcola(sl []int) (somma int) {
	/* riceve in input un valore []int nel parametro sl e restituisce un valore di tipo int pari alla
	somma dei prodotti pari associati alle coppie non ordinate di numeri che si possono definire a
	partire dai numeri presenti in sl */

	// facciamo la lista delle coppie non ordinate possibile a partire da sl
	for i1, n1 := range sl {
		for i2 := i1 + 1; i2 < len(sl); i2++ {
			n2 := sl[i2]
			// if the product of the two numbers of the pair is even, add it to the sum
			prodotto := n1 * n2
			if prodotto%2 == 0 {
				somma += prodotto
			}
		}
	}
	return somma
}
