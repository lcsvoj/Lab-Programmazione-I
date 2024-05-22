/*
Scrivere un programma che:
	1. Genera in modo casuale un numero intero compreso nell'intervallo che va da 1 a 100, estremi
		inclusi (sia numeroGenerato la variabile intera in cui viene memorizzato il numero generato.
	2. Chiede iterativamente all'utente di inserire da standard input un numero intero; ad ogni
		iterazione il programma controlla se il numero inserito è uguale al numero memorizzato:
			- se sono uguali, il programma termina;
			- se sono diversi, il programma fornisce un suggerimento specificando se il numero inserito 
				è più alto o più basso di quello memorizzato
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	
	/* inizializzazione del generatore di numeri casuali */
	rand.Seed(int64(time.Now().Nanosecond()))
	/* generazione di un numero casuale compreso nell'intervallo
	che va da 0 a 99 (estremi inclusi) */
	var numeroGenerato int = rand.Intn(100)
	
	var contatore, tentativo int
	
	for {
		
		contatore++;

		fmt.Print("Tentativo n° ", contatore, ": ")
		fmt.Scan(&tentativo)
		
		if tentativo == numeroGenerato {
			fmt.Println("Hai indovinato in ", contatore, " tentativi!")
			break
		} else {
			if tentativo > numeroGenerato {
				fmt.Println("Troppo alto! Riprova!")
				continue
			} else {
				fmt.Println("Troppo basso! Riprova!")
				continue
			}
		}
	}
}
