/* Scrivere un programma che, dopo aver richiesto all'utente di inserire da standard input un numero
intero n , stampi a video la corrispondente tabellina (moltiplicando n per i numeri naturali da 1 a
10) */

package main

import "fmt"

func main() {

	var n int

	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)

	for i := 1; i < 11; i++ {
		fmt.Println(i, "x", n, "=", i*n)
	}

}
