/* Definizione: Se a , b e c sono numeri naturali e a² + b² = c² , si dice che la terna di numeri a ,
b e c è una terna pitagorica.

Scrivere un programma che legga da standard input un intero soglia>0 e stampi a video tutte le
terne pitagorighe tali che a<soglia , b<soglia e c<soglia .

Oltre alla funzione main() , devono essere definite ed utilizzate almeno le seguenti funzioni:
	1. ÈTernaPitagoriga(a int, b int, c int) bool che riceve in input tre valori interi nei parametri
		a , b e c , e restituisce true se c² è uguale a a² + b² e false altrimenti;
	2. TernePitagoriche(soglia int) che riceve in input un valore intero nel parametro soglia e
		stampa tutte le terne pitagoriche inferiori a soglia ; la funzione deve utilizzare la funzione
		ÈTernaPitagoriga() . */

package main

import "fmt"

func main() {
	var limite int
	fmt.Print("Inserisci un numero intero non negativo: ")
	fmt.Scan(&limite)

	if limite < 0 {
		fmt.Println("La soglia inserita non è positiva")
	} else {
		TernePitagoriche(limite)
	}
}

func ÈTernaPitagoriga(a int, b int, c int) bool {
	if a*a == b*b+c*c {
		return true
	}
	return false
}

func TernePitagoriche(limite int) {
	for a := 1; a < limite; a++ {
		for b := 1; b < limite; b++ {
			for c := 1; c < limite; c++ {
				if ÈTernaPitagoriga(a, b, c) {
					fmt.Printf("(%d, %d, %d)\n", c, b, a)
				}
			}
		}
	}
}
