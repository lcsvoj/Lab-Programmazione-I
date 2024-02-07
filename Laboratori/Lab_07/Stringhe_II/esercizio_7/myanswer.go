package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	testo := LeggiTesto()

	fmt.Printf("Risultato:\n%s\n", Garibaldi(testo))

}

func LeggiTesto() string {
	/* legge da standard input un testo su più righe e terminatodall'indicatore EOF,
	restituendo un valore string in cui è memorizzato il testo letto */
	fmt.Println("Inserisci un testo su più righe (termina con CTRL+D): ")
	scanner := bufio.NewScanner(os.Stdin)
	var result string

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Errore nella lettura dell'input:", err)
		}

		result += line + "\n"
	}

	return result
}

func TrasformaCarattere(c rune) rune {
	/* riceve in input un valore rune nel parametro c e restituisce un valore string che
	rappresenta la traduzione in linguaggio farfallino di c */
	if strings.ContainsRune("aeiouèéòóùúàáíì", c) || strings.ContainsRune("AEIOUÀÁÈÉÌÍÓÒÚÙ", c) {
		return 'u'
	} else {
		return c
	}
}

func Garibaldi(s string) string {
	/* riceve in input un valore string nel parametro s e restituisce un valore string che
	rappresenta la traduzione in linguaggio farfallino di s */
	var testoTrasformato string
	for _, c := range s {
		testoTrasformato += string(TrasformaCarattere(c))
	}
	return testoTrasformato
}
