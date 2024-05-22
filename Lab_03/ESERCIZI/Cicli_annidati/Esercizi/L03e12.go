/* Scrivere un programma che legga da standard input un numero intero n e che, come mostrato
nell'Esempio di esecuzione, stampi a video un quadrato di lato n in cui:
una diagonale è rappresentata utilizzando il carattere o (lettera o);
l'area del quadrato al di sotto della diagonale è rappresentata con il carattere * (asterisco);
l'area del quadrato al di sopra della diagonale è rappresentata con il carattere + (più);
i caratteri di ogni riga del quadrato sono separati tra di loro da un carattere spazio. */

package main

import "fmt"

func main() {

	var n int
	fmt.Print("Inserisci un intero: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				fmt.Print("o")
			}
			if j > i {
				fmt.Print("+")
			}
			if j < i {
				fmt.Print("*")
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
