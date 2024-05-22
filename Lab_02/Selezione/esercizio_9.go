/*

Scrivere un programma che legge da standard input un intero n e stampa a video se il numero è pari o
dispari.

*/


package main

import "fmt"

func main() {
	
	var numero int
	fmt.Print("Inserisci numero: ")
	fmt.Scan(&numero)

	if numero % 2 == 0 {
		fmt.Println(numero, "è pari")
	} else {
		fmt.Println(numero, "è dispari")
	}

}
