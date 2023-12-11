package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Inserisci un numero intero: ")
	fmt.Scan(&n)
	if n <= 0 {
		fmt.Println("Dimensione non sufficiente")
	} else {

		var stringaSpazi string
		var stringaAsterischi string = "*"

		for i := 0; i < n; i++ {
			stringaSpazi += " "
		}

		for i := 0; i < 2*n+1; i++ {
			if i < n {
				fmt.Print(stringaSpazi[:len(stringaSpazi)-i])
				fmt.Print(stringaAsterischi, "\n")
				stringaAsterischi += "**"
			} else if i == n {
				fmt.Print(stringaAsterischi, "\n")
			} else {
				fmt.Print(stringaSpazi[:(i - len(stringaSpazi))])
				fmt.Print(stringaAsterischi[:(len(stringaAsterischi)+(n-i)*2)], "\n")
			}
		}
	}
}
