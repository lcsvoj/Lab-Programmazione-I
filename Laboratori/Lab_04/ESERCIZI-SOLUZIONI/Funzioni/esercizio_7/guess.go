/* Scrivere un programma che legga da standard input un numero intero e stampi la tabellina
corrispondente solo se il numero inserito è compreso tra 1 e 9 (estremi inclusi).
Oltre alla funzione main() , il programma deve definire e utilizzare le seguenti funzioni:
Tabellina(numero int) che riceve in input un valore intero nel parametro numero e, se numero
è compreso tra 1 e 9 (estremi inclusi), stampa la tabellina associata, altrimenti non stampa
nulla. */

package main

import "fmt"

func main() {

	var numero int
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&numero)
	Tabellina(numero)
}

func Tabellina(numero int) {
	if numero > 0 && numero < 10 {
		fmt.Print("Tabellina del ", numero, ": ")
		for i := 0; i <= 10; i++ {
			fmt.Print(numero*i, " ")
		}
		fmt.Print("\n")
	}
}
