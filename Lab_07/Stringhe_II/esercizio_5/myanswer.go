package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	testo := LeggiTesto()
	NumeroNascosto, _ := NumeroNascosto(testo)

	if NumeroNascosto == 0 {
		fmt.Println(("Nessun numero nascosto."))
	} else {
		fmt.Printf("Doppio del numero nascosto: %d (%d * 2)\n", NumeroNascosto*2, NumeroNascosto)
	}
}

func NumeroNascosto(testo string) (int, error) {
	var NumeroNascosto string
	for _, c := range testo {
		if unicode.IsNumber(c) {
			NumeroNascosto += string(c)
		}
	}
	return strconv.Atoi(NumeroNascosto)
}

func LeggiTesto() string {
	scanner := bufio.NewScanner(os.Stdin)
	var result string

	if scanner.Scan() {
		line := scanner.Text()

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Errore nella lettura dell'input:", err)
		}

		result += line + "\n"
	}

	return result
}
