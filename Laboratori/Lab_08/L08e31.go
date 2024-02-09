package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	// legga da riga di comando tre numeri interi positivi, rispettivamente N , DIVISORIMIN , e DIVISORIMAX ;
	N, DIVISORIMIN, DIVISORIMAX := LeggiNumeri()

	daStampare := make(map[int]int)

	// stampi a video tutte le coppie di interi positivi a e b , con a <= N e b <= N
	for a := 2; a <= N; a++ {
		for b := a; b <= N; b++ {
			// se a e b abbiano in comune un numero di divisori propri maggiore o uguale a DIVISORIMIN e minore o uguale a DIVISORIMAX ;
			divisoriComuni := DivisoriComuni(a, b)
			if len(divisoriComuni) <= DIVISORIMAX && len(divisoriComuni) >= DIVISORIMIN {
				// se almeno uno tra a e b sia un numero perfetto (uguale alla somma dei suoi divisori propri).
				if èPerfetto(a) || èPerfetto(b) {
					if daStampare[a] != b {
						fmt.Println(a, b)
						daStampare[a] = b
					}
				}
			}

		}
	}

}

// LeggiNumeri() restituice 3 interi positivi letti dalla riga di comando
func LeggiNumeri() (n, DIVISORIMIN, DIVISORIMAX int) {
	input := os.Args[1:]

	numeri := make([]int, 0)
	for _, el := range input {
		n, _ := strconv.Atoi(el)
		numeri = append(numeri, n)
	}
	return numeri[0], numeri[1], numeri[2]
}

// DivisoriPropri restituisce tutti i divisori propri di n (tutti i suoi divisori, tranne n)
func DivisoriPropri(n int) []int {
	divisori := make([]int, 0)
	for i := 1; i < n; i++ {
		if n%i == 0 {
			divisori = append(divisori, i)
		}
	}
	return divisori
}

// DivisoriComuni restituice tutti i divisori propri che a e b hanno in comune
func DivisoriComuni(a, b int) (divisoriComuni []int) {
	divisoriMap := make(map[int]bool)
	for _, div := range DivisoriPropri(a) {
		divisoriMap[div] = true
	}
	for _, div := range DivisoriPropri(b) {
		if _, found := divisoriMap[div]; found {
			divisoriComuni = append(divisoriComuni, div)
		}
	}
	return divisoriComuni
}

// èPerfetto restituice true se il numero dato è perfetto
func èPerfetto(n int) bool {
	sommaDivisori := 0
	for _, div := range DivisoriPropri(n) {
		sommaDivisori += div
	}
	risultato := (n == sommaDivisori)
	return risultato
}
