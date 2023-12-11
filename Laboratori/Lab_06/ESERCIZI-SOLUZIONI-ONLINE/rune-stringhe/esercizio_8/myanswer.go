package main

import (
	"fmt"
)

func main() {
	var stringa string
	var maiuscole, minuscole string

	fmt.Print("Inserisci la stringa: ")
	fmt.Scan(&stringa)

	for _, c := range stringa {
		maiuscole += string(inMaiuscolo(c))
		minuscole += string(inMinuscolo(c))
	}

	fmt.Printf("Testo maiuscolo: %s \nTesto minuscolo: %s\n", maiuscole, minuscole)

}

func inMaiuscolo(c rune) rune {
	if c >= 'a' && c <= 'z' {
		return c - 32
	}
	return c
}

func inMinuscolo(c rune) rune {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}
