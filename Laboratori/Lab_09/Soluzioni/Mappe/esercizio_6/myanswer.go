package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	// Prendi l'input
	testo := LeggiTesto()

	// Calcola le ocorrenze di ogni lettera del'input
	ocorrenze := Ocorrenze(testo)

	// Stampa l'istogramma risultante
	stampaIstogrammaOrdinato(ocorrenze)

}

func LeggiTesto() (testo string) {
	// Leggi da standard input un testo su più righe (che possono essere vuote) fino a EOF (CTRL+D)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		riga := scanner.Text()
		testo += riga
	}
	return testo
}

func Ocorrenze(s string) map[rune]int {
	// Restituice una mappa con le lettere e i suoi rispettivo numero di ocorrenze nella stringa fornita

	ocorrenze := make(map[rune]int)
	for _, c := range s {
		if unicode.IsLetter(c) {
			ocorrenze[c]++
		}
	}
	return ocorrenze
}

func ordinaSequenza(ocorrenze map[rune]int) (ordine []rune) {
	// Restituice una slice con la sequenza ordinata delle rune presenti nell'input

	ordine = make([]rune, len(ocorrenze))
	for key, _ := range ocorrenze {
		ordine = append(ordine, key)
	}
	for i := 1; i < len(ordine); i++ {
		if ordine[i] < ordine[i-1] {
			ordine[i-1], ordine[i] = ordine[i], ordine[i-1]
		}
	}

	return ordine
}

func stampaIstogrammaOrdinato(ocorrenze map[rune]int) {
	// Stampa l'istogramma con le ocorrenze delle lettere nel testo

	// Una slice guida per l'ordine di stampa:
	ordine := ordinaSequenza(ocorrenze)

	// Stampa l'istogramma ordinato
	fmt.Println("Istogramma:")
	for _, c := range ordine {
		fmt.Printf("%c: ", c)
		for i := 0; i < ocorrenze[c]; i++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
