package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	testo := LeggiTesto()
	fmt.Println("Testo invertito:", InvertiStringa(testo))
}

func LeggiTesto() string {

	fmt.Println("Inserisci un testo su più righe (termina con riga vuota): ")
	scanner := bufio.NewScanner(os.Stdin)

	var righe string
	for scanner.Scan() {
		riga := scanner.Text()
		if riga == "" {
			break
		}
		righe += riga + "\n"
	}
	return righe
}

func InvertiStringa(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
