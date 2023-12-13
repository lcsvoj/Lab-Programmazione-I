package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	testo := LeggiTesto()
	var dateFormattate []string

	for _, c := range testo {
		s := string(c)
		fmt.Printf("\nAnalizing %v\n", s)
		if !daInvertire(s) {
			fmt.Printf("\nS non è da invertire\n")
			dateFormattate = append(dateFormattate, s)
			fmt.Printf("\nNuova dateOrdinate: %v\n", dateFormattate)
		} else {
			fmt.Printf("\nS da invertire\n")
			dateFormattate = append(dateFormattate, Inverti(s))
			fmt.Printf("\nNuova dateOrdinate: %v\n", dateFormattate)
		}
	}

	sort.Strings(dateFormattate)
	for _, data := range dateFormattate {
		fmt.Println(data)
	}

}

func LeggiTesto() (result []string) {
	// Leggi da standard input un testo su più righe fino a quando viene inserita una riga vuota ( "" )
	fmt.Printf("\nIniziando LeggiTesto()\n")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		if line == " " {
			break
		}
		result = append(result, line)
	}
	fmt.Printf("\nReturn of LeggiTesto: %v\n", result)
	return result
}

func trovaLeBarre(s string) (posizioneDelleBarre []int) {
	fmt.Printf("\nIniziando trovaLeBarre()\n")
	for i, c := range s {
		fmt.Println("trovaLeBarre i ==", i, "c ==", c)
		if c == '/' {
			posizioneDelleBarre = append(posizioneDelleBarre, i)
		}
	}
	fmt.Printf("\nReturn of trovaLeBarre: %v\n", posizioneDelleBarre)
	return posizioneDelleBarre
}

func daInvertire(s string) bool {
	// riceve in input un valore string nel parametro s in cui è codificata una data nel formato
	// aaaa/m/g, aaaa/m/gg, aaaa/mm/g, aaaa/mm/gg, g/m/aaaa, gg/m/aaaa, g/mm/aaaa, o gg/mm/aaaa
	// true se in s è codificata una data nel formato g/m/aaaa, gg/m/aaaa, g/mm/aaaa, o gg/mm/aaaa
	// false altrimenti

	fmt.Printf("\nIniziando daInvertire()\n")

	for i, c := range s {
		if c == '/' {
			if i == 4 {
				fmt.Printf("\nReturn of daInvertire: false\n")
				return false
			} else {
				fmt.Printf("\nReturn of daInvertire: true\n")
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

	fmt.Printf("\nIniziando Inverti()\n")
	posizioneDelleBarre := trovaLeBarre(s)
	var result string

	// concatenate the year
	result += s[posizioneDelleBarre[1]+1:] + "/"

	// concatenate the month
	result += s[posizioneDelleBarre[0]+1:posizioneDelleBarre[1]] + "/"

	// concatenate the day
	result += s[:posizioneDelleBarre[0]]

	fmt.Printf("Return of inverti: %v", result)
	return result
}
