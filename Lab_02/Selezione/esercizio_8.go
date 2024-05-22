/*

Scrivere un programma che legge da standard input un numero intero e:
- stampa "Fizz" se il numero è multiplo di 3,
- stampa "Buzz" se il numero è multiplo di 5,
- stampa "Fizz Buzz" se è multiplo sia di 3 sia di 5,
- niente altrimenti.

*/

package main

import "fmt"

func main() {

	var numero int
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&numero)

	if numero%3 == 0 {
		fmt.Println("Fizz ")
	}
	if numero%5 == 0 {
		fmt.Println("Buzz")
	}

}
