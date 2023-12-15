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

func ordinaSequenza(ocorrenze map[rune]int) []rune {
	// Restituice una slice con la sequenza ordinata delle rune presenti nell'input

	// Create and populate a slice containing the runes of the input
	var ordine []rune
	for key, _ := range ocorrenze {
		fmt.Printf("\nLoop in 'ocorrenze', key: %c\n", key)
		fmt.Printf("Appending key to 'ordine' slice:\n")
		ordine = append(ordine, key)
		fmt.Printf("The slice is now: %v\n", ordine)
	}

	// Sort the runes inside of it
	fmt.Printf("\nStarting the loop in 'ordine', goes until i < %v\n", len(ordine)-1)
	for i := 0; i < len(ordine)-1; i++ {
		fmt.Printf("### index: %v ###\n", i)
		fmt.Printf("Checking if %v > %v...\n", ordine[i], ordine[i+1])
		if ordine[i] > ordine[i+1] {
			fmt.Printf("It is, so let's switch their places.\n")
			ordine[i+1], ordine[i] = ordine[i], ordine[i+1]
			fmt.Printf("The slice is now: %v\n", ordine)
		}
		fmt.Printf("It's not, so let's move on to the next index.\n")
	}

	fmt.Printf("\nReturning 'ordine' = %v\n", ordine)
	return ordine
}

func stampaIstogrammaOrdinato(ocorrenze map[rune]int) {
	// Stampa l'istogramma con le ocorrenze delle lettere nel testo

	// Una slice guida per l'ordine di stampa:
	ordine := ordinaSequenza(ocorrenze)

	// Stampa l'istogramma ordinato
	fmt.Println("Istogramma:")
	fmt.Printf("\nStarting the loop range in 'ordine'\n")
	for i, c := range ordine {
		fmt.Printf("### index: %v, rune: %c ###\n", i, c)
		fmt.Printf("%c: ", c)
		for i := 0; i < ocorrenze[c]; i++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
