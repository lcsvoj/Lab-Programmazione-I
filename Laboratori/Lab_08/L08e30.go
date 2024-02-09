package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	nStr := (os.Args[1])
	numeri := RimuoveCifre(nStr)

	sort.Ints(numeri)

	for _, n := range numeri {
		if ÈPrimo(n) {
			fmt.Println(n)
		}
	}
}

func RimuoveCifre(n string) []int {
	numeri := make([]int, 0)

	// Inserisci il numero stesso nell'elenco
	numero, _ := strconv.Atoi(n)
	numeri = append(numeri, numero)

	// In ogni iterazione: dal indice p1, si rimuove una sottostringa lunga p2
	for p2 := 1; p2 <= 3; p2++ {
		for p1 := 0; p1 <= len(n)-p2; p1++ {
			numero, _ := strconv.Atoi(n[:p1] + n[p1+p2:])
			numeri = append(numeri, numero)
		}
	}

	return numeri
}

func ÈPrimo(n int) bool {

	if n <= 1 {
		return false
	}

	for i := n - 1; i > 1; i-- {
		if n%i == 0 {
			return false
		}
	}

	ÈPrimo(n - 1)

	return true
}
