package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// leggi da riga di comando una sequenza di valori positivi
	var numeri []int
	for i := 1; i < len(os.Args); i++ {
		n, _ := strconv.Atoi(os.Args[i])
		numeri = append(numeri, n)
	}

	// trova il minimo valore dispari nella sequenza
	minDispari := MinimoDispari(numeri)
	// stampa i valori interi pari superiori al minimo dispari presente nella sequeza letta
	for i := 0; i < len(numeri); i++ {
		if numeri[i]%2 != 0 {
			continue
		}
		if numeri[i] > minDispari {
			fmt.Print(numeri[i], " ")
		}
	}
	fmt.Println()

}

func MinimoDispari(sl []int) (minDispari int) {
	// restituisce il minimo valore dispari presente in sl

	// trova il primo numero dispari per usare come base
	for i := 0; i < len(sl); i++ {
		if sl[i]%2 != 0 {
			minDispari = sl[i]
			break
		}
	}

	// cerca un altro dispari minore del primo trovato
	for i := 0; i < len(sl); i++ {
		if sl[i]%2 == 0 {
			continue
		}
		if sl[i] < minDispari {
			minDispari = sl[i]
		}
	}
	return minDispari
}
