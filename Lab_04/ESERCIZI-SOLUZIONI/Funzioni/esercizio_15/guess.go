/* Scrivere un programma che legga da standard input un n intero soglia e stampi tutti i
numeri abbondanti inferiori a soglia . Se soglia <= 0 il programma deve stampare La soglia
inserita non è positiva.

Definizione: Un n naturale è abbondante se è inferiore alla somma dei suoi divisori propri (per
esempio, 12 è abbondante perché la somma dei suoi divisori propri ( 1 , 2 , 3 , 4 , 6 ) è 16 ).

Oltre alla funzione main() , il programma deve definire e utilizzare le seguenti funzioni:
- una funzione SommaDivisoriPropri(n int) int che riceve in input un valore intero nel parametro
n e restituisce la somma dei divisori propri di n . Se n non ha divisori propri la funzione
restituisce 0;
- una funzione ÈAbbondante(n int) bool che riceve in input un valore intero nel parametro n e
restituisce true se n è abbondante, false altrimenti; la funzione deve utilizzare la funzione
SommaDivisoriPropri();
- una funzione NumeriAbbondanti(limite int) che riceve in input un valore intero nel parametro
limite e stampa tutti i numeri abbondanti inferiori a limite ; la funzione deve utilizzare la
funzione ÈAbbondante(). */

package main

import "fmt"

func main() {
	var limite int
	fmt.Print("Inserisci un numero intero non negativo: ")
	fmt.Scan(&limite)

	if limite < 0 {
		fmt.Println("La soglia inserita non è positiva")
	} else {
		NumeriAbbondanti(limite)
	}

}

func NumeriAbbondanti(limite int) {
	fmt.Print("Numeri abbondanti: ")
	for n := 1; n < limite; n++ {
		if ÈAbbondante(n) {
			fmt.Print(n, " ")
		}
	}
	fmt.Print("\n")
}

func ÈAbbondante(n int) bool {
	if n < SommaDivisoriPropri(n) {
		return true
	}
	return false
}

func SommaDivisoriPropri(n int) int {
	var somma int
	for divisore := 1; divisore < n; divisore++ {
		if n%divisore == 0 {
			somma += divisore
		}
	}
	return somma
}
