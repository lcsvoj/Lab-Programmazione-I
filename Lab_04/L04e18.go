package main

import (
	"fmt"
	"strings"
)

func main() {

	var gradini int
	fmt.Print("Inserisci il numero di gradini: ")
	fmt.Scan(&gradini)

	// Se n è negativo o nullo, stampare il messaggio "Dimensione non sufficiente"
	if gradini <= 0 {
		fmt.Println("Dimensione non sufficiente")
	}

	for i := 0; i < gradini; i++ {
		StampaGradino(i)
	}
}

func StampaGradino(gradino int) {
	// Stampa a video un singolo gradino della scala, opportunamente traslato verso destra

	// Se gradino < 0 non stampa nulla
	if gradino < 0 {
		return
	}

	// Se gradino = 0 non stampa la rientranza
	if gradino == 0 {
		fmt.Println("***")
		fmt.Println("  *")
	}

	// Se gradino > 0 stampa il gradino traslato di (gradino * 2) spazi verso destra
	if gradino > 0 {
		fmt.Print(strings.Repeat("  ", gradino))
		fmt.Println("***")
		fmt.Print(strings.Repeat("  ", gradino))
		fmt.Println("  *")
	}
}
