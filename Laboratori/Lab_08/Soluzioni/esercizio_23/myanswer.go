package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	soglia, _ := strconv.Atoi(os.Args[1])
	sequenza := Genera(soglia)
	fmt.Printf("Valori generati %v\nValori sopra soglia %v\n", sequenza, sequenza[:len(sequenza)-1])
}

func Genera(soglia int) (sequenza []int) {

	// generazione di un numero casuale compreso nell'intervallo che va da 0 a 99 (estremi inclusi)
	for {
		numeroGenerato := rand.Intn(100)
		sequenza = append(sequenza, numeroGenerato)
		if numeroGenerato < soglia {
			break
		}
	}
	return sequenza
}
