package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	chiave := LeggiNumero()
	testo := LeggiTesto()

	testoCifrato := CifraTesto(testo, chiave)
	fmt.Printf("Testo cifrato:\n%s\n", testoCifrato)
}

func LeggiNumero() int {
	fmt.Print("Inserisci un numero: ")
	var n int
	fmt.Scan(&n)
	return n
}

func LeggiTesto() string {
	fmt.Print("Inserisci un testo su più righe (termina con CTRL D):\n")
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

func CifraCarattere(c rune, chiave int) rune {
	var carattereCrifrato, base rune

	if !unicode.IsLetter(c) {
		return c
	} else {
		if unicode.IsUpper(c) {
			base = 'A'
		} else {
			base = 'a'
		}
	}

	adjustedKey := (chiave%26 + 26) % 26
	carattereCrifrato = base + (c-base+rune(adjustedKey))%26

	return carattereCrifrato
}

func CifraTesto(s string, chiave int) string {
	var testoCifrato string
	for _, c := range s {
		testoCifrato += string(CifraCarattere(c, chiave))
	}
	return testoCifrato
}
