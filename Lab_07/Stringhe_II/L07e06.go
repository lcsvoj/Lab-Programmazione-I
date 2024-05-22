package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {

	input := LeggiTesto()
	fmt.Println(TraduciTestoInFarfallino(input))
}

func LeggiTesto() string {
	fmt.Println("Inserisci un testo su più righe (termina con riga vuota): ")
	scanner := bufio.NewScanner(os.Stdin)
	var result string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		result += line + "\n"
	}
	return result[:len(result)-1]
}

func TraduciTestoInFarfallino(s string) string {
	testoFarfalino := ""
	for _, c := range s {
		if unicode.IsLetter(c) {
			if unicode.IsUpper(c) {
				if strings.ContainsRune("AEIOUÀÁÈÉÌÍÓÒÚÙ", c) {
					testoFarfalino += string(c) + "F" + string(c)
				} else {
					testoFarfalino += string(c)
				}
			} else {
				if strings.ContainsRune("aeiouèéòóùúàáíì", c) {
					testoFarfalino += string(c) + "f" + string(c)
				} else {
					testoFarfalino += string(c)
				}
			}
		} else {
			testoFarfalino += string(c)
		}
	}
	return testoFarfalino

}
