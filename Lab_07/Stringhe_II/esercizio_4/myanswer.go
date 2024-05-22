package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	testo := LeggiTesto()
	fmt.Println(invertiStringa(testo))
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

func invertiStringa(s string) string {
	var stringaInvertita string

	for _, c := range s {
		stringaInvertita = string(c) + stringaInvertita
	}

	return stringaInvertita
}
