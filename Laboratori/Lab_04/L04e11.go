/* Scrivere un programma che legga da standard input un numero intero e, come mostrato
nell'Esempio d'esecuzione, stampi a video una schacchiera di dimensione n x n utilizzando i
caratteri * (asterisco) e + (più). */

package main

import "fmt"

func main() {
	var n int
	fmt.Print("Inserisci un intero: ")
	fmt.Scan(&n)
	StampaScacchiera(n)
}

func StampaRigaInizioAsterisco(lunghezza int) {
	// stampa in modo alternato i caratteri * e + (partendo da *) nella lunghezza data
	caratteri := []string{"*", "+"}
	for j := 0; j < lunghezza; j++ {
		fmt.Print(caratteri[j%2])
	}
	fmt.Println()
}

func StampaRigaInizioPiù(lunghezza int) {
	// stampa in modo alternato i caratteri * e + (partendo da +) nella lunghezza data
	caratteri := []string{"+", "*"}
	for j := 0; j < lunghezza; j++ {
		fmt.Print(caratteri[j%2])
	}
	fmt.Println()
}

func StampaScacchiera(dimensione int) {
	// se dimensione <= 0 , non stampa nulla
	if dimensione <= 0 {
		return
	}
	// stampa la scacchiera nella dimensione data
	for i := 0; i < dimensione; i++ {
		if i%2 == 0 {
			StampaRigaInizioAsterisco(dimensione)
		} else {
			StampaRigaInizioPiù(dimensione)
		}
	}
}
