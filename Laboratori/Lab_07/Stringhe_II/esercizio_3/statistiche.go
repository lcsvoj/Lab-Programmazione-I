package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	fmt.Println("Inserisci un testo su più righe (termina con Ctrl+D):")
	numeroParole, lunghezza := StatisticheParole(LeggiTesto())

	fmt.Println("Statistiche:")
	fmt.Println("Numero parole:", numeroParole)
	fmt.Println("Lunghezza media:", float64(lunghezza)/float64(numeroParole))
}

func LeggiTesto() (testo string) {
	// leggo del testo dallo standard input (tastiera) e lo concateno nella variabile testo
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		testo += scanner.Text() + "\n"
	}
	return
}

func StatisticheParole(testo string) (numeroParole int, lunghezzaTotale int) {
	// Questa funzione deve svolgere due compiti:
	// 1) contare il numero di parole
	// 2) contare il numero totale di lettere
	// Il punto 2 è semplice: all'interno di un ciclo for in cui scorro ogni carattere del testo controllo se il
	// carattere corrente è una lettera o no; se è una lettera incremento di 1 la variabile lunghezzaTotale
	// Il punto 1 è invece più complesso.
	// Partiamo da questa domanda: quando devo contare una nuova parola?
	// Una possibilità è di contare una nuova parola quando troviamo una lettera e il carattere che
	// la precedeva NON era una lettera (ad esempio uno spazio o un carattere di punteggiatura).
	// Seguendo questa idea possiamo mantenere una variabile booleana che memorizzerà se l'ultimo carattere
	// preso in esame era una lettera o meno. Quando ultimoCarattereLettera è false e incontriamo una lettera,
	// allora sta iniziando una nuova parola (e dobbiamo incrementare il contatore numeroParole).
	// La variabile ultimoCarattereLettera va poi modificata all'interno del ciclo for in modo tale che sia consistente
	// e che prenda un valore true se incontriamo una lettera, e un valore false altrimenti.
	ultimoCarattereLettera := false
	for _, carattere := range testo {

		if unicode.IsLetter(carattere) {
			lunghezzaTotale++
			if !ultimoCarattereLettera {
				numeroParole++
			}
			ultimoCarattereLettera = true
		} else {
			ultimoCarattereLettera = false
		}

	}
	return
}
