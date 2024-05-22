package main

import (
	"fmt"
	"math/rand"
	"time"
)

const FACCEDADO = 6

func main() {
	// inizializzazione del generatore di numeri casuali
	rand.Seed(int64(time.Now().Nanosecond()))

	var numeroGiocatori, totaleTurni int
	var giocatori, vincitori []string

	// lettura del numero di giocatori
	fmt.Print("Inserisci il numero di giocatori: ")
	fmt.Scan(&numeroGiocatori)

	// lettura dei nomi dei giocatori
	giocatori = make([]string, numeroGiocatori)
	fmt.Printf("Inserisci i nomi dei %d giocatori:\n", numeroGiocatori)
	for indiceGiocatore := range giocatori {
		fmt.Scan(&giocatori[indiceGiocatore])
	}

	// lettura del numero di turni di gioco
	fmt.Print("Inserisci il numero di turni: ")
	fmt.Scan(&totaleTurni)

	// slice utilizzata per memorizzare il nome del vincitore di ogni turno
	vincitori = make([]string, totaleTurni)

	for turno := 0; turno < totaleTurni; turno++ {
		// per ogni turno di gioco devo simulare il lancio di tutti i giocatori e
		// trovare chi ha effettuato il punteggio migliore
		fmt.Printf("=== Turno %d ===\n", turno+1)
		var lancioMigliore int
		for _, giocatore := range giocatori {
			// per ogni giocatore simulo due lanci di dato (con la funzione rand.Intn)
			lancio1, lancio2 := rand.Intn(FACCEDADO)+1, rand.Intn(FACCEDADO)+1
			fmt.Printf("\tGiocatore %s, lanci di valore: %d e %d\n", giocatore, lancio1, lancio2)
			// verifico se il totale ottenuto dai lanci effettuati è migliore di quello dei giocatori precedenti
			if totale := lancio1 + lancio2; totale > lancioMigliore {
				// se il giocatore ha effettuato dei lanci migliori allora salvo il suo nome nella slice dei vincitori
				vincitori[turno] = giocatore
				lancioMigliore = totale
			}
		}
		fmt.Printf("\tTurno %d, vincitore: %s\n", turno+1, vincitori[turno])
	}

	// stampa del riepilogo finale
	fmt.Println("Vittorie:")
	for indice, vincitore := range vincitori {
		fmt.Printf("Vincitore turno %d: %s\n", indice+1, vincitore)
	}
}
