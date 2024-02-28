package main

import "fmt"

func main() {
	var n int
	fmt.Print("Inserisci un intero: ")
	fmt.Scan(&n)
	if èPrimo(n) {
		fmt.Println("Il numero", n, "è primo.")
	} else {
		fmt.Println("Il numero", n, "non è primo.")
	}
}

func èPrimo(n int) bool {

	if n <= 1 {
		return false
	}

	for i := n - 1; i >= 2; i-- {
		if n%i == 0 {
			return false
		}
	}
	return true
}
