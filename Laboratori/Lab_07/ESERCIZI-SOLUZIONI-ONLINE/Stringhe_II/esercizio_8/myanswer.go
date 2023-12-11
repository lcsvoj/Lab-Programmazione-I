package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {

	testo := LeggiTesto()

	fmt.Printf("Risultato:\n%s\n", Spazia(testo))

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

func Spazia(s string) string {
	var stringaSpaziata string

	for _, c := range s {
		if c == '\n' {
			stringaSpaziata += "\n"
			continue
		}

		if !unicode.IsSpace(c) {
			stringaSpaziata += string(c) + " "
		}
	}
	return stringaSpaziata
}
