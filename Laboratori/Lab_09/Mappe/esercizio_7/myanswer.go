package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	// legga da standard input un testo su più righe
	testoInput := LeggiTesto()

        // Lista il numero di occorrenze di ogni parola
        parole := ContaRipetizioni(SeparaParole(testoInput))
        
	// stampa a video il numero di parole distinte presenti nel testo
        fmt.Println("Parole distinte:", len(parole))        

	// stampa a video la lista di parole distinte con il relativo numero di occorrenze nel testo
        for key, value := range parole {
                fmt.Printf("%s: %d\n", key, value)
        }
}

func LeggiTesto() (testoInput string) {
	/* legge da standard input un testo su più righe (alcune delle quali possono essere delle righe vuote) e terminato dall'indicatore EOF, restituendo un valore string in cui è memorizzato il testo letto */

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		testoInput += line + " "
	}
	return testoInput
}

func SeparaParole(s string) (parole []string) {
	/* riceve in input un valore string nel parametro s e restituisce un valore []string in cui sono memorizzate tutte le parole presenti in s */

        var ÈParola bool
	var inizioParola int
	for i, c := range s {
		// se non abbiamo trovato l'inizio di una parola:
		if !ÈParola {
		        if unicode.IsLetter(c) {
		                inizioParola = i
		                ÈParola = true
		        }
	                continue
		}
		// se abbiamo trovato l'inizio di una parola:
		if ÈParola {
		        if (!unicode.IsLetter(c)) || (i == len(s) - 1) {
		                
		                if unicode.IsLetter(c) && i == len(s)-1 {
		                        parole = append(parole, s[inizioParola:])
		                        ÈParola = false
		                        continue
		                }
		                
		                parole = append(parole, s[inizioParola: i])
		                ÈParola = false
		        }
		}
	}
	return parole
}

func ContaRipetizioni(sl []string) map[string]int {
/* riceve in input un valore []string nel parametro sl e restituisce un valore map[string]int in cui, per ogni parola presente in sl , è memorizzato il numero di occorrenze della parola in sl */

        parole := make(map[string]int)
        for _, c := range sl {
                parole[c]++
        }        
        return parole
}

