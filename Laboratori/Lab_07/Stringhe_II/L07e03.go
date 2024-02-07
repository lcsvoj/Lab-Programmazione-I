/* Scrivere un programma che:
1. legga da standard input un testo su più righe;
2. termini la lettura quando, premendo la combinazione di tasti Ctrl+D , viene inserito da standard input l'indicatore End-Of-File (EOF);
3. stampi a video le seguenti statistiche relative al testo letto:
	3.1. il numero di parole presenti nel testo;
	3.2. la lunghezza media delle parole presenti nel testo. */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {

	var numeroParole, lunghezzaTotale float64

	// leggi il testo e separa le loro righe
	testo := LeggiTesto()
	righe := strings.Split(testo, "\n")

	// calcola le statistiche di ogni riga
	for _, riga := range righe {
		numeroParoleRiga, lunghezzaTotaleRiga := StatisticheParole(riga)
		numeroParole += float64(numeroParoleRiga)
		lunghezzaTotale += float64(lunghezzaTotaleRiga)
	}

	// stampa i risultati
	fmt.Printf("Statistiche:\nNumero parole: %.0f\nLunghezza media: %f", numeroParole, lunghezzaTotale/numeroParole)
}

func LeggiTesto() string {
	fmt.Println("Inserisci un testo su più righe (termina con Ctrl+D):")

	// legge da standard input un testo su più righe terminato dall'indicatore EOF ( CTRL+D )
	scanner := bufio.NewScanner(os.Stdin)

	var testo string
	for scanner.Scan() {
		testo += scanner.Text() + "\n"
	}
	return testo
}

func StatisticheParole(s string) (int, int) {
	// restituisce il numero di parole presenti e la loro lunghezza media
	// numeri, punteggiatura e caratteri di spaziatura intervallano parole diverse

	parole := TrovaParole(s)

	// conta le parole
	numeroParole := len(parole)

	// calcola la media della loro lunghezze
	lunghezzaTotale := 0
	for _, parola := range parole {
		lunghezzaTotale += len(parola)
	}

	// ritorna i risultati
	return numeroParole, lunghezzaTotale
}

func TrovaParole(s string) []string {
	// restituice una slice con le parole della stringa
	// numeri, punteggiatura e caratteri di spaziatura intervallano parole diverse

	parole := make([]string, 0)
	var inizioParola, fineParola int
	var tracciandoParola bool = false

	for i, c := range s {
		if !tracciandoParola {
			if unicode.IsLetter(c) {
				tracciandoParola = true
				inizioParola = i
				continue
			} else {
				continue
			}
		} else {
			if unicode.IsLetter(c) {
				continue
			} else {
				fineParola = i
				parole = append(parole, s[inizioParola:fineParola])
				tracciandoParola = false
			}
		}
	}

	if tracciandoParola {
		parole = append(parole, s[inizioParola:])
	}

	return parole
}
