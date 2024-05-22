package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	testo := LeggiTesto()
	var dateOrdinate []string

	for _, c := range testo {
		s := string(c)
		dateOrdinate = append(dateOrdinate, Inverti(s))

	}

	sort.Strings(dateOrdinate)
	for _, data := range dateOrdinate {
		fmt.Println(data)
	}

}

func LeggiTesto() (result []string) {
	// Leggi da standard input un testo su più righe fino a quando viene inserita una riga vuota ( "" )
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		if line == " " {
			break
		}
		result = append(result, line)
	}
	return result
}

func trovaLeBarre(s string) (posizioneDelleBarre []int) {
	for i, c := range s {
		if c == '/' {
			posizioneDelleBarre = append(posizioneDelleBarre, i)
		}
	}
	return posizioneDelleBarre
}

func daInvertire(s string) bool {
	// riceve in input un valore string nel parametro s in cui è codificata una data nel formato
	// aaaa/m/g, aaaa/m/gg, aaaa/mm/g, aaaa/mm/gg, g/m/aaaa, gg/m/aaaa, g/mm/aaaa, o gg/mm/aaaa
	// true se in s è codificata una data nel formato g/m/aaaa, gg/m/aaaa, g/mm/aaaa, o gg/mm/aaaa
	// false altrimenti

	for i, c := range s {
		if c == '/' {
			if i == 4 {
				return false
			} else {
				return true
			}
		}
	}
	return false
}

func Inverti(s string) string {
	// riceve in input un valore string nel parametro s e restituisce un valore
	// string in cui il primo carattere è l'ultimo che definisce s , il secondo carattere è il
	// penultimo che definisce s , ... e l'ultimo carattere è il primo che definisce s

	posizioneDelleBarre := trovaLeBarre(s)

	primoCampo := s[:posizioneDelleBarre[0]]
	if len(primoCampo) == 1 {
		primoCampo = "0" + primoCampo
	}

	secondoCampo := s[posizioneDelleBarre[0]+1 : posizioneDelleBarre[1]]
	if len(secondoCampo) == 1 {
		secondoCampo = "0" + secondoCampo
	}

	terzoCampo := s[posizioneDelleBarre[1]+1:]
	if len(terzoCampo) == 1 {
		terzoCampo = "0" + terzoCampo
	}

	if !daInvertire(s) {
		return (primoCampo + "-" + secondoCampo + "-" + terzoCampo)
	}

	return (terzoCampo + "-" + secondoCampo + "-" + primoCampo)
}
