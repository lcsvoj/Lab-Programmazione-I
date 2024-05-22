/*
Scrivere un programma che legga da standard input un intero n > 1 e stampi, utilizzando
il carattere * , il perimetro di due triangoli rettangoli con base e altezza di lunghezza n.
*/

package main

import "fmt"

func main() {

	var n int
	fmt.Print("Inserisci un numero intero maggiore di 1: ")
	fmt.Scan(&n)

	for i := 0; i < (2 * n); i++ {
		for j := 0; j < (2*n - 1); j++ {
			// The center column (height) is always printed
			if j == n-1 {
				fmt.Print("*")
				continue
			}

			// if ((on the upper-left quarter) and (base or diagonal)) or ((on the bottom-right quarter) and (base or diagonal))
			if ((i < n && j < n-1) && (i == 0 || i == j)) || ((i >= n+1 && j >= n) && ((i == 2*n-1) || i == j+1)) {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}
