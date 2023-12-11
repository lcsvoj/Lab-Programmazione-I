/* Scrivere un programma che legga da standard input un numero intero e stampi a video una schacchiera di dimensione n x n utilizzando i
caratteri * (asterisco) e + (più).

Oltre alla funzione main() , il programma deve definire e utilizzare le seguenti funzioni:
- StampaRigaInizioAsterisco(lunghezza int) che riceve in input la lunghezza della riga da
stampare nel parametro lunghezza e stampa in modo alternato i caratteri * e + (partendo dal
carattere * );
- StampaRigaInizioPiù(lunghezza int) che riceve in input la lunghezza della riga da stampare nel
parametro lunghezza e stampa in modo alternato i caratteri + e * (partendo dal carattere + );
- StampaScacchiera(dimensione int) che riceve in input la dimensione della scacchiera da
stampare nel parametro dimensione e stampa la scacchiera utilizzando le funzioni StampaRigaInizioAsterisco() e StampaRigaInizioPiù() .
Se dimensione <= 0 , non stampa nulla. */

package main

import "fmt"

func main() {
	var dimensione int
	fmt.Print("Inserisci la dimensione: ")
	fmt.Scan(&dimensione)

	StampaScacchiera(dimensione)
}

func StampaRigaInizioAsterisco(lunghezza int) {
	for i := 0; i < lunghezza; i++ {
		if i%2 != 0 {
			fmt.Print("*")
		} else {
			fmt.Print("+")
		}
	}
	fmt.Println()
}

func StampaRigaInizioPiù(lunghezza int) {
	for i := 0; i < lunghezza; i++ {
		if i%2 != 0 {
			fmt.Print("+")
		} else {
			fmt.Print("*")
		}
	}
	fmt.Println()
}

func StampaScacchiera(dimensione int) {
	if dimensione <= 0 {
		return
	}

	for i := 0; i < dimensione; i++ {
		if i%2 == 0 {
			StampaRigaInizioPiù(dimensione)
		} else {
			StampaRigaInizioAsterisco(dimensione)
		}
	}
}
