package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	testo := LeggiTesto()
	paroleSl := SeparaParole(testo)
	paroleMap := ContaRipetizioni(paroleSl)

	fmt.Println("Parole distinte:", len(paroleMap))
	for parola, n := range paroleMap {
		fmt.Printf("%s: %d\n", parola, n)
	}

}

func LeggiTesto() string {
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

func SeparaParole(s string) []string {
	words := strings.FieldsFunc(s, func(c rune) bool {
		// True for every character that is not a letter, signaling a split point:
		return !unicode.IsLetter(c)
	})
	return words
}

func ContaRipetizioni(sl []string) map[string]int {
	parole := make(map[string]int)
	for _, parola := range sl {
		parole[parola]++
	}
	return parole
}
