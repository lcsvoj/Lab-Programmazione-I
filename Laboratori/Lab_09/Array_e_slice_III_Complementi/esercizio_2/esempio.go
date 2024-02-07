package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var slc [][]int
	fmt.Printf("Tipo di slc: %T\n", slc)

	if slc == nil {
		fmt.Println("'nil' è lo zero-value di una variabile di tipo 'reference type' non inizializzata.")
	}
	fmt.Printf("slc: %v %d %d\n", slc, len(slc), cap(slc))
	fmt.Println()

	n := 4
	slc = CreaSliceBidimensionale(n)

	InizializzaSliceBidimensionale(slc)

	fmt.Println("slc:")
	for _, r := range slc {
		fmt.Printf("%v\n", r)
	}

	slc_soprasoglia := FiltraSliceBidimensionale(slc, 2)

	fmt.Println("slc_soprasoglia:")
	for _, r := range slc_soprasoglia {
		fmt.Printf("%v\n", r)
	}

}

/* 'CreaSliceBidimensionale' riceve in input un valore int nel parametro l e restituisce un valore s di tipo [][]int (una slice bi-dimensionale
di interi) con lunghezza/capacità pari a l in cui s[i], con 0 <= i < l, è un valore di tipo []int (una slice
di interi) con lunghezza/capacità pari a l
*/
func CreaSliceBidimensionale(l int) [][]int {
	var s [][]int
	/* s è una slice bi-dimensionale di tipo [][]int*/
	fmt.Printf("Tipo di s: %T\n", s)
	if s == nil {
		// 'nil' è lo zero-value di una variabile di tipo 'reference type' non inizializzata
		fmt.Printf("s == nil \t=>\t")
		fmt.Printf("s: %v %d %d\n", s, len(s), cap(s))
	}
	s = make([][]int, l)
	/* s è una slice bi-dimensionale di tipo [][]int con lunghezza/capacità pari a l */
	fmt.Printf("s: %v %d %d\n\n", s, len(s), cap(s))

	for i := 0; i < l; i++ {
		/* s[i], con 0 <= i < dim, è di tipo []int */
		fmt.Printf("Tipo di s[%d]: %T\n", i, s[i])
		if s[i] == nil {
			// 'nil' è lo zero-value di una variabile di tipo 'reference type' non inizializzata
			fmt.Printf("s[%d] == nil \t=>\t", i)
			fmt.Printf("s[%d]: %v %d %d\n", i, s[i], len(s[i]), cap(s[i]))
		}
		s[i] = make([]int, l)
		/* s[i] è una slice con lunghezza/capacità pari a l */
		fmt.Printf("s[%d]: %v %d %d\n\n", i, s[i], len(s[i]), cap(s[i]))
	}
	return s
}

/* 'InizializzaSliceBidimensionale' riceve in input un valore [][]int nel parametro s ed inizializza s[i][j]
(con 0 <= i < len(s) e 0 <= j <len(s[i])) con un valore intero estratto in maniera casuale nell'insieme {0,1}
 ... ossia ...
'InizializzaSliceBidimensionale' inizializza una slice bi-dimensionale di interi con valori interi estratti in maniera casuale
nell'insieme {0,1} */
func InizializzaSliceBidimensionale(s [][]int) {

	rand.Seed(int64(time.Now().Nanosecond()))

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			s[i][j] = rand.Intn(2)
		}
	}

}

/* 'FiltraSliceBidimensionale' riceve in input un valore [][]int ed un valore int nei parametri s e soglia, rispettivamente, e
restituisce un valore [][]int definito dai valori s[i] di s (con 0 <= i < len(s)) che includono
valori interi la cui somma è maggiore o uguale a soglia */
func FiltraSliceBidimensionale(s [][]int, soglia int) [][]int {

	var s_soprasoglia [][]int

	for i := 0; i < len(s); i++ {
		somma := 0
		for j := 0; j < len(s[i]); j++ {
			somma += s[i][j]
		}
		if somma >= soglia {
			s_soprasoglia = append(s_soprasoglia, s[i])
		}
	}

	return s_soprasoglia
}
