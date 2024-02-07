package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	input := LeggiTesto()
	testoFiltrato := FiltraTesto(input)

	fmt.Println("Testo filtrato:")
	for _, s := range testoFiltrato {
		fmt.Println(s)
	}
}

func LeggiTesto() (testo []string) {
	/* Legge da standard input un testo su più righe (alcune delle quali possono essere delle
	righe vuote ( "" ))	e terminato dall'indicatore EOF, restituendo un valore []string in cui
	sono memorizzate le stringhe corrispondenti alle righe del testo letto */

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		testo = append(testo, line)
	}
	return testo
}

func ContieneCifre(s string) bool {
	/* Riceve in input un valore string nel parametro s e restituisce true se almeno un
	carattere in s è una cifra, false altrimenti */
	for _, c := range s {
		if unicode.IsDigit(c) {
			return true
		}
	}
	return false
}

func FiltraTesto(s1 []string) (testoFiltrato []string) {
	/* Riceve in input un valore []string nel parametro sl e restituisce un valore []string in cui
	sono memorizzate le stringhe di sl in cui compaiono cifre; deve utilizzare ContieneCifre() */
	for _, s := range s1 {
		if ContieneCifre(s) {
			testoFiltrato = append(testoFiltrato, s)
		}
	}
	return testoFiltrato
}
