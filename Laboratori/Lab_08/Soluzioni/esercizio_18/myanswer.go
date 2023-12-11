package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	date := LeggiTesto()

	var output string
	for i := range date {
		output += Formatta(date[i]) + "\n"
	}
	fmt.Print(output)
}

func LeggiTesto() []string {
	// legge da standard input un testo su più righe e terminato da una riga vuota ( "" ),
	// restituendo un valore []string in cui sono memorizzate le righe del testo letto

	scanner := bufio.NewScanner(os.Stdin)
	var result []string

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Errore nella lettura dell'input:", err)
		}

		result = append(result, line)
	}

	return result

}

func Formatta(s string) string {
	// riceve in input un valore string nel parametro s in cui è codificata una data nel
	// formato aaaa/m/g, aaaa/m/gg, aaaa/mm/g o aaaa/mm/gg, e restituisce un valore
	// string in cui la data in input è codificata nel formato aaaa-mm-gg

	var doveSonoLeBarre []int
	for i, c := range s {
		if c == '/' {
			doveSonoLeBarre = append(doveSonoLeBarre, i)
		}
	}

	anno := s[:doveSonoLeBarre[0]]

	mese := s[doveSonoLeBarre[0]+1 : doveSonoLeBarre[1]]
	if len(mese) == 1 {
		mese = "0" + mese
	}

	giorno := s[doveSonoLeBarre[1]+1:]
	if len(giorno) == 1 {
		giorno = "0" + giorno
	}

	return anno + "-" + mese + "-" + giorno

}
