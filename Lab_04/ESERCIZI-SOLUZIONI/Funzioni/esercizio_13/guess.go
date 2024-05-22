/* Scrivere un programma che legga da standard input un numero intero soglia e stampi tutti i
numeri primi gemelli tali che p sia inferiore a soglia .
Se soglia <= 0 il programma deve stampare La soglia inserita non è positiva .

Oltre alla funzione main() , il programma deve definire e utilizzare le seguenti funzioni:
- una funzione ÈPrimo(n int) bool che riceve in input un valore intero nel parametro n e restituisce true se n è primo e false altrimenti;
- una funzione NumeriPrimiGemelliGemelli(limite int) che riceve in input un valore intero nel parametro
limite e stampa tutte le coppie di numeri primi gemelli tali che p sia inferiore a limite (cfr.
Definizione); la funzione deve utilizzare la funzione ÈPrimo() . */

/* Scrivere un programma che legga da standard input un numero intero `soglia` e stampi tutti i
numeri primi inferiori a soglia . Se soglia <= 0 il programma deve stampare La soglia inserita
non è positiva .

Oltre alla funzione main() , il programma deve definire e utilizzare le seguenti funzioni:
1. una funzione ÈPrimo(n int) bool che riceve in input un valore intero nel parametro n e
restituisce true se n è primo e false altrimenti;
2. una funzione NumeriPrimiGemelli(limite int) che riceve in input un valore intero nel parametro
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

	NumeriPrimiGemelli(soglia)
}

func ÈPrimo(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func NumeriPrimiGemelli(limite int) {
	fmt.Println("Numeri primi gemelli inferiori a", limite)

	for p := 3; p+2 < limite; p++ {
		if ÈPrimo(p) && ÈPrimo(p+2) {
			fmt.Print("(", p, ", ", p+2, ") ")
		}
	}
	fmt.Println()
}
