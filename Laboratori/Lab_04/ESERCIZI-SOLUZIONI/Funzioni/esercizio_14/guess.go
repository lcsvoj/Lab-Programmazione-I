/* Scrivere un programma che legga da standard input un numero intero n e stampi se n è perfetto
oppure no.

Oltre alla funzione main() , il programma deve definire e utilizzare le seguenti funzioni:
- una funzione SommaDivisoriPropri(n int) int che riceve in input un valore intero nel parametro
n e restituisce la somma dei divisori propri di n . Se n non ha divisori propri la funzione
restituisce 0 ;
- una funzione ÈPerfetto(n int) bool che riceve in input un valore intero nel parametro n e
restituisce true se n è perfetto e false altrimenti; la funzione deve utilizzare la funzione
SommaDivisoriPropri() . */

package main

import "fmt"

func main() {

	var numero int
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&numero)

	if ÈPerfetto(numero) && numero != 0 {
		fmt.Println(numero, "è perfetto")
	} else {
		fmt.Println(numero, "non è perfetto")
	}

}

func SommaDivisoriPropri(n int) int {
	var somma int
	for i := 1; i < n; i++ {
		if n%i == 0 {
			somma += i
		}
	}
	return somma
}

func ÈPerfetto(n int) bool {
	return SommaDivisoriPropri(n) == n
}
