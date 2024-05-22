package main

import "fmt"

func main() {
	var n int
	fmt.Print("Inserisci un numero intero: ")
	fmt.Scan(&n)

	if n > 0 {
		var stringaAsterischi string = ""
		for i := 0; i < n; i++ {
			stringaAsterischi += "*"
			fmt.Println(stringaAsterischi)
		}
	} else {
		fmt.Println("Dimensione non sufficiente")
	}
}
