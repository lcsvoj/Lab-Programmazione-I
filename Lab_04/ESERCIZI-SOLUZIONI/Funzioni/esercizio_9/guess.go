/* Scrivere un programma che legga da standard input una sequenza di numeri interi compresi tra 1 e
9 (estremi inclusi) e stampi per ognuno di essi la tabellina corrispondente. Il programma si
interrompe quando viene inserito in input un numero non valido (inferiore a 1 o superiore a 9 )
stampando Programma terminato. .

Oltre alla funzione main() , il programma deve definire e utilizzare le seguenti funzioni:
Tabellina(numero int) bool che riceve in input un valore intero nel parametro numero . Se
numero è compreso tra 1 e 9 (estremi inclusi), la funzione stampa la tabellina associata e
restituisce un valore logico true . Altrimenti, la funzione non stampa nulla e restituisce un valore
logico false. */

package main

import "fmt"

func main() {

	var numero int

	for {

		fmt.Print("Inserisci un numero: ")
		fmt.Scan(&numero)

		if !Tabellina(numero) {
			break
		}
	}

	fmt.Println("Programma terminato.")
}

func Tabellina(numero int) bool {
	if numero > 0 && numero < 10 {
		fmt.Print("Tabellina del ", numero, ": ")
		for i := 0; i <= 10; i++ {
			fmt.Print(numero*i, " ")
		}
		fmt.Print("\n")
		return true
	}
	return false
}
