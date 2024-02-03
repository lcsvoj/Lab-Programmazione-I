package main

import "fmt"

func main() {

	var n int
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)

	for riga := 0; riga < n; riga++ {
		for colonna := 0; colonna <= 2*n; colonna++ {
			if colonna%2 == 0 {
				// Caso perimetro
				if riga == 0 || riga == n-1 || colonna == 0 || colonna == 2*n {
					fmt.Print("*")
				} else {
					// Caso area interna
					fmt.Print("+")
				}
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

}
