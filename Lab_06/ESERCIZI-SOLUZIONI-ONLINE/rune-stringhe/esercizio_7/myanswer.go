package main

import (
	"fmt"
)

func main() {
	var stringa string
	var maiuscole, minuscole int
	var vocale, consonanti int

	fmt.Print("Inserisci la stringa: ")
	fmt.Scan(&stringa)

	for _, c := range stringa {
		if èLetteraValida(c) {
			if èMaiuscola(c) {
				maiuscole++
			} else {
				minuscole++
			}
			if èVocale(c) {
				vocale++
			} else {
				consonanti++
			}
		}
	}

	fmt.Printf("Maiuscole: %d \nMinuscole: %d \nVocali: %d \nConsonanti: %d \n", maiuscole, minuscole, vocale, consonanti)

}

func èLetteraValida(c rune) bool {
	if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {
		return true
	}
	return false
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
