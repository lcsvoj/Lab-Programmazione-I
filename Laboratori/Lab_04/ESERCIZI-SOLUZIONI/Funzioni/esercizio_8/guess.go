/* Scrivere un programma che legga da standard input un numero intero e stampi la tabellina
corrispondente solo se il numero inserito è compreso tra 1 e 9 (estremi inclusi).
Se il numero
inserito non è valido (inferiore di 1 o superiore a 9 ), il programma deve stampare `Numero non
valido.` .

Oltre alla funzione main() , il programma deve definire e utilizzare le seguenti funzioni:
Tabellina(numero int) bool che riceve in input un valore intero nel parametro numero . Se
numero è compreso tra 1 e 9 (estremi inclusi), la funzione stampa la tabellina associata e
restituisce un valore logico true . Altrimenti, la funzione non stampa nulla e restituisce un valore
logico false . */

package main

import "fmt"

func main() {

	var numero int
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&numero)

	Tabellina(numero)
}

func Tabellina(numero int) bool {
	if numero > 0 && numero < 10 {
		fmt.Print("Tabellina del ", numero, ": ")
		for i := 0; i <= 10; i++ {
			fmt.Print(numero*i, " ")
		}
		fmt.Print("\n")
		return true
	} else {
		fmt.Println("Numero non valido.")
		return false
	}
}
