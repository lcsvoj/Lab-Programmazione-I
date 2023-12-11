/* Scrivere un programma che legga da standard input un numero intero `soglia` e stampi tutti i
numeri primi inferiori a soglia . Se soglia <= 0 il programma deve stampare La soglia inserita
non è positiva .

Oltre alla funzione main() , il programma deve definire e utilizzare le seguenti funzioni:
1. una funzione ÈPrimo(n int) bool che riceve in input un valore intero nel parametro n e
restituisce true se n è primo e false altrimenti;
2. una funzione NumeriPrimi(limite int) che riceve in input un valore intero nel parametro
limite e stampa tutti i numeri primi inferiori a limite ; la funzione deve utilizzare la funzione
ÈPrimo() . */

package main

import "fmt"

func main() {

	var soglia int
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&soglia)

	if soglia <= 0 {
		fmt.Println("La soglia inserita non è positiva")
		return
	}

	NumeriPrimi(soglia)
}

func ÈPrimo(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func NumeriPrimi(limite int) {
	fmt.Println("Numeri primi inferiori a", limite)
	for i := 2; i < limite; i++ {
		if ÈPrimo(i) {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}
