package main

import "fmt"

func main() {
	fmt.Print("Inserisci un intero: ")
	var n int
	fmt.Scan(&n)

	s := make([]int, n)
	fmt.Printf("Inserisci %d numeri:\n", n)
	for i := len(s) - 1; i >= 0; i-- {
		fmt.Scan(&s[i])
	}

	fmt.Println("Numeri in ordine inverso:")
	for i := range s {
		fmt.Print(s[i], " ")
	}
	fmt.Println()
}
