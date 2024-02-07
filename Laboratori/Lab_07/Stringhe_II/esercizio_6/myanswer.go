package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	testo := LeggiTesto()

	fmt.Printf("Risultato:\n%s\n", TraduciTestoInFarfallino(testo))

}

func LeggiTesto() string {
	/* legge da standard input un testo su più righe e terminatodall'indicatore EOF,
	restituendo un valore string in cui è memorizzato il testo letto */
	fmt.Println("Inserisci un testo su più righe (termina con CTRL+D): ")
	scanner := bufio.NewScanner(os.Stdin)
	var result string

	for scanner.Scan() {
		line := scanner.Text()

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Errore nella lettura dell'input:", err)
		}

		result += line + "\n"
	}

	return result
}

func TraduciCarattereInFarfallino(c rune) string {
	/* riceve in input un valore rune nel parametro c e restituisce un valore string che
	rappresenta la traduzione in linguaggio farfallino di c */
	var carattereTrasformato string = string(c)
	if strings.ContainsRune("aeiouèéòóùúàáíì", c) {
		return (carattereTrasformato + "f" + carattereTrasformato)
	} else if strings.ContainsRune("AEIOUÀÁÈÉÌÍÓÒÚÙ", c) {
		return (carattereTrasformato + "F" + carattereTrasformato)
	} else {
		return carattereTrasformato
	}
}

func TraduciTestoInFarfallino(s string) string {
	/* riceve in input un valore string nel parametro s e restituisce un valore string che
	rappresenta la traduzione in linguaggio farfallino di s */
	var testoTrasformato string
	for _, c := range s {
		testoTrasformato += TraduciCarattereInFarfallino(c)
	}
	return testoTrasformato
}
