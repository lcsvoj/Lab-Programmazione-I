/* Scrivere un programma che legga da standard input un numero intero n e stampi a video una scala
utilizzando il carattere * (asterisco):
	- La scala è formata da n numeroDiGradini.
	- Ciascun gradino è profondo 3 caratteri e alto 2.
	- Il gradino più in alto deve essere stampato senza alcuna rientranza verso destra; considerando
		poi i successivi numeroDiGradini dall'alto verso il basso, ciascuno di essi rientra (è traslato) verso
		destra rispetto al precedente di due caratteri (spazio) .

Se n è negativo o nullo, anziché stampare la scala il programma deve stampare il messaggio
d'errore Dimensione non sufficiente .

Oltre alla funzione main() , devono essere definite ed utilizzate almeno le seguenti funzioni:
	- StampaGradino(gradino int) che riceve in input un valore intero nel parametro gradino e
		stampa a video un singolo gradino della scala, opportunamente traslato verso destra. Se gradino
		< 0 non stampa nulla. Se gradino = 0 non stampa la rientranza. Se gradino > 0 stampa il
		gradino traslato di gradino * 2 spazi verso destra;
	- StampaScala(numeroDiGradini int) che riceve in input un valore intero nel parametro numeroDiGradini e stampa
		una scala composta da numeroDiGradini numeroDiGradini. */

package main

import "fmt"

func main() {
	var numeroDiGradini int
	fmt.Print("Inserisci il numero dei gradini: ")
	fmt.Scan(&numeroDiGradini)

	if numeroDiGradini <= 0 {
		fmt.Println("Dimensione non sufficiente")
	} else {
		StampaScala(numeroDiGradini)
		fmt.Println()
	}

}

func StampaScala(numeroDiGradini int) {
	for livello := numeroDiGradini; livello > 0; livello-- {
		StampaGradino(livello)
	}
}

func StampaGradino(livelloDelGradino int) {
	// Stampa il gradino padronizzato:
	var profonditaDelGradino, altezzaDelGradino int
	profonditaDelGradino = 3
	altezzaDelGradino = 2

	StampaRientro(livelloDelGradino)
	for p := 0; p < profonditaDelGradino; p++ {
		fmt.Print("*")
	}
	fmt.Println()

	StampaRientro(livelloDelGradino)
	for a := 0; a < altezzaDelGradino-1; a++ {
		fmt.Print("*  ")
	}
	fmt.Println()
}

func StampaRientro(livelloDelGradino int) {
	for i := 0; i < livelloDelGradino; i++ {
		fmt.Print("  ")
	}
}
