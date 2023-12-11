package main

import (
	"fmt"
	"unicode"
)

func main() {
	var stringa string
	var vocaleMaiuscole, vocaleMinuscole int
	var consonantiMaiuscole, consonantiMinuscole int

	fmt.Print("Inserisci la stringa: ")
	fmt.Scan(&stringa)

	for _, c := range stringa {
		if !unicode.IsLetter(c) {
			continue
		}

		if èMaiuscola(c) {
			if èVocale(c) {
				vocaleMaiuscole++
			} else {
				consonantiMaiuscole++
			}
		} else {
			if èVocale(c) {
				vocaleMinuscole++
			} else {
				consonantiMinuscole++
			}
		}
	}

	fmt.Printf("Vocali maiuscole: %d \nConsonanti maiuscoli: %d \nVocali minuscole: %d \nConsonanti minuscole: %d \n", vocaleMaiuscole, consonantiMaiuscole, vocaleMinuscole, consonantiMinuscole)

}

func èMaiuscola(c rune) bool {
	if c >= 'A' && c <= 'Z' {
		return true
	}
	return false
}

func èVocale(c rune) bool {
	switch {
	case c == 'a' || c == 'A':
		return true
	case c == 'e' || c == 'E':
		return true
	case c == 'i' || c == 'I':
		return true
	case c == 'o' || c == 'O':
		return true
	case c == 'u' || c == 'U':
		return true
	default:
		return false
	}
}
