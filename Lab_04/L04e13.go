/* Scrivere un programma che legga da standard input un numero intero soglia e stampi tutti i
numeri primi gemelli (due numeri primi p e q sono gemelli se p = q + 2) tali che p sia inferiore a soglia.
Se soglia <= 0 il programma deve stampare "La soglia inserita non è positiva". */

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

	NumeriPrimiGemelli(soglia)

}

func ÈPrimo(n int) bool {
	// restituisce true se n è primo e false altrimenti

	for m := n - 1; m > 1; m-- {
		if n%m == 0 {
			return false
		}
	}
	return true
}

func NumeriPrimiGemelli(limite int) {
	// stampa tutte le coppie di numeri primi gemelli tali che p sia inferiore a limite

	for p := 2; p < limite; p++ {
		for q := 2; q < p; q++ {
			if ÈPrimo(p) && ÈPrimo(q) {
				if p == q+2 {
					fmt.Printf("(%d, %d) ", q, p)
				}
			}
		}
	}
}
