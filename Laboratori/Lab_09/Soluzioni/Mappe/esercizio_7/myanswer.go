package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	testo := LeggiTesto()
	paroleContate := (ContaRipetizioni(SeparaParole(testo)))

	fmt.Printf("Parole distinte: %d\n", len(paroleContate))
	for parola, ocorrenze := range paroleContate {
		fmt.Printf("%s: %d\n", parola, ocorrenze)
	}
}

func LeggiTesto() string {
	/* legge da standard input un testo su più righe (alcune delle quali possono
	essere delle righe vuote ( "" )) e terminato dall'indicatore EOF, restituendo
	un valore string in cui è memorizzato il testo letto*/
	scanner := bufio.NewScanner(os.Stdin)

	var testoLetto string
	for scanner.Scan() {
		riga := scanner.Text()
		testoLetto += riga + " "
	}
	return testoLetto
}

func SeparaParole(s string) []string {
	/* restituisce un valore []string in cui sono memorizzate tutte le parole presenti in s */
	var parole []string
	var inizioParola, fineParola int
	var èParola bool

	for i, c := range s {
		if !èParola {
			if unicode.IsLetter(c) {
				inizioParola = i
				èParola = true
			}
			continue
		} else {
			if !unicode.IsLetter(c) {
				fineParola = i
				parole = append(parole, s[inizioParola:fineParola])
				èParola = false
			}
			continue
		}
	}
	return parole
}

func ContaRipetizioni(sl []string) map[string]int {
	ocorrenze := make(map[string]int)
	for _, c := range sl {
		ocorrenze[c]++
	}
	return ocorrenze
}
