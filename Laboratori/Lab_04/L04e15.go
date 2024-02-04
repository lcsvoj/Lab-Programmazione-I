/* Scrivere un programma che legga da standard input un numero intero soglia e stampi tutti i
numeri abbondanti (inferiori alla somma dei suoi divisori propri) inferiori a soglia .
Se soglia <= 0 il programma deve stampare "La soglia inserita non è positiva" .*/

package main

import "fmt"

func main() {

	var soglia int
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&soglia)
	if soglia <= 0 {
		fmt.Println("La soglia inserita non è positiva.")
		return
	}

	NumeriAbbondanti(soglia)

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

func ÈAbbondante(n int) bool {
	// restituisce true se n è abbondante
	return n < SommaDivisoriPropri(n)
}

func NumeriAbbondanti(limite int) {
	// stampa tutti i numeri abbondanti inferiori a limite
	fmt.Print("Numeri abbondanti: ")
	for i := 0; i < limite; i++ {
		if ÈAbbondante(i) {
			fmt.Print(i, " ")
		}
	}
}
