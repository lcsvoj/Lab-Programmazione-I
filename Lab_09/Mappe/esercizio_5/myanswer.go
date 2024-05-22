package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	testo := LeggiTesto()
	ocorrenze := Ocorrenze(testo)

	// Stampa l'istogramma con le ocorrenze delle lettere nel testo
	fmt.Println("Istogramma:")
	for key, value := range ocorrenze {
		fmt.Printf("%c: ", key)
		for i := 0; i < value; i++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
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
			ocorrenze[c] += 1
		}
	}
	return ocorrenze
}
