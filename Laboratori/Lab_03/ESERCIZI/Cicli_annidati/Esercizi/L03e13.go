package main

import "fmt"

func main() {

	var n int
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)

	for i := 0; i < 2*n; i++ {
		for j := 0; j < 2*n; j++ {
			if i == j {
				fmt.Print("*")
			} else if j == n-1 {
				fmt.Print("*")
			} else if (i == 0 && j < n-1) || (i == 2*n-1 && j > n-1) {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

}
