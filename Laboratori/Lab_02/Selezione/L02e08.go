/*Scrivere un programma che legge da standard input un numero intero e stampa "Fizz" se il numero è
multiplo di 3, "Buzz" se il numero è multiplo di 5, "Fizz Buzz" se è multiplo sia di 3 sia di 5, niente
altrimenti.*/

package main

import "fmt"

func main() {

	var n int
	fmt.Print("Inserisci un intero: ")
	fmt.Scan(&n)

	switch {
	case n%3 == 0 && n%5 == 0:
		fmt.Println("Fizz Buzz")
	case n%3 == 0:
		fmt.Println("Fizz")
	case n%5 == 0:
		fmt.Println("Buzz")
	default:
		return
	}

}
