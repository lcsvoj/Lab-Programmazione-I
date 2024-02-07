package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	// legga da riga di comando un numero intero n ;
	n := LeggiNumero()

	// inizializza una variabile s di tipo [][]int con lunghezza/capacità pari a
	// n in cui ogni elemento s[i] , con 0 <= i < l , è un valore di tipo []int
	// con lunghezza/capacità pari a n ;
	s := CreaSliceBidimensionale(n)
	InizializzaSliceBidimensionale(s)

	// stampi a video tutte le coppie di indici (i, j) , con 0 <= i < l e
	// 0 <= j < l , tali che s[i][j] è uguale a 1
	for _, el := range Coppie(s) {
		fmt.Println(el)
	}

}

func CreaSliceBidimensionale(l int) [][]int {
	/* Riceve in input un valore int nel parametro l e restituisce un valore s
	di tipo [][]int (una slice bi-dimensionale di interi) con lunghezza/capacità
	pari a l in cui s[i], con 0 <= i < l, è un valore di tipo []int (una slice
	di interi) con lunghezza/capacità pari a l */

	var s [][]int
	/* s è una slice bi-dimensionale di tipo [][]int */

	s = make([][]int, l)

	for i := 0; i < l; i++ {
		s[i] = make([]int, l)
	}

	return s
}

func InizializzaSliceBidimensionale(s [][]int) {
	/* 'InizializzaSliceBidimensionale' inizializza una slice bi-dimensionale di
	interi con valori interi estratti in maniera casuale nell'insieme {0,1}	*/

	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			s[i][j] = rand.Intn(2)
		}
	}
}

func Coppie(s [][]int) (coppie [][]int) {
	/* Riceve in input un valore [][]int nel parametro s e restituisce il valore
	di tipo [][]int nella variabile coppie in cui sono memorizzate tutte le
	coppie di indici (i, j) , con 0 <= i < l e 0 <= j < l , tali che s[i][j]
	è uguale a 1 ( coppie[k] , con 0 <= k < len(coppie) , è un valore di tipo
	[]int di lunghezza 2 ).	*/

	coppie = make([][]int, 0)
	for i, el := range s {
		for j, el_j := range el {
			if el_j == 1 {
				coppie = append(coppie, []int{i, j})
			}
		}
	}
	return coppie
}

func LeggiNumero() (n int) {
	n, _ = strconv.Atoi(os.Args[1])
	return n
}
