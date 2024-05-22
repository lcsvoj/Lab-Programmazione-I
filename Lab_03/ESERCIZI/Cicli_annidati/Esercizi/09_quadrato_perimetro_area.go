/*
Scrivere un programma che legga legga da standard input un numero intero n e che stampi a video un quadrato di lato n in cui:
- il perimetro è rappresentato con il carattere * (asterisco);
- l'area interna è rappresentata con il carattere + (più);
- i caratteri di ogni riga del quadrato sono separati tra di loro da un carattere spazio.
*/

package main

import "fmt"

func main() {

	var n int
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		if (i == 0) || (i == n-1) {
			for j := 0; j < n; j++ {
				fmt.Print("* ")
			}
		} else {
			for j := 0; j < n; j++ {
				if (j == 0) || (j == n-1){
					fmt.Print("* ")
				} else { 
					fmt.Print("+ ")
				}
			}
		}
	fmt.Print("\n")
	}
}
