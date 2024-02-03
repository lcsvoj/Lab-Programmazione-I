/*
	Scrivere un programma che legga da standard input un numero intero n e stampi se n è perfetto

[se è uguale alla somma dei suoi divisori propri (per esempio, 6 è perfetto perché 6 = 1 + 2 + 3)]
*/
package main

import "fmt"

func main() {

	var n int
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)

	if ÈPerfetto(n) {
		fmt.Println(n, "è perfetto.")
	} else {
		fmt.Println(n, "non è perfetto.")
	}
}

func SommaDivisoriPropri(n int) int {
	// restituisce la somma dei divisori propri di n
	// se n non ha divisori propri la funzione restituisce 0

	if n <= 1 {
		return 0
	}

	var somma int
	for i := n - 1; i >= 1; i-- {
		if n%i == 0 {
			somma += i
		}
	}
	return somma
}

func ÈPerfetto(n int) bool {
	// restituisce true se n è perfetto e false altrimenti
	if n > 1 && n == SommaDivisoriPropri(n) {
		return true
	} else {
		return false
	}
}
