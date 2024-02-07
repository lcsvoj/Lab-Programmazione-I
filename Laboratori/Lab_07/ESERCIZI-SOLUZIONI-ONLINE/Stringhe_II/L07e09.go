package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {

	chiave, _ := strconv.Atoi(os.Args[1])
	input := LeggiTesto()
	fmt.Println(CifraTesto(input, chiave))

}

func CifraCarattere(c rune, chiave int) rune {
    if !unicode.IsLetter(c) {
        return c
    }

    var offset rune
    if unicode.IsUpper(c) {
        offset = 'A'
    } else {
        offset = 'a'
    }

    // Calcolo corretto per gestire sia chiavi positive che negative
    nuovaPosizione := int(c) - int(offset) + chiave
    if nuovaPosizione < 0 {
        nuovaPosizione += 26 // Assicura che la nuovaPosizione sia positiva
    }
    carattereCifrato := int(offset) + (nuovaPosizione % 26)

    return rune(carattereCifrato)
}


func CifraTesto(s string, chiave int) string {
	testoCifrato := ""
	for _, c := range s {
		testoCifrato += string(CifraCarattere(c, chiave))
	}
	return testoCifrato
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
